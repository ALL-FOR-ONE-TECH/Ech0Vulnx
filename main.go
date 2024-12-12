package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"runtime"
	"strings"
	"sync"
)

// --> RunCommand executes a shell command
func RunCommand(cmd string) (string, error) {
	fmt.Printf("Running command: %s\n", cmd)
	output, err := exec.Command("bash", "-c", cmd).CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("error running command: %v, output: %s", err, string(output))
	}
	return string(output), nil
}

// --> Detects the package manager based on the system
func DetectPackageManager() string {
	packageManagers := []string{"apt", "yum", "zypper", "dnf", "apk", "pacman"}
	for _, pm := range packageManagers {
		if _, err := exec.LookPath(pm); err == nil {
			return pm
		}
	}
	return ""
}

// --> Installs a package using the detected package manager
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

// --> Installs a Go tool using `go install`
func InstallGoTool(tool, packageName string) error {
	fmt.Printf("Installing Go tool: %s\n", tool)
	_, err := RunCommand(fmt.Sprintf("go install %s", packageName))
	if err != nil {
		return fmt.Errorf("failed to install Go tool %s: %w", tool, err)
	}
	return nil
}

// --> Installs a tool either through a package manager or directly from source
func InstallTool(name, installCmd, checkCmd string) {
	if _, err := exec.LookPath(checkCmd); err != nil {
		fmt.Printf("Installing %s...\n", name)
		if _, err := RunCommand(installCmd); err != nil {
			fmt.Printf("Failed to install %s: %v\n", name, err)
		} else {
			fmt.Printf("%s installed successfully\n", name)
		}
	} else {
		fmt.Printf("Found %s\n", name)
	}
}

// --> Checks if the system is WSL (Windows Subsystem for Linux)
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

// --> Updates and upgrades the system using the detected package manager
func UpdateUpgradeSystem(packageManager string) error {
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
	default:
		return fmt.Errorf("unsupported package manager: %s", packageManager)
	}
	fmt.Println("System updated and upgraded successfully.")
	return nil
}

// -->  pip is installed on the system
func EnsurePipInstalled(packageManager string) error {
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
		return nil
	}
	fmt.Println("pip is already installed.")
	return nil
}

func main() {
	isWSL := CheckWSL()
	if isWSL {
		fmt.Println("Detected Windows Subsystem for Linux (WSL)")
	}

	// --> Detect the package manager
	packageManager := DetectPackageManager()
	if packageManager == "" {
		fmt.Println("Unable to detect package manager. Please install packages manually.")
		return
	}

	fmt.Printf("Detected package manager: %s\n", packageManager)

	if isWSL {
		if err := UpdateUpgradeSystem(packageManager); err != nil {
			fmt.Println("Error updating system:", err)
			return
		}
	}

	if err := EnsurePipInstalled(packageManager); err != nil {
		fmt.Println("Error ensuring pip installation:", err)
		return
	}

	// --> Install tools (using InstallTool function for each)
	tools := map[string][2]string{
		"go":          {"sudo apt install -y golang", "go"},
		"node":        {"sudo apt install -y nodejs", "node"},
		"npm":         {"sudo apt install -y npm", "npm"},
		"jq":          {"sudo apt install -y jq", "jq"},
		"shodan":      {"pip3 install shodan", "shodan"},
		"paramspider": {"sudo apt install -y paramspider"},
	}

	for name, cmd := range tools {
		InstallTool(name, cmd[0], cmd[1])
	}

	// --> Go tools
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

	// --> Wait for tools to finish
	wg.Wait()

	//  -->  paramspider
	InstallTool("paramspider", "git clone https://github.com/devanshbatham/paramspider && cd paramspider && python3 setup.py install", "paramspider")

	// --> httpx
	InstallTool("httpx", "go get -u github.com/projectdiscovery/httpx/cmd/httpx", "httpx")
}
