# TweekIT Universal Media Converter MCP Server

## Overview

TweekIT MCP Server **eliminates agentic workflow blockers** by providing on-demand file conversion and media optimization for any format. It converts **400+ file formats** to AI-compatible outputs (PDF, PNG, JPG, WebP) and optimizes media content for websites, applications, and AI workflows—enabling agents to work without interruption.

**Key Capabilities:**
- **Removes workflow blockers:** Converts any file format agents encounter, preventing "unsupported file type" failures
- **On-demand optimization:** Resize, crop, format conversion for websites and media applications in real-time
- **400+ supported formats:** DOC, XLS, PSD, DWG, TIFF, CAD files, legacy Office docs, Adobe files, and more
- **Enterprise-grade engine:** 20+ years in production, trusted by Fortune 500 companies
- **Fast, stateless processing:** <2 second average conversion time

**Perfect for agentic workflows that need to:**
- Process user-uploaded files in any format without failing
- Optimize images and media for web delivery on-demand
- Convert legacy documents (DOC, XLS, PPT) to modern formats
- Handle design files (PSD, AI, INDD) in content pipelines
- Extract and normalize content from CAD files (DWG, DXF)
- Transform any media type for downstream AI processing

## Documentation

**Full Documentation:** https://github.com/equilibrium-team/tweekit-mcp/blob/main/README.md

**Website:** https://www.tweekit.io

**Live Demo:** https://www.tweekit.io/demo

**Use Cases:** https://www.tweekit.io/use-case

## Getting Started

1. **Sign up for TweekIT** at https://www.tweekit.io
   - Free tier includes 10,000 conversions
   - Get your API Key and API Secret from the account portal

2. **Configure credentials** (via environment variables or tool parameters):
   ```bash
   export TWEEKIT_API_KEY="your-key-here"
   export TWEEKIT_API_SECRET="your-secret-here"
   ```

3. **Use in your AI agent:**
   ```python
   # Example: Convert a legacy DOC file to PDF (removes workflow blocker)
   result = await mcp.call_tool('convert', {
       'inext': 'doc',
       'outfmt': 'pdf',
       'blob': base64_encoded_file,
       'noRasterize': True  # Preserve text
   })

   # Example: Optimize image for web on-demand
   result = await mcp.call_tool('convert', {
       'inext': 'tiff',
       'outfmt': 'webp',
       'blob': base64_encoded_image,
       'width': 800,  # Optimize for web
       'height': 600
   })

   # Example: Convert a file from URL (no manual download)
   result = await mcp.call_tool('convert_url', {
       'url': 'https://example.com/design.psd',
       'outfmt': 'png',
       'width': 1200  # Optimize dimensions
   })

   # Example: Check supported formats (prevent failures)
   formats = await mcp.call_tool('doctype', {
       'extension': '*'  # List all formats
   })
   ```

## Supported Formats

**Office Documents:** DOC, DOCX, DOT, DOCM, XLS, XLSX, XLSM, PPT, PPTX, PPTM, ODT, ODS, ODP, RTF

**Adobe Files:** PSD, AI, INDD, EPS, PDF

**CAD/Design:** DWG, DXF, SKP

**Images:** TIFF, PNG, JPEG, BMP, GIF, WebP, RAW formats, ICO, SVG

**And 350+ more formats**

Use the `doctype` tool to get the complete list.

## Architecture

```
AI Agent (E2B/Local)
    ↓
TweekIT MCP Server (Cloud Run - https://mcp.tweekit.io/mcp)
    ↓
MediaRich Backend (ISO 27001/SAS 70 compliant data centers)
```

- **Stateless processing**: Files are not stored persistently
- **Secure**: SSL encryption in transit, automatic cleanup after processing
- **Fast**: <2 second average conversion time
- **Scalable**: Enterprise infrastructure handling millions of conversions
- **On-demand**: No pre-processing or batch jobs required

## Use Cases

### 1. Agentic Workflow Blocker Removal
**Problem:** AI agent receives user-uploaded DOC file → Agent fails with "unsupported format"
**Solution:** Agent calls TweekIT convert → DOC becomes PDF → Workflow continues

### 2. On-Demand Website Optimization
**Problem:** CMS receives high-res TIFF images → Too large for web
**Solution:** Agent calls TweekIT convert with width/height → Optimized WebP delivered

### 3. Legacy Document Modernization
**Problem:** Scanning old XLS files for data extraction
**Solution:** Agent converts XLS → CSV → Processes with standard tools

### 4. Design Asset Pipeline
**Problem:** Marketing uploads PSD mockups → Need PNGs for web
**Solution:** Agent converts PSD → PNG with transparency → Ready for website

## Rate Limits & Pricing

- **Free Tier:** 10,000 conversions/month (no credit card required)
- **Paid Plans:** Available for higher volume
- **Details:** https://www.tweekit.io/pricing

## Support

- **Email:** support@tweekit.io
- **GitHub Issues:** https://github.com/equilibrium-team/tweekit-mcp/issues
- **Documentation:** https://github.com/equilibrium-team/tweekit-mcp

## License

MIT License - see [LICENSE](https://github.com/equilibrium-team/tweekit-mcp/blob/main/LICENSE)

## About Equilibrium

TweekIT is built on Equilibrium's MediaRich Server technology, trusted by Fortune 500 companies and major media portals since 2000. The conversion engine draws on the pedigree of DeBabelizer, the industry-standard batch processing tool. Over two decades of production use ensures reliability for enterprise-grade agentic workflows.
