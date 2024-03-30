package config

import (
	"gopkg.in/yaml.v3"
	"os"
	"os/user"
	"strings"
)

type Config struct {
	PubKey   []byte     `yaml:"pub_key"`
	Host     string     `yaml:"host"`
	Port     uint       `yaml:"port"`
	Projects []*Project `yaml:"projects"`
}

type Project struct {
	ID       string `yaml:"id"`
	Repo     string `yaml:"repo"`
	Location string `yaml:"location"`
	Script   string `yaml:"script"`
}

func (p *Project) ResolveLocation() {
	usr, _ := user.Current()
	dir := usr.HomeDir
	p.Location = strings.Replace(p.Location, "~", dir, 1)
}

func Read(path string) (*Config, error) {
	var config Config
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(file, &config)
	for _, p := range config.Projects {
		p.ResolveLocation()
	}

	return &config, err
}

func (c *Config) Write(path string) error {
	buf, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	f, err := os.Create(path)
	defer f.Close()

	_, err = f.Write(buf)
	return err
}
