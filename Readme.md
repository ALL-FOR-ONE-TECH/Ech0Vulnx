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

### **Step 3: Run Installation Script**

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
usage: Ech0Vulnx.py [-h] [-sv filename.txt | -wl filename.txt] [-th 25]
                    [-s domain.com] [-d domains.txt] [-p domains.txt]
                    [-r domains.txt] [-b domains.txt] [-pspider domain.com]
                    [-w https://domain.com] [-j domain.com]
                    [-wc https://domain.com] [-fi https://domain.com]
                    [-fm https://domain.com] [-na https://domain.com] [-ri IP]
                    [-rim IP] [-sc domain.com] [-ph domain.txt]
                    [-co domains.txt] [-hh domain.com] [-sh domain.com]
                    [-ed domain.com] [-smu domain.com] [-ips domain list]
                    [-dinfo domain list] [-isubs domain list]
                    [-nft domains.txt] [-n domain.com or IP] [-api domain.com]
                    [-sho domain.com] [-fp domain.com] [-db domain.com]
                    [-cidr IP/24] [-ps 80,443,8443] [-pai IP/24]
                    [-xss https://example.com/page?param=value]
                    [-sqli https://example.com/page?param=value] [-shodan KEY]
                    [-webserver domain.com] [-javascript domain.com] [-dp 10]
                    [-je file.txt] [-hibp password] [-pm domain.com]
                    [-ch domain.com] [-or domain.com] [-asn AS55555]
                    [-st subdomains.txt] [-ar domain.com] [-jwt token]
                    [-jwt-modify token] [-heapds heapdump.txt]
                    [-heapts domain.com] [-f_p domain.com] [-v]
                    [-c CONCURRENCY] [-nl] [-gs] [-e EXTENSIONS] [-x EXCLUDE]
                    [-u] [--shodan-api SHODAN_API] [--proxy PROXY]
                    [--proxy-file PROXY_FILE] [--heapdump HEAPDUMP]
                    [--output-dir OUTPUT_DIR] [-aws domain.com]
                    [-az domain.com] [--s3-scan S3_SCAN] [-gcp domain.com]
                    [-zt domain.com] [--ipinfo TARGET] [--token TOKEN]
                    [--save-ranges FILENAME]
                    [--forbidden_domains FORBIDDEN_DOMAINS]
                    [--brute-user-pass domain.com]
                    [--username_wordlist domain.com]
                    [--password_wordlist domain.com]


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

