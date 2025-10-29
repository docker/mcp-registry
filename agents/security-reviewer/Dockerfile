# syntax=docker/dockerfile:1.7

# Use a golang image to build the reviewer and proxy binaries (with cgo disabled).
FROM golang:1.25-trixie AS gobuilder
ENV CGO_ENABLED=0

# Create and set a working directory for the code.
RUN mkdir -p /security-reviewer
WORKDIR /security-reviewer

# Copy dependency specifications and ensure they're downloaded.
COPY go.mod go.sum ./
RUN go mod download

# Copy reviewer and proxy sources, then build both binaries.
COPY reviewer ./reviewer
COPY proxy ./proxy
RUN mkdir -p /security-reviewer/build
RUN go build -o /security-reviewer/build/reviewer ./reviewer
RUN go build -o /security-reviewer/build/proxy ./proxy


# Use a debian image to host the agent.
FROM debian:trixie-slim AS reviewer-image

# Set up configuration.
ARG CLAUDE_CODE_VERSION=latest
ARG CODEX_CLI_VERSION=latest
ARG AGENT_UID=1000
ARG AGENT_GID=1000

# Install core packages, developer tooling, and runtime dependencies.
ENV DEBIAN_FRONTEND=noninteractive
RUN apt-get update \
    && apt-get install -y --no-install-recommends \
        ca-certificates \
        curl \
        git \
        gnupg \
        jq \
        nodejs \
        npm \
        ripgrep \
        unzip \
    && rm -rf /var/lib/apt/lists/*

# Install Claude Code and Codex CLIs for headless operation.
RUN npm install --global \
        @anthropic-ai/claude-code@${CLAUDE_CODE_VERSION} \
        @openai/codex@${CODEX_CLI_VERSION} \
    && npm cache clean --force

# Create a dedicated runtime user and working directories.
# Use build args to allow matching host UID/GID to avoid permission issues.
RUN groupadd --gid ${AGENT_GID} agent \
    && useradd --uid ${AGENT_UID} --gid ${AGENT_GID} --create-home --shell /bin/bash agent \
    && install -d -o agent -g agent /workspace \
    && install -d -o agent -g agent /workspace/input \
    && install -d -o agent -g agent /workspace/output

# Copy the reviewer entrypoint and assets.
COPY --from=gobuilder /security-reviewer/build/reviewer /opt/security-reviewer/reviewer
COPY --chown=agent:agent prompt-template.md /opt/security-reviewer/prompt-template.md
COPY --chown=agent:agent report-template.md /opt/security-reviewer/report-template.md

# Set the reviewer user and working directory.
USER agent
WORKDIR /workspace

# Set the reviewer entrypoint.
ENTRYPOINT ["/opt/security-reviewer/reviewer"]


# Use a debian image to host the proxy.
FROM debian:trixie-slim AS proxy-image

# Install dependencies, including those needed for healthchecks.
ENV DEBIAN_FRONTEND=noninteractive
RUN apt-get update \
    && apt-get install -y --no-install-recommends \
        ca-certificates \
        wget \
    && rm -rf /var/lib/apt/lists/*

# Copy the proxy entrypoint.
COPY --from=gobuilder /security-reviewer/build/proxy /opt/security-reviewer/proxy

# Expose the proxy's default port.
EXPOSE 4000

# Set the proxy entrypoint.
ENTRYPOINT ["/opt/security-reviewer/proxy"]
