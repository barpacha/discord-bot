package config

type Main struct {
	Discord Discord `yaml:"discord"`
}

type Discord struct {
	Token string `yaml:"token"`
}
