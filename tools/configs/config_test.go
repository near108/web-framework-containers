package configs_test

import (
	"log"
	"os"
	"testing"

	"github.com/jusplat/web-fw-test-container/configs"
)

const (
	ORACLE     = "oracle"
	SQLSERVER  = "sqlserver"
	MYSQL      = "mysql"
	POSTGRESQL = "postgresql"
	APACHE     = "apache"
	NGINX      = "nginx"
	LIGHTSPEED = "lightspeed"
	RAILS      = "rails"
	LARAVEL    = "laravel"
	SPRING     = "spring"
	DJANGO     = "django"
	FALSE      = "FALSE"
	TRUE       = "TRUE"
	BRANCH     = "main"
	REPOS      = "http://example.com/sample.git"
)

var (
	color    = "31" // red
	testfile = "./testdata/env_test.toml"
	logfile  = "./testdata/config_test.log"
)

func TestConfig(t *testing.T) {
	// テストログの設定
	setLogfile(logfile)

	// 設定ファイル読み込み
	config := configs.GetConfig(testfile)
	assertEqual(t, "env", config.Envfile, "./testdata/.env_test")

	// 各種設定値読み込み
	for _, container := range config.Container {
		switch container.Name {
		case ORACLE:
			assertEqual(t, ORACLE, container.Enable, FALSE)
		case SQLSERVER:
			assertEqual(t, SQLSERVER, container.Enable, FALSE)
		case MYSQL:
			assertEqual(t, MYSQL, container.Enable, TRUE)
		case POSTGRESQL:
			assertEqual(t, POSTGRESQL, container.Enable, FALSE)
		case APACHE:
			assertEqual(t, APACHE, container.Enable, FALSE)
		case NGINX:
			assertEqual(t, NGINX, container.Enable, FALSE)
		case LIGHTSPEED:
			assertEqual(t, LIGHTSPEED, container.Enable, FALSE)
		case RAILS:
			assertEqual(t, RAILS, container.Branch, BRANCH)
			assertEqual(t, RAILS, container.Repos, REPOS)
			assertEqual(t, RAILS, container.Enable, FALSE)
		case LARAVEL:
			assertEqual(t, LARAVEL, container.Branch, BRANCH)
			assertEqual(t, LARAVEL, container.Repos, REPOS)
			assertEqual(t, LARAVEL, container.Enable, FALSE)
		case SPRING:
			assertEqual(t, SPRING, container.Branch, BRANCH)
			assertEqual(t, SPRING, container.Repos, REPOS)
			assertEqual(t, SPRING, container.Enable, FALSE)
		case DJANGO:
			assertEqual(t, DJANGO, container.Branch, BRANCH)
			assertEqual(t, DJANGO, container.Repos, REPOS)
			assertEqual(t, DJANGO, container.Enable, TRUE)
		}
	}
}

func setLogfile(fileName string) {
	f, err := os.OpenFile(logfile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0664)
	if err != nil {
		panic(err)
	}
	log.SetOutput(f)
}

func assertEqual(t *testing.T, name string, got string, expected string) {
	if got != expected {
		log.Printf("[NG] \x1b[%smtest_name:%s \"got %s, expected: %s\"\x1b[0m", color, name, got, expected)
		t.Fatalf("\x1b[%smtest_name:%s \"got %s, expected: %s\"\x1b[0m", color, name, got, expected)
	} else {
		log.Printf("[OK] test %s is ok. got %s, expected %s.", name, got, expected)
	}
}
