package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func main() {
	fmt.Println("NeoGO")
	fmt.Println("======")

	goVersionCmd := exec.Command("go", "version")
	goVersionOutput, err := goVersionCmd.Output()
	if err != nil {
		fmt.Println("Error retrieving Go version:", err)
		os.Exit(1)
	}

	goVersion := strings.TrimSpace(string(goVersionOutput))
	fmt.Println("Go version:", goVersion)

	osName := getOperatingSystem()
	fmt.Println("Operating System:", osName)

	arch := runtime.GOARCH
	fmt.Println("Architecture:", arch)
}

func getOperatingSystem() string {
	if runtime.GOOS == "linux" {
		etcOSReleaseCmd := exec.Command("cat", "/etc/os-release")
		etcOSReleaseOutput, err := etcOSReleaseCmd.Output()
		if err != nil {
			return runtime.GOOS
		}

		osRelease := string(etcOSReleaseOutput)
		if strings.Contains(osRelease, "Arch Linux") {
			return "archlinux"
		}
	}

	return runtime.GOOS
}
