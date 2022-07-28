package config

type Config struct {
	NumCerts int `yaml:"num_certs"`
}

var DefaultConfig = Config{
	NumCerts: 3,
}
