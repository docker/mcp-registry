# Verificate Gate

Documentation: https://verificate.ai/docs/gate/

Verificate Gate is a hosted review gate for AI-written code, plans and documents. Deterministic
checks run first and can veto: placeholder or mock implementations presented as finished, tests
that cannot fail, and unsupported completion claims. A model review then checks for calls to APIs
that do not exist and grades correctness, security, performance and maintainability and returns a fix plan.
If the review cannot run, the gate fails closed rather than approving. Submitted content is read,
never executed.

Tools: `validate_ai_output`, `validate_plan`, `analyze_code`, `generate_code`, `validate_artifact`.

Authentication: none required — 100 free validations with no signup. For continued use, create an
account at https://verificate.ai/auth/signup (30-day trial, no card) and send the token as
`Authorization: Bearer <token>`.

Source and client configs: https://github.com/VerificateAI/verificate-mcp-quickstart
