package valeria

import (
	"io/ioutil"
	"testing"

	"gopkg.in/yaml.v2"
)

var sampleConfig = &Config{
	BaseURL:   "http:/example.com",
	DataDir:   "_data",
	PluginDir: "_plugins",
	SourceDir: "_src",
	OutputDir: "web",
	ThemesDir: "_themes",
	Plugins: []string{
		"load_plugins", "load_files", "load_data", "register_handlers",
	},
	Up: []string{
		"load_plugins", "load_files", "load_data", "register_handlers",
	},
	Down: []string{
		"gh_release",
	},
}

func TestConfig(t *testing.T) {
	data, err := yaml.Marshal(sampleConfig)
	if err != nil {
		t.Fatal(err)
	}
	ioutil.WriteFile("fixture/_valeria.yml", data, 0600)
}

func TestLoadConfig(t *testing.T) {
	_, err := loadConfig("fixtures")
	if err != nil {
		t.Fatal(err)
	}
}
