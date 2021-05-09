package git

import (
	"log"
	"os"
	"os/exec"
)

const (
	BASE_COMMAND = "git"
)

func Clone(path string, branch string, repos string) error {
	if branch == "" {
		branch = "main"
	}
	args := []string{}
	args = append(args, "clone")
	args = append(args, "-b")
	args = append(args, branch)
	args = append(args, repos)
	args = append(args, path)
	cmd := exec.Command(BASE_COMMAND, args...)
	log.Println(cmd.String())
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		return err
	}
	if err := cmd.Wait(); err != nil {
		return err
	}
	return nil
}
