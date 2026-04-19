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

package gitutil

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sync"
)

var (
	resolvedGit     string
	resolvedGitOnce sync.Once
	resolvedGitErr  error
)

// Executable returns the path to the git binary. On Windows it probes common
// Git for Windows install locations and honours the CLAUDE_CODE_GIT_BASH_PATH
// environment variable when git is not directly available in PATH.
// The result is cached after the first successful lookup.
func Executable() (string, error) {
	resolvedGitOnce.Do(func() {
		resolvedGit, resolvedGitErr = findGit()
	})
	return resolvedGit, resolvedGitErr
}

func findGit() (string, error) {
	if p, err := exec.LookPath("git"); err == nil {
		return p, nil
	}

	if runtime.GOOS != "windows" {
		return "", fmt.Errorf("git executable not found in PATH")
	}

	// On Windows, git may be installed but not on PATH. Try common locations.
	var candidates []string

	// Derive git.exe from CLAUDE_CODE_GIT_BASH_PATH when available.
	// The variable typically points to bash.exe inside a Git for Windows
	// installation (e.g. C:\Program Files\Git\bin\bash.exe). Walk up to
	// the installation root and check for cmd\git.exe.
	if bashPath := os.Getenv("CLAUDE_CODE_GIT_BASH_PATH"); bashPath != "" {
		dir := filepath.Dir(bashPath)
		for dir != filepath.Dir(dir) {
			candidates = append(candidates, filepath.Join(dir, "cmd", "git.exe"))
			dir = filepath.Dir(dir)
		}
	}

	for _, env := range []string{"ProgramFiles", "ProgramFiles(x86)", "LocalAppData"} {
		if base := os.Getenv(env); base != "" {
			candidates = append(candidates, filepath.Join(base, "Git", "cmd", "git.exe"))
		}
	}

	for _, c := range candidates {
		if _, err := os.Stat(c); err == nil {
			return c, nil
		}
	}

	return "", fmt.Errorf(
		"git is required but was not found in PATH or common install locations; " +
			"install Git for Windows (https://git-scm.com/download/win) or set " +
			"CLAUDE_CODE_GIT_BASH_PATH to the bash.exe in your Git for Windows installation",
	)
}
