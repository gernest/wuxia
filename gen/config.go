package gen

import (
	"errors"
	"path/filepath"
)

//Config is the settings needed by the static generatot. It is inspired by the
//jekyll configuration options.
//
// The format can either be json, yaml or toml
// TODO: add yaml and toml support.
type Config struct {
	Name         string   `json:"name" yaml:"name"`
	Source       string   `json:"source" yaml:"source"`
	Destination  string   `json:"destination" yaml:"destination"`
	StaticDir    string   `json:"statc_dir" yaml:"statc_dir"`
	TemplatesDir string   `json:"templates_dir" yaml:"templates_dir"`
	ThemeDir     string   `json:"theme_dir" yaml:"theme_dir"`
	DefaultTheme string   `json:"default_theme" yaml:"default_theme"`
	PluginDir    string   `json:"plugin_dir" yaml:"plugin_dir"`
	Safe         bool     `json:"safe" yaml:"safe"`
	Excluede     []string `json:"exclude" yaml:"exclude"`
	Include      []string `json:"include" yaml:"include"`
}

//DefaultConfig retruns *Config with default settings
func DefaultConfig() *Config {
	return &Config{
		Source:       "src",
		Destination:  "dest",
		StaticDir:    "static",
		ThemeDir:     "themes",
		TemplatesDir: "templates",
		DefaultTheme: "doxsey",
		PluginDir:    "plugins",
		Safe:         true,
		Excluede: []string{
			"CONTRIBUTING.md",
		},
		Include: []string{
			"LICENCE.md",
		},
	}
}

//System configuration for the whole static generator system. The information
//consost of things necessary to bootup, configure and plan the execution of the
//static project..
type System struct {
	Boot   *Boot   `json:"boot"`
	Config *Config `json:"config"`
	Plan   *Plan   `json:"plan"`
}

//Boot necessary info to bootstrap the Generator.
type Boot struct {
	ConfigiFile string            `json:"config_file"`
	PlanFile    string            `json:"plan_file"`
	ENV         map[string]string `json:"env"`
}

//Theme discreption of a theme.
type Theme struct {
	Name   string   `json:"name"`
	Author []Author `json:"author"`
}

//Author description of the author of the project being built.
type Author struct {
	Name     string `json:"name"`
	Github   string `json:"github"`
	Twitter  string `json:"twitter"`
	Linkedin string `json:"linkedin"`
	Email    string `json:"email"`
	Website  string `json:"website"`
}

// Plan is the execution planner object. It states the steps and stages on which
// the execution process should take.
type Plan struct {
	Title string `json:"title"`

	// Modules that are supposed to be loaded before the execution starts. The
	// execution process wont start if one of the dependencies is missing.
	Dependency []string `json:"dependencies"`

	TemplateEngine string      `json:"template_engine"`
	Strategies     []*Strategy `json:"strategies"`
	After          []string    `json:"after"`
}

//Strategy defines what should be done when a file matches a certain pattern.
//the pattern is glob like and matches only on the filenames
//
// The FullMatch field specifies the pattern to be matchig the full file path.
// if set to  false is means the pattern only matches the base of the file path(
// ignoring the directories).
type Strategy struct {
	Title     string   `json:"title"`
	Patterns  []string `json:"patterns"`
	FullMatch bool     `json:"fullMatch"`
	Before    []string `json:"before"`
	Exec      []string `json:"exec"`
	After     []string `json:"after"`
}

//FindStrategy searches for a strategy matching the given filePath that is
//defined in the Plan.
//
//TODO: Support multiple return strategies, since more than one strategy can be
//defined for a given file path.
func (p *Plan) FindStrategy(filePath string) (*Strategy, error) {
	for i := range p.Strategies {
		s := p.Strategies[i]
		m := filePath
		if !s.FullMatch {
			m = filepath.Base(filePath)
		}
		for _, patt := range s.Patterns {
			ok, err := filepath.Match(patt, m)
			if err != nil {
				return nil, err
			}
			if !ok {
				continue
			}
			return s, nil
		}
	}
	return nil, errors.New("can't find strategy for " + filePath)
}

//File is a representation of a file unit as it is passed arouund for
//processing.
// File content is passed as a string so as to allow easy trasition between Go
// and javascript boundary.
type File struct {
	Name     string                 `json:"name"`
	Meta     map[string]interface{} `json:"meta"`
	Contents string                 `json:"contents"`
}

//FileList is a list of files, which can be ordered based on various criterias
type FileList []*File
