# Ech0Vulnx


Ech0Vulnx is a powerful multi-purpose network scanner and vulnerability assessment tool built for bug bounty hunters, penetration testers, and security researchers. It supports subdomain enumeration, vulnerability detection, port scanning, cloud bucket hunting, JavaScript analysis, and much more!

---

## 🚀 Features

- Subdomain enumeration & takeover detection
- DNS scanning & WHOIS information
- Port scanning with CIDR notation support
- Vulnerability scanning (XSS, SQLi, LFI, Open Redirect, etc.)
- JavaScript file analysis & endpoint fuzzing
- Directory brute-forcing & parameter fuzzing
- AWS S3, Azure, and GCP bucket enumeration
- Shodan, Wayback, Nuclei integration
- Customizable headers, concurrency, depth options
- Auto Recon & multipurpose scans

---

## 📥 Installation

### **Supported OS:**
- ✅ Linux (Debian, Ubuntu, Kali, Parrot OS, Arch)
- ✅ macOS
- ✅ Windows (via WSL / Python)

### **Step 1: Install System Dependencies**

#### Debian/Ubuntu/Kali/Parrot:
```bash
sudo apt update && sudo apt upgrade -y
sudo apt install python3 python3-pip git nmap -y
```

#### Arch Linux:
```bash
sudo pacman -Syu
sudo pacman -S python python-pip git nmap
```

#### macOS:
```bash
brew update
brew install python3 git nmap
```

#### Windows:
1. Install **Python 3.x** from [python.org](https://www.python.org/downloads/).
2. Install **Nmap** from [nmap.org](https://nmap.org/download.html).
3. Optionally, enable **WSL** for better Linux compatibility.

---

### **Step 2: Clone Repository**

```bash
git clone https://github.com/whoamikiddie/Ech0Vulnx.git
cd Ech0Vulnx
```

---

### **Step 3: Install Python Dependencies**

```bash
pip3 install -r requirements.txt
```

---

### **Step 4: Run Installation Script**

```bash
chmod +x install.sh
./install.sh
```

---


---

## ✅ Verify Installation

```bash
python3 Ech0Vulnx.py -h
```

---

## 📖 USAGE

```
usage: Ech0Vulnx.py [-h] [-sv filename.txt | -wl filename.txt] [-th 25] [-s domain.com]
                  [-t domain.com] [-d domains.txt] [-p domains.txt] [-r domains.txt]
                  ... (shortened for brevity) ...
                  [-jwt-modify token] [--s3-scan S3_SCAN] [-v] [-c CONCURRENCY] [-nl] [-gs]
                  [-e EXTENSIONS] [-x EXCLUDE] [-u] [--shodan-api SHODAN_API]
```

**Main Options Overview:**

| Option | Description |
|-------|------------|
| `-s`  | Scan subdomains |
| `-d`  | Scan domains for DNS records |
| `-fi` | Get favicon hashes |
| `-wc` | Web crawler |
| `-b`  | Broken links check |
| `-co` | CORS misconfiguration scan |
| `-hh` | Host header injection scan |
| `-db` | Directory brute force |
| `-ph` | Directory traversal vulnerability scan |
| `-sqli`| SQL Injection scan |
| `-xss` | XSS vulnerability scan |
| `-javascript` | JavaScript file sensitive info scan |
| `--jwt` | Analyze JWT tokens |
| `--s3-scan` | AWS S3 bucket enumeration |
| `--cidr` | CIDR notation port scan |
| `-ar` | Auto recon |
| Many more... (Full list available with `-h` flag)

---

## 📌 EXAMPLES

Scan for subdomains and save output:
```
python3 Ech0Vulnx.py -s yahoo.com --save filename.txt
```

Scan subdomains + Shodan extraction:
```
python3 Ech0Vulnx.py -s yahoo.com --shodan API_KEY --save filename.txt
```

JavaScript file discovery:
```
python3 Ech0Vulnx.py -j yahoo.com --depth 4 --save jsfiles.txt -c 20
```

DNS record scan:
```
python3 Ech0Vulnx.py -d domains.txt
```

Favicon hash scan:
```
python3 Ech0Vulnx.py -fi domain.com
```

Web crawler:
```
python3 Ech0Vulnx.py -wc https://www.domain.com
```

Broken links scan:
```
python3 Ech0Vulnx.py -b https://www.domain.com
```

CORS Misconfiguration:
```
python3 Ech0Vulnx.py -co domains.txt
```

Host Header Injection:
```
python3 Ech0Vulnx.py -hh domains.txt
```

Directory brute force:
```
python3 Ech0Vulnx.py --directorybrute domain.com --wordlist list.txt --threads 50 -e php,txt,html -x 404,403
```

Subnet scan:
```
python3 Ech0Vulnx.py --cidr_notation IP/24 --ports 80,443 --threads 200
```

Directory traversal:
```
python3 Ech0Vulnx.py -ph domain.com?id=
```

SQL Injection:
```
python3 Ech0Vulnx.py -sqli domain.com?id=1
```

XSS Scan:
```
python3 Ech0Vulnx.py -xss domain.com?id=1
```

Sensitive JavaScript analysis:
```
python3 Ech0Vulnx.py -javascript domain.com
```

JWT Token scan:
```
python3 Ech0Vulnx.py -jwt Token
```

AWS S3 Bucket Enumeration:
```
python3 Ech0Vulnx.py --s3-scan bucket.com
```

Subdomain Takeover:
```
python3 Ech0Vulnx.py -st domains.txt --save vuln_subs.txt -c 50 
```

Auto Recon:
```
python3 Ech0Vulnx.py -ar domain.com
```

---

## 🌐 Optional Shodan API Key

Register and get your API key at:

- [Shodan.io](https://account.shodan.io/register)

Usage:
```
python3 Ech0Vulnx.py -s domain.com --shodan-api YOUR_API_KEY
```

---

## 🛠️ Troubleshooting

| Issue                                   | Solution                                         |
|----------------------------------------|--------------------------------------------------|
| Missing `requests` module              | Run: `pip3 install requests`                     |
| Nmap/Nuclei not found                  | Install & add to PATH                           |
| Permission denied                      | Ensure: `chmod +x install.sh` & run with `sudo`  |
| Python version issues                  | Requires **Python 3.12.9**                          |

