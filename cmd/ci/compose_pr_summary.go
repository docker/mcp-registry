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
	"errors"
	"flag"
	"os"
	"strings"
)

// runComposePRSummary merges per-category summaries into a single Markdown
// document. It requires --pins-summary, --new-summary, and --output flags and
// tolerates missing summary files by emitting nothing.
func runComposePRSummary(args []string) error {
	flags := flag.NewFlagSet("compose-pr-summary", flag.ContinueOnError)
	pinsSummary := flags.String("pins-summary", "", "summary file for updated pins")
	newSummary := flags.String("new-summary", "", "summary file for new servers")
	output := flags.String("output", "", "path to write merged summary")
	if err := flags.Parse(args); err != nil {
		return err
	}

	if *output == "" {
		return errors.New("output is required")
	}

	var sections []string

	if *pinsSummary != "" {
		if content, err := os.ReadFile(*pinsSummary); err == nil {
			if len(strings.TrimSpace(string(content))) > 0 {
				sections = append(sections, string(content))
			}
		}
	}

	if *newSummary != "" {
		if content, err := os.ReadFile(*newSummary); err == nil {
			if len(strings.TrimSpace(string(content))) > 0 {
				sections = append(sections, string(content))
			}
		}
	}

	if len(sections) == 0 {
		removeIfPresent(*output)
		return nil
	}

	builder := strings.Builder{}
	builder.WriteString("# Security Review Targets\n\n")
	for _, section := range sections {
		builder.WriteString(section)
		if !strings.HasSuffix(section, "\n") {
			builder.WriteRune('\n')
		}
		builder.WriteRune('\n')
	}

	return os.WriteFile(*output, []byte(builder.String()), 0o644)
}
