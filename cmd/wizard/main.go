package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"github.com/docker/mcp-registry/internal/licenses"
	"github.com/docker/mcp-registry/pkg/github"
	"github.com/docker/mcp-registry/pkg/servers"
	"gopkg.in/yaml.v3"
)

var (
	categories = []string{
		"ai",
		"data-visualization",
		"database",
		"devops",
		"ecommerce",
		"finance",
		"games",
		"communication",
		"monitoring",
		"productivity",
		"search",
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

type Config struct {
	Description string      `yaml:"description"`
	Secrets     []SecretVar `yaml:"secrets,omitempty"`
	Env         []EnvVar    `yaml:"env,omitempty"`
	Parameters  *Parameters `yaml:"parameters,omitempty"`
}

type Run struct {
	Volumes []string `yaml:"volumes,omitempty"`
	Command []string `yaml:"command,omitempty"`
}

type Volumes struct {
	HostPath      Volume `yaml:"hostPath"`
	ContainerPath Volume `yaml:"containerPath"`
}

type Volume struct {
	Name        string `yaml:"name"`
	Value       string `yaml:"value"`
	Description string `yaml:"description"`
}
type SecretVar struct {
	Name    string `yaml:"name"`
	Env     string `yaml:"env"`
	Example string `yaml:"example"`
}

type EnvVar struct {
	Name    string `yaml:"name"`
	Example string `yaml:"example"`
	Value   string `yaml:"value"`
}

type Parameters struct {
	Type       string                 `yaml:"type"`
	Properties map[string]interface{} `yaml:"properties"`
	Required   []string               `yaml:"required,omitempty"`
}

type WizardData struct {
	ServerName  string
	GitHubRepo  string
	Branch      string
	Category    string
	Title       string
	Description string
	Icon        string
	Image       string
	Secrets     []SecretInput
	EnvVars     []EnvInput
	AddSecrets  bool
	AddEnvVars  bool
	Volumes     []Volumes
	Command     []string
	AddVolumes  bool
}

type SecretInput struct {
	Name    string
	EnvName string
	Example string
}

type EnvInput struct {
	Name    string
	Example string
	Value   string
}

func main() {
	fmt.Print(titleStyle.Render("🐳 MCP Server Registry Wizard"))
	fmt.Print(headerStyle.Render("Welcome! Let's add your MCP server to the registry."))
	fmt.Println()
	fmt.Println()

	var data WizardData

	// Basic Information Form
	repoForm := huh.NewForm(
		huh.NewGroup(

			huh.NewInput().
				Description("Enter the GitHub repository URL (e.g., 'https://github.com/user/repo').\n\t\t⚠️ Remember that your repository needs to have a Dockerfile.").
				Value(&data.GitHubRepo).
				Validate(func(s string) error {
					s = strings.TrimSpace(s)
					if s == "" {
						return fmt.Errorf("repository URL is required")
					}
					if !strings.HasPrefix(s, "https://") {
						s = "https://" + s
					}
					data.GitHubRepo = s
					return validateGithubRepo(&data)
				}),
		).Title("📋 Repo Information"),
	).WithTheme(huh.ThemeCharm())

	if err := repoForm.Run(); err != nil {
		log.Fatal(err)
	}

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
				Title("Branch (optional)").
				Description("Enter the branch name (leave empty for default branch)").
				Value(&data.Branch),

			huh.NewSelect[string]().
				Title("Category").
				Description("Select the category that best describes your MCP server").
				Options(huh.NewOptions(categories...)...).
				Value(&data.Category),
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
				Title("Icon URL (optional)").
				Description("Enter an icon URL (or leave the default)").
				Value(&data.Icon),

			huh.NewInput().
				Title("Docker Image (optional)").
				Description("Enter custom Docker image (or leave the default mcp/NAME)").
				Value(&data.Image),
		).Title("📝 Server Details"),
	).WithTheme(huh.ThemeCharm())

	if err := detailsForm.Run(); err != nil {
		log.Fatal(err)
	}

	// Configuration Options Form
	configForm := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title("Add Secrets?").
				Description("Does your server require any secret variables (passwords, API keys, etc.)?").
				Value(&data.AddSecrets),

			huh.NewConfirm().
				Title("Add Environment Variables?").
				Description("Does your server require any environment variables for configuration?").
				Value(&data.AddEnvVars),
		).Title("⚙️ Configuration"),
	).WithTheme(huh.ThemeCharm())

	if err := configForm.Run(); err != nil {
		log.Fatal(err)
	}

	// Secrets Configuration
	if data.AddSecrets {
		if err := collectSecrets(&data); err != nil {
			log.Fatal(err)
		}
	}

	// Environment Variables Configuration
	if data.AddEnvVars {
		if err := collectEnvVars(&data); err != nil {
			log.Fatal(err)
		}
	}

	configForm = huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title("Add Volumes?").
				Description("Do you want to add any volumes to your MCP server?").
				Value(&data.AddVolumes),
		).Title("⚙️ Configuration"),
	).WithTheme(huh.ThemeCharm())

	if err := configForm.Run(); err != nil {
		log.Fatal(err)
	}

	if data.AddVolumes {
		if err := collectVolumes(&data); err != nil {
			log.Fatal(err)
		}
	}

	// Generate and save the configuration
	if err := generateAndSave(&data); err != nil {
		log.Fatal(err)
	}

	fmt.Print(headerStyle.Render("✅ Success! Your MCP server configuration has been generated."))
	fmt.Println()
	fmt.Printf("📁 Generated at: servers/%s/server.yaml\n", data.ServerName)
	fmt.Println()
	fmt.Println("🚀 Next steps:")
	fmt.Println("1. Review the generated server.yaml file")
	fmt.Println("2. Build your server locally with: task build -- " + data.ServerName)
	fmt.Println("3. Generate the catalog with: task catalog -- " + data.ServerName)
	fmt.Println("4. Test your server locally in Docker Desktop with: task import -- " + data.ServerName)
	fmt.Println("5. Reset your catalog in Docker Desktop with: task reset")
	fmt.Println("6. Create a pull request to add it to the registry")

}

func collectVolumes(data *WizardData) error {
	fmt.Print(headerStyle.Render("🔐 Configure Volumes"))
	fmt.Println()

	for {
		var volume Volumes
		var addAnother bool

		volumeForm := huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Title("Volume Host Name").
					Description("Enter the volume host name (e.g., 'data', 'config')").
					Value(&volume.HostPath.Name),
				huh.NewInput().
					Title("Volume Host Path Description").
					Description("Enter the volume host path description (e.g., 'Data directory')").
					Value(&volume.HostPath.Description),
				huh.NewInput().
					Title("Volume Host Path Value").
					Description("Enter the volume host path value (e.g., '/data', '/config') or leave empty to let the user choose the path").
					Value(&volume.HostPath.Value),
				huh.NewInput().
					Title("Volume Container Name").
					Description("Enter the volume container name (e.g., 'data', 'config')").
					Value(&volume.ContainerPath.Name),
				huh.NewInput().
					Title("Volume Container Path Description").
					Description("Enter the volume container path description (e.g., 'Data directory')").
					Value(&volume.ContainerPath.Description),
				huh.NewInput().
					Title("Volume Container Path Value").
					Description("Enter the volume container path value (e.g., '/data', '/config') or leave empty to let the user choose the path").
					Value(&volume.ContainerPath.Value),
				huh.NewConfirm().
					Title("Add Another Volume?").
					Description("Do you want to add another volume?").
					Value(&addAnother),
			).Title("Volume Configuration"),
		).WithTheme(huh.ThemeCharm())

		if err := volumeForm.Run(); err != nil {
			return err
		}

		data.Volumes = append(data.Volumes, volume)

		if !addAnother {
			break
		}
	}

	return nil
}

func collectSecrets(data *WizardData) error {
	fmt.Print(headerStyle.Render("🔐 Configure Secrets"))
	fmt.Println()

	for {
		var secret SecretInput
		var addAnother bool

		secretForm := huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Title("Secret Name").
					Description("Enter the secret name (e.g., 'api_key', 'password')").
					Value(&secret.Name).
					Validate(func(s string) error {
						if strings.TrimSpace(s) == "" {
							return fmt.Errorf("secret name is required")
						}
						return nil
					}),

				huh.NewInput().
					Title("Environment Variable Name").
					Description("Enter the environment variable name (e.g., 'API_KEY', 'PASSWORD')").
					Value(&secret.EnvName).
					Validate(func(s string) error {
						if strings.TrimSpace(s) == "" {
							return fmt.Errorf("environment variable name is required")
						}
						return nil
					}),

				huh.NewInput().
					Title("Example Value").
					Description("Enter an example value (for documentation)").
					Value(&secret.Example).
					Validate(func(s string) error {
						if strings.TrimSpace(s) == "" {
							return fmt.Errorf("example value is required")
						}
						return nil
					}),

				huh.NewConfirm().
					Title("Add Another Secret?").
					Description("Do you want to add another secret variable?").
					Value(&addAnother),
			).Title("Secret Configuration"),
		).WithTheme(huh.ThemeCharm())

		if err := secretForm.Run(); err != nil {
			return err
		}

		data.Secrets = append(data.Secrets, secret)

		if !addAnother {
			break
		}
	}

	return nil
}

func collectEnvVars(data *WizardData) error {
	fmt.Print(headerStyle.Render("🌍 Configure Environment Variables"))
	fmt.Println()

	for {
		var envVar EnvInput
		var addAnother bool

		envForm := huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Title("Environment Variable Name").
					Description("Enter the environment variable name (e.g., 'HOST', 'PORT')").
					Value(&envVar.Name).
					Validate(func(s string) error {
						if strings.TrimSpace(s) == "" {
							return fmt.Errorf("environment variable name is required")
						}
						return nil
					}),

				huh.NewInput().
					Title("Example Value").
					Description("Enter an example value").
					Value(&envVar.Example).
					Validate(func(s string) error {
						if strings.TrimSpace(s) == "" {
							return fmt.Errorf("example value is required")
						}
						return nil
					}),

				huh.NewInput().
					Title("Template Value").
					Description("Enter the template value (e.g., '{{server.host}}') or leave empty to use variable name").
					Value(&envVar.Value),

				huh.NewConfirm().
					Title("Add Another Environment Variable?").
					Description("Do you want to add another environment variable?").
					Value(&addAnother),
			).Title("Environment Variable Configuration"),
		).WithTheme(huh.ThemeCharm())

		if err := envForm.Run(); err != nil {
			return err
		}

		if envVar.Value == "" {
			envVar.Value = fmt.Sprintf("{{%s.%s}}", strings.ToLower(data.ServerName), strings.ToLower(envVar.Name))
		}

		data.EnvVars = append(data.EnvVars, envVar)

		if !addAnother {
			break
		}
	}

	return nil
}

func generateAndSave(data *WizardData) error {
	// Set defaults
	if data.Image == "" {
		data.Image = "mcp/" + data.ServerName
	}
	if data.Icon == "" {
		data.Icon = "mcp/" + data.ServerName
	}

	// Build the server configuration
	config := servers.Server{
		Name:  data.ServerName,
		Image: data.Image,
		Type:  "server",
		Meta: servers.Meta{
			Category: data.Category,
			Tags:     []string{data.Category},
		},
		About: servers.About{
			Title:       data.Title,
			Description: data.Description,
			Icon:        data.Icon,
		},
		Source: servers.Source{
			Project: data.GitHubRepo,
		},
	}

	if data.Branch != "" {
		config.Source.Branch = data.Branch
	}

	// Add configuration if needed
	if len(data.Secrets) > 0 || len(data.EnvVars) > 0 {
		config.Config = servers.Config{
			Description: fmt.Sprintf("Configure the connection to %s", data.Title),
		}

		// Add secrets
		for _, secret := range data.Secrets {
			config.Config.Secrets = append(config.Config.Secrets, servers.Secret{
				Name:    fmt.Sprintf("%s.%s", data.ServerName, secret.Name),
				Env:     secret.EnvName,
				Example: secret.Example,
			})
		}

		// Add environment variables
		for _, envVar := range data.EnvVars {
			config.Config.Env = append(config.Config.Env, servers.Env{
				Name:    envVar.Name,
				Example: envVar.Example,
				Value:   envVar.Value,
			})
		}

		// Add parameters if we have env vars
		if len(data.EnvVars) > 0 {
			config.Config.Parameters = servers.Schema{
				Type:       "object",
				Properties: make(servers.SchemaList, 0),
			}

			for _, envVar := range data.EnvVars {
				paramName := strings.ToLower(strings.ReplaceAll(envVar.Name, "_", ""))
				config.Config.Parameters.Properties = append(config.Config.Parameters.Properties, servers.SchemaEntry{
					Schema: servers.Schema{
						Type: "string",
					},
					Name: paramName,
				})
			}
		}
	}

	if len(data.Volumes) > 0 {
		config.Run = servers.Run{
			Volumes: make([]string, 0),
		}
		for _, vol := range data.Volumes {
			fmt.Println(vol.HostPath)
			fmt.Println(vol.ContainerPath)
			host, container := "", ""
			//
			if vol.HostPath.Value != "" {
				host = vol.HostPath.Value
			} else {
				host = fmt.Sprintf("{{%s.%s}}", data.ServerName, vol.HostPath.Name)
			}
			if vol.ContainerPath.Value != "" {
				container = vol.ContainerPath.Value
			} else {
				container = fmt.Sprintf("{{%s.%s}}", data.ServerName, vol.ContainerPath.Name)
			}
			fmt.Println(host, container)
			config.Run.Volumes = append(config.Run.Volumes, fmt.Sprintf("%s:%s", host, container))
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

	// Write to file
	configPath := filepath.Join(serverDir, "server.yaml")
	if err := os.WriteFile(configPath, yamlData, 0644); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}

func validateGithubRepo(data *WizardData) error {

	ctx := context.Background()
	client := github.New()
	repository, err := client.GetProjectRepository(ctx, data.GitHubRepo)
	if err != nil {
		return err
	}
	if !licenses.IsValid(repository.License) {
		fmt.Println("[WARNING] Project", data.GitHubRepo, "is licensed under", repository.License.GetName(), "which may be incompatible with some tools")
	}

	detectedInfo, err := github.DetectBranchAndDirectory(data.GitHubRepo, repository)
	if err != nil {
		return err
	}

	parts := strings.Split(strings.ToLower(data.GitHubRepo), "/")
	name := parts[len(parts)-1]

	data.ServerName = name
	data.Image = "mcp/" + name

	data.Title = strings.ToUpper(name[0:1]) + strings.ReplaceAll(name[1:], "-", " ")

	if detectedInfo.Branch == repository.GetDefaultBranch() {
		data.Branch = ""
	} else {
		data.Branch = detectedInfo.Branch
	}

	refProjectURL := detectedInfo.ProjectURL
	if repository.GetParent() != nil {
		refProjectURL = repository.GetParent().GetHTMLURL()
	}
	icon, err := client.FindIcon(ctx, refProjectURL)
	if err != nil {
		return err
	}
	data.Icon = icon

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
