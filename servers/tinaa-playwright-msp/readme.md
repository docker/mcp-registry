# TINAA Playwright MCP

Browser automation and testing MCP server powered by Playwright.

## Overview

TINAA (Testing Intelligence Network Automation Assistant) Playwright MCP is a comprehensive browser automation server that enables AI assistants to control web browsers, run end-to-end tests, perform accessibility audits, conduct security checks, and automate web interactions. Built on Microsoft's Playwright framework, it provides 15 powerful tools for web testing and automation.

This MCP server is ideal for developers and QA engineers who want to leverage AI assistants for automated browser testing, web scraping, accessibility validation, and responsive design verification. It supports exploratory testing with intelligent heuristics, automated form filling, screenshot capture, and comprehensive test reporting.

Key capabilities include:
- Full browser automation (navigate, click, type, screenshot)
- Form detection and automated filling
- Accessibility testing (WCAG compliance checks)
- Responsive design testing across breakpoints
- Security vulnerability scanning
- Exploratory testing with guided heuristics
- Test report generation
- Playwright LSP server integration for IDE support

## Tools

### `start_lsp_server`
Start TINAA's Playwright Language Server Protocol (LSP) server for IDE integration. Provides intelligent code completion, diagnostics, and hover information for Playwright test code.

**Parameters:**
- `tcp` (boolean, optional): Whether to use TCP mode instead of stdio mode. TCP mode is recommended for MCP integration. Default: false
- `port` (integer, optional): Port number for TCP mode server. Default: 8765

### `check_browser_connectivity`
Check TINAA's browser automation capabilities by visiting a URL and capturing basic information. Verifies that Playwright is working correctly.

**Parameters:**
- `url` (string, optional): The URL to visit for testing browser connectivity. Default: "https://example.com"

### `analyze_script`
Analyze a Playwright test script for common issues, errors, and improvement opportunities. Identifies missing await statements, syntax issues, and potential improvements.

**Parameters:**
- `script_path` (string, required): Absolute path to the Playwright test script file to analyze

### `navigate_to_url`
Navigate to a specific URL in the browser. Opens the specified URL and waits for the page to load.

**Parameters:**
- `url` (string, required): The URL to navigate to

### `take_page_screenshot`
Take a screenshot of the current page. Captures either the visible viewport or the entire scrollable page.

**Parameters:**
- `full_page` (boolean, optional): Whether to take a full page screenshot. Default: false

### `take_element_screenshot`
Take a screenshot of a specific element on the page. Captures only the specified element.

**Parameters:**
- `selector` (string, required): CSS selector for the element to screenshot

### `fill_login_form`
Fill and submit a login form. Automatically fills username and password fields, then clicks the submit button.

**Parameters:**
- `username_selector` (string, required): CSS selector for the username field
- `password_selector` (string, required): CSS selector for the password field
- `submit_selector` (string, required): CSS selector for the submit button
- `username` (string, required): Username value to enter
- `password` (string, required): Password value to enter

### `detect_form_fields`
Detect form fields on the current page. Automatically discovers all input fields, textareas, and selects.

**Parameters:**
- `form_selector` (string, optional): CSS selector for a specific form to analyze. If not provided, analyzes all forms

### `fill_form_fields`
Fill form fields on the current page. Fills multiple form fields with specified values and optionally submits the form.

**Parameters:**
- `fields` (object, required): Dictionary mapping CSS selectors to values
- `submit_selector` (string, optional): CSS selector for submit button to click after filling fields

### `run_accessibility_test`
Run accessibility tests on the current page. Checks for WCAG compliance issues including missing alt text, improper heading structure, color contrast problems, and ARIA attribute issues.

**Parameters:** None

### `run_responsive_test`
Run responsive design tests on the current page. Tests the page at multiple viewport sizes (mobile, tablet, desktop) to identify layout issues and responsive design failures.

**Parameters:** None

### `run_security_test`
Run basic security tests on the current page. Checks for security issues including mixed content, exposed sensitive data, insecure forms, missing security headers, and potential XSS vulnerabilities.

**Parameters:** None

### `generate_test_report`
Generate a comprehensive test report. Creates a formatted report summarizing test results, issues found, recommendations, and next steps.

**Parameters:**
- `test_type` (string, required): Type of test report to generate (e.g., "exploratory", "accessibility", "responsive", "security")
- `url` (string, required): URL that was tested

### `prompt_for_credentials`
Prompt the user for credentials. Generates a credential request prompt for sites requiring authentication.

**Parameters:**
- `site` (string, required): Site requiring authentication
- `username_field` (string, optional): Selector for username field
- `password_field` (string, optional): Selector for password field

### `run_exploratory_test`
Run an exploratory test on a website. Navigates to the URL, captures initial screenshots, and provides testing strategies and heuristics for guided exploration.

**Parameters:**
- `url` (string, required): URL to test
- `focus_area` (string, optional): Area to focus exploration on (e.g., "general", "forms", "navigation", "performance"). Default: "general"

## Configuration

This server uses stdio transport and does not require environment variables for basic operation. Playwright browsers are installed automatically on first run.

## Usage Examples

### Basic Browser Automation
```
Navigate to https://example.com and take a full page screenshot
```

### Form Testing
```
Go to the login page, detect all form fields, then fill the login form with test credentials
```

### Accessibility Audit
```
Navigate to https://myapp.com and run accessibility tests to check for WCAG compliance issues
```

### Responsive Design Testing
```
Test the homepage at mobile, tablet, and desktop breakpoints and report any layout issues
```

### Security Scanning
```
Run security tests on the current page to check for vulnerabilities and missing security headers
```

### Exploratory Testing
```
Run an exploratory test on https://example.com focusing on forms and navigation
```

## Links

- [Source Repository](https://github.com/aj-geddes/tinaa-playwright-msp)
- [Playwright Documentation](https://playwright.dev)
- [Issue Tracker](https://github.com/aj-geddes/tinaa-playwright-msp/issues)
