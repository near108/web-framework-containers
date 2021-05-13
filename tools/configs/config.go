package configs

import (
	"log"
	"os"
	"path"
	"sync"

	"github.com/BurntSushi/toml"
	"github.com/joho/godotenv"
)

type Config struct {
	Envfile        string       `toml:"env"`
	ClearWorkspace bool         `toml:"clearWorkspace"`
	ComposeFile    string       `toml:"composefile"`
	Container      []*Container `toml:"Container"`
}

var config *Config
var once sync.Once

func newConfig(filepath string) *Config {
	var c Config

	// TOMLファイル読み込み
	if _, err := toml.DecodeFile(filepath, &c); err != nil {
		panic(err)
	} else {
		log.Printf("read toml file at %s\n", filepath)
	}

	// .envに記載がある場合は.envの値を読み込む
	if c.Envfile != "" {
		readDotEnv(&c)
		for _, container := range c.Container {
			setEnv(container)
		}
	}
	return &c
}

func setEnv(c *Container) {
	if name := os.Getenv(c.Name); name != "" {
		c.Name = name
	}
	if branch := os.Getenv(c.Branch); branch != "" {
		c.Branch = branch
	}
	if repos := os.Getenv(c.Repos); repos != "" {
		c.Repos = repos
	}
	if enable := os.Getenv(c.Enable); enable != "" {
		c.Enable = enable
	}
	if workspace := os.Getenv(c.Workspace); workspace != "" {
		c.Workspace = path.Join(c.Name, workspace)
		log.Println(c.Workspace)
	}
}

// .envの値を読み込む
func readDotEnv(c *Config) error {
	if err := godotenv.Load(c.Envfile); err != nil {
		return err
	} else {
		log.Printf("read env file at %s\n", c.Envfile)
		return nil
	}
}

func GetConfig(filepath string) *Config {
	once.Do(func() {
		config = newConfig(filepath)
	})
	return config
}
