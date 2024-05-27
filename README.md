# rogue-mac-check
Compare authorized [BSSIDs](https://en.wikipedia.org/wiki/BSSID) (access point MAC addresses) against found BSSIDs to discover rogue wireless access points.

## Syntax
`rogue-mac-check.bin /path/to/allowed_macs.txt /path/to/found_macs.txt`

## Input
Both input files must be plaintext and newline-separated. Each line should contain a MAC address, however surrounding text is allowed. The **first file** contains the MAC addresses that are known and authorized, and are used as the whitelist to compare against. The **second file** contains the MAC addresses that have been found throughout the course of reconnaissance, using tools such as Kismet, airodump-ng, iwlist, etc.

## Output
All lines from the second file that contained unmatched MAC addresses.

## Compiling
### One-off
Use "go run": `go run /path/to/rogue-mac-check.go`
### Compile Binary
Use "go build": `go build -o /path/to/output/binary /path/to/rogue-mac-check.go`
### Cross-compile
Use "go build", specifying the desired [operating system and architecture](https://golang.org/doc/install/source#environment): `GOOS=windows GOARCH=386 go build -o /path/to/binary.exe /path/to/rogue-mac-check.go`
