Docs: https://api.purecipher.com/fixtures/omniseal-claude-guide.html

Omniseal exposes seal_files, retrieve_secret_file, verify_sealed_file, and
remove_hidden_file. No account, API key, or OAuth is required for this public
endpoint. The platform=claude query selects the compatible MCP file schema;
it does not send files to Anthropic.

Inputs are limited to 6 MB per file and 10 MB combined for sealing. Outputs use
binary MCP resources capped at 140,000 serialized characters (roughly 100 KB of
file bytes). Local paths and chat attachments are not automatically uploaded.
Sealing, retrieval, and removal require an Omniseal private seed, not an account
password. Original files are preserved.

Privacy: https://purecipher.com/privacy
Terms: https://purecipher.com/terms
Support: support@purecipher.com

For review, use the synthetic review-cover.png, review-secret.pdf, and
review-sealed.png fixtures under https://api.purecipher.com/fixtures/ with
public demo seed omniseal-review-2026. These fixtures contain no personal data.
