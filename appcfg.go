package qaecommon

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Task struct {
	Type string
	Name string
	Url  string
}

type Config struct {
	Local  string
	Remote string
}

type AppCfg struct {
	AppName string `yaml:"appname"`
	Image   string `yaml:"image"`
	Tasks   []Task
	Config  []Config
}

func (a *AppCfg) check() error {
	if a.AppName == "" {
		return fmt.Errorf("appname is empty")
	}

	if a.Image == "" {
		return fmt.Errorf("image is empty")
	}

	if len(a.Tasks) == 0 {
		return fmt.Errorf("tasks is nil")
	}

	return nil
}

//LoadAppCfg Load App Config from file
func LoadAppCfg(filename string) (*AppCfg, error) {
	var cfg = &AppCfg{}

	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(content, &cfg)
	fmt.Println(err)
	if err != nil {
		return nil, err
	}

	err = cfg.check()
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
