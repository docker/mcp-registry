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
