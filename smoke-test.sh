# Smoke test pour BioLab MCP Research Server
# - Pull l'image GHCR
# - Lance un conteneur temporaire (mapping host:8080 -> container:8000)
# - Vérifie l'endpoint /health

#!/usr/bin/env bash
set -euo pipefail

IMAGE="ghcr.io/aguennoune/biolab-mcp:v0.1.0"
CONTAINER_NAME="biolab-mcp-smoke-test"
PORT="8080"
HEALTH_PATH="/health"
TIMEOUT=30

echo "🔍 Starting smoke test for BioLab MCP Research Server..."
echo "Image: $IMAGE"

# Clean up any existing container
docker rm -f "$CONTAINER_NAME" >/dev/null 2>&1 || true

echo "📦 Pulling image..."
if ! docker pull "$IMAGE"; then
    echo "❌ Failed to pull image $IMAGE"
    exit 1
fi

echo "🚀 Starting container..."
CONTAINER_ID=$(docker run -d --name "$CONTAINER_NAME" -p "${PORT}:8000" "$IMAGE")
if [ -z "$CONTAINER_ID" ]; then
    echo "❌ Failed to start container"
    exit 1
fi

echo "Container ID: $CONTAINER_ID"

echo "⏳ Waiting for service to be ready (timeout: ${TIMEOUT}s)..."
for i in $(seq 1 $TIMEOUT); do
    if curl -sSf "http://localhost:${PORT}${HEALTH_PATH}" >/dev/null 2>&1; then
        echo "✅ Health check passed!"
        echo "📋 Health check response:"
        curl -sS "http://localhost:${PORT}${HEALTH_PATH}" || true

        echo "🧹 Cleaning up..."
        docker stop "$CONTAINER_NAME" >/dev/null 2>&1 || true
        docker rm "$CONTAINER_NAME" >/dev/null 2>&1 || true

        echo "✅ Smoke test PASSED!"
        exit 0
    fi

    if [ $((i % 5)) -eq 0 ]; then
        echo "⏳ Still waiting... ($i/${TIMEOUT}s)"
    fi
    sleep 1
done

echo "❌ Health check failed after ${TIMEOUT}s timeout"
echo "📋 Container logs:"
docker logs "$CONTAINER_NAME" || true

echo "🧹 Cleaning up..."
docker stop "$CONTAINER_NAME" >/dev/null 2>&1 || true
docker rm "$CONTAINER_NAME" >/dev/null 2>&1 || true

echo "❌ Smoke test FAILED!"
exit 1
#!/usr/bin/env bash
set -euo pipefail

# Configuration
IMAGE="ghcr.io/aguennoune/biolab-mcp:v0.1.0"
CONTAINER_NAME="biolab-mcp-smoke-test"
PORT="8080"
HEALTH_PATH="/health"
TIMEOUT=30

echo "🔍 Starting smoke test for BioLab MCP Research Server..."
echo "Image: $IMAGE"

# Clean up any existing container
docker rm -f "$CONTAINER_NAME" >/dev/null 2>&1 || true

echo "📦 Pulling image..."
if ! docker pull "$IMAGE"; then
    echo "❌ Failed to pull image $IMAGE"
    exit 1
fi

echo "🚀 Starting container..."
CONTAINER_ID=$(docker run -d --name "$CONTAINER_NAME" -p "${PORT}:8000" "$IMAGE")
if [ -z "$CONTAINER_ID" ]; then
    echo "❌ Failed to start container"
    exit 1
fi

echo "Container ID: $CONTAINER_ID"

# Wait for service to be ready
echo "⏳ Waiting for service to be ready (timeout: ${TIMEOUT}s)..."
for i in $(seq 1 $TIMEOUT); do
    if curl -sSf "http://localhost:${PORT}${HEALTH_PATH}" >/dev/null 2>&1; then
        echo "✅ Health check passed!"

        #!/usr/bin/env bash
        set -euo pipefail

        # Configuration
        IMAGE="ghcr.io/aguennoune/biolab-mcp:v0.1.0"
        CONTAINER_NAME="biolab-mcp-smoke-test"
        PORT="8080"
        HEALTH_PATH="/health"
        TIMEOUT=30

        echo "🔍 Starting smoke test for BioLab MCP Research Server..."
        echo "Image: $IMAGE"

        # Clean up any existing container
        docker rm -f "$CONTAINER_NAME" >/dev/null 2>&1 || true

        echo "📦 Pulling image..."
        if ! docker pull "$IMAGE"; then
            echo "❌ Failed to pull image $IMAGE"
            exit 1
        fi

        echo "🚀 Starting container..."
        CONTAINER_ID=$(docker run -d --name "$CONTAINER_NAME" -p "${PORT}:8000" "$IMAGE")
        if [ -z "$CONTAINER_ID" ]; then
            echo "❌ Failed to start container"
            exit 1
        fi

        echo "Container ID: $CONTAINER_ID"

        # Wait for service to be ready
        echo "⏳ Waiting for service to be ready (timeout: ${TIMEOUT}s)..."
        for i in $(seq 1 $TIMEOUT); do
            if curl -sSf "http://localhost:${PORT}${HEALTH_PATH}" >/dev/null 2>&1; then
                echo "✅ Health check passed!"

                # Get the actual response
                echo "📋 Health check response:"
                curl -sS "http://localhost:${PORT}${HEALTH_PATH}" | jq . 2>/dev/null || curl -sS "http://localhost:${PORT}${HEALTH_PATH}"
                echo "🧹 Cleaning up..."
                docker stop "$CONTAINER_NAME" >/dev/null 2>&1
                docker rm "$CONTAINER_NAME" >/dev/null 2>&1

                echo "✅ Smoke test PASSED!"
                exit 0
            fi

            if [ $((i % 5)) -eq 0 ]; then
                echo "⏳ Still waiting... ($i/${TIMEOUT}s)"
            fi

            sleep 1
        done

        echo "❌ Health check failed after ${TIMEOUT}s timeout"
        echo "📋 Container logs:"
        docker logs "$CONTAINER_NAME" || true

        echo "🧹 Cleaning up..."
        docker stop "$CONTAINER_NAME" >/dev/null 2>&1 || true
        docker rm "$CONTAINER_NAME" >/dev/null 2>&1 || true

        echo "❌ Smoke test FAILED!"
        exit 1
