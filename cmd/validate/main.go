package main

import (
	"context"
	"flag"
	"fmt"
	"image"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	_ "image/jpeg"
	_ "image/png"

	"github.com/docker/mcp-registry/internal/licenses"
	"github.com/docker/mcp-registry/pkg/github"
	"github.com/docker/mcp-registry/pkg/servers"
	"gopkg.in/yaml.v3"
)

func main() {
	name := flag.String("name", "", "Name of the mcp server, name is guessed if not provided")
	flag.Parse()

	if err := run(*name); err != nil {
		log.Fatal(err)
	}
}

func run(name string) error {
	if err := isNameValid(name); err != nil {
		return err
	}

	if err := isDirectoryValid(name); err != nil {
		return err
	}

	if err := areSecretsValid(name); err != nil {
		return err
	}

	if err := IsLicenseValid(name); err != nil {
		return err
	}
	if err := isIconValid(name); err != nil {
		return err
	}

	return nil
}

// check if the name is a valid
func isNameValid(name string) error {
	// check if name has only letters, numbers, and hyphens
	if !regexp.MustCompile(`^[a-z0-9-]+$`).MatchString(name) {
		return fmt.Errorf("name is not valid. It must be a lowercase string with only letters, numbers, and hyphens")
	}

	fmt.Println("✅ Name is valid")
	return nil
}

// check if the directory is valid
// servers/<NAME>/server.yaml exists
func isDirectoryValid(name string) error {
	_, err := os.Stat(filepath.Join("servers", name, "server.yaml"))
	if err != nil {
		return err
	}
	server, err := readServerYaml(name)
	if err != nil {
		return err
	}

	// check if the server.yaml file has a valid name
	if server.Name != name {
		return fmt.Errorf("server.yaml file has a invalid name. It must be %s", name)
	}

	fmt.Println("✅ Directory is valid")
	return nil
}

// check if the secrets are valid
// secrets must be prefixed with the name of the server
func areSecretsValid(name string) error {
	// read the server.yaml file
	server, err := readServerYaml(name)
	if err != nil {
		return err
	}

	// check if the server.yaml file has a valid secrets
	if len(server.Config.Secrets) > 0 {
		for _, secret := range server.Config.Secrets {
			if !strings.HasPrefix(secret.Name, name+".") {
				return fmt.Errorf("secret %s is not valid. It must be prefixed with the name of the server", secret.Name)
			}
		}
	}

	fmt.Println("✅ Secrets are valid")
	return nil
}

// check if the license is valid
// the license must be valid
func IsLicenseValid(name string) error {
	ctx := context.Background()
	client := github.New()
	server, err := readServerYaml(name)
	if err != nil {
		return err
	}
	repository, err := client.GetProjectRepository(ctx, server.Source.Project)
	if err != nil {
		return err
	}

	if !licenses.IsValid(repository.License) {
		return fmt.Errorf("project %s is licensed under %s which may be incompatible with some tools", server.Source.Project, repository.License.GetName())
	}
	fmt.Println("✅ License is valid")

	return nil
}

func isIconValid(name string) error {
	server, err := readServerYaml(name)
	if err != nil {
		return err
	}

	if server.About.Icon == "" {
		fmt.Println("🛑 No icon found")
		return nil
	}
	// fetch the image and check the size
	resp, err := http.Get(server.About.Icon)
	if err != nil {
		fmt.Println("🛑 Icon could not be fetched")
		return nil
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return fmt.Errorf("icon is not valid. It must be a valid image")
	}
	if resp.ContentLength > 2*1024*1024 {
		return fmt.Errorf("icon is too large. It must be less than 2MB")
	}
	img, format, err := image.DecodeConfig(resp.Body)
	if err != nil {
		return err
	}
	if format != "png" {
		return fmt.Errorf("icon is not a png. It must be a png")
	}

	if img.Width > 512 || img.Height > 512 {
		return fmt.Errorf("image is too large. It must be less than 512x512")
	}

	fmt.Println("✅ Icon is valid")
	return nil
}

func readServerYaml(name string) (servers.Server, error) {
	serverYaml, err := os.ReadFile(filepath.Join("servers", name, "server.yaml"))
	if err != nil {
		return servers.Server{}, err
	}
	var server servers.Server
	err = yaml.Unmarshal(serverYaml, &server)
	if err != nil {
		return servers.Server{}, err
	}
	return server, nil
}
