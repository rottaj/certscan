# certscan

<p>Certscan is a command-line tool written in Go that fetches SSL/TLS certificates from "crt.sh" and performs enumeration on the URLs found in those certificates. This tool is useful for discovering subdomains and identifying potential attack vectors.</p>

## Installation

<p>Certscan can be installed using the go get command:</p>

```bash
go get github.com/rottaj/certscan
```

<p>Alternatively, you can clone the repository and build the binary manually:</p>

```bash
git clone https://github.com/yourusername/certscan.git
cd certscan
go build
```

## Usage

```bash
certscan [OPTIONS] DOMAIN
```

### Options

-h, --help: Show the help message and exit
-o, --output FILE: Write output to a file instead of STDOUT
-d, --depth DEPTH: Set the depth of enumeration (default: 2)
-t, --timeout TIMEOUT: Set the timeout for HTTP requests (default: 5s)
-v, --verbose: Enable verbose output

## Example

```bash
certscan -u example.com -v -o output.txt
```

<p>This command will fetch SSL/TLS certificates for example.com and enumerate URLs found within those certificates. The output will be written to output.txt.</p>

## Contributing

<p>Contributions are welcome! Please create a pull request with your changes.</p>

