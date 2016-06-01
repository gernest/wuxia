package base

import (
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

const configFile = "_valeria.yml"

//Config is the configuration object for valor. The supported format is yaml.
type Config struct {
	WorkingDir string   `yaml:"working_dir"`
	BaseURL    string   `yaml:"base_url"`
	DataDir    string   `yaml:"data_dir"`
	SourceDir  string   `yaml:"source_dir"`
	OutputDir  string   `yaml:"output_dir"`
	PluginDir  string   `yaml:"plugin_dir"`
	ThemesDir  string   `yaml:"themes_dir"`
	Plugins    []string `yaml:"plugins"`
	Up         []string `yaml:"up"`
	Down       []string `yaml:"down"`
}

func loadConfig(root string) (*Config, error) {
	cfgFile := filepath.Join(root, configFile)
	data, err := ioutil.ReadFile(cfgFile)
	cfg := &Config{}
	err = yaml.Unmarshal(data, cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
