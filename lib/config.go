// config
package lib

import (
	"encoding/json"
	"os"
)

type Config struct {
	Listen                  string `json:"listen"`
	ReloadCMD               string `json:"reload-cmd"`
	EntryTemplate           string `json:"config-template-path"`
	DomainCfgSaveDir        string `json:"save-path"`
	WebRoot                 string `json:"web_root"`
	DomainCfgFileNameFormat string `json:"filename-format"`
	savePath                string
}

func NewConfig(path string) (*Config, error) {
	c := &Config{}
	c.savePath = path
	err := c.load(path)
	return c, err
}

func (c *Config) load(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		log.Error(err)
		return err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(c)
	if err != nil {
		log.Error(err)
	}
	return err
}

func (c *Config) Save() error {
	file, err := os.Create(c.savePath)
	if err != nil {
		log.Error(err)
		return err
	}
	defer file.Close()
	data, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		log.Error(err)
		return err
	}
	_, err = file.Write(data)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
