package wuxia

//Config is the settings needed by the static generatot. It is inspired by the
//jekyll configuration options.
//
// The format can either be json, yaml or toml
// TODO: add yaml and toml support.
type Config struct {
	Source      string   `json:"source"`
	Destination string   `json:"destination"`
	Safe        bool     `json:"safe"`
	Excluede    []string `json:"exclude"`
	Include     []string `json:"include"`
}

//DefaultConfig retruns *Config with default settings
func DefaultConfig() *Config {
	return &Config{
		Source:      "src",
		Destination: "dest",
		Safe:        true,
		Excluede: []string{
			"CONTRITUTING", "CONTRIBUTING.md",
		},
		Include: []string{
			"LICENCE.md",
		},
	}
}

//System configuration for the whole static generator system.
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

	TemplateEngine string   `json:"template_engine"`
	Before         []string `json:"before"`
	Exec           []string `json:"exec"`
	After          []string `json:"after"`
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
