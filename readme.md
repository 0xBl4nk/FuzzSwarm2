# FuzzSwarm

**FuzzSwarm** is an advanced fuzzing tool designed to brute-force HTTP endpoints and uncover vulnerabilities in web applications and APIs. It efficiently handles large-scale testing through multi-threaded fuzzing, offering support for both GET and POST requests. 

Users can inject custom headers and even load headers from a file for more controlled tests. FuzzSwarm also allows precise brute-forcing by specifying numeric ranges, filtering responses by size to focus on significant results, and controlling rate limits and timeouts to ensure safe and measured testing. 

Additionally, it offers optional proxy usage and SSL configuration for secure environments, along with wordlist support for custom payloads. This makes FuzzSwarm a robust and versatile tool for pentesters, bug bounty hunters, and security professionals looking to expose security flaws in various web environments.

## 📃 Requirements

- `Go 1.23.1` or higher

## Installation

1. Clone this repository:
    ```bash
    git clone https://github.com/0xBl4nk/FuzzSwarm2.git
    ```

2. Install the dependencies:
    ```bash
    make build
    ```

## Usage

To run FuzzSwarm, use the following syntax:

```bash
# Use "FUZZ" placeholder where you want to fuzz
./FuzzSwarm --url <url> --range/--wordlist [Flags]
```

### Example Usage:

```bash
./FuzzSwarm -X POST --url http://example.com --range 1-10000,1 --data '{"number": FUZZ}' -H "Content-Type: application/json" -f 34 -v
```
<img src="https://i.imgur.com/guvTo1Y.png">

### Example Scripts Usage:
<img src="https://i.imgur.com/1rjekSu.png">

### Available Parameters:

```
Flags:
  -d, --data string           POST data.
  -f, --filter-size int       Filter responses by size (skip responses with this size).
  -H, --headers string        Custom header.
      --headers-file string   Path to the headers file.
  -h, --help                  help for FuzzSwarm
  -X, --method string         HTTP method to use: GET or POST. (Default: GET) (default "GET")
  -R, --range string          Range of numbers to use, format start-end,digits (e.g., 1-10000,3).
      --rate-limit int        Rate limit in milliseconds between requests.
  -t, --threads int           Number of threads to use for fuzzing. (default 10)
      --timeout int           Set timeout. (default 10)
  -u, --url string            The target URL.
      --use-proxy             Display verbose output including response preview.
      --use-ssl string        Enable SSL certificate from .env file.
  -v, --verbose               Enable proxy configuration from .env file.
  -W, --wordlist string       Path to the wordlist file.

```

# How to Generate a Valid SSL Certificate with OpenSSL
fuzzswarm uses unified certificates, i.e. key and certificate in the same file.

```bash
cat certificate.pem privatekey.pem > fullcert.pem
```
If your unified certificate is .c12, you will need to convert to pem.
```bash
openssl pkcs12 -in yourfile.p12 -clcerts -nokeys -out certificate.pem
openssl pkcs12 -in yourfile.p12 -nocerts -out privatekey.pem
openssl rsa -in privatekey.pem -out privatekey.pem
cat certificate.pem privatekey.pem > fullcert.pem
```

## Contributing

1. Fork this repository.
2. Create a new branch: `git checkout -b <branch_name>`.
3. Make your changes and commit: `git commit -m '<commit_message>'`.
4. Push to your branch: `git push origin <branch_name>`.
5. Open a pull request.
