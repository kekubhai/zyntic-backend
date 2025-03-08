//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func Build() error {
	fmt.Println("Building project...")
	outputPath := filepath.Join("bin", "main")
	if runtime.GOOS == "windows" {
		outputPath = filepath.Join("bin", "main.exe")
	}

	cmd := exec.Command("go", "build", "-o", outputPath, filepath.Join("cmd", "api", "main.go"))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func Start() error {
	fmt.Println("Starting project with Docker...")
	cmd := exec.Command("docker", "compose", "up", "--build")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func Stop() error {
	fmt.Println("Stopping Docker containers...")
	cmd1 := exec.Command("docker", "compose", "rm", "-v", "--force", "--stop")
	cmd1.Stdout = os.Stdout
	cmd1.Stderr = os.Stderr
	err1 := cmd1.Run()
	if err1 != nil {
		return err1
	}

	fmt.Println("Removing Docker images...")
	cmd2 := exec.Command("docker", "rmi", "zyntic")
	cmd2.Stdout = os.Stdout
	cmd2.Stderr = os.Stderr
	return cmd2.Run()
}
