package main

import (
	"log"
	"os"
	"path"
	"strconv"

	"github.com/jusplat/web-fw-test-container/configs"
	"github.com/jusplat/web-fw-test-container/pkg/dockercompose"
	"github.com/jusplat/web-fw-test-container/pkg/git"
)

var (
	ENV_FILE  = "ENV_FILE"
	workspace = "work"
)

func main() {

	if os.Getenv(ENV_FILE) == "" {
		os.Setenv(ENV_FILE, "./tools/env.toml")
	}

	// 環境変数の読み込み
	config := configs.GetConfig(os.Getenv(ENV_FILE))

	// 起動対象コンテナ取得
	var containerNames []string
	for _, container := range config.Container {
		if enable, err := strconv.ParseBool(container.Enable); err == nil && enable {
			containerNames = append(containerNames, container.Name)
		}
	}
	log.Printf("target containers: %v", containerNames)

	// Dockerコンテナ起動
	dcCient := &dockercompose.DockerCompose{
		Files: []string{config.ComposeFile},
	}
	if err := dcCient.Up([]string{"-d"}, containerNames); err != nil {
		handleError(err)
	}

	// git clone
	// TODO: 要リファクタリング
	for _, container := range config.Container {
		enable, _ := strconv.ParseBool(container.Enable)
		if enable && container.Repos != "" {
			if container.Workspace != "" {
				workspace = container.Workspace
			}
			if err := git.Clone(path.Join(container.Name, workspace), container.Branch, container.Repos); err != nil {
				handleError(err)
			}
		}
	}
}

// 共通エラー処理
func handleError(err error) {
	log.Fatal(err.Error())
	os.Exit(1)
}
