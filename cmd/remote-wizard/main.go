package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"github.com/docker/mcp-registry/pkg/servers"
	"gopkg.in/yaml.v3"
)

var (
	transportTypes = []string{
		"streamable-http",
		"sse",
	}

	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#4ECDC4")).
			Bold(true).
			Padding(1, 2)

	headerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#4ECDC4")).
			Bold(true).
			Margin(1, 0).
			Padding(1, 4)
)

type RemoteWizardData struct {
	ServerName    string
	Category      string
	Title         string
	Description   string
	Icon          string
	TransportType string
	URL           string
	DocsURL       string
	UseOAuth      bool
}

func main() {
	fmt.Print(titleStyle.Render("🐳 MCP Remote Server Registry Wizard"))
	fmt.Print(headerStyle.Render("Welcome! Let's add your remote MCP server to the registry."))
	fmt.Println()
	fmt.Println()

	var data RemoteWizardData

	// Basic Information Form
	basicForm := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Server Name").
				Description("Enter the name for your MCP server (e.g., 'my-awesome-server')").
				Value(&data.ServerName).
				Validate(func(s string) error {
					if strings.TrimSpace(s) == "" {
						return fmt.Errorf("server name is required")
					}
					if strings.Contains(s, " ") {
						return fmt.Errorf("server name cannot contain spaces")
					}
					exists, err := checkLocalServerExists(s)
					if err != nil {
						return err
					}
					if exists {
						return fmt.Errorf("server name %s already exists", s)
					}
					return nil
				}),

			huh.NewInput().
				Title("Category").
				Description("Enter the category that best describes your MCP server\n\t\t(e.g., ai, database, devops, productivity, search, communication, etc.)").
				Value(&data.Category).
				Validate(func(s string) error {
					if strings.TrimSpace(s) == "" {
						return fmt.Errorf("category is required")
					}
					return nil
				}),
		).Title("📋 Basic Information"),
	).WithTheme(huh.ThemeCharm())

	if err := basicForm.Run(); err != nil {
		log.Fatal(err)
	}

	// Server Details Form
	detailsForm := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Server Title").
				Description("Enter a descriptive title for your MCP server").
				Value(&data.Title).
				Validate(func(s string) error {
					if strings.TrimSpace(s) == "" {
						return fmt.Errorf("title is required")
					}
					return nil
				}),

			huh.NewText().
				Title("Description").
				Description("Enter a detailed description of what your MCP server does").
				Value(&data.Description).
				Validate(func(s string) error {
					if strings.TrimSpace(s) == "" {
						return fmt.Errorf("description is required")
					}
					return nil
				}),

			huh.NewInput().
				Title("Icon URL").
				Description("Enter an icon URL (e.g., https://example.com/icon.png or use Google's favicon service)").
				Value(&data.Icon).
				Validate(func(s string) error {
					if strings.TrimSpace(s) == "" {
						return fmt.Errorf("icon URL is required")
					}
					return nil
				}),

			huh.NewInput().
				Title("Documentation URL").
				Description("Enter the URL to your server's documentation (will be saved in readme.md)").
				Value(&data.DocsURL).
				Validate(func(s string) error {
					if strings.TrimSpace(s) == "" {
						return fmt.Errorf("documentation URL is required")
					}
					if !strings.HasPrefix(s, "http://") && !strings.HasPrefix(s, "https://") {
						return fmt.Errorf("URL must start with http:// or https://")
					}
					return nil
				}),
		).Title("📝 Server Details"),
	).WithTheme(huh.ThemeCharm())

	if err := detailsForm.Run(); err != nil {
		log.Fatal(err)
	}

	// Remote Configuration Form
	remoteForm := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Transport Type").
				Description("Select the transport protocol your remote server uses").
				Options(huh.NewOptions(transportTypes...)...).
				Value(&data.TransportType),

			huh.NewInput().
				Title("Server URL").
				Description("Enter the full URL of your remote MCP server (e.g., https://mcp.example.com/mcp)").
				Value(&data.URL).
				Validate(func(s string) error {
					if strings.TrimSpace(s) == "" {
						return fmt.Errorf("server URL is required")
					}
					if !strings.HasPrefix(s, "http://") && !strings.HasPrefix(s, "https://") {
						return fmt.Errorf("URL must start with http:// or https://")
					}
					return nil
				}),
		).Title("🌐 Remote Configuration"),
	).WithTheme(huh.ThemeCharm())

	if err := remoteForm.Run(); err != nil {
		log.Fatal(err)
	}

	// OAuth Configuration Form
	oauthForm := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title("Use OAuth?").
				Description("Does your server require OAuth authentication?").
				Value(&data.UseOAuth),
		).Title("🔐 Authentication"),
	).WithTheme(huh.ThemeCharm())

	if err := oauthForm.Run(); err != nil {
		log.Fatal(err)
	}

	// Generate and save the configuration
	if err := generateAndSave(&data); err != nil {
		log.Fatal(err)
	}

	fmt.Print(headerStyle.Render("✅ Success! Your remote MCP server configuration has been generated."))
	fmt.Println()
	fmt.Printf("📁 Generated files in servers/%s/:\n", data.ServerName)
	fmt.Printf("   - server.yaml (server configuration)\n")
	fmt.Printf("   - tools.json (empty, for dynamic tool discovery)\n")
	fmt.Printf("   - readme.md (documentation link)\n")
	fmt.Println()
	fmt.Println("🚀 Next steps:")
	fmt.Println("1. Review the generated server.yaml file")
	fmt.Println("2. Test your server:")
	fmt.Println("   task catalog -- " + data.ServerName)
	fmt.Println("   docker mcp catalog import $PWD/catalogs/" + data.ServerName + "/catalog.yaml")
	fmt.Println("   docker mcp server enable " + data.ServerName)
	if data.UseOAuth {
		fmt.Println("   docker mcp oauth authorize " + data.ServerName + " (for OAuth)")
	}
	fmt.Println("3. Reset catalog when done: docker mcp catalog reset")
	fmt.Println("4. Create a pull request to add it to the registry")
}

func generateAndSave(data *RemoteWizardData) error {
	// Build the server configuration
	config := servers.Server{
		Name: data.ServerName,
		Type: "remote",
		Dynamic: &servers.Dynamic{
			Tools: true,
		},
		Meta: servers.Meta{
			Category: data.Category,
			Tags:     []string{data.Category, "remote"},
		},
		About: servers.About{
			Title:       data.Title,
			Description: data.Description,
			Icon:        data.Icon,
		},
		Remote: servers.Remote{
			TransportType: data.TransportType,
			URL:           data.URL,
		},
	}

	// Add OAuth configuration if needed
	if data.UseOAuth {
		config.OAuth = []servers.OAuthProvider{
			{
				Provider: data.ServerName,
				Secret:   fmt.Sprintf("%s.personal_access_token", data.ServerName),
				Env:      fmt.Sprintf("%s_PERSONAL_ACCESS_TOKEN", strings.ToUpper(strings.ReplaceAll(data.ServerName, "-", "_"))),
			},
		}
	}

	// Create directory
	serverDir := filepath.Join("servers", data.ServerName)
	if err := os.MkdirAll(serverDir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	// Marshal to YAML
	yamlData, err := yaml.Marshal(&config)
	if err != nil {
		return fmt.Errorf("failed to marshal YAML: %w", err)
	}

	// Write server.yaml file
	configPath := filepath.Join(serverDir, "server.yaml")
	if err := os.WriteFile(configPath, yamlData, 0644); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	// Create empty tools.json file (always empty for remote servers with dynamic tools)
	toolsPath := filepath.Join(serverDir, "tools.json")
	if err := os.WriteFile(toolsPath, []byte("[]"), 0644); err != nil {
		return fmt.Errorf("failed to write tools file: %w", err)
	}

	// Create readme.md file with documentation link
	readmePath := filepath.Join(serverDir, "readme.md")
	readmeContent := fmt.Sprintf("Docs: %s\n", data.DocsURL)
	if err := os.WriteFile(readmePath, []byte(readmeContent), 0644); err != nil {
		return fmt.Errorf("failed to write readme file: %w", err)
	}

	return nil
}

func checkLocalServerExists(name string) (bool, error) {
	entries, err := os.ReadDir("servers")
	if err != nil {
		return false, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			if entry.Name() == name {
				return true, nil
			}
		}
	}

	return false, nil
}
