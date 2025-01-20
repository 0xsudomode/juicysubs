# JuicySubs

**JuicySubs** is a simple and efficient tool for filtering juicy subdomains from a list of subdomains. It takes input from piped data or a file and matches it against a configurable list of patterns stored in `~/.config/juicysubs/config.yaml`.

## Features

- Filters subdomains based on predefined or user-configured juicy patterns.
- Automatically creates a configuration file in `~/.config/juicysubs/config.yaml` if it doesn't exist.
- Supports input from both piped data and files.
- Customizable juicy subdomain patterns using YAML configuration.

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/0xsudomode/juicysubs.git
   cd juicysubs
   ```

2. Build the tool:
   ```bash
   go build -o juicysubs juicysubs.go
   ```

3. (Optional) Move the binary to go directory in your PATH:
   ```bash
   mv juicysubs ~/go/bin
   ```

## Usage

### Piped Input

You can pipe subdomains directly into JuicySubs:
```bash
cat subdomains.txt | juicysubs
```

### File Input

Provide a file containing subdomains as an argument:
```bash
juicysubs subdomains.txt
```

### Output

The tool will output a list of juicy subdomains that match the patterns defined in the configuration file.

## Configuration

The configuration file is stored at:
```
~/.config/juicysubs/config.yaml
```

If the file doesn't exist, JuicySubs will create it automatically with a default set of juicy subdomain patterns.

### Default Configuration

```yaml
juicy_subdomains:
  - api
- dev
- stg
- test
- admin
- demo
- stage
- pre
- vpn
- uat
- sandbox
- panel
- dashboard
- internal
- intranet
- backend
- secure
- login
- auth
- pay
- billing
- wordpress
- blog
- shop
- forum
- wiki
- monitor
- status
- analytics
- logs
- backup
- archive
- old
- legacy
- files
- cdn
- assets
- media
- mail
- smtp
- beta
- preview
- staging
- qa
- support
- helpdesk
- portal
- services
- client
- customer
- user
- account
- manage
- update
- db
- database
- sys
- config
- settings
- uploads
- download
- downloads
- upload
- signin
- signup
- register
- verify
- validation
- checkout
- cart
- purchase
- order
- invoice
- app
- application
- gateway
- api-gateway
- cache
- docs
- documentation
- report
- reporting
- rest
- graphql
- v1
- v2
- v3
- static
- public
- private
- session
- token
- oauth
- sso
- saml
- directory
- dir
- login2
- auth2
- cert
- certificates
- key
- keys
- encryption
```

You can edit this file to add, remove, or modify patterns as needed.

## Example

Input file (`subdomains.txt`):
```
api.example.com
dev.example.com
www.example.com
secure.example.com
```

Command:
```bash
juicysubs subdomains.txt
```

Output:
```
api.example.com
dev.example.com
secure.example.com
```
