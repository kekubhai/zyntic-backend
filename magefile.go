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
	fmt.Println("Starting project...")
	executablePath := filepath.Join(".", "bin", "main")
	if runtime.GOOS == "windows" {
		executablePath = filepath.Join("bin", "main.exe")
	}

	cmd := exec.Command(executablePath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
