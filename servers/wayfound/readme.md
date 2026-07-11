# Wayfound MCP Server

[See description](https://www.wayfound.ai/pages/wayfound-mcp)

## What is the Wayfound MCP Remote Server?

The Wayfound MCP remote server is a specialized Model Context Protocol server that provides AI agents with access to organizational data and tools. It enables agents to:

- List and interact with agents in your organization
- Get details about agents including their role, goal, guidelines, etc.
- Get performance analysis of agents including guideline violations, knowledge gaps, user ratings, sentiment and more.
- Get agent improvement suggestions based on performance analysis and feedback.

MCP (Model Context Protocol) is an open standard that allows AI applications to connect to external data sources and tools in a secure, standardized way. Think of it as a universal connector that lets your AI agents interact with real-world services and data.

## Available MCP Tools

The Wayfound MCP server provides the following tools for interacting with your organization's agents:

### `list_agents`
**Description**: Get the list of all Agents in your Wayfound organization.

**Usage**: Ask questions like "What agents are in my organization?" or "List all available agents"

**Returns**: A comprehensive list of all agents configured in your Wayfound organization, including their names and basic information.

### `get_agent_details`
**Description**: Get the details of a specific Agent in your Wayfound organization.

**Usage**: Ask questions like "Tell me about the Customer Support agent" or "What are the details for agent X?"

**Returns**: Detailed information about a specific agent including:
- Agent role and purpose
- Goals and objectives
- Guidelines and constraints
- Configuration settings
- Other relevant metadata

### `get_manager_analysis_for_agent`
**Description**: Get Wayfound Manager analysis for a specific Agent. This includes top topics, potential issues, tool call data, knowledge gaps, user ratings, sentiment, and guideline issues.

**Usage**: Ask questions like "What's the performance analysis for my Sales agent?" or "Show me the manager analysis for agent X"

**Returns**: Comprehensive performance analysis including:
- **Top Topics**: Most frequently discussed subjects
- **Potential Issues**: Identified problems or concerns
- **Tool Call Data**: Usage statistics and patterns
- **Knowledge Gaps**: Areas where the agent lacks information
- **User Ratings**: Customer satisfaction scores
- **Sentiment Analysis**: Overall user sentiment trends
- **Guideline Issues**: Violations or deviations from established guidelines

### `get_improvement_suggestions_for_agent`
**Description**: Get improvement suggestions for a specific Agent. This includes suggested system prompt updates and additional knowledge needed.

**Usage**: Ask questions like "How can I improve my Customer Service agent?" or "What suggestions do you have for agent X?"

**Returns**: Actionable improvement recommendations including:
- **System Prompt Updates**: Suggested modifications to the agent's instructions
- **Additional Knowledge**: Recommended knowledge base additions
- **Training Suggestions**: Areas for further development
- **Best Practice Recommendations**: Industry-standard improvements

## Example Queries

Here are some example questions you can ask your agent:

```python
# Basic agent information
"What are all the agents in my organization?"
"Give me details about the Customer Support agent"

# Performance analysis
"Show me the manager analysis for my Sales agent"
"What are the top topics and issues for the Marketing agent?"

# Improvement recommendations
"How can I improve my Customer Service agent's performance?"
"What knowledge gaps exist for the Technical Support agent?"

# Combined queries
"List all agents and show me which ones have the most issues"
"Compare the performance of my Sales and Marketing agents"
```


## Project Structure

```
wayfound-mcp-example/
├── main.py          # Main example script
├── .env             # Environment variables (create this file based on .env.example)
├── README.md        # This file
└── requirements.txt # Python dependencies
```

## Setup Instructions

### 1. Clone and Navigate

```bash
git clone https://github.com/Wayfound-AI/wayfound-mcp-example
cd wayfound-mcp-example
```

### 2. Install Dependencies

Make sure you have Python 3.10+ installed, then install the required packages:

```bash
pip install -r requirements.txt
```

### 3. Configure Environment Variables

Create a `.env` file in the project root with the following configuration:

```env
# Wayfound MCP Server Configuration
WAYFOUND_MCP_API_KEY=your_mcp_api_key_here

# OpenAI Configuration (required for the Agents SDK)
OPENAI_API_KEY=your_openai_api_key_here
```

**Important**:
- Replace `your_api_key_here` with your actual Wayfound MCP API key
- Replace `your_openai_api_key_here` with your OpenAI API key
- Keep the `.env` file secure and never commit it to version control

## Running the Example

Run the main example script:

```bash
python main.py
```

The script will:

1. Connect to the Wayfound MCP remote server using SSE (Server-Sent Events)
2. Create an AI agent with access to MCP tools
3. Execute example queries to demonstrate the available functionality
4. Display the results and provide a trace URL for debugging

### Example Output

```
View trace: https://platform.openai.com/traces/trace?trace_id=<trace_id>

Running: What are all the agents in my organization?
[Agent response with list of organizational agents]
```

## Customizing the Example

You can modify `main.py` to:

- Change the agent instructions
- Add new queries or tool calls
- Adjust model settings
- Add error handling and logging

## Troubleshooting

### Debug Mode

Enable debug logging to see more detailed information:

```python
import logging
logging.basicConfig(level=logging.DEBUG)
```

## Learn More

- [OpenAI Agents SDK Documentation](https://openai.github.io/openai-agents-python/)
- [Model Context Protocol Specification](https://modelcontextprotocol.io/)
