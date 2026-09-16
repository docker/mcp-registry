# Google Sheets MCP Server

A fully-featured MCP server that lets AI agents interact with Google Sheets via the Google Sheets API v4. Supports **three authentication methods** — auto-detected based on what's available.

## Authentication

The server detects which method to use in this priority order:

| Priority | Method | How | Best for |
|----------|--------|-----|----------|
| 1st | `GOOGLE_SERVICE_ACCOUNT_JSON` env var | Docker Desktop config UI | ✅ Easiest — Docker users |
| 2nd | `service_account.json` file | Volume mount | ✅ Teams, automation |
| 3rd | `credentials.json` file | Volume mount + browser login | ✅ Personal, local dev |

---

### ⭐ Option A — Docker Desktop Config UI (Easiest)

If you're using **Docker Desktop with the MCP Toolkit**, paste your Service Account JSON directly in the Configuration tab — no file mounting needed.

**Step 1 — Create a Google Cloud Project**
1. Go to [console.cloud.google.com](https://console.cloud.google.com/)
2. Click **Select a project → New Project**, give it a name, click **Create**

**Step 2 — Enable the Google Sheets API**
1. Go to **APIs & Services → Library**
2. Search for **Google Sheets API** and click **Enable**

**Step 3 — Create a Service Account**
1. Go to **IAM & Admin → Service Accounts**
2. Click **Create Service Account**, give it a name, click **Create and Continue**, then **Done**

**Step 4 — Download the JSON Key**
1. Click on your new service account → **Keys** tab
2. Click **Add Key → Create new key → JSON**
3. Open the downloaded file and copy the entire contents

**Step 5 — Share your Sheet with the Service Account**
> ⚠️ The service account can only access sheets explicitly shared with it.
1. In the JSON file, find the `client_email` value (e.g. `sheets-mcp@my-project.iam.gserviceaccount.com`)
2. Open your Google Sheet → click **Share**
3. Paste the `client_email` and give it **Editor** access

**Step 6 — Configure in Docker Desktop**
1. Open **Docker Desktop → MCP Toolkit → Catalog → Google Sheets**
2. Go to the **Configuration** tab
3. Paste the full JSON contents into the **Google Service Account JSON** field
4. Click **Save** — done! 🎉

---

### ✅ Option B — Service Account File (Docker / Teams)

Same as Option A but using a mounted file instead of the config UI.

**Steps 1–5** — Same as Option A above.

**Step 6 — Run**
```bash
mkdir -p data
cp /path/to/service_account.json data/

docker run -i --rm \
  -v $(pwd)/data:/data \
  mcp/google-sheets
```

No browser login needed! 🎉

---

### ✅ Option C — OAuth 2.0 (Personal / Local Use)

Best for personal Google accounts and local development. Requires a one-time browser login.

**Step 1 — Create a Google Cloud Project**
1. Go to [console.cloud.google.com](https://console.cloud.google.com/)
2. Click **Select a project → New Project**, give it a name, click **Create**

**Step 2 — Enable the Google Sheets API**
1. Go to **APIs & Services → Library**
2. Search for **Google Sheets API** and click **Enable**

**Step 3 — Configure OAuth Consent Screen**
1. Go to **APIs & Services → OAuth consent screen**
2. Choose **External**, click **Create**
3. Fill in App name and email, click **Save and Continue** through all steps

**Step 4 — Create OAuth 2.0 Credentials**
1. Go to **APIs & Services → Credentials**
2. Click **Create Credentials → OAuth 2.0 Client ID**
3. Choose **Desktop app**, give it a name, click **Create**
4. Download the JSON file and rename it to `credentials.json`

**Step 5 — Run**
```bash
mkdir -p data
cp /path/to/credentials.json data/

docker run -i --rm \
  -v $(pwd)/data:/data \
  mcp/google-sheets
```

A browser window will open for Google login on first run. A `token.json` is saved in `data/` for future runs.

---

## Tools (28 total)

### 📋 Values — Reading & Writing
| Tool | Description |
|------|-------------|
| `read_sheet` | Read data from a single range |
| `write_sheet` | Write or update cell values |
| `append_rows` | Append new rows to a sheet |
| `batch_get_values` | Read multiple ranges in a single API call |
| `batch_update_values` | Write to multiple ranges in a single API call |
| `batch_clear_values` | Clear multiple ranges in a single API call |
| `clear_range` | Clear all values in a single range |

### 🗂️ Spreadsheet & Sheet Management
| Tool | Description |
|------|-------------|
| `create_spreadsheet` | Create a new Google Spreadsheet |
| `get_spreadsheet_info` | Get title, sheet names, row/column counts |
| `list_sheets` | List all sheet tabs |
| `add_sheet` | Add a new sheet tab |
| `delete_sheet` | Delete a sheet tab |
| `rename_sheet` | Rename a sheet tab |
| `duplicate_sheet` | Duplicate a sheet tab within the same spreadsheet |
| `copy_sheet_to` | Copy a sheet tab to a different spreadsheet |

### 🔢 Rows & Columns
| Tool | Description |
|------|-------------|
| `insert_rows_columns` | Insert empty rows or columns at a specific position |
| `delete_rows` | Delete rows by index |
| `delete_columns` | Delete columns by index |
| `move_dimension` | Move rows or columns to a new position |
| `insert_range` | Insert a blank range and shift cells down or right |
| `delete_range` | Delete a range and shift cells up or left |

### 🧹 Data Cleaning
| Tool | Description |
|------|-------------|
| `delete_duplicates` | Remove duplicate rows from a range |
| `trim_whitespace` | Trim extra spaces from cells in a range |
| `find_and_replace` | Search and replace text across sheets |
| `sort_range` | Sort a range by a specific column (asc or desc) |

### 🎨 Formatting
| Tool | Description |
|------|-------------|
| `format_cells` | Apply bold, font size, background and text colors |
| `freeze_rows_columns` | Freeze header rows and/or columns |
| `add_conditional_formatting` | Highlight cells based on rules (e.g. red if negative) |

## Source

[github.com/adalpan/google-sheets-mcp](https://github.com/adalpan/google-sheets-mcp)

## License

MIT — Copyright (c) 2026 Adalberto Perez Salas (pan)
