Gemini Prompt
YOU are an expert AI assistant creating an interactive demo of the Docker Model Context Protocol (MCP) Toolkit for a developer. Your goal is to help them understand the power of these tools and figure out which ones make the most sense to integrate into their current project.

Use the comprehensive quickstart guide below as your single source of truth.

Your Task:

Initiate the Conversation: Greet the user and explain that you're here to walk them through the Docker MCP ecosystem.

Present the Tools: Based on the guide, introduce the different categories of tools (e.g., Infrastructure, Observability, Data). Briefly explain what MCP is and why it's useful, drawing from Part I of the guide.

Engage the User: Ask the user about their current project, their tech stack, and their biggest challenges or repetitive tasks.

Recommend and Discuss: Based on their answers, recommend specific tools from the guide. For example, if they work with Kubernetes, dive deep into the mcp/kubernetes tool. If they're focused on data analysis, introduce them to mcp/vizro or mcp/sec-edgar. Use the examples and configuration details from the guide to make it concrete.

Facilitate Integration: Discuss how they could start integrating a recommended tool into their workflow. Talk about the setup process, security best practices (like handling secrets and filesystem access from Part III), and what a first practical use case might look like.

Maintain a Conversational Tone: Be helpful, interactive, and focus on solving the user's specific problems with the tools at your disposal.

KNOWLEDGE BASE: The Definitive Quickstart Guide to the Docker MCP Ecosystem
Part I: Foundations of the MCP Ecosystem
This initial section establishes the fundamental concepts, setting the stage for the detailed tool compendium. It explains the "why" behind the Model Context Protocol (MCP) before diving into the "what" and "how" of the Docker MCP Toolkit.

Understanding the Model Context Protocol (MCP)
The Model Context Protocol, or MCP, is an open standard designed to create a secure and standardized translation layer between AI agents (clients) and external tools (servers).[1, 2, 3] At its heart is a foundational client-server model. Clients, such as the Claude Desktop App, Cursor, or Docker's own AI agent, are applications that request resources or actions. MCP servers are specialized, often containerized, applications that are launched to perform the requested tasks, using any necessary tools, languages, or processes within a controlled environment.[4] This protocol is the critical component that allows an AI to evolve beyond simple text generation to performing tangible, real-world actions on behalf of a developer.[3]

Before the advent of MCP, integrating AI with developer tools was a fragmented and complex endeavor. Developers had to wrestle with disparate protocols, write significant amounts of boilerplate code, and manage various transport layers for each tool integration.[1] MCP was created to solve this critical fragmentation problem. It provides a unified interface for AI agents to discover and access external tools, a secure framework for managing credentials, and standard patterns for tool documentation and interaction.[1, 3] By establishing a common language between AI and tools, MCP dramatically lowers the barrier to building powerful, agentic applications.

The Docker MCP Toolkit: Your Secure Gateway to Tooling
The Docker MCP Toolkit is engineered to be the fastest and most secure path from tool discovery to local execution, effectively bridging the gap between the conceptual power of AI agents and their practical application in a developer's workflow.[4] It offers a suite of features designed to remove friction and build trust, including one-click setup directly from Docker Desktop, broad cross-LLM compatibility with popular clients like Claude and Cursor, and integrated tool discovery through the Docker MCP Catalog.[4, 5]

The toolkit's architecture revolves around a central Gateway MCP Server. When a developer enables individual tools from the catalog—for instance, mcp/kubernetes and mcp/grafana—these tools are not exposed individually. Instead, they are dynamically routed through this single gateway. This model vastly simplifies the client-side configuration; an AI agent only needs to connect to the one gateway endpoint to gain access to a multitude of powerful, ready-to-use tools.[5]

This streamlined usability is built on a foundation of security, which is a first principle of the Docker MCP offering. The security model is twofold:

Passive Security (At Rest): Trust begins before a tool is ever run. All official MCP server images published under the mcp/ namespace on Docker Hub are built and digitally signed by Docker. Each image also includes a comprehensive Software Bill of Materials (SBOM) for full transparency into its components.[4] Developers can independently verify any image's signature using cosign and Docker's public key, ensuring the tool's origin and integrity have not been compromised.[6, 7] This system is not merely about providing tools; it's about providing a trusted and verifiable supply chain for AI capabilities, positioning Docker Hub as a secure "App Store" for developers' AI agents.

Active Security (At Runtime): Once a tool is running, security is enforced through strict, non-negotiable constraints. Every MCP tool is executed in its own isolated container, which is restricted to a maximum of 1 CPU and 2 GB of memory to limit the impact of potential misuse.[4] Crucially, MCP servers have no access to the host filesystem by default. The user must explicitly and deliberately grant file system access by mounting specific directories, a critical safeguard against unauthorized data access. Furthermore, the toolkit actively intercepts requests to and from tools, blocking any that contain sensitive information such as secrets, providing an additional layer of protection at runtime.[4]

This robust, multi-layered security architecture directly addresses the primary concerns a developer would have before granting an AI agent the ability to execute code on their machine: "Is it safe?" and "Is it easy to set up?". By solving these core friction points, the Docker MCP Toolkit is strategically designed to solve the "last mile" problem of AI agent adoption, enabling widespread, secure experimentation and integration of agentic workflows into the daily development cycle.

Enabling and Navigating the MCP Catalog
Getting started with the MCP Toolkit is a straightforward process designed for rapid adoption.

Enable the Toolkit: The first step is to activate the feature within Docker Desktop. Navigate to Settings, select the Beta features tab, and check the box for Enable Docker MCP Toolkit. A quick restart via the Apply & restart button will complete the activation.[4]

Discover and Install Servers: With the toolkit enabled, a new MCP Toolkit item appears in the Docker Desktop sidebar. Selecting it reveals the Catalog tab, which is a curated collection of available MCP servers.[1, 4] Here, you can browse tools, view their descriptions, publishers, and the specific, callable sub-tools they provide. To install a server, you simply click the plus icon (+) next to it. For servers that require configuration, such as API keys or file paths, a Config tab allows for easy input of these parameters.[4]

Connect a Client: After installing the desired servers, navigate to the Clients tab. This interface lists supported AI clients like Claude Desktop and Cursor. Clicking Connect next to a client establishes the link between that application and the MCP Toolkit's gateway. If the client application is already running, a restart is required to make the newly enabled tools available to the AI agent.[4]

This streamlined process allows a developer to go from discovery to a fully functional, multi-tool AI assistant in a matter of minutes.

Part II: The Complete MCP Tool Compendium
The true power of the MCP ecosystem lies in the breadth and depth of the available tools. This section provides a comprehensive directory of the most prominent and useful MCP servers available through the Docker MCP Catalog. The tools are categorized by their primary function to help you quickly identify solutions relevant to your workflow.

The ecosystem is characterized by a healthy, competitive spirit, which is a hallmark of a successful open standard. For certain core tasks, such as Kubernetes management, you will find multiple implementations from different publishers.[8, 9, 10] This is not a sign of confusion but of innovation. It means different developers are tackling the same problem with unique approaches—for example, one server might be a simple wrapper around the kubectl command-line tool, while another might be a native Go application that interacts directly with the Kubernetes API, eliminating external dependencies.[11] This guide will highlight these distinctions, empowering you to make an informed choice based on your specific needs for performance, security, and features. Your decision shifts from "should I use a Kubernetes MCP?" to "which Kubernetes MCP is the best fit for my project?".

The following table provides a high-level overview of the key tools covered in this guide.

Tool Image

Category

Primary Function

Requires Secrets?

mcp/kubernetes

Infrastructure & DevOps

Manage Kubernetes clusters using natural language.

No (uses kubeconfig)

mcp/aws-terraform

Infrastructure & DevOps

Provision AWS resources with Terraform and Checkov.

Yes (AWS Credentials)

mcp/pulumi

Infrastructure & DevOps

Manage cloud infrastructure with the Pulumi IaC platform.

Yes (Pulumi Token)

mcp/grafana

Observability & Code Quality

Query datasources and manage Grafana dashboards.

Yes (Grafana API Key)

mcp/sonarqube

Observability & Code Quality

Analyze code quality and security with SonarQube.

Yes (SonarQube Token)

mcp/playwright

Web Automation & Data

Automate browser interactions for testing and scraping.

No

mcp/tavily

Web Automation & Data

Perform AI-native web searches and data extraction.

Yes (Tavily API Key)

mcp/stripe

Finance & Payments

Interact with the Stripe API for payments and billing.

Yes (Stripe API Key)

mcp/github

Productivity & Collaboration

Interact with the GitHub API for repository management.

Yes (GitHub PAT)

mcp/vizro

Data & Databases

Create data visualizations and dashboards with Vizro.

No

mcp/sec-edgar

Data & Databases

Retrieve financial data from the US SEC EDGAR database.

Yes (User Agent)

mcp/everything

Utilities & Miscellaneous

A reference server for testing and demonstrating MCP features.

No

Infrastructure & DevOps
This category includes tools essential for managing infrastructure, deployments, and cloud-native environments. They empower AI agents to become active participants in the entire software delivery lifecycle.

mcp/kubernetes: Conversational Cluster Management
Purpose & Use Case: The mcp/kubernetes server acts as an intelligent bridge between AI assistants and Kubernetes clusters. It is designed to translate natural language requests—like "show me the logs for the broken pod"—into the precise kubectl and helm operations required to fulfill them.[9, 10, 12] This tool directly addresses a significant pain point for developers and operators: the complexity of managing Kubernetes, which often involves juggling multiple terminal windows, memorizing intricate commands with numerous flags, and constantly referencing documentation.[3]

Key Features: This server provides a comprehensive suite of over 20 dedicated tools for Kubernetes management. The toolset covers the full spectrum of operations, including resource creation (create_deployment), inspection (describe, logs), execution (exec_in_pod), and even Helm chart management (install_helm_chart).[3, 8] As noted previously, several community versions exist with different underlying architectures, offering choices between simple wrappers and high-performance native binaries.[11]

Quickstart & Configuration (Docker Desktop):

Ensure that the Kubernetes engine is enabled within your Docker Desktop settings.[3]

From the MCP Toolkit > Catalog, locate the mcp/kubernetes server and click the plus icon to enable it.

The server is pre-configured to automatically use the kubeconfig file managed by Docker Desktop, providing instant, zero-configuration access to your local cluster.

Manual Configuration (e.g., Claude Desktop): To connect to a remote or non-Docker-Desktop-managed cluster, you must mount your kubeconfig file into the container.

{
  "mcpServers": {
    "kubernetes": {
      "command": "docker",
      "args": [
        "run", "-i", "--rm",
        "--mount", "type=bind,src=/path/to/your/.kube/config,dst=/home/mcp/.kube/config,readonly",
        "mcp/kubernetes:latest"
      ]
    }
  }
}

(Note: The src path must be the absolute path to your local kubeconfig file. Using readonly is a security best practice. [9, 13])

Usage Example (Natural Language Prompt):

"Describe the 'nginx-ingress' deployment in the 'networking' namespace and then show me its logs from the last 10 minutes."

Tool Reference Table: mcp/kubernetes

Tool Name

Description

Key Parameters

Read-Only?

kubectl_get

Get or list Kubernetes resources.

resource_type, name, namespace

Yes

kubectl_describe

Describe a specific resource in detail.

resource_type, name, namespace

Yes

kubectl_logs

Get logs from a resource like a pod or deployment.

resource_type, name, namespace

Yes

kubectl_apply

Apply a Kubernetes YAML manifest.

yaml_content

No

kubectl_delete

Delete a resource by type and name.

resource_type, name, namespace

No

exec_in_pod

Execute a command in a pod or container.

pod_name, command, namespace

No

explain_resource

Get documentation for a resource or field.

resource_name

Yes

install_helm_chart

Install a Helm chart into the cluster.

chart_name, release_name, namespace

No

mcp/aws-terraform: Secure and Compliant IaC for AWS
Purpose & Use Case: This server is designed for developers who want to provision and manage AWS infrastructure using Terraform through natural language, while ensuring adherence to organizational best practices and security standards.[14, 15] It integrates security scanning directly into the IaC workflow, making it an ideal tool for security-conscious development teams.

Key Features: As part of the official AWS Labs suite of MCP servers, this tool provides robust capabilities for Terraform workflow execution (plan, apply), integrated security compliance scanning with Checkov, and direct access to AWS and AWS Cloud Control (AWSCC) provider documentation.[15, 16]

Quickstart & Configuration (Docker Desktop):

In the MCP Toolkit > Catalog, find mcp/aws-terraform and enable it.

Configuration requires providing AWS credentials. The most secure method is to navigate to the Config tab for the server and mount your local ~/.aws directory into the container. This allows the server to use your existing AWS CLI profiles without exposing secret keys directly.

Manual Configuration (e.g., Cursor):

{
  "mcpServers": {
    "aws-terraform": {
      "command": "docker",
      "args": [
        "run", "-i", "--rm",
        "-v", "/Users/your-user/.aws:/root/.aws:ro",
        "-v", "/path/to/your/terraform/project:/app",
        "mcp/aws-terraform:latest"
      ]
    }
  }
}

(Note: Mounting the AWS credentials directory as read-only (:ro) is a recommended security measure.)

Usage Example (Natural Language Prompt):

"Using the Terraform project in the /app directory, create a plan to deploy a new t3.micro EC2 instance in us-east-1 with an Ubuntu 22.04 AMI. Ensure it has a security group that allows SSH access from my IP only. After planning, run a Checkov scan on the proposed changes and report any critical vulnerabilities."

mcp/pulumi: Universal Infrastructure as Code
Purpose & Use Case: The mcp/pulumi server provides a powerful MCP interface to the Pulumi CLI and Cloud API. It empowers AI agents to perform the full range of Pulumi operations—such as previewing infrastructure changes, deploying updates, and retrieving stack outputs—without requiring the Pulumi CLI to be installed in the client's environment.[17] A key use case is its ability to analyze application code in a directory and automatically generate the corresponding Pulumi infrastructure code for deployment to AWS.[17]

Key Features: This server interacts directly with the Pulumi Automation API. Its core tools include preview, up, and stack-output. It also features tools for resource discovery, allowing an agent to query the Pulumi Registry for available providers, modules, and resources (get-resource, list-resources).[17]

Quickstart & Configuration (Docker Desktop):

Enable mcp/pulumi from the MCP Toolkit > Catalog.

In the Config tab, you must mount your local Pulumi project directory to the container's working directory (e.g., /app/project).

You must also provide your PULUMI_ACCESS_TOKEN as an environment variable for authentication with the Pulumi service.

Manual Configuration (e.g., Claude Desktop):

{
  "mcpServers": {
    "pulumi": {
      "command": "docker",
      "args":
    }
  }
}

(As documented in [17])

Usage Example (Natural Language Prompt):

"Inside my project at /app/project, run pulumi preview for the 'staging' stack. If there are no errors, proceed to run pulumi up and then show me all the resulting stack outputs."

Observability & Code Quality
These tools are designed to give AI agents visibility into application performance, logs, and code health, enabling them to act as intelligent assistants for debugging and maintenance.

mcp/grafana: Interactive Dashboards and Querying
Purpose & Use Case: The mcp/grafana server provides an MCP interface for your Grafana instance, giving AI agents programmatic access to the entire Grafana ecosystem.[18] This enables powerful, natural language-driven interactions for querying monitoring data from datasources like Prometheus and Loki, managing dashboards, and even interacting with operational tools like Grafana OnCall and Sift for incident investigation.[18]

Key Features: The server boasts a rich and extensive toolset that covers Prometheus PromQL and Loki LogQL querying, dashboard creation and updates, Sift log and trace analysis, and management of Grafana OnCall schedules.[18, 19] A particularly thoughtful feature is its configurability; you can disable entire categories of tools (e.g., --disable-loki, --disable-oncall) when starting the server. This allows you to tailor the tools exposed to the LLM, which can improve its accuracy by reducing the size of the context window and preventing it from attempting to use tools you don't have configured.[18]

Quickstart & Configuration:

Prerequisites: You must have a running Grafana instance accessible from where you are running Docker.[20] You also need to create a Grafana Service Account with the appropriate permissions and generate a Service Account Token.[18]

Docker Desktop: Enable mcp/grafana from the catalog. In the server's Config tab, you must provide two environment variables: GRAFANA_URL (the full URL to your Grafana instance, e.g., http://localhost:3000) and GRAFANA_API_KEY (the service account token you generated).

Manual Configuration (Docker Run): The following command would be embedded within your AI client's JSON configuration file.

docker run --rm -i \
  -e GRAFANA_URL="http://your-grafana-instance:3000" \
  -e GRAFANA_API_KEY="<your_service_account_token>" \
  mcp/grafana

(This command runs the server interactively, passing the necessary credentials as environment variables.[18])

Usage Example (Natural Language Prompt):

"Query my Prometheus datasource. What is the 95th percentile of HTTP request latency for the 'frontend-api' service over the last hour? After that, find any elevated error patterns in the Loki logs for the same service during the same time period."

Tool Reference Table: mcp/grafana

Tool Name

Description

Key Parameters

Read-Only?

query_prometheus

Execute a PromQL query against a Prometheus datasource.

query, datasource_uid

Yes

query_loki

Execute a LogQL query against a Loki datasource.

query, datasource_uid

Yes

get_dashboard_by_uid

Retrieve a complete dashboard definition.

uid

Yes

update_dashboard

Update or create a dashboard from a JSON model.

dashboard_json

No

find_error_pattern_logs

Use Sift to find elevated error patterns in Loki logs.

labels, start, end

Yes

get_oncall_users

Get the list of users currently on-call for a schedule.

scheduleId

Yes

create_incident

Create a new Grafana incident.

title, severity

No

mcp/sonarqube: AI-Assisted Code Quality Analysis
Purpose & Use Case: This server integrates with your SonarQube instance (supporting SonarQube Cloud, Server, and Community editions) to grant AI assistants deep insights into code quality and security.[14, 21, 22] It allows developers to use natural language to query for code smells, find security vulnerabilities, check quality gate status, and even analyze code snippets on the fly, making code reviews and quality assurance more interactive and efficient.

Key Features: The server provides a powerful set of tools, including the ability to analyze a code snippet directly without a full project scan, list all projects, search for issues with an extensive array of filters (by severity, status, type, etc.), retrieve source code with issues highlighted, and programmatically change an issue's status (e.g., mark as "false-positive" or "reopen").[21, 23]

Quickstart & Configuration:

Prerequisites: You need access to a SonarQube instance and an authentication token with sufficient permissions.[22]

Docker Desktop: Enable mcp/sonarqube from the catalog. In the Config tab, you must provide your SONARQUBE_TOKEN. If you are using SonarQube Cloud, you must also provide your SONARQUBE_ORG. If you are using a self-hosted SonarQube Server, you must provide its SONARQUBE_URL.[21]

Manual Configuration (Docker for SonarQube Cloud):

{
  "mcpServers": {
    "sonarqube": {
      "command": "docker",
      "args":,
      "env": {
        "SONARQUBE_TOKEN": "<your_sonarqube_token>",
        "SONARQUBE_ORG": "<your_sonarcloud_organization>"
      }
    }
  }
}

(As documented in [21])

Usage Example (Natural Language Prompt):

"In the 'payment-gateway' project, list all open 'CRITICAL' and 'BLOCKER' vulnerabilities that were introduced in the last week. For the most critical one, show me the affected lines of code."

Tool Reference Table: mcp/sonarqube

Tool Name

Description

Key Parameters

Read-Only?

list_issues

Search for issues with extensive filtering.

projects, severities, types, statuses

Yes

analyze_code_snippet

Analyze a snippet of code for issues.

codeSnippet, language

Yes

get_raw_source

Get the raw source code of a file.

key, branch

Yes

change_sonar_issue_status

Change the status of an issue.

key, status (accept, reopen, etc.)

No

get_measures

Get quality measures/metrics for a component.

component, metricKeys

Yes

Web Automation & Data Retrieval
This category focuses on tools that allow AI agents to interact with the web, performing tasks from automated testing and scraping to intelligent, real-time search.

mcp/playwright: Advanced Browser Automation
Purpose & Use Case: The mcp/playwright server provides powerful browser automation capabilities, enabling LLMs to interact with web pages in a structured and reliable way.[24] Instead of relying on fragile, pixel-based vision models to interpret screenshots, this server uses Playwright's accessibility tree. This allows the AI to "see" a web page as a structured set of elements (buttons, inputs, links), leading to more deterministic and reliable interactions for tasks like end-to-end testing, web scraping, and form filling.[24]

Key Features: The server is fast and lightweight due to its use of the accessibility tree rather than screenshots.[24] It offers a comprehensive set of over 25 tools for browser control, including navigation (browser_navigate), interaction (browser_click, browser_type), file uploads (browser_file_upload), and capturing snapshots (browser_snapshot).[25, 26] It can operate in a standard "snapshot" mode or a "vision" mode that uses screenshots if needed.[24]

Quickstart & Configuration (Docker Desktop):

Enable mcp/playwright from the MCP Toolkit > Catalog.

No specific secrets are required for basic operation. Configuration options like viewport size can be set in the Config tab via command-line arguments.

Manual Configuration (e.g., Cursor):

{
  "mcpServers": {
    "playwright": {
      "command": "docker",
      "args": [
        "run", "-i", "--rm", "--init",
        "mcr.microsoft.com/playwright/mcp"
      ]
    }
  }
}

(Note: The official Docker image is from Microsoft's container registry (mcr.microsoft.com).[24] The --init flag is recommended to prevent zombie processes.[27])

Usage Example (Natural Language Prompt):

"Navigate to docs.docker.com. Take an accessibility snapshot of the page. Then, find the search bar, type 'MCP Toolkit', and press Enter. Wait for the results to load and take another snapshot."

Tool Reference Table: mcp/playwright

Tool Name

Description

Key Parameters

Read-Only?

browser_navigate

Navigate the browser to a specific URL.

url

No

browser_snapshot

Capture an accessibility snapshot of the current page.

None

Yes

browser_click

Click on an element on the page.

ref (element reference), element (description)

No

browser_type

Type text into an editable element.

ref, element, text, submit (optional)

No

browser_file_upload

Upload one or more files to an input element.

ref, element, file_paths

No

browser_scroll

Scroll the page up, down, or to an element.

direction, ref (optional)

No

browser_close

Close the current page or browser context.

None

No

mcp/tavily: AI-Native Web Search
Purpose & Use Case: The mcp/tavily server provides seamless interaction with Tavily's AI-native search engine, which is specifically designed to provide comprehensive, real-time, and relevant results for LLMs.[28, 29] It's the ideal tool for gathering current information, news, and performing detailed web content analysis, overcoming the limitation of an LLM's static training data.[2]

Key Features: The server exposes four powerful tools: tavily-search for real-time web search with advanced filtering; tavily-extract for intelligently pulling structured content from web pages; tavily-map for creating a structured map of a website's URLs; and tavily-crawl for systematically exploring a website.[28, 29]

Quickstart & Configuration:

Prerequisites: You must sign up for a Tavily account to get a free API key.[2]

Docker Desktop: Enable mcp/tavily from the catalog. In the Config tab, you must add a secret for tavily.api_token and provide your API key.[30]

Manual Configuration (e.g., Claude Desktop with npx):

{
  "mcpServers": {
    "tavily-mcp": {
      "command": "npx",
      "args": ["-y", "tavily-mcp@latest"],
      "env": {
        "TAVILY_API_KEY": "your-tavily-api-key-here"
      }
    }
  }
}

(As documented in [2])

Usage Example (Natural Language Prompt):

"Search for news articles about developments in generative AI from the last 7 days. From the top 3 results, extract the main content of each article and provide a concise summary."

Tool Reference Table: mcp/tavily

Tool Name

Description

Key Parameters

Read-Only?

tavily-search

Performs a real-time web search.

query, search_depth, topic, max_results

Yes

tavily-extract

Extracts structured content from a list of URLs.

urls, extract_depth, format

Yes

tavily-map

Creates a structured map of a website's URLs.

url, max_depth, allow_external

Yes

tavily-crawl

Initiates a structured web crawl from a base URL.

base_url, max_depth, max_breadth

Yes

Data & Databases
This section covers tools for interacting with various data storage systems, from object storage and NoSQL databases to specialized financial data sources.

mcp/tigris: S3-Compatible Object Storage
Purpose & Use Case: The mcp/tigris server provides an MCP interface to Tigris, a globally distributed, S3-compatible object storage service.[14, 31] It enables AI agents to manage buckets and objects, making it useful for workflows that involve storing and retrieving files, such as AI model artifacts, training data, or application backups.[31, 32]

Key Features: The server offers a full suite of tools for bucket management (tigris_create_bucket, tigris_delete_bucket) and object management (tigris_put_object, tigris_get_object, tigris_delete_object, tigris_get_signed_url_object).[32]

Quickstart & Configuration:

Prerequisites: Sign up for a Tigris account at storage.new and generate an access key pair (AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY).[31]

Docker Desktop: Enable mcp/tigris. In the Config tab, provide your access key ID and secret key as environment variables.

Manual Configuration (Docker):

{
  "mcpServers": {
    "tigris-mcp-server": {
      "command": "docker",
      "args":,
      "env": {
        "AWS_ACCESS_KEY_ID": "YOUR_TIGRIS_ACCESS_KEY_ID",
        "AWS_SECRET_ACCESS_KEY": "YOUR_TIGRIS_SECRET_ACCESS_KEY"
      }
    }
  }
}

(Adapted from [31])

Usage Example (Natural Language Prompt):

"Create a new private Tigris bucket named 'project-alpha-logs'. Then, upload the local file /path/to/app.log to this bucket with the key 'daily-log-2024-07-03.txt'."

mcp/mongodb: NoSQL Database Interaction
Purpose & Use Case: This server connects AI agents to MongoDB databases and MongoDB Atlas clusters, allowing for natural language interaction with NoSQL data stores.[14] It's useful for developers building applications on top of MongoDB who want to query data, check records, or perform simple data analysis without writing manual queries.

Key Features: Provides tools for connecting to databases, listing collections, finding documents, and performing aggregations.

Quickstart & Configuration:

Prerequisites: A MongoDB connection string (URI).

Docker Desktop: Enable mcp/mongodb. In the Config tab, provide the MongoDB connection string as an environment variable (e.g., MONGODB_URI).

Usage Example (Natural Language Prompt):

"In my 'e-commerce' database, connect to the 'users' collection and find the user with the email 'test@example.com'."

mcp/sec-edgar: US Financial Data Retrieval
Purpose & Use Case: The mcp/sec-edgar server provides a direct line for AI agents to the U.S. Securities and Exchange Commission's EDGAR database, which contains a vast trove of public company financial filings.[14, 33, 34] This tool is invaluable for financial analysts, researchers, and developers who need to programmatically access official financial data for tasks like competitive analysis, investment research, or building fintech applications.

Key Features: The server exposes tools to retrieve company facts (get_company_facts), specific financial concepts (get_company_concepts), submission records (get_submissions), and XBRL data frames (get_xbrl_frames).[33] It uses the ticker symbol or CIK to identify companies.[34]

Quickstart & Configuration:

Prerequisites: The SEC API requires a User-Agent string for all requests, which should identify your application or contact information (e.g., "Your Name (your.email@example.com)").[33, 34]

Docker Desktop: Enable mcp/sec-edgar. In the Config tab, set the SEC_EDGAR_USER_AGENT environment variable to your identifier.

Manual Configuration (Docker):

{
  "mcpServers": {
    "sec-edgar": {
      "command": "docker",
      "args":,
      "env": {
        "SEC_EDGAR_USER_AGENT": "Your Name (your.email@example.com)"
      }
    }
  }
}

(As documented in [33])

Usage Example (Natural Language Prompt):

"Using the SEC EDGAR database, retrieve the total revenue and net income for Apple (AAPL) and Microsoft (MSFT) for the fiscal year 2023."

Productivity & Collaboration
These tools integrate AI with common developer platforms for version control, project management, and communication.

mcp/github: Advanced Version Control and Repository Management
Purpose & Use Case: The official GitHub MCP Server provides seamless integration with the GitHub API, enabling advanced automation and interaction capabilities for developers.[35] It allows an AI agent to become an active participant in the development workflow, capable of automating tasks like analyzing repository data, managing issues and pull requests, and even interacting with GitHub Actions.[35]

Key Features: The server is available in two modes: a remote server hosted by GitHub for easy setup, and a local server for environments that require it.[35] It features a highly granular set of tools that can be enabled or disabled via --toolsets flags (e.g., repos, issues, pull_requests, actions). This allows for fine-tuned control over the agent's capabilities and helps optimize the LLM's context window.[35]

Quickstart & Configuration (Local Server):

Prerequisites: A GitHub Personal Access Token (PAT) with the necessary scopes for the actions you want to perform.

Docker Desktop: Enable mcp/github. In the Config tab, provide your GITHUB_PERSONAL_ACCESS_TOKEN as an environment variable.

Manual Configuration (Local Go Binary):

{
  "mcp": {
    "servers": {
      "github": {
        "command": "/path/to/github-mcp-server",
        "args": ["stdio", "--toolsets", "repos,issues"],
        "env": {
          "GITHUB_PERSONAL_ACCESS_TOKEN": "<YOUR_GITHUB_PAT>"
        }
      }
    }
  }
}

(As documented in [35])

Usage Example (Natural Language Prompt):

"In the 'docker/mcp-registry' repository, list the 5 most recently updated open issues that have the 'bug' label. For the most recent one, retrieve its title, body, and all comments."

mcp/line: Real-time Messaging Integration
Purpose & Use Case: This server integrates the LINE Messaging API, allowing an AI agent to connect to a LINE Official Account.[14] This is useful for building chatbots, notification systems, or any application that needs to send or receive messages via the LINE platform.

Key Features: Provides tools for sending messages, receiving webhooks, and managing user interactions within the LINE ecosystem.

Quickstart & Configuration:

Prerequisites: A LINE Official Account and the associated channel access token and channel secret.

Docker Desktop: Enable mcp/line. In the Config tab, provide your LINE credentials as environment variables.

Usage Example (Natural Language Prompt):

"Send a message to the user with ID 'U12345...' on LINE that says: 'Your build has completed successfully.'"

mcp/keboola-mcp: Enterprise Data Platform Bridge
Purpose & Use Case: The Keboola MCP Server is an open-source bridge connecting AI agents to a Keboola project.[14, 36] It turns Keboola features—like Storage access, SQL transformations, and job triggers—into callable tools for AI clients, enabling natural language data exploration and analysis on an enterprise data platform.[36]

Key Features: Provides tools for querying tables, managing metadata, creating SQL transformations from natural language, and running Keboola jobs.[36]

Quickstart & Configuration:

Prerequisites: Access to a Keboola project with an admin or custom storage token.[36]

Docker Desktop: Enable mcp/keboola-mcp. In the Config tab, provide your Keboola Storage API URL (KBC_URL) and Token (KBC_STORAGE_API_TOKEN).

Usage Example (Natural Language Prompt):

"In my Keboola project, what tables contain customer information? Then, run a SQL query to find the top 10 customers by revenue from the 'sales_final' table."

Finance & Payments
This category includes tools for interacting with payment gateways and financial data providers, enabling AI-driven e-commerce and financial analysis.

mcp/stripe: Comprehensive Payment and Billing API
Purpose & Use Case: The mcp/stripe server allows an AI agent to interact directly with Stripe services over the Stripe API.[14, 37] This is extremely powerful for building e-commerce applications, managing subscriptions, handling customer billing, and processing refunds through natural language commands.

Key Features: The server provides a comprehensive set of over 20 tools covering the entire customer lifecycle, including creating and listing customers, products, prices, coupons, payment links, and subscriptions. It also includes tools for processing refunds and even searching the Stripe documentation.[38, 39]

Quickstart & Configuration:

Prerequisites: A Stripe account and a secret API key.[40]

Docker Desktop: Enable mcp/stripe. In the Config tab, you must add a secret for stripe.secret_key and provide your Stripe API key.[41]

Manual Configuration (Docker):

{
  "mcpServers": {
    "stripe": {
      "command": "docker",
      "args":,
      "env": {
        "STRIPE_SECRET_KEY": "YOUR_STRIPE_SECRET_KEY"
      }
    }
  }
}

(Adapted from [40])

Usage Example (Natural Language Prompt):

"Create a new customer in Stripe named 'Jane Doe' with the email 'jane.doe@example.com'. Then, create a new product called 'Premium Subscription'. Finally, create a recurring monthly price for this product at $25 USD and create a payment link for it."

Tool Reference Table: mcp/stripe

Tool Name

Description

Key Parameters

Read-Only?

create_customer

Creates a new customer in Stripe.

name (str), email (str)

No

list_customers

Fetches a list of customers.

limit (int), email (str)

Yes

create_product

Creates a new product.

name (str), description (str)

No

create_price

Creates a new price for a product.

product (ID), unit_amount (int), currency

No

create_payment_link

Creates a new payment link for a price.

price (ID), quantity (int)

No

create_refund

Refunds a payment intent.

payment_intent (ID)

No

list_subscriptions

Lists all subscriptions.

customer (ID), price (ID), status

Yes

search_stripe_documentation

Searches the official Stripe documentation.

question (str)

Yes

mcp/nasdaq-data-link: Real-time Financial Market Data
Purpose & Use Case: This server, developed by the community, interacts with the data feeds provided by the Nasdaq Data Link.[14, 42] It's a powerful tool for developers and analysts needing to integrate real-time and historical financial market data into their applications, such as stock prices, company fundamentals, and economic indicators.

Key Features: Provides a wide array of tools for retrieving data from various Nasdaq databases, including Equities 360 (balance sheets, cash flow), Nasdaq Fund Network, and World Bank data.[42, 43]

Quickstart & Configuration:

Prerequisites: An API key from the Nasdaq Data Link website (data.nasdaq.com).[43]

Docker Desktop: Enable mcp/nasdaq-data-link. In the Config tab, provide your NASDAQ_DATA_LINK_API_KEY as an environment variable.

Manual Configuration (Cloning from source):

# 1. Clone the repo
git clone https://github.com/stefanoamorelli/nasdaq-data-link-mcp.git
cd nasdaq-data-link-mcp
# 2. Install dependencies
uv init mcp
uv add "mcp[cli]"
# 3. Create.env file with your API key
echo "NASDAQ_DATA_LINK_API_KEY=your_api_key_here" >.env
# 4. Install the server with the mcp CLI
uv run mcp install nasdaq_data_link_mcp_os/server.py --env-file.env --name "Nasdaq Data Link MCP"

(As documented in [43])

Usage Example (Natural Language Prompt):

"Retrieve the annual balance sheet data for Microsoft (MSFT) for the fiscal year ending in 2022. Specifically, I need the values for total assets and total liabilities."

Utilities & Miscellaneous
This category contains a variety of useful tools, from reference implementations to specialized servers for visualization and SDK interaction.

mcp/everything: The Reference Server
Purpose & Use Case: The mcp/everything server is an official reference implementation from the MCP team.[6, 44] Its primary purpose is not to perform a specific real-world task, but to demonstrate the various features of the Model Context Protocol itself. It's an essential tool for developers who are building their own MCP servers, as it provides clear examples of how to implement features like long-running operations, progress updates, resource references, and annotations.[6]

Key Features: Includes a set of simple but illustrative tools: echo (echoes back input), add (adds two numbers), longRunningOperation (demonstrates progress updates), getResourceReference (shows how to return references), and printEnv (helps debug server configuration).[6]

Quickstart & Configuration: This server is designed to work out-of-the-box. Simply enable it from the MCP Toolkit > Catalog in Docker Desktop. No secrets or special configuration are required.

Usage Example (Natural Language Prompt):

"Use the 'longRunningOperation' tool and show me the progress updates as it runs."

mcp/vizro: Step-by-Step Data Visualization
Purpose & Use Case: The mcp/vizro server, from McKinsey, provides tools and templates to create functioning Vizro charts and dashboards step-by-step.[7, 14, 45] This allows an AI agent to act as a data visualization assistant, taking a user's request and a dataset and turning it into Python code for a chart or dashboard.

Key Features: The server provides a structured workflow for visualization: load_and_analyze_data to understand a dataset, get_vizro_chart_or_dashboard_plan to get instructions, and validate_chart_code or validate_model_config to validate the generated code and even open it in a browser via a PyCafe link.[7, 46] It also includes sample datasets like iris and stocks for easy experimentation.[7]

Quickstart & Configuration: Enable mcp/vizro from the Docker MCP Catalog. It requires no special configuration.

Usage Example (Natural Language Prompt):

"Using the sample 'stocks' dataset, create a line chart showing the change in the 'GOOG' stock price over time."

Tool Reference Table: mcp/vizro

Tool Name

Description

Key Parameters

Read-Only?

load_and_analyze_data

Load data from a file/URL and analyze its structure.

path_or_url

Yes

get_sample_data_info

Get information about available sample datasets.

data_name

Yes

get_vizro_chart_or_dashboard_plan

Get instructions for creating a Vizro visualization.

user_plan

Yes

validate_chart_code

Validate the generated chart code.

chart_config, data_info

Yes

validate_model_config

Validate a complete dashboard configuration.

dashboard_config, data_infos

Yes

mcp/dart: Dart & Flutter SDK Interaction
Purpose & Use Case: This server exposes the Dart SDK commands as MCP tools, bridging the gap between AI coding assistants and Dart/Flutter development workflows.[14, 47] It allows a developer to ask their AI agent to perform tasks like analyzing code, creating new projects, or running tests directly.

Key Features: Provides tools that map directly to Dart commands, such as dart-analyze, dart-compile, dart-create, dart-fix, and dart-test.[47] It features intelligent path handling and auto-detects Dart/Flutter projects.[47]

Quickstart & Configuration: Can be run directly via npx without installation.

{
  "mcpServers": {
    "dart": {
      "command": "npx",
      "args": [ "-y", "@egyleader/dart-mcp-server" ]
    }
  }
}

(As documented in [47])

Usage Example (Natural Language Prompt):

"Analyze the Dart code in my current project directory and report any errors or critical warnings."

mcp/npm-sentinel: Intelligent NPM Package Analysis
Purpose & Use Case: mcp/npm-sentinel is a powerful server that revolutionizes NPM package analysis by leveraging AI.[14, 48] It provides real-time intelligence on package security, dependencies, and performance, helping developers make faster and safer decisions about the packages they use in their projects.[48, 49]

Key Features: Offers a wide range of tools for deep package analysis, including npmVulnerabilities (security scanning), npmDeps (dependency tree analysis), npmTrends (download statistics), npmCompare (compare multiple packages), and npmAlternatives (find similar packages).[48]

Quickstart & Configuration (VS Code):

{
  "servers": {
    "npm-sentinel": {
      "type": "stdio",
      "command": "npx",
      "args": ["-y", "@nekzus/mcp-server@latest"]
    }
  }
}

(As documented in [48])

Usage Example (Natural Language Prompt):

"Compare the packages 'express' and 'fastify'. Show me their latest versions, weekly download trends, and a summary of any known critical vulnerabilities."

Part III: Advanced Concepts and Best Practices
Moving beyond individual tools, this final section explores how to strategically leverage the MCP ecosystem to build sophisticated workflows and maintain a secure and manageable setup.

Architecting Complex Workflows with MCP
As you become more familiar with individual MCP tools, the next step is to combine them into powerful, multi-step agentic workflows that can automate complex development tasks from start to finish.

The Rise of the Meta-Server: Managing Complexity
The rapid growth of the MCP ecosystem, with over 130 servers available in the Docker Hub mcp namespace alone, presents a new challenge: configuration management.[14] Manually adding and maintaining dozens of server configurations in a single JSON file is inefficient and error-prone. This very success has created a need for a higher level of abstraction.

In response, the ecosystem has seen the emergence of "meta-servers" or proxies. These are MCP servers whose primary function is not to provide a direct tool, but to manage all your other MCP servers. Tools like mcp/pluggedin-mcp-proxy act as a unified interface, aggregating multiple underlying servers and routing requests intelligently.[14, 50] Similarly, community projects like ggoodman/mcp envision a command-line tool that acts as a central broker for discovering, installing, and managing all other servers.[51]

For a power user, the most effective "quickstart" might not be to install ten individual tools, but to install a single meta-server like mcp/pluggedin-mcp-proxy.[50] This allows you to manage your entire toolset from a single dashboard or API, connecting it once to your AI client. This represents a significant maturation of the ecosystem, moving from individual tool integrations to a more holistic platform approach.

Chaining Tools for Agentic Workflows: A Practical Example
The true potential of MCP is unlocked when an AI agent can intelligently chain calls to multiple tools from different servers to accomplish a complex goal. Consider the following scenario:

User Prompt: "Research the latest financial performance of Microsoft, create a visualization of its revenue over the last 4 quarters, and post a summary to my team's '#general' Slack channel."

An advanced AI agent connected to the MCP gateway could execute the following workflow:

Tool: mcp/tavily -> tavily-search

Action: The agent first performs a broad search for "Microsoft investor relations" and "Microsoft latest quarterly earnings report" to find the official sources of financial data.

Tool: mcp/sec-edgar -> get_company_facts

Action: Using Microsoft's CIK (0000789019), the agent calls this tool to retrieve structured financial data, specifically looking for the Revenues fact for the last four fiscal quarters.

Tool: mcp/vizro -> validate_chart_code

Action: The agent takes the structured revenue data and generates Python code to create a bar chart using the Vizro library. It then passes this code to the validate_chart_code tool to ensure it's correct and to get a shareable link to the visualization.

Tool: mcp/slack -> post_message [44]

Action: Finally, the agent synthesizes the information it has gathered. It drafts a summary of Microsoft's quarterly revenue and uses the post_message tool to send the summary, along with the link to the Vizro chart, to the specified Slack channel.

This example demonstrates how MCP enables an AI to orchestrate a sequence of actions across search, financial data retrieval, data visualization, and communication tools, automating a task that would have previously required significant manual effort.

Security and Credential Management
While the Docker MCP Toolkit provides a secure-by-default environment, the ultimate responsibility for maintaining a secure setup lies with the user. Adhering to best practices for handling secrets and filesystem access is paramount.

Best Practices for Handling Secrets
The most critical security practice is to never hardcode secrets like API keys, tokens, or passwords directly in the args array of your MCP server configuration. The toolkit may attempt to block these, but the correct approach is to pass them as environment variables.

Almost every MCP server that requires secrets supports this method. The client configuration file has a dedicated env block for this purpose.

Correct (Secure) Method:

{
  "mcpServers": {
    "stripe": {
      "command": "npx",
      "args": ["-y", "@stripe/mcp", "--tools=all"],
      "env": { "STRIPE_SECRET_KEY": "sk_test_..." }
    }
  }
}

(Example adapted from [2, 39, 40])

Incorrect (Insecure) Method:

// DO NOT DO THIS
{
  "mcpServers": {
    "stripe": {
      "command": "npx",
      "args": ["-y", "@stripe/mcp", "--tools=all", "--api-key=sk_test_..."]
    }
  }
}

By using the env block, you ensure that secrets are handled by the secure environment management systems of the client and the Docker runtime, rather than being exposed as plain text arguments.

Understanding Filesystem Access
As emphasized in Part I, MCP servers run in a highly restrictive sandbox and have no access to the host filesystem by default.[4] This is a powerful security feature that prevents an agent from reading or modifying your local files without your explicit permission.

When a tool legitimately needs access to a project—for example, mcp/pulumi needing to read your .ts or .go files [17]—you must grant this access by mounting a local directory into the container. When doing so, always follow the principle of least privilege:

Mount Specific Directories: Never mount your entire home directory (~/) or other broad, sensitive locations. Only mount the specific project folder that the tool needs to access.

Use Read-Only Mounts When Possible: If a tool only needs to read files (e.g., for analysis), append :ro to the volume mount to make it read-only. This prevents the agent from making any modifications.

Example of a secure mount for a Pulumi project:

"args": [
  "run", "-i", "--rm",
  "-v", "/Users/dev/projects/my-pulumi-app:/app/project:ro", // Mount specific project read-only
  "pulumi/mcp-server:latest"
]

By carefully managing secrets and filesystem access, you can confidently leverage the full power of the Docker MCP ecosystem while maintaining a secure and controlled development environment.