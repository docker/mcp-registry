/*
Copyright © 2025 Docker, Inc.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

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
