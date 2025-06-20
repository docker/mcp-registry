package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/docker/mcp-registry/pkg/catalog"
	"github.com/docker/mcp-registry/pkg/servers"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: catalog <server-name>")
		os.Exit(1)
	}

	name := os.Args[1]

	if err := run(name); err != nil {
		log.Fatal(err)
	}
}

func run(name string) error {
	serverFile := filepath.Join("servers", name, "server.yaml")
	server, err := servers.Read(serverFile)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "server.yaml for %s not found. Run `task create -- <your_repo>` to create a new server definition first.\n", name)
		}
		return err
	}

	tile, err := catalog.ToTile(context.Background(), server)
	if err != nil {
		return err
	}

	catalogDir := filepath.Join("catalogs", name)
	if err := os.MkdirAll(catalogDir, 0755); err != nil {
		return err
	}

	if err := writeCatalog(name, catalogDir, tile); err != nil {
		return err
	}

	return nil
}

func writeCatalog(name, catalogDir string, tile catalog.Tile) error {
	catalogFile := filepath.Join(catalogDir, "catalog.yaml")

	if err := catalog.WriteYaml(catalogFile, catalog.TopLevel{
		Version:     catalog.Version,
		Name:        "docker-mcp", // overwrite the default catalog
		DisplayName: "Local Test Catalog",
		Registry: catalog.TileList{
			{
				Name: name,
				Tile: tile,
			},
		},
	}); err != nil {
		return err
	}

	fmt.Printf("Catalog written to %s\n", catalogFile)

	return nil
}
