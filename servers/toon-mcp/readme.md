# TOON - Token-Optimized Object Notation

MCP server for reducing token consumption by up to 60% when working with JSON data in AI contexts.

## Overview

TOON (Token-Optimized Object Notation) is a compact representation format for JSON data that dramatically reduces token consumption in AI-assisted workflows. By using intelligent key abbreviations, schema compression, and pattern detection, TOON minimizes the tokens needed to represent structured data while maintaining full reversibility.

Key benefits:
- **40-60% Token Reduction**: Significantly reduce API costs and context usage
- **Fully Reversible**: Convert back to original JSON with zero data loss
- **Pattern Detection**: Automatically identifies optimization opportunities
- **Batch Processing**: Efficiently handle arrays of similar objects
- **Zero Configuration**: Works out of the box with any JSON data

## Tools

### `convert_to_toon`
Convert JSON to TOON format for reduced token consumption.

**Parameters:**
- `json_data` (string, required): JSON data to convert
- `aggressive` (boolean, optional): Enable aggressive compression (default: false)

**Returns:** TOON formatted data with savings statistics

### `convert_to_json`
Convert TOON format back to standard JSON.

**Parameters:**
- `toon_data` (string, required): TOON data to convert

**Returns:** Original JSON data

### `analyze_patterns`
Detect compression patterns in JSON data.

**Parameters:**
- `json_data` (string, required): JSON data to analyze

**Returns:** Detected patterns and compression recommendations

### `get_compression_strategy`
Get optimal compression settings for specific data.

**Parameters:**
- `json_data` (string, required): JSON data to analyze

**Returns:** Recommended strategy and expected savings percentage

### `calculate_savings`
Estimate token savings without converting.

**Parameters:**
- `json_data` (string, required): JSON data to analyze

**Returns:** Estimated token reduction and savings percentage

### `batch_convert`
Convert multiple JSON objects efficiently.

**Parameters:**
- `json_array` (string, required): Array of JSON objects
- `aggressive` (boolean, optional): Enable aggressive compression

**Returns:** Batch conversion results with total savings

## How TOON Works

### Key Abbreviations
Common JSON keys are shortened:
- `id` → `i`
- `name` → `n`
- `type` → `t`
- `status` → `s`

### Schema Compression
Arrays of similar objects use schema-based encoding to eliminate redundancy.

### Example

**Original JSON (156 tokens):**
```json
{
  "id": 123,
  "name": "John Doe",
  "type": "user",
  "status": "active",
  "metadata": {
    "created": "2024-01-01",
    "updated": "2024-06-15"
  }
}
```

**TOON Format (62 tokens):**
```json
{"_toon":"1.0","d":{"i":123,"n":"John Doe","t":"user","s":"active","m":{"c":"2024-01-01","u":"2024-06-15"}}}
```

**Savings: 60%**

## Use Cases

- **API Response Processing**: Compress large API responses before analysis
- **Context Window Optimization**: Fit more data in limited context windows
- **Cost Reduction**: Lower token usage = lower API costs
- **Batch Data Analysis**: Process large datasets efficiently

## Links

- [Source Repository](https://github.com/aj-geddes/toon-context-mcp)
- [Issue Tracker](https://github.com/aj-geddes/toon-context-mcp/issues)
