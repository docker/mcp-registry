# OCR MCP Server

Local OCR (Optical Character Recognition) server for extracting text from PDFs, Word documents, and images.

## Features

- Extract text from PDF files (both text-based and scanned)
- Extract text from Word documents (.docx)
- Extract text from images (PNG, JPG, etc.)
- Uses Tesseract OCR for image-based text extraction
- Runs completely locally - no external API calls

## Tools

### extract_text

Extract text from PDF, Word documents, and images using OCR.

**Arguments:**
- `file_path` (string, required): Path to the file to extract text from
- `file_type` (string, required): Type of file - one of: `pdf`, `docx`, or `image`

## Usage

The server processes documents locally using Tesseract OCR and various document processing libraries.

## Source

GitHub: [https://github.com/Tanuj-kaswa23/ocr-mcp](https://github.com/Tanuj-kaswa23/ocr-mcp)

