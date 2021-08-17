package configs

type Container struct {
	Name       string `toml:"name"`
	Workspace  string `toml:"workspace"`
	Branch     string `toml:"branch"`
	Repos      string `toml:"repos"`
	Enable     string `toml:"enable"`
	Entrypoint string `toml:"entrypoint"`
}
