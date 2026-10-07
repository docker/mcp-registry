# Bagel

Bagel analyzes robot logs locally with SQL, metadata inspection, visualization exports, and data-reduction pipelines. This entry runs the publisher's ROS 2 Kilted image over native MCP stdio on AMD64 or ARM64. No hosted account or API key is required for local analysis.

## Configure

- **data_directory:** an existing absolute host directory containing your logs. Bagel sees it read-only at `/data/input`.
- **workspace_directory:** an existing writable absolute host directory. Bagel persists artifacts, saved pipelines, and custom capabilities here. On Linux, ensure the container's `ubuntu` user can write to this directory.

The two directories should be separate. Only configure directories you want this integration to access. Generated visualizations are saved under the workspace's `artifacts` directory. Query caches are temporary inside each container so read-only tool calls can operate while the gateway mounts host directories read-only. Cloud integrations, live subscriptions, and optional external services require their own configuration and are not configured by this catalog entry.

## Try it

Ask: "Summarize the metadata of `./data/sample/ros2/mcap`." This bundled sample contains 15 messages. For your own data, use paths such as `/data/input/run.mcap`.

Inspect source metadata and topic schemas before querying messages. Preview data-reduction pipelines before running them. For standing subscriptions or recurring jobs that must remain active, keep the gateway session running; this entry does not create a background service.

Documentation: https://github.com/Extelligence-ai/bagel#readme

Website: https://www.trybagel.com
