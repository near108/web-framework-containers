package main

import (
	"fmt"
	"log"
	"os"
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

	// ワークスペース作成
	for _, container := range config.Container {
		if _, err := os.Stat(container.Workspace); os.IsNotExist(err) {
			os.Mkdir(container.Workspace, 0775)
		}
	}

	// Dockerコンテナ起動
	dc := &dockercompose.DockerCompose{
		Files: []string{config.ComposeFile},
	}
	if err := dc.Up([]string{"-d"}, containerNames); err != nil {
		handleError(err)
	}

	// git clone
	// TODO: 要リファクタリング
	for _, container := range config.Container {
		enable, _ := strconv.ParseBool(container.Enable)
		if enable && container.Repos != "" {
			if err := git.Clone(container.Workspace, container.Branch, container.Repos); err != nil {
				fmt.Printf("Worning: git clone skipped %s\n", container.Name)
			}
		}
	}
}

// 共通エラー処理
func handleError(err error) {
	log.Fatal(err.Error())
	os.Exit(1)
}
