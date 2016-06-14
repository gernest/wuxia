package gen

type Config struct {
	Source      string
	Destination string
	Safe        bool
	Excluede    []string
	Include     string
	KeepFiles   string
	TimeZone    string
	Encoding    string
	Port        int
	Host        string
	BaseURL     string
}
