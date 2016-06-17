package gen

type File struct {
	Name     string                 `json:"name"`
	Meta     map[string]interface{} `json:'meta"`
	Contents string                 `json:"contents"`
}
