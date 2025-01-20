package main

import (
	"bufio"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

// Config structure to hold juicy subdomain patterns
type Config struct {
	JuicySubdomains []string `yaml:"juicy_subdomains"`
}

// GetConfigFilePath returns the path to the config file in ~/.config/juicysubs/config.yaml
func GetConfigFilePath() string {
	usr, err := user.Current()
	if err != nil {
		fmt.Println("Error fetching user home directory:", err)
		os.Exit(1)
	}

	configDir := filepath.Join(usr.HomeDir, ".config", "juicysubs")
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		if err := os.MkdirAll(configDir, 0755); err != nil {
			fmt.Println("Error creating config directory:", err)
			os.Exit(1)
		}
	}

	return filepath.Join(configDir, "config.yaml")
}

// LoadConfig loads the configuration file, creating it if it doesn't exist
func LoadConfig(configFilePath string) Config {
	var config Config
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		// Create default config if it doesn't exist
		defaultConfig := Config{
			JuicySubdomains: []string{
				"api", "dev", "stg", "test", "admin", "demo", "stage", "pre",
				"vpn", "uat", "sandbox", "panel", "dashboard", "internal",
				"intranet", "backend", "secure", "login", "auth", "pay", "billing",
				"wordpress", "blog", "shop", "forum", "wiki", "monitor", "status",
				"analytics", "logs", "backup", "archive", "old", "legacy", "files",
				"cdn", "assets", "media", "mail", "smtp", "beta", "preview", "staging",
				"qa", "support", "helpdesk", "portal", "services", "client", "customer",
				"user", "account", "manage", "update", "db", "database", "sys", "config",
				"settings", "uploads", "download", "downloads", "upload", "signin",
				"signup", "register", "verify", "validation", "checkout", "cart",
				"purchase", "order", "invoice", "app", "application", "gateway",
				"api-gateway", "cache", "docs", "documentation", "report", "reporting",
				"rest", "graphql", "v1", "v2", "v3", "static", "public", "private",
				"session", "token", "oauth", "sso", "saml", "directory", "dir",
				"login2", "auth2", "cert", "certificates", "key", "keys", "encryption",
			},
		}

		data, _ := yaml.Marshal(defaultConfig)
		if err := ioutil.WriteFile(configFilePath, data, 0644); err != nil {
			fmt.Println("Error writing default config file:", err)
			os.Exit(1)
		}
		fmt.Println("Default config file created at:", configFilePath)
		return defaultConfig
	}

	// Load existing config
	data, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		fmt.Println("Error reading config file:", err)
		os.Exit(1)
	}

	if err := yaml.Unmarshal(data, &config); err != nil {
		fmt.Println("Error parsing config file:", err)
		os.Exit(1)
	}

	return config
}

// FilterJuicySubdomains filters subdomains based on juicy patterns
func FilterJuicySubdomains(subdomains []string, juicyPatterns []string) []string {
	var filtered []string
	for _, subdomain := range subdomains {
		for _, pattern := range juicyPatterns {
			if strings.Contains(subdomain, pattern) {
				filtered = append(filtered, subdomain)
				break
			}
		}
	}
	return filtered
}

func main() {
	configFilePath := GetConfigFilePath()
	config := LoadConfig(configFilePath)

	// Check for piped input or a file argument
	info, err := os.Stdin.Stat()
	if err != nil {
		fmt.Println("Error checking stdin:", err)
		os.Exit(1)
	}

	var subdomains []string
	if (info.Mode() & os.ModeCharDevice) == 0 {
		// Read from piped input
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			subdomains = append(subdomains, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading piped input:", err)
			os.Exit(1)
		}
	} else if len(os.Args) > 1 {
		// Read from file argument
		filePath := os.Args[1]
		data, err := ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Println("Error reading file:", err)
			os.Exit(1)
		}
		subdomains = strings.Split(string(data), "\n")
	} else {
		// Show usage if no input
		fmt.Println("Usage: juicysubs < file_with_subdomains.txt")
		fmt.Println("       cat subdomains.txt | juicysubs")
		os.Exit(1)
	}

	// Filter juicy subdomains
	juicySubdomains := FilterJuicySubdomains(subdomains, config.JuicySubdomains)

	// Print results
	for _, subdomain := range juicySubdomains {
		fmt.Println(subdomain)
	}
}
