package gen

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
