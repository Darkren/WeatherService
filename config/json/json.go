// Package json encapsulates structure and methods for
// parsing and getting values from json configuration files
package json

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

// Config represents data type for configuration
type Config map[string]*json.RawMessage

// New parses json string and gets config structure
func New(jsonStr string) (Config, error) {
	return new([]byte(jsonStr))
}

// Load reads file from filePath, parses json and
// gets config structure
func Load(filePath string) (Config, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return new(data)
}

// Section gets config se
func (c Config) Section(key string) (Config, error) {
	section := make(map[string]*json.RawMessage)

	if err := json.Unmarshal(*c[key], &section); err != nil {
		return nil, err
	}

	return section, nil
}

func (c Config) GetString(key string, defaultVal string) string {
	value, err := c.getString(key)
	if err != nil {
		return defaultVal
	}

	return value
}

func (c Config) MustGetString(key string) string {
	value, err := c.getString(key)
	if err != nil {
		panic(err)
	}

	return value
}

func (c Config) getString(key string) (string, error) {
	var value string

	if err := json.Unmarshal(*c[key], &value); err != nil {
		return "", err
	}

	return value, nil
}

func (c Config) GetInt(key string, defaultVal int) int {
	value, err := c.getInt(key)
	if err != nil {
		return defaultVal
	}

	return value
}

func (c Config) MustGetInt(key string) int {
	value, err := c.getInt(key)
	if err != nil {
		panic(err)
	}

	return value
}

func (c Config) getInt(key string) (int, error) {
	var value int

	if err := json.Unmarshal(*c[key], &value); err != nil {
		return 0, err
	}

	return value, nil
}

func (c Config) GetTime(key string, defaultVal time.Time) time.Time {
	value, err := c.getTime(key)
	if err != nil {
		return defaultVal
	}

	return value
}

func (c Config) MustGetTime(key string) time.Time {
	value, err := c.getTime(key)
	if err != nil {
		panic(err)
	}

	return value
}

func (c Config) getTime(key string) (time.Time, error) {
	valueStr, err := c.getString(key)
	if err != nil {
		return time.Now(), err
	}

	value, err := time.Parse("2.1.2006", valueStr)
	if err != nil {
		return time.Now(), err
	}

	return value, nil
}

func new(jsonData []byte) (Config, error) {
	config := make(map[string]*json.RawMessage)

	if err := json.Unmarshal(jsonData, &config); err != nil {
		return nil, err
	}

	return config, nil
}
