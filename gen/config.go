package gen

type Config struct {
	Source      string   `json:"source"`
	Destination string   `json:'destination"`
	Safe        bool     `json:"safe"`
	Excluede    []string `json:"exclude"`
	Include     string   `json""include"`
	KeepFiles   string   `json:"keep_files"`
	TimeZone    string   `json:"timezone"`
	Encoding    string   `json:"encoding"`
	Port        int      `json:"port"`
	Host        string   `json:"host"`
	BaseURL     string   `json:"base_url"`
}
