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
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/docker/mcp-registry/pkg/servers"
)

// main processes the provided server names and cleans build artifacts for each.
func main() {
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Fprintln(os.Stderr, "Usage: task clean -- <server> [server...]")
		os.Exit(1)
	}

	var failed bool
	for _, name := range flag.Args() {
		if err := cleanServer(name); err != nil {
			fmt.Fprintf(os.Stderr, "cleanup failed for %s: %v\n", name, err)
			failed = true
		}
	}

	if failed {
		os.Exit(1)
	}
}

// cleanServer removes generated artifacts and Docker images for the server.
func cleanServer(name string) error {
	serverPath := filepath.Join("servers", name, "server.yaml")
	server, err := servers.Read(serverPath)
	if err != nil {
		return fmt.Errorf("reading server file: %w", err)
	}

	removeCatalog(name)
	removeDockerImage(server.Image)
	removeDockerImage("check")
	pruneDockerBuilder()
	pruneDockerImages()

	return nil
}

// removeCatalog deletes the generated catalog directory if it exists.
func removeCatalog(name string) {
	path := filepath.Join("catalogs", name)
	if err := os.RemoveAll(path); err != nil {
		fmt.Fprintf(os.Stderr, "warning: removing %s: %v\n", path, err)
	}
}

// removeDockerImage removes the specified Docker image, ignoring missing images.
func removeDockerImage(image string) {
	if image == "" {
		return
	}

	out, err := exec.Command("docker", "image", "rm", "-f", image).CombinedOutput()
	if err != nil {
		msg := string(out)
		if strings.Contains(msg, "No such image") {
			return
		}
		fmt.Fprintf(os.Stderr, "warning: removing image %s: %v\n%s", image, err, msg)
	} else {
		fmt.Print(string(out))
	}
}

// pruneDockerBuilder removes unused builder cache entries.
func pruneDockerBuilder() {
	cmd := exec.Command("docker", "builder", "prune", "--force")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "warning: pruning builder cache: %v\n", err)
	}
}

// pruneDockerImages removes dangling Docker images.
func pruneDockerImages() {
	cmd := exec.Command("docker", "image", "prune", "--force")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "warning: pruning images: %v\n", err)
	}
}
