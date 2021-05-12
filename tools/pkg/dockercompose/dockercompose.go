package dockercompose

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
)

const (
	BASE_COMMAND = "docker-compose"
)

type DockerCompose struct {
	Files   []string
	Options []string
}

type Options struct {
	// ProjectName       string
	// Verbose           bool
	// Version           string
	// Host              string
	// TLS               bool
	// TLSCacert         string
	// TLSCert           string
	// TLSKey            string
	// TLSVerify         bool
	// SkipHostnameCheck bool
}

// func (dc *DockerCompose) Build()           {}
// func (dc *DockerCompose) Config()           {}
// func (dc *DockerCompose) Create()           {}

func (dc *DockerCompose) Down() {
	dc.dockerCompose("down", nil)
}

// func (dc *DockerCompose) Events()            {}
// func (dc *DockerCompose) Help()            {}
// func (dc *DockerCompose) Kill()            {}
// func (dc *DockerCompose) Logs()            {}
// func (dc *DockerCompose) Pause()           {}
// func (dc *DockerCompose) Port()            {}
// func (dc *DockerCompose) Ps()              {}
// func (dc *DockerCompose) Pull()            {}
// func (dc *DockerCompose) Restart()         {}
// func (dc *DockerCompose) Rm()              {}
// func (dc *DockerCompose) Run() {}
// func (dc *DockerCompose) Scale()           {}
// func (dc *DockerCompose) Start()           {}
// func (dc *DockerCompose) Stop()            {}
// func (dc *DockerCompose) Unpause()         {}

func (dc *DockerCompose) Up(options []string, services []string) error {
	args := []string{}
	if options != nil {
		args = append(args, options...)
	}
	if services != nil {
		args = append(args, services...)
	}
	return dc.dockerCompose("up", args)
}

// func (dc *DockerCompose) Version()         {}

func (dc *DockerCompose) dockerCompose(command string, commandArgs []string) error {
	// コマンド
	args := []string{}
	// -f=<引数>...
	if dc.Files != nil {
		for _, file := range dc.Files {
			if file != "" {
				args = append(args, "-f")
				args = append(args, file)
			}
		}
	}
	// [オプション]
	if dc.Options != nil {
		args = append(args, dc.Options...)
	}
	// [コマンド]
	if command != "" {
		args = append(args, command)
	} else {
		return errors.New(fmt.Sprintf("invalid args: docker-compose command is %s\n", command))
	}
	// [引数///]
	if commandArgs != nil {
		args = append(args, commandArgs...)
	}

	// コマンド実行
	cmd := exec.Command(BASE_COMMAND, args...)
	log.Printf("exec command: %s\n", cmd.String())
	// cmd.Stdout = os.Stdin
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
