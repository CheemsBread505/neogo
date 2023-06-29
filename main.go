package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// Colors
const (
	cyanBold    = "\x1b[36;1;4m"
	darkBold    = "\x1b[30;1m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorBlue   = "\033[34m"
	colorYellow = "\033[33m"
	colorReset  = "\033[0m"
)

func main() {
	fmt.Println(cyanBold + "NeoGO" + colorReset)
	fmt.Println(darkBold + "======" + colorReset)

	osName := getOperatingSystem()
	fmt.Println(cyanBold+"Operating System:"+colorReset, osName)

	hostname := getHostname()
	fmt.Println(cyanBold+"Hostname:"+colorReset, hostname)

	kernelVersion := getKernelVersion()
	fmt.Println(cyanBold+"Kernel Version:"+colorReset, kernelVersion)

	uptime := getUptime()
	fmt.Println(cyanBold+"Uptime:"+colorReset, uptime)

	shell := getShell()
	fmt.Println(cyanBold+"Shell:"+colorReset, shell)

	fmt.Println(darkBold + "======" + colorReset)
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

