#!/bin/bash

# Color definitions
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
CYAN='\033[0;36m'
MAGENTA='\033[0;35m'
NC='\033[0m'

TOTAL_TOOLS=21  
CURRENT_TOOL=0

# --> Function to display main progress bar
show_main_progress() {
    local percentage=$((CURRENT_TOOL * 100 / TOTAL_TOOLS))
    local completed=$((percentage / 2))
    local remaining=$((50 - completed))
    local bar=$(printf '█%.0s' $(seq 1 $completed))
    local spaces=$(printf '─%.0s' $(seq 1 $remaining))
    echo -e "${CYAN}╔════════════════════════════════════════════════════╗${NC}"
    echo -e "${CYAN}║ ${MAGENTA}Overall Progress: [${GREEN}${bar}${NC}${spaces}${MAGENTA}] ${percentage}%${NC} ${CYAN}║${NC}"
    echo -e "${CYAN}╚════════════════════════════════════════════════════╝${NC}"
}

# --> Function to display tool-specific progress bar
show_tool_progress() {
    local tool="$1"
    local stage="$2"
    local total_stages="$3"
    local percentage=$((stage * 100 / total_stages))
    local completed=$((percentage / 2))
    local remaining=$((50 - completed))
    local bar=$(printf '▓%.0s' $(seq 1 $completed))
    local spaces=$(printf '░%.0s' $(seq 1 $remaining))
    echo -e "${BLUE}╭────────────────────────────────────────────────────╮${NC}"
    echo -e "${BLUE}│ ${YELLOW}$tool: [${GREEN}${bar}${NC}${spaces}${YELLOW}] ${percentage}%${NC} ${BLUE}│${NC}"
    echo -e "${BLUE}╰────────────────────────────────────────────────────╯${NC}"
}

# --> Function to run commands with progress
run_command() {
    local cmd="$1"
    local tool="$2"
    echo -e "\n${BLUE}► Executing: ${cmd}${NC}"
    echo -e "${CYAN}────────────────────────────────────────────────────${NC}"
    show_tool_progress "$tool" 1 3
    output=$(eval "$cmd" 2>&1)
    local status=$?
    show_tool_progress "$tool" 2 3
    if [ $status -eq 0 ]; then
        echo -e "${GREEN}✓ Success:${NC} $output"
    else
        echo -e "${RED}✗ Error:${NC} $output"
        return 1
    fi
    show_tool_progress "$tool" 3 3
    return $status
}

# --> Detect package manager
detect_package_manager() {
    for pm in apt yum zypper dnf apk pacman; do
        if command -v "$pm" &>/dev/null; then
            echo "$pm"
            return 0
        fi
    done
    echo ""
    return 1
}

# --> Install package based on package manager
install_package() {
    local package="$1"
    local manager="$2"
    local tool="$3"
    case "$manager" in
        "apt") run_command "sudo apt install -y $package" "$tool" ;;
        "dnf"|"yum") run_command "sudo $manager install -y $package" "$tool" ;;
        "pacman") run_command "sudo pacman -S --noconfirm $package" "$tool" ;;
        "zypper") run_command "sudo zypper install -y $package" "$tool" ;;
        "apk") run_command "sudo apk add $package" "$tool" ;;
        "brew") run_command "brew install $package" "$tool" ;;
        "pip") run_command "pip3 install $package" "$tool" ;;
        "npm") run_command "sudo npm install -g $package" "$tool" ;;
        "go") run_command "go install $package" "$tool" ;;
    esac
}

# --> Install tool if not present
install_tool() {
    local name="$1"
    local install_cmd="$2"
    local check_cmd="${3:-$name}"
    
    echo -e "\n${YELLOW}⚙ Checking $name...${NC}"
    if ! command -v "$check_cmd" &>/dev/null; then
        echo -e "${MAGENTA}⬇ Installing $name...${NC}"
        show_tool_progress "$name" 1 2
        if $install_cmd "$name"; then
            show_tool_progress "$name" 2 2
            echo -e "${GREEN}✓ $name installed${NC}"
        else
            echo -e "${RED}✗ Failed to install $name${NC}"
            return 1
        fi
    else
        show_tool_progress "$name" 1 1
        echo -e "${GREEN}✓ $name already present${NC}"
    fi
    ((CURRENT_TOOL++))
    show_main_progress
}

# --> Install Go tool with detailed progress
go_tool() {
    local tool="$1"
    local package="$2"
    
    echo -e "\n${YELLOW}⚙ Installing $tool...${NC}"
    show_tool_progress "$tool" 1 4
    if run_command "go install $package" "$tool" &>/dev/null; then
        go_path=$(go env GOPATH)
        bin_path="$go_path/bin/$tool"
        show_tool_progress "$tool" 3 4
        if [ -f "$bin_path" ]; then
            run_command "sudo mv $bin_path /usr/local/bin/" "$tool"
            echo -e "${GREEN}✓ $tool installed${NC}"
        else
            echo -e "${RED}✗ Failed to find $tool${NC}"
            return 1
        fi
    else
        echo -e "${RED}✗ Failed to install $tool${NC}"
        return 1
    fi
    ((CURRENT_TOOL++))
    show_main_progress
}

# --> Install Python dependencies from requirements.txt
python_deps() {
    echo -e "\n${YELLOW}⚙ Checking Python dependencies...${NC}"
    if [ -f "requirements.txt" ]; then
        echo -e "${MAGENTA}⬇ Installing Python dependencies...${NC}"
        show_tool_progress "Python Deps" 1 2
        if run_command "pip3 install -r requirements.txt" "Python Deps" || \
           run_command "pip install -r requirements.txt --break-system-packages" "Python Deps"; then
            show_tool_progress "Python Deps" 2 2
            echo -e "${GREEN}✓ Python dependencies installed${NC}"
        else
            echo -e "${RED}✗ Failed to install Python dependencies${NC}"
        fi
    else
        show_tool_progress "Python Deps" 1 1
        echo -e "${BLUE}ℹ No requirements.txt found, skipping${NC}"
    fi
    ((CURRENT_TOOL++))
    show_main_progress
}
# -->  Check if running in WSL
check_wsl() {
    if [ -f /proc/version ] && grep -qi microsoft /proc/version; then
        echo "true"
    else
        echo "false"
    fi
}

# --> Update and upgrade system
update_upgrade_system() {
    local package_manager="$1"
    echo -e "\n${YELLOW}⚙ Updating system...${NC}"
    case "$package_manager" in
        "apt") run_command "sudo apt update && sudo apt upgrade -y" "System Update" ;;
        "dnf"|"yum") run_command "sudo $manager update -y" "System Update" ;;
        "pacman") run_command "sudo pacman -Syu --noconfirm" "System Update" ;;
        "zypper") run_command "sudo zypper update -y" "System Update" ;;
        "apk") run_command "sudo apk update && sudo apk upgrade" "System Update" ;;
    esac
    echo -e "${GREEN}✓ System updated${NC}"
    ((CURRENT_TOOL++))
    show_main_progress
}

# --> Ensure pip is installed
check_pip() {
    local package_manager="$1"
    echo -e "\n${YELLOW}⚙ Checking pip...${NC}"
    if ! command -v pip3 &>/dev/null && ! command -v pip &>/dev/null; then
        echo -e "${MAGENTA}⬇ Installing pip...${NC}"
        show_tool_progress "pip" 1 2
        case "$package_manager" in
            "apt") run_command "sudo apt install -y python3-pip" "pip" ;;
            "dnf"|"yum") run_command "sudo $package_manager install -y python3-pip" "pip" ;;
            "pacman") run_command "sudo pacman -S --noconfirm python-pip" "pip" ;;
            "zypper") run_command "sudo zypper install -y python3-pip" "pip" ;;
            "apk") run_command "sudo apk add py3-pip" "pip" ;;
            "brew") run_command "brew install python" "pip" ;;
        esac
        show_tool_progress "pip" 2 2
        echo -e "${GREEN}✓ pip installed${NC}"
    else
        show_tool_progress "pip" 1 1
        echo -e "${GREEN}✓ pip already present${NC}"
    fi
    ((CURRENT_TOOL++))
    show_main_progress
}

main() {
    local system=$(uname -s)
    local is_wsl=$(check_wsl)

    clear
    echo -e "${CYAN}╔════════════════════════════════════╗${NC}"
    echo -e "${CYAN}║ ${MAGENTA}Security Tools Installer${NC}         ${CYAN}║${NC}"
    echo -e "${CYAN}╚════════════════════════════════════╝${NC}\n"

    if [ "$is_wsl" = "true" ]; then
        echo -e "${BLUE}ℹ Running on WSL${NC}"
    fi

    if [ "$system" = "Linux" ]; then
        package_manager=$(detect_package_manager)
        if [ -z "$package_manager" ]; then
            echo -e "${RED}✗ No package manager found${NC}"
            exit 1
        fi
        echo -e "${BLUE}ℹ Using: $package_manager${NC}"
        
        if [ "$is_wsl" = "true" ]; then
            update_upgrade_system "$package_manager"
        fi
    elif [ "$system" = "Darwin" ]; then
        package_manager="brew"
        if ! command -v brew &>/dev/null; then
            echo -e "${RED}✗ Homebrew required${NC}"
            exit 1
        fi
    else
        echo -e "${RED}✗ Unsupported OS: $system${NC}"
        exit 1
    fi

    python_deps  # Added Python dependencies installation
    check_pip "$package_manager"
    install_tool "go" "install_package golang $package_manager"
    install_tool "node" "install_package nodejs $package_manager"
    install_tool "npm" "install_package npm $package_manager"
    install_tool "blc" "install_package broken-link-checker npm"
    go_tool "nuclei" "github.com/projectdiscovery/nuclei/v2/cmd/nuclei@latest"

    if [ ! -d "$HOME/nuclei-templates" ]; then
        echo -e "\n${MAGENTA}⬇ Cloning nuclei-templates...${NC}"
        run_command "git clone https://github.com/projectdiscovery/nuclei-templates.git $HOME/nuclei-templates" "nuclei-templates"
        ((CURRENT_TOOL++))
        show_main_progress
    else
        echo -e "\n${GREEN}✓ nuclei-templates present${NC}"
        show_tool_progress "nuclei-templates" 1 1
        ((CURRENT_TOOL++))
        show_main_progress
    fi

    go_tools=(
        "dnsx github.com/projectdiscovery/dnsx/cmd/dnsx@latest"
        "subfinder github.com/projectdiscovery/subfinder/v2/cmd/subfinder@latest"
        "waybackurls github.com/tomnomnom/waybackurls@latest"
        "httprobe github.com/tomnomnom/httprobe@latest"
        "httpx github.com/projectdiscovery/httpx/cmd/httpx@latest"
        "anew github.com/tomnomnom/anew@latest"
        "gau github.com/lc/gau/v2/cmd/gau@latest"
        "gauplus github.com/bp0lr/gauplus@latest"
        "hakrawler github.com/hakluke/hakrawler@latest"
        "assetfinder github.com/tomnomnom/assetfinder@latest"
        "asnmap github.com/projectdiscovery/asnmap/cmd/asnmap@latest"
        "naabu github.com/projectdiscovery/naabu/v2/cmd/naabu@latest"
    )

    for tool in "${go_tools[@]}"; do
        IFS=' ' read -r tool_name package <<< "$tool"
        go_tool "$tool_name" "$package"
    done

    install_tool "jq" "install_package jq $package_manager"
    install_tool "shodan" "install_package shodan pip"
    install_tool "paramspider" "run_command 'git clone https://github.com/devanshbatham/paramspider && cd paramspider && python3 setup.py install'"

    echo -e "\n${GREEN}╔════════════════════════════════════╗${NC}"
    echo -e "${GREEN}║ ${MAGENTA}Installation Complete!${NC}           ${GREEN}║${NC}"
    echo -e "${GREEN}╚════════════════════════════════════╝${NC}"
}

trap 'echo -e "\n${RED}╔════════════════════════════════════╗${NC}\n${RED}║ ${MAGENTA}Installation Interrupted${NC}        ${RED}║${NC}\n${RED}╚════════════════════════════════════╝${NC}"; exit 1' INT

main