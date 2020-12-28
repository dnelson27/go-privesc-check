# Simple Linux Privesc Scanner in Go
## About
This is an experiment to achieve some of the basic functionality found in tools like https://github.com/carlospolop/privilege-escalation-awesome-scripts-suite in a compiled language. The end goal is to increase speed, decrease noise, and increase flexibility in automatically scanning for privilege escalation vulnerabilities.

## Current Functionality
- List files in a directory with SUID bit set

## Upcoming Functionality
- Check files for SUID bits across a target host that line up with common exploitable binaries, as listed in https://gtfobins.github.io/.
- Scan for vulnerable software associated with privilege escalation from user-provided software list.
- Scan a host's ecosystem environment and neighbor containers for misconfigurations that could result in information disclosure, lateral movement, or privilege escalation
- Scan a host's AWS ecosystem and neighbor containers for misconfigurations that could result in information disclosure, lateral movement, or privilege escalation

## Usage
### Binary
`./suidcheck target-directory {stdout | file} [output file]`

### Build From Source (From Respective Directory)
Keep Debugging Information (Increases Size): `go build -o suidcheck main.go suidcheck.go writer.go`
Strip Debugging Information (Reduces Size): `go build -o suidcheck -ldflags="-w -s" main.go suidcheck.go writer.go`