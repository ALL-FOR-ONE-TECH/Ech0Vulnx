# Ech0Vulnx

<div align="center">
  <img src="https://raw.githubusercontent.com/whoamikiddie/Ech0Vulnx/main/assets/ech0vulnx.gif" alt="Ech0Vulnx Animation" width="600"/>
</div>

<div align="center">
  <a href="https://github.com/whoamikiddie/Ech0Vulnx/stargazers"><img src="https://img.shields.io/github/stars/whoamikiddie/Ech0Vulnx?style=for-the-badge" alt="Stars"></a>
  <a href="https://github.com/whoamikiddie/Ech0Vulnx/network/members"><img src="https://img.shields.io/github/forks/whoamikiddie/Ech0Vulnx?style=for-the-badge" alt="Forks"></a>
  <a href="https://github.com/whoamikiddie/Ech0Vulnx/issues"><img src="https://img.shields.io/github/issues/whoamikiddie/Ech0Vulnx?style=for-the-badge" alt="Issues"></a>
  <a href="https://github.com/whoamikiddie/Ech0Vulnx/blob/main/LICENSE"><img src="https://img.shields.io/github/license/whoamikiddie/Ech0Vulnx?style=for-the-badge" alt="License"></a>
</div>

<div align="center">
  <h3>Advanced Network Scanner & Vulnerability Assessment Tool</h3>
</div>

Ech0Vulnx is a powerful multi-purpose network scanner and vulnerability assessment tool built for bug bounty hunters, penetration testers, and security researchers. It supports subdomain enumeration, vulnerability detection, port scanning, cloud bucket hunting, JavaScript analysis, and much more!

## 🚀 Features

<div align="center">
  <table>
    <tr>
      <td>
        <ul>
          <li>Subdomain enumeration & takeover detection</li>
          <li>DNS scanning & WHOIS information</li>
          <li>Port scanning with CIDR notation support</li>
          <li>Vulnerability scanning (XSS, SQLi, LFI, Open Redirect, etc.)</li>
          <li>JavaScript file analysis & endpoint fuzzing</li>
        </ul>
      </td>
      <td>
        <ul>
          <li>Directory brute-forcing & parameter fuzzing</li>
          <li>AWS S3, Azure, and GCP bucket enumeration</li>
          <li>Shodan, Wayback, Nuclei integration</li>
          <li>Customizable headers, concurrency, depth options</li>
          <li>Auto Recon & multipurpose scans</li>
        </ul>
      </td>
    </tr>
  </table>
</div>

## 📥 Installation

### **Supported OS:**
- Linux (Debian, Ubuntu, Kali, Parrot OS, Arch)
- macOS
- Windows (via WSL / Python)

### **Step 1: Install System Dependencies**

<details>
<summary><b>Debian/Ubuntu/Kali/Parrot</b></summary>

```bash
sudo apt update && sudo apt upgrade -y
sudo apt install python3 python3-pip git nmap -y
```
</details>

<details>
<summary><b>Arch Linux</b></summary>

```bash
sudo pacman -Syu
sudo pacman -S python python-pip git nmap
```
</details>

<details>
<summary><b>macOS</b></summary>

```bash
brew update
brew install python3 git nmap
```
</details>

<details>
<summary><b>Windows</b></summary>

1. Install **Python 3.12.9** from [python.org](https://www.python.org/downloads/).
2. Install **Nmap** from [nmap.org](https://nmap.org/download.html).
3. Optionally, enable **WSL** for better Linux compatibility.
</details>

### **Step 2: Clone Repository**

```bash
git clone https://github.com/whoamikiddie/Ech0Vulnx.git
cd Ech0Vulnx
```

### **Step 3: Run Installation Script**

```bash
chmod +x install.sh
./install.sh
```

## ✅ Verify Installation

```bash
python3 Ech0Vulnx.py -h
```

## 📖 Usage

<div align="center">
  <img src="https://raw.githubusercontent.com/whoamikiddie/Ech0Vulnx/main/assets/usage.png" alt="Usage Example"/>
</div>

```bash
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

### **Main Options Overview:**

<div align="center">
  <table>
    <tr>
      <th>Option</th>
      <th>Description</th>
    </tr>
    <tr>
      <td><code>-s</code></td>
      <td>Scan subdomains</td>
    </tr>
    <tr>
      <td><code>-d</code></td>
      <td>Scan domains for DNS records</td>
    </tr>
    <tr>
      <td><code>-fi</code></td>
      <td>Get favicon hashes</td>
    </tr>
    <tr>
      <td><code>-wc</code></td>
      <td>Web crawler</td>
    </tr>
    <tr>
      <td><code>-b</code></td>
      <td>Broken links check</td>
    </tr>
    <tr>
      <td><code>-co</code></td>
      <td>CORS misconfiguration scan</td>
    </tr>
    <tr>
      <td><code>-hh</code></td>
      <td>Host header injection scan</td>
    </tr>
    <tr>
      <td><code>-db</code></td>
      <td>Directory brute force</td>
    </tr>
    <tr>
      <td><code>-ph</code></td>
      <td>Directory traversal vulnerability scan</td>
    </tr>
    <tr>
      <td><code>-sqli</code></td>
      <td>SQL Injection scan</td>
    </tr>
    <tr>
      <td><code>-xss</code></td>
      <td>XSS vulnerability scan</td>
    </tr>
    <tr>
      <td><code>-javascript</code></td>
      <td>JavaScript file sensitive info scan</td>
    </tr>
    <tr>
      <td><code>--jwt</code></td>
      <td>Analyze JWT tokens</td>
    </tr>
    <tr>
      <td><code>--s3-scan</code></td>
      <td>AWS S3 bucket enumeration</td>
    </tr>
    <tr>
      <td><code>--cidr</code></td>
      <td>CIDR notation port scan</td>
    </tr>
    <tr>
      <td><code>-ar</code></td>
      <td>Auto recon</td>
    </tr>
  </table>
</div>

## 📌 Examples

<details>
<summary><b>Subdomain Scanning</b></summary>

```bash
python3 Ech0Vulnx.py -s yahoo.com --save filename.txt
```
</details>

<details>
<summary><b>Shodan Integration</b></summary>

```bash
python3 Ech0Vulnx.py -s yahoo.com --shodan API_KEY --save filename.txt
```
</details>

<details>
<summary><b>JavaScript Analysis</b></summary>

```bash
python3 Ech0Vulnx.py -j yahoo.com --depth 4 --save jsfiles.txt -c 20
```
</details>

<details>
<summary><b>DNS Scanning</b></summary>

```bash
python3 Ech0Vulnx.py -d domains.txt
```
</details>

<details>
<summary><b>Favicon Analysis</b></summary>

```bash
python3 Ech0Vulnx.py -fi domain.com
```
</details>

<details>
<summary><b>Web Crawling</b></summary>

```bash
python3 Ech0Vulnx.py -wc https://www.domain.com
```
</details>

<details>
<summary><b>Broken Links Check</b></summary>

```bash
python3 Ech0Vulnx.py -b https://www.domain.com
```
</details>

<details>
<summary><b>CORS Testing</b></summary>

```bash
python3 Ech0Vulnx.py -co domains.txt
```
</details>

<details>
<summary><b>Host Header Injection</b></summary>

```bash
python3 Ech0Vulnx.py -hh domains.txt
```
</details>

<details>
<summary><b>Directory Brute Force</b></summary>

```bash
python3 Ech0Vulnx.py --directorybrute domain.com --wordlist list.txt --threads 50 -e php,txt,html -x 404,403
```
</details>

<details>
<summary><b>Subnet Scanning</b></summary>

```bash
python3 Ech0Vulnx.py --cidr_notation IP/24 --ports 80,443 --threads 200
```
</details>

<details>
<summary><b>Directory Traversal</b></summary>

```bash
python3 Ech0Vulnx.py -ph domain.com?id=
```
</details>

<details>
<summary><b>SQL Injection</b></summary>

```bash
python3 Ech0Vulnx.py -sqli domain.com?id=1
```
</details>

<details>
<summary><b>XSS Scanning</b></summary>

```bash
python3 Ech0Vulnx.py -xss domain.com?id=1
```
</details>

<details>
<summary><b>JavaScript Analysis</b></summary>

```bash
python3 Ech0Vulnx.py -javascript domain.com
```
</details>

<details>
<summary><b>JWT Analysis</b></summary>

```bash
python3 Ech0Vulnx.py -jwt Token
```
</details>

<details>
<summary><b>AWS S3 Bucket Enumeration</b></summary>

```bash
python3 Ech0Vulnx.py --s3-scan bucket.com
```
</details>

<details>
<summary><b>Subdomain Takeover</b></summary>

```bash
python3 Ech0Vulnx.py -st domains.txt --save vuln_subs.txt -c 50 
```
</details>

<details>
<summary><b>Auto Recon</b></summary>

```bash
python3 Ech0Vulnx.py -ar domain.com
```
</details>

## 🌐 Optional Shodan API Key

Register and get your API key at: [Shodan.io](https://www.shodan.io/)

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## ⭐ Support

If you find this tool useful, please consider giving it a ⭐ on GitHub!

<div align="center">
  <img src="https://raw.githubusercontent.com/whoamikiddie/Ech0Vulnx/main/assets/ech0vulnx-footer.png" alt="Ech0Vulnx Footer" width="600"/>
</div>

