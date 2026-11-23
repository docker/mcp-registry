# Sitelemetry

Docs: https://sitelemetry.com/mcp-guide

Sitelemetry audits websites you own or are authorized to test across six areas: security, technical SEO, AI visibility (GEO/AEO), analytics and tracking integrations, accessibility with WCAG 2.2-focused checks, and performance with Core Web Vitals (LCP, INP, CLS) plus FCP and TTFB from Google PageSpeed Insights. Each result lists findings with a severity, the evidence behind them, why they matter and a recommended fix, plus a score when the area could be measured. Reports are available in nine languages.

## Tools

| Tool | What it covers |
| --- | --- |
| `audit_security` | The security checks your plan includes: DNS, email, domain, TLS, HTTP headers, HTTPS and technology on public targets, plus protected modules on verified sites |
| `audit_seo` | Indexation, metadata, headings, canonicals, structured data, internal links, redirects and AI-crawler policy |
| `audit_ai_visibility` | Entity clarity, answerability, source signals, prompt coverage, llms.txt and robots.txt rules for AI crawlers |
| `audit_integrations` | Analytics, tag managers, marketing pixels, site-verification signals, consent systems and potential personal data in public data-layer output |
| `audit_accessibility` | WCAG 2.2-focused checks of alternative text, labels, headings, landmarks, language declarations, duplicate IDs and iframe titles |
| `audit_performance` | Core Web Vitals (LCP, INP, CLS) plus FCP and TTFB from Google PageSpeed Insights |
| `audit_full` | All six audit areas in one call, each within the connected plan or the workspace's audit credits for that area, with one blended score and a summary per area |

## Connect

1. Enable **Sitelemetry** in the MCP Toolkit catalog, or run `docker mcp server enable sitelemetry`.
2. Authorize it from the server's **OAuth** setting, or run `docker mcp oauth authorize sitelemetry`. Your browser opens: sign in to Sitelemetry or create a free account, then approve access.
3. Ask your connected client for an audit, for example: "Run a Sitelemetry security audit of example.com and list the findings by severity, each with its fix." Replace example.com with a site you own or are authorized to audit.

Long audits run as one background job, and the client checks that same job until the result is ready.

## Plans and verification

- Free includes a monthly allowance of security checks: 10 of the 27 security checks, with 7 public checks that run without verification and 3 more after the site is verified. Every Free workspace also starts with 1 audit credit in each of the six audit areas, and one-time credit packs add more credits with no subscription. A completed audit uses 1 credit from its own area; security credits apply only to verified sites. Paid plans add monitoring, regular audits and report exports: https://sitelemetry.com/pricing
- Protected security modules need a site verified in your Sitelemetry workspace, and so does a six-area full audit on a paid plan. Verify a site with a DNS TXT record, an HTTP file or Google Search Console.

## Links

- Website: https://sitelemetry.com
- Privacy Policy: https://sitelemetry.com/privacy
- Terms of Service: https://sitelemetry.com/terms
- Security contact: https://sitelemetry.com/.well-known/security.txt
- Support: support@sitelemetry.com
