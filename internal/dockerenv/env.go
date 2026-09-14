// Package dockerenv builds the environment used when shelling out to the
// Docker CLI.
//
// The environment is deliberately narrow rather than inherited wholesale, so a
// build cannot pick up unrelated variables from the caller. It still has to
// carry the variables the Docker CLI needs to find its own configuration,
// otherwise the CLI cannot locate its config directory or its plugins.
package dockerenv

import "os"

// passthrough lists the variables the Docker CLI needs to resolve its config
// directory, its CLI plugins, and which daemon to talk to.
//
//   - HOME / USERPROFILE: the CLI looks for ~/.docker/config.json and for CLI
//     plugins such as buildx under ~/.docker/cli-plugins. Without these the
//     plugin lookup fails and "docker buildx build" exits 125 with the top
//     level usage text, which is how this surfaces on Docker Desktop for macOS
//     and Linux.
//   - DOCKER_CONFIG: explicit override of that config directory.
//   - DOCKER_HOST, DOCKER_CONTEXT, DOCKER_TLS_VERIFY, DOCKER_CERT_PATH: which
//     daemon to use. Needed by anyone not on the default socket, for example
//     Colima, Rancher Desktop, or a remote daemon.
//   - ProgramW6432: Windows only, see issue #79.
var passthrough = []string{
	"HOME",
	"USERPROFILE",
	"DOCKER_CONFIG",
	"DOCKER_HOST",
	"DOCKER_CONTEXT",
	"DOCKER_TLS_VERIFY",
	"DOCKER_CERT_PATH",
	"ProgramW6432",
}

// Env returns the environment for a Docker CLI invocation: PATH, the
// passthrough variables above that are actually set, and anything extra the
// caller supplies (such as GIT_AUTH_TOKEN).
func Env(additionalEnv ...string) []string {
	env := []string{"PATH=" + os.Getenv("PATH")}

	for _, key := range passthrough {
		if value, ok := os.LookupEnv(key); ok {
			env = append(env, key+"="+value)
		}
	}

	return append(env, additionalEnv...)
}
