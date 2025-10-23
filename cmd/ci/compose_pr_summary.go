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
