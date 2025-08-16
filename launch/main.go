package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	targets := []struct {
		goos      string
		goarch    string
		unameos   string
		unamearch string
	}{
		{"linux", "amd64", "Linux", "x86_64"},
		{"linux", "arm64", "Linux", "aarch64"},
		{"darwin", "amd64", "Darwin", "x86_64"},
		{"darwin", "arm64", "Darwin", "arm64"},
	}

	for _, target := range targets {
		cmd := exec.Command("go", "build", "-o", fmt.Sprintf("dist/%s-%s", target.unameos, target.unamearch), "./uni")
		cmd.Env = os.Environ()
		cmd.Env = append(cmd.Env, "GOOS="+target.goos)
		cmd.Env = append(cmd.Env, "GOARCH="+target.goarch)
		cmd.Env = append(cmd.Env, "CGO_ENABLED=0")
		cmd.Run()
	}
}
