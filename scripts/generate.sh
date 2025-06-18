#!/bin/bash

set -e

# Output file
OUTPUT_FILE="mcp-servers.yaml"

# Remove existing output file if it exists
rm -f "$OUTPUT_FILE"

# Counter for tracking progress
count=0

# Walk through all directories in ./servers
for server_dir in ./servers/*/; do
    # Extract the directory name
    dir_name=$(basename "$server_dir")
    
    # Check if server.yaml exists in this directory
    if [[ -f "$server_dir/server.yaml" ]]; then
        echo "Processing: $dir_name"
        
        if [[ $count -eq 0 ]]; then
            # First server: add dash and content with no leading newline
            sed 's/^/- /' "$server_dir/server.yaml" | sed '2,$s/^- /  /' >> "$OUTPUT_FILE"
        else
            # Subsequent servers: add newline, then dash and content
            echo "" >> "$OUTPUT_FILE"
            sed 's/^/- /' "$server_dir/server.yaml" | sed '2,$s/^- /  /' >> "$OUTPUT_FILE"
        fi
        
        ((count++))
    else
        echo "Warning: No server.yaml found in $dir_name"
    fi
done

echo ""
echo "✅ Generated $OUTPUT_FILE with $count servers"
echo "📁 Output location: $(pwd)/$OUTPUT_FILE"

