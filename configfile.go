package goconfig

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type ConfigFile struct {
	data map[string]map[string]string
}

var (
	DefaultSection string = "default"

	BoolStrings = map[string]bool{
		"0":     false,
		"1":     true,
		"false": false,
		"true":  true,
		"n":     false,
		"y":     true,
		"no":    false,
		"yes":   true,
		"off":   false,
		"on":    true,
	}
)

func (c *ConfigFile) AddSection(section string) bool {
	section = strings.ToLower(section)
	if _, ok := c.data[section]; ok {
		return false // section exists
	}
	c.data[section] = make(map[string]string)
	return true
}

func (c *ConfigFile) AddOption(section, option, value string) bool {
	c.AddSection(section)
	section = strings.ToLower(section)
	option = strings.ToLower(option)
	if _, ok := c.data[section][option]; ok {
		//return false	// option exists
		//we need update vale, so do not return
	}
	c.data[section][option] = value
	return true
}

func (c *ConfigFile) GetRawString(section, option string) (string, error) {
	section = strings.ToLower(section)
	option = strings.ToLower(option)

	if _, ok := c.data[section]; ok {
		if value, ok := c.data[section][option]; ok {
			return value, nil
		}
		return "", errors.New(fmt.Sprintf("Option not found: %s", option))
	}
	return "", errors.New(fmt.Sprintf("Section not found: %s", section))
}

func (c *ConfigFile) GetString(section, option string) (string, error) {
	value, err := c.GetRawString(section, option)
	if err != nil {
		return "", err
	}
	return value, nil
}

func (c *ConfigFile) GetInt(section, option string) (int, error) {
	value, err := c.GetInt64(section, option)
	if err != nil {
		return 0, err
	}
	return int(value), nil
}

func (c *ConfigFile) GetInt64(section, option string) (int64, error) {
	value, err := c.GetRawString(section, option)
	if err != nil {
		return 0, err
	}
	iv, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, err
	}
	return iv, nil
}

func (c *ConfigFile) GetFloat(section, option string) (float64, error) {
	value, err := c.GetRawString(section, option)
	if err != nil {
		return float64(0), err
	}
	fv, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return float64(0), err
	}
	return fv, nil
}

func (c *ConfigFile) GetBool(section, option string) (bool, error) {
	value, err := c.GetRawString(section, option)
	if err != nil {
		return false, err
	}
	bv, ok := BoolStrings[strings.ToLower(value)]
	if ok == false {
		return false, errors.New(fmt.Sprintf("Cound not parse bool value: %s", value))
	}
	return bv, nil
}

// init a new ConfigFile
func NewConfigFile() *ConfigFile {
	c := new(ConfigFile)
	c.data = make(map[string]map[string]string)
	c.AddSection(DefaultSection) // deafult section always exists
	return c
}

// find delimiter first occur position
func firstIndex(l string, delimiter []byte) int {
	for i := 0; i < len(delimiter); i++ {
		if j := strings.Index(l, string(delimiter[i])); j != -1 {
			return j
		}
	}
	return -1
}

// strip comment in value
func stripComments(l string) string {
	for _, c := range []string{" ;", "\t;", " #", "\t#"} {
		if i := strings.Index(l, c); i != -1 {
			l = l[0:i]
		}
	}
	return l
}

func (c *ConfigFile) read(buf *bufio.Reader) error {
	var section, option string
	section = DefaultSection
	for {
		l, err := buf.ReadString('\n') // parse line-by-line
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		l = strings.TrimSpace(l)
		//switch written for readability
		switch {
		case len(l) == 0: //empty line
			continue
		case l[0] == '#': //comment
			continue
		case l[0] == ';': //comment
			continue
		case len(l) >= 3 && strings.ToLower(l[0:3]) == "rem": // comment for windows
			continue
		case l[0] == '[' && l[len(l)-1] == ']': // new section
			option = "" // reset multi-line value
			section = strings.TrimSpace(l[1 : len(l)-1])
			c.AddSection(section)
		case section == "": // not new section and no sectiondefined so far
			return errors.New("Section not found: must start with section")
		default: // other alternatives
			i := firstIndex(l, []byte{'=', ':'})
			switch {
			case i > 0:
				option = strings.TrimSpace(l[0:i])
				value := strings.TrimSpace(stripComments(l[i+1:]))
				value = strings.Trim(value, "\"")
				value = strings.Trim(value, "'")
				c.AddOption(section, option, value)
			case section != "" && option != "":
				// continuation of multi-line value
				prev, _ := c.GetRawString(section, option)
				value := strings.TrimSpace(stripComments(l))
				c.AddOption(section, option, prev+"\n"+value)
			default:
				return errors.New(fmt.Sprintf("Cound not parse line: %s", l))
			}

		}
	}
	return nil
}

func ReadConfigFile(f string) (*ConfigFile, error) {
	file, err := os.Open(f)
	if err != nil {
		return nil, err
	}
	c := NewConfigFile()
	if err = c.read(bufio.NewReader(file)); err != nil {
		return nil, err
	}
	if err = file.Close(); err != nil {
		return nil, err
	}
	return c, nil
}
