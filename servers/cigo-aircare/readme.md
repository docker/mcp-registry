# Cigo Aircare and Carsafe

Docs: https://cigo-mcp-system-wzoy.onrender.com/docs

Connect over Streamable HTTP at https://cigo-mcp-system-wzoy.onrender.com/mcp/. No account, API key, or local server installation is required. Tools are discovered dynamically.

- `get_hvac_filter_info`: FFU/BFU dimensions, published performance, pressure-loss information, and airflow/power comparisons. Start with `filter_type=all`, `FFU`, or `BFU`.
- `search_carsafe`: vehicle-and-year cabin filter matching and purchase links.
- `check_environment`: regional weather and air-quality guidance.
- `set_crm_lifecycle`: stores a hashed user identifier, vehicle/driving information, and a suggested replacement date. It does not itself send email or SMS.
- `submit_b2b_inquiry`: stores company/contact information and inquiry details for follow-up.

Ask the user for confirmation before the two write tools and explain what information will be sent. Published measurements have specific test conditions; power comparisons do not guarantee electricity-bill savings.

Support and security contact: cigo240920@gmail.com.
