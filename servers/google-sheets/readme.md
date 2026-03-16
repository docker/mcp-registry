# Google Sheets MCP Server

A fully-featured MCP server that lets AI agents interact with Google Sheets via the Google Sheets API v4 and OAuth 2.0. Covers reading, writing, formatting, data cleaning, sheet management, and more.

## Authentication

This server uses **OAuth 2.0**. On first run, a browser window opens for you to authorize your Google account. A `token.json` is saved in the `/data` volume for future runs — no re-login needed.

You must provide a `credentials.json` file (OAuth 2.0 Desktop App type) from [Google Cloud Console](https://console.cloud.google.com/).

## Configuration

Mount a local folder containing your `credentials.json`:

```bash
docker run -i --rm \
  -v /path/to/data:/data \
  mcp/google-sheets
```

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
