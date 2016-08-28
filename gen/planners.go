package gen

// strategy for processing markdown files. The files may contain front matter,
// so front matter is extracted from the file content before the contents of the
// file are rendered by a markdown rendering engine.
//
// Matches for files with .md extension.
func markdownStrategy() *Strategy {
	return &Strategy{
		Title:    "markdown",
		Patterns: []string{"*.md"},
		Before:   []string{"front"},
		Exec:     []string{"markdown"},
	}
}

// copies files.
func copyStrategy() *Strategy {
	return &Strategy{
		Title:    "copy",
		Patterns: []string{"*.md"},
		Exec:     []string{"copy"},
	}
}

// is the default plan offered by the generator. It contains default strategis
// which are wired to sensible conventions.
func defaultPlan() *Plan {
	return &Plan{
		Title: "default_plan",
		Dependency: []string{
			"front", "markdown",
		},
		TemplateEngine: "go",
		Strategies: []*Strategy{
			copyStrategy(),
			markdownStrategy(),
		},
	}
}
