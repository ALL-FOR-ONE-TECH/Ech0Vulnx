package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"runtime"
	"strings"
	"sync"
)

// RunCommand executes a shell command and returns its output
func RunCommand(cmd string) (string, error) {
	fmt.Printf("Running command: %s\n", cmd)
	output, err := exec.Command("bash", "-c", cmd).CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("error running command: %s, output: %s", err, string(output))
	}
	return string(output), nil
}

// DetectPackageManager detects the package manager based on the system
func DetectPackageManager() string {
	packageManagers := []string{"apt", "yum", "zypper", "dnf", "apk", "pacman"}

	for _, pm := range packageManagers {
		if _, err := exec.LookPath(pm); err == nil {
			return pm
		}
	}
	return ""
}

// InstallPackage installs a package using the detected package manager
func InstallPackage(packageName, manager string) error {
	var cmd string
	switch manager {
	case "apt":
		cmd = fmt.Sprintf("sudo apt install -y %s", packageName)
	case "dnf", "yum":
		cmd = fmt.Sprintf("sudo %s install -y %s", manager, packageName)
	case "pacman":
		cmd = fmt.Sprintf("sudo pacman -S --noconfirm %s", packageName)
	case "zypper":
		cmd = fmt.Sprintf("sudo zypper install -y %s", packageName)
	case "apk":
		cmd = fmt.Sprintf("sudo apk add %s", packageName)
	}
	_, err := RunCommand(cmd)
	return err
}

// InstallGoTool installs a Go tool using `go install`
func InstallGoTool(tool, packageName string) error {
	fmt.Printf("Installing Go tool: %s\n", tool)
	_, err := RunCommand(fmt.Sprintf("go install %s", packageName))
	if err != nil {
		return err
	}
	return nil
}

// InstallTool installs a tool either through a package manager or directly from source
func InstallTool(name, installCmd string, checkCmd string) {
	if _, err := exec.LookPath(checkCmd); err != nil {
		fmt.Printf("Installing %s...\n", name)
		_, err := RunCommand(installCmd)
		if err != nil {
			fmt.Printf("Failed to install %s\n", name)
		} else {
			fmt.Printf("%s installed successfully\n", name)
		}
	} else {
		fmt.Printf("Found %s\n", name)
	}
}

// CheckWSL checks if the system is WSL
func CheckWSL() bool {
	if runtime.GOOS == "linux" {
		data, err := ioutil.ReadFile("/proc/version")
		if err != nil {
			return false
		}
		return strings.Contains(string(data), "Microsoft")
	}
	return false
}

func UpdateUpgradeSystem(packageManager string) {
	fmt.Println("Updating and upgrading the system...")
	switch packageManager {
	case "apt":
		_, _ = RunCommand("sudo apt update && sudo apt upgrade -y")
	case "dnf", "yum":
		_, _ = RunCommand(fmt.Sprintf("sudo %s update -y", packageManager))
	case "pacman":
		_, _ = RunCommand("sudo pacman -Syu --noconfirm")
	case "zypper":
		_, _ = RunCommand("sudo zypper update -y")
	case "apk":
		_, _ = RunCommand("sudo apk update && sudo apk upgrade")
	}
	fmt.Println("System updated and upgraded successfully.")
}

func EnsurePipInstalled(packageManager string) {
	if _, err := exec.LookPath("pip3"); err != nil {
		fmt.Println("pip is not installed. Installing pip...")
		switch packageManager {
		case "apt":
			_, _ = RunCommand("sudo apt install -y python3-pip")
		case "dnf", "yum":
			_, _ = RunCommand(fmt.Sprintf("sudo %s install -y python3-pip", packageManager))
		case "pacman":
			_, _ = RunCommand("sudo pacman -S --noconfirm python-pip")
		case "zypper":
			_, _ = RunCommand("sudo zypper install -y python3-pip")
		case "apk":
			_, _ = RunCommand("sudo apk add py3-pip")
		}
		fmt.Println("pip installed successfully.")
	} else {
		fmt.Println("pip is already installed.")
	}
}

func main() {
	isWSL := CheckWSL()
	if isWSL {
		fmt.Println("Detected Windows Subsystem for Linux (WSL)")
	}

	packageManager := DetectPackageManager()
	if packageManager == "" {
		fmt.Println("Unable to detect package manager. Please install packages manually.")
		return
	}

	fmt.Printf("Detected package manager: %s\n", packageManager)

	// Update and upgrade the system (if on WSL)
	if isWSL {
		UpdateUpgradeSystem(packageManager)
	}

	EnsurePipInstalled(packageManager)

	// Install Go and other tools
	InstallTool("go", "sudo apt install -y golang", "go")
	InstallTool("node", "sudo apt install -y nodejs", "node")
	InstallTool("npm", "sudo apt install -y npm", "npm")
	InstallTool("jq", "sudo apt install -y jq", "jq")
	InstallTool("shodan", "pip3 install shodan", "shodan")

	// Install Go tools concurrently
	var wg sync.WaitGroup
	goTools := map[string]string{
		"nuclei":      "github.com/projectdiscovery/nuclei/v2/cmd/nuclei@latest",
		"dnsx":        "github.com/projectdiscovery/dnsx/cmd/dnsx@latest",
		"subfinder":   "github.com/projectdiscovery/subfinder/v2/cmd/subfinder@latest",
		"waybackurls": "github.com/tomnomnom/waybackurls@latest",
	}

	for tool, pkg := range goTools {
		wg.Add(1)
		go func(tool, pkg string) {
			defer wg.Done()
			if err := InstallGoTool(tool, pkg); err != nil {
				fmt.Printf("Failed to install %s: %v\n", tool, err)
			} else {
				fmt.Printf("%s installed successfully.\n", tool)
			}
		}(tool, pkg)
	}

	// Wait for all Go tools to finish
	wg.Wait()

	// Install other tools like broken-link-checker, paramspider, etc.
	InstallTool("paramspider", "git clone https://github.com/devanshbatham/paramspider && cd paramspider && python3 setup.py install", "paramspider")
}
