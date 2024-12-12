# DESCRIPTION

St3althGHost is comprehensive network scanning and vulnerability assessment tool. This tool is designed for security professionals and penetration testers to perform comprehensive reconnaissance and vulnerability assessment on target networks and web applications. It combines multiple scanning techniques and integrates various external tools to provide a wide range of information about the target.

![StealthGhost](https://github.com/whoamikiddie/StealthGhost/blob/main/image.png)
##  Overview 

1. It imports various libraries for network operations, web scraping, and parallel processing.

2. The script defines a colorful banner and sets up command-line argument parsing for different scanning options.

3. It includes multiple scanning functions for different purposes:
   - Subdomain enumeration
   - Technology detection
   - DNS record scanning
   - Web crawling and URL extraction
   - Favicon hash calculation
   - Host header injection testing
   - Security header analysis
   - Network vulnerability analysis
   - Wayback machine URL retrieval
   - JavaScript file discovery
   - Broken link checking
   - HTTP request smuggling detection
   - IP address extraction
   - Domain information gathering
   - API endpoint fuzzing
   - Shodan integration for additional recon
   - 403 Forbidden bypass attempts
   - Directory and file brute-forcing
   - Local File Inclusion (LFI) scanning with Nuclei
   - Google dorking
   - Directory Traversal
   - SQL Injection
   - XSS
   - Subdomain Takeover
   - Web Server Detection
   - JavaScript file scanning for sensitive info
   - Auto Recon
   - Port Scanning
   - CIDR Notation Scanning
   - Custom Headers
   - API Fuzzing
   - AWS S3 Bucket Enumeration
   - JSON Web Token Scanning
  


4. The script uses multithreading and multiprocessing to perform scans efficiently.

5. It includes options to save results to files and customize scan parameters.

6. The tool integrates with external tools and APIs like Shodan, Nmap, and various web-based services.

7. It implements various techniques to bypass restrictions and discover vulnerabilities.

8. The script includes a CIDR notation scanner for port scanning across IP ranges.


# INSTALLATION

```bash

git clone https://github.com/whoamikiddie/StealthGhost.git

cd StealthGhost

pip3 install -r requirements.txt

sudo python3 install.py

```
# Fast Installation dependencies
```bash

sudo go run main.go 

```


# USAGE 

```

usage: stealthghost.py [-h] [-sv filename.txt | -wl filename.txt] [-th 25] [-s domain.com]
                  [-t domain.com] [-d domains.txt] [-p domains.txt] [-r domains.txt]
                  [-b domains.txt] [-pspider domain.com] [-w https://domain.com]
                  [-j domain.com] [-wc https://domain.com] [-fi https://domain.com]
                  [-fm https://domain.com] [-na https://domain.com] [-ri IP] [-rim IP]
                  [-sc domain.com] [-ph domain.txt] [-co domains.txt] [-hh domain.com]
                  [-sh domain.com] [-ed domain.com] [-smu domain.com] [-ips domain list]
                  [-dinfo domain list] [-isubs domain list] [-nft domains.txt]
                  [-n domain.com or IP] [-api domain.com] [-sho domain.com] [-fp domain.com]
                  [-db domain.com] [-cidr IP/24] [-ps 80,443,8443] [-pai IP/24]
                  [-xss https://example.com/page?param=value]
                  [-sqli https://example.com/page?param=value] [-shodan KEY]
                  [-webserver domain.com] [-javascript domain.com] [-dp 10] [-je file.txt]
                  [-hibp password] [-pm domain.com] [-ch domain.com] [-or domain.com]
                  [-asn AS55555] [-st subdomains.txt] [-ar domain.com] [-jwt token]
                  [-jwt-modify token] [--s3-scan S3_SCAN] [-v] [-c CONCURRENCY] [-nl] [-gs]
                  [-e EXTENSIONS] [-x EXCLUDE] [-u] [--shodan-api SHODAN_API]

options:
  -h, --help            show this help message and exit
  -sv, --save filename.txt
                        save output to file
  -wl, --wordlist filename.txt
                        wordlist to use
  -th, --threads 25     default 25
  -p, --probe domains.txt
                        probe domains.
  -r, --redirects domains.txt
                        links getting redirected
  -fi, --favicon https://domain.com
                        get favicon hashes
  -fm, --faviconmulti https://domain.com
                        get favicon hashes
  -ri, --reverseip IP   reverse ip lookup
  -rim, --reverseipmulti IP
                        reverse ip lookup for multiple ips
  -sc, --statuscode domain.com
                        statuscode
  -sh, --securityheaders domain.com
                        scan for security headers
  -ed, --enumeratedomain domain.com
                        enumerate domains
  -isubs, --importantsubdomains domain list
                        extract interesting subdomains from a list like dev, admin, test and
                        etc..
  -webserver, --webserver_scan domain.com
                        webserver scan
  --s3-scan S3_SCAN     Scan for exposed S3 buckets
  -v, --verbose         Increase output verbosity
  -c, --concurrency CONCURRENCY
                        Maximum number of concurrent requests
  --shodan-api SHODAN_API
                        Shodan API key for subdomain enumeration

Update:
  -u, --update          Update the script

Nuclei Scans:
  -nl, --nuclei_lfi     Find Local File Inclusion with nuclei

Vulnerability:
  -b, --brokenlinks domains.txt
                        search for broken links
  -ph, --pathhunt domain.txt
                        check for directory traversal
  -co, --corsmisconfig domains.txt
                        cors misconfiguration
  -hh, --hostheaderinjection domain.com
                        host header injection
  -smu, --smuggler domain.com
                        enumerate domains
  -fp, --forbiddenpass domain.com
                        Bypass 403 forbidden
  -xss, --xss_scan https://example.com/page?param=value
                        scan for XSS vulnerabilities
  -sqli, --sqli_scan https://example.com/page?param=value
                        scan for SQLi vulnerabilities
  -or, --openredirect domain.com
                        open redirect
  -st, --subdomaintakeover subdomains.txt
                        subdomain takeover
  -jwt, --jwt_scan token
                        analyze JWT token for vulnerabilities
  -jwt-modify, --jwt_modify token
                        modify JWT token

Crawlers:
  -pspider, --paramspider domain.com
                        extract parameters from a domain
  -w, --waybackurls https://domain.com
                        scan for waybackurls
  -j domain.com         find javascript files
  -wc, --webcrawler https://domain.com
                        scan for urls and js files
  -javascript, --javascript_scan domain.com
                        scan for sensitive info in javascript files
  -dp, --depth 10       depth of the crawl
  -je, --javascript_endpoints file.txt
                        extract javascript endpoints
  -hibp, --haveibeenpwned password
                        check if the password has been pwned

Passive Recon:
  -s domain.com         scan for subdomains
  -t, --tech domain.com
                        find technologies
  -d, --dns domains.txt
                        scan a list of domains for dns records
  -na, --networkanalyzer https://domain.com
                        net analyzer
  -ips, --ipaddresses domain list
                        get the ips from a list of domains
  -dinfo, --domaininfo domain list
                        get domain information like codes,server,content length
  -sho, --shodan domain.com
                        Recon with shodan
  -shodan, --shodan_api KEY
                        shodan api key
  -gs, --google         Google Search

Fuzzing:
  -nft, --not_found domains.txt
                        check for 404 status code
  -api, --api_fuzzer domain.com
                        Look for API endpoints
  -db, --directorybrute domain.com
                        Brute force filenames and directories
  -pm, --param_miner domain.com
                        param miner
  -ch, --custom_headers domain.com
                        custom headers
  -asn, --automoussystemnumber AS55555
                        asn
  -ar, --autorecon domain.com
                        auto recon
  -e, --extensions EXTENSIONS
                        Comma-separated list of file extensions to scan
  -x, --exclude EXCLUDE
                        Comma-separated list of status codes to exclude

Port Scanning:
  -n, --nmap domain.com or IP
                        Scan a target with nmap
  -cidr, --cidr_notation IP/24
                        Scan an ip range to find assets and services
  -ps, --ports 80,443,8443
                        Port numbers to scan
  -pai, --print_all_ips IP/24
                        Print all ip
```