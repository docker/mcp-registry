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

// pinTarget describes a server that updated its commit pin within a pull request.
type pinTarget struct {
	// Server is the registry entry name (directory) that was updated.
	Server string `json:"server"`
	// File is the relative YAML path that changed for the server.
	File string `json:"file"`
	// Image is the Docker image identifier associated with the server.
	Image string `json:"image"`
	// Project is the upstream repository URL for the server source.
	Project string `json:"project"`
	// Directory points to the subdirectory inside the upstream repository, when set.
	Directory string `json:"directory,omitempty"`
	// OldCommit contains the previously pinned commit SHA.
	OldCommit string `json:"old_commit"`
	// NewCommit contains the newly pinned commit SHA.
	NewCommit string `json:"new_commit"`
}

// newServerTarget captures metadata for a newly added local server.
type newServerTarget struct {
	// Server is the registry entry name for the newly added server.
	Server string `json:"server"`
	// File is the YAML file that defines the server in the registry.
	File string `json:"file"`
	// Image is the Docker image identifier associated with the new server.
	Image string `json:"image"`
	// Project is the upstream repository URL that hosts the server code.
	Project string `json:"project"`
	// Commit is the pinned commit SHA for the newly added server.
	Commit string `json:"commit"`
	// Directory specifies a subdirectory inside the upstream repository, when present.
	Directory string `json:"directory,omitempty"`
}

// auditTarget represents a server selected for a manual full audit.
type auditTarget struct {
	// Server is the registry entry name included in the manual audit.
	Server string `json:"server"`
	// Project is the upstream repository URL for the audited server.
	Project string `json:"project"`
	// Commit is the pinned commit SHA to audit.
	Commit string `json:"commit"`
	// Directory is the subdirectory within the upstream repo to inspect, when applicable.
	Directory string `json:"directory,omitempty"`
}
