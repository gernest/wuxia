package wuxia

type Config struct {
	Source      string   `json:"source"`
	Destination string   `json:'destination"`
	Safe        bool     `json:"safe"`
	Excluede    []string `json:"exclude"`
	Include     []string `json""include"`
	KeepFiles   []string `json:"keep_files"`
	TimeZone    string   `json:"timezone"`
	Encoding    string   `json:"encoding"`
	Port        int      `json:"port"`
	Host        string   `json:"host"`
	BaseURL     string   `json:"base_url"`
}

type System struct {
	Boot    *Boot   `json:"boot"`
	Config  *Config `json:"config"`
	Plan    *Plan   `json:"plan"`
	WorkDir string  `json:"work_dir"`
}

type Boot struct {
	ConfigiFile string            `json:"config_file"`
	PlanFile    string            `json:"plan_file"`
	ENV         map[string]string `json:"env"`
}

type Theme struct {
	Name   string   `json:"name"`
	Author []Author `json:"author"`
}

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
