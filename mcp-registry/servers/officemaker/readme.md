Docs: https://officemaker.ai

## Overview

OfficeMaker is a remote MCP server that enables AI agents to create Microsoft Office documents — Word (.docx), Excel (.xlsx), and PowerPoint (.pptx) — from structured JSON specifications. Documents are generated server-side and returned as downloadable files or preview images.

## Authentication

An API key is required to use the authenticated endpoint (`https://docs.officemaker.ai/mcp`). Set your key via the `OFFICEMAKER_API_KEY` environment variable.

A free unauthenticated endpoint is also available at `https://free.officemaker.ai/mcp` with a subset of features (document creation, schema retrieval, and preview generation — no conversion or automations).

## Tools

Tools are discovered dynamically. Key capabilities include:

- **create_document_structured** — Create Word, Excel, or PowerPoint documents from a structured JSON payload
- **create_preview_structured** — Generate a preview image from a stored document
- **get_minimal_schema** — Retrieve a compact machine-readable schema for building document JSON
- **get_schema** — Retrieve the full or partial document schema
- **convert_document** — Convert an existing Office or PDF file to the OfficeMaker JSON format *(authenticated only)*
- **macros_list / macros_get / macros_execute** — List and run saved document automations *(authenticated only)*

## Usage

Before creating a document, call `get_minimal_schema` to retrieve the schema and `schemaAttestation` JWT, then pass both into `create_document_structured` along with your document JSON.
