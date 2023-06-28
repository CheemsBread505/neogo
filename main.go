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

	goVersion := getGoVersion()
	fmt.Println("Go version:", goVersion)

	osName := getOperatingSystem()
	fmt.Println("Operating System:", osName)

	hostname := getHostname()
	fmt.Println("Hostname:", hostname)

	kernelVersion := getKernelVersion()
	fmt.Println("Kernel Version:", kernelVersion)

	uptime := getUptime()
	fmt.Println("Uptime:", uptime)

	shell := getShell()
	fmt.Println("Shell:", shell)

	fmt.Println("======")
}

func getGoVersion() string {
	goVersionCmd := exec.Command("go", "version")
	goVersionOutput, err := goVersionCmd.Output()
	if err != nil {
		fmt.Println("Error retrieving Go version:", err)
		os.Exit(1)
	}

	goVersion := strings.TrimSpace(string(goVersionOutput))
	return goVersion
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
			return "Arch Linux"
		}
	}

	return runtime.GOOS
}

func getHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("Error retrieving hostname:", err)
		os.Exit(1)
	}

	return hostname
}

func getKernelVersion() string {
	unameCmd := exec.Command("uname", "-r")
	unameOutput, err := unameCmd.Output()
	if err != nil {
		fmt.Println("Error retrieving kernel version:", err)
		os.Exit(1)
	}

	kernelVersion := strings.TrimSpace(string(unameOutput))
	return kernelVersion
}

func getUptime() string {
	uptimeCmd := exec.Command("uptime", "-p")
	uptimeOutput, err := uptimeCmd.Output()
	if err != nil {
		fmt.Println("Error retrieving uptime:", err)
		os.Exit(1)
	}

	uptime := strings.TrimSpace(string(uptimeOutput))
	return uptime
}

func getShell() string {
	shell := os.Getenv("SHELL")
	return shell
}
