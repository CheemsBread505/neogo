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

func getOperatingSystem() string {
	if runtime.GOOS == "linux" {
		etcOSReleaseCmd := exec.Command("cat", "/etc/os-release")
		etcOSReleaseOutput, err := etcOSReleaseCmd.Output()
		if err != nil {
			return runtime.GOOS
		}

		osRelease := string(etcOSReleaseOutput)
		lines := strings.Split(osRelease, "\n")
		for _, line := range lines {
			if strings.HasPrefix(line, "PRETTY_NAME=") {
				dist := strings.TrimPrefix(line, "PRETTY_NAME=")
				dist = strings.Trim(dist, `"`)
				return dist
			}
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
