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
	Server    string `json:"server"`
	File      string `json:"file"`
	Image     string `json:"image"`
	Project   string `json:"project"`
	Directory string `json:"directory,omitempty"`
	OldCommit string `json:"old_commit"`
	NewCommit string `json:"new_commit"`
}

// newServerTarget captures metadata for a newly added local server.
type newServerTarget struct {
	Server    string `json:"server"`
	File      string `json:"file"`
	Image     string `json:"image"`
	Project   string `json:"project"`
	Commit    string `json:"commit"`
	Directory string `json:"directory,omitempty"`
}

// auditTarget represents a server selected for a manual full audit.
type auditTarget struct {
	Server    string `json:"server"`
	Project   string `json:"project"`
	Commit    string `json:"commit"`
	Directory string `json:"directory,omitempty"`
}

// serverDocument is the decoded structure of a server YAML definition.
type serverDocument struct {
	Type   string `yaml:"type"`
	Image  string `yaml:"image"`
	Source struct {
		Project   string `yaml:"project"`
		Commit    string `yaml:"commit"`
		Directory string `yaml:"directory"`
	} `yaml:"source"`
}
