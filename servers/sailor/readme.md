# Sailor - Mermaid Diagram Generator

MCP server for generating and rendering professional Mermaid diagrams from natural language descriptions.

## Overview

Sailor transforms how developers create technical diagrams. Instead of manually writing Mermaid syntax, simply describe what you want in plain English and get production-ready diagram code. The server also validates, renders, and optimizes existing diagrams.

Key features:
- **Natural Language Generation**: Describe your diagram in plain text and get valid Mermaid code
- **Multi-Format Rendering**: Export to PNG, SVG, or PDF with customizable themes
- **Code Analysis**: Generate diagrams directly from source code
- **Diagram Intelligence**: Analyze, improve, and optimize existing diagrams
- **Comprehensive Support**: Flowcharts, sequence, class, state, ER, Gantt, pie charts, and mindmaps

## Tools

### `request_mermaid_generation`
Generate a Mermaid diagram from a natural language description.

**Parameters:**
- `description` (string, required): Natural language description of the diagram
- `diagram_type` (string, optional): Type of diagram (flowchart, sequence, class, state, er, gantt, pie, mindmap, auto)
- `style_preferences` (object, optional): Style customization options

### `validate_and_render_mermaid`
Validate Mermaid code and render to an image format.

**Parameters:**
- `mermaid_code` (string, required): Mermaid diagram code
- `output_format` (string, optional): png, svg, or pdf (default: png)
- `theme` (string, optional): default, dark, forest, or neutral

### `get_mermaid_examples`
Get example diagrams for reference and learning.

**Parameters:**
- `diagram_type` (string, optional): Filter by diagram type
- `complexity` (string, optional): simple, medium, or complex
- `industry` (string, optional): Filter by domain/industry

### `analyze_diagram`
Extract structure, relationships, and meaning from a diagram.

**Parameters:**
- `mermaid_code` (string, required): Diagram code to analyze
- `analysis_depth` (string, optional): basic, detailed, or comprehensive

### `suggest_improvements`
Get recommendations for improving diagram quality.

**Parameters:**
- `mermaid_code` (string, required): Diagram code to analyze
- `focus_areas` (array, optional): Areas to focus on (readability, style, structure, accessibility)

### `convert_diagram_style`
Apply a new visual style or theme to a diagram.

**Parameters:**
- `mermaid_code` (string, required): Original diagram code
- `target_style` (string, required): Target style to apply

### `generate_from_code`
Generate diagrams from source code analysis.

**Parameters:**
- `source_code` (string, required): Source code to analyze
- `language` (string, optional): Programming language
- `diagram_type` (string, optional): class, sequence, or flowchart

### `generate_from_data`
Create diagrams from structured JSON or CSV data.

**Parameters:**
- `data` (string, required): Structured data
- `data_format` (string, optional): json or csv
- `visualization_type` (string, optional): Type of visualization

### `modify_diagram`
Apply targeted modifications to an existing diagram.

**Parameters:**
- `mermaid_code` (string, required): Original diagram code
- `modifications` (array, required): List of modifications

### `merge_diagrams`
Combine two diagrams into one cohesive diagram.

**Parameters:**
- `diagram1` (string, required): First diagram
- `diagram2` (string, required): Second diagram
- `merge_strategy` (string, optional): append, interleave, or smart

### `extract_subgraph`
Extract a portion of a diagram as a standalone subgraph.

**Parameters:**
- `mermaid_code` (string, required): Original diagram
- `node_ids` (array, required): Nodes to include
- `include_connected` (boolean, optional): Include connected nodes

### `optimize_layout`
Optimize diagram layout for better readability.

**Parameters:**
- `mermaid_code` (string, required): Diagram to optimize
- `optimization_goals` (array, optional): Goals like minimize_crossings, balance_layout, compact

## Usage Examples

### Generate a System Architecture Diagram
```
Create a flowchart showing a web application architecture with a load balancer,
two web servers, a database cluster, and a cache layer
```

### Analyze an Existing Diagram
```
Analyze this class diagram and suggest improvements for better readability
```

### Generate from Code
```
Generate a class diagram from this Python code showing the inheritance hierarchy
```

## Links

- [Source Repository](https://github.com/aj-geddes/sailor)
- [Issue Tracker](https://github.com/aj-geddes/sailor/issues)
