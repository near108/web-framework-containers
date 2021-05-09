package main

import (
	"log"
	"os"

	"github.com/jusplat/web-fw-test-container/configs"
	"github.com/jusplat/web-fw-test-container/pkg/dockercompose"
)

var (
	ENV_FILE = ".env"
)

func main() {
	if os.Getenv(ENV_FILE) == "" {
		os.Setenv(ENV_FILE, "./tools/env.toml")
	}

	// 環境変数の読み込み
	config := configs.GetConfig(os.Getenv(ENV_FILE))

	// Dockerコンテナ停止
	dcClient := &dockercompose.DockerCompose{
		Files: []string{config.ComposeFile},
	}
	dcClient.Down()

}

// 共通エラー処理
func handleError(err error) {
	log.Fatal(err.Error())
	os.Exit(1)
}
