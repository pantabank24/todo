package config

type Configs struct {
	App Fiber
}

type Fiber struct {
	Host string
	Port string
}
