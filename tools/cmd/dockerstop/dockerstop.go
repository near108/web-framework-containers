package main

import (
	"log"
	"os"

	"github.com/jusplat/web-fw-test-container/configs"
	"github.com/jusplat/web-fw-test-container/pkg/dockercompose"
)

var (
	ENV_FILE = ".env"
	config   *configs.Config
)

func main() {
	if os.Getenv(ENV_FILE) == "" {
		os.Setenv(ENV_FILE, "./tools/env.toml")
	}

	// 環境変数の読み込み
	config = configs.GetConfig(os.Getenv(ENV_FILE))

	// Dockerコンテナ停止
	dc := &dockercompose.DockerCompose{
		Files: []string{config.ComposeFile},
	}
	dc.Down()

	// Workディレクトリ内の削除
	if config.ClearWorkspace {
		for _, container := range config.Container {
			if container.Workspace != "" {
				log.Printf("remove workspace at %s\n", container.Workspace)
				if err := os.RemoveAll(container.Workspace); err != nil {
					log.Println(err)
				}
			}
		}
	}
}

// 共通エラー処理
func handleError(err error) {
	log.Fatal(err.Error())
	os.Exit(1)
}
