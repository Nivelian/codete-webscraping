package model

type Config struct {
	Port       string
	StaticPath string `yaml:"static_path"`
}

type Record struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
