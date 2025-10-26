# syntax=docker/dockerfile:1.7

# Use a golang image to build the agent entrypoint (with cgo disabled).
FROM golang:1.25-trixie AS gobuilder
ENV CGO_ENABLED=0

# Copy dependency specifications and ensure they're downloaded.
COPY go.mod go.sum ./
RUN go mod download

# Copy entrypoint sources and build the entrypoint
COPY entrypoint ./entrypoint
RUN go build -o /build/security-reviewer ./entrypoint


# Use a debian image to host the agent.
FROM debian:trixie-slim

# Set up configuration.
ARG CLAUDE_CODE_VERSION=latest
ARG CODEX_CLI_VERSION=latest

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
RUN useradd --create-home --shell /bin/bash agent \
    && install -d -o agent -g agent /workspace \
    && install -d -o agent -g agent /input \
    && install -d -o agent -g agent /output

# Copy the entrypoint and report template.
COPY --from=gobuilder /build/security-reviewer /opt/security-reviewer/security-reviewer
COPY --chown=agent:agent report-template.md /opt/security-reviewer/report-template.md

# Set the entrypoint user and working directory.
USER agent
WORKDIR /workspace

# Set the entrypoint.
ENTRYPOINT ["/opt/security-reviewer/security-reviewer"]
