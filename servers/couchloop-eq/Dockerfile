# CouchLoop EQ MCP Server - Production Dockerfile
FROM node:20-alpine AS builder

# Install build dependencies
RUN apk add --no-cache python3 make g++

WORKDIR /app

# Copy package files
COPY package*.json ./
COPY tsconfig.json ./

# Install ALL dependencies (including dev) for building TypeScript
RUN npm ci && \
    npm cache clean --force

# Copy source code
COPY src/ ./src/

# Build TypeScript
RUN npm run build

# Production image
FROM node:20-alpine

# Install runtime dependencies
RUN apk add --no-cache \
    curl \
    tini \
    && addgroup -g 1001 nodejs \
    && adduser -S nodejs -u 1001

WORKDIR /app

# Copy package files first for production dependencies
COPY --chown=nodejs:nodejs package*.json ./

# Install only production dependencies
RUN npm ci --only=production && \
    npm cache clean --force

# Copy built application
COPY --from=builder --chown=nodejs:nodejs /app/dist ./dist

# Security hardening
ENV NODE_ENV=production \
    NODE_OPTIONS="--max-old-space-size=2048 --enable-source-maps" \
    NPM_CONFIG_LOGLEVEL=error

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=40s --retries=3 \
    CMD curl -f http://localhost:3000/health || exit 1

# Switch to non-root user
USER nodejs

# Use tini for proper signal handling
ENTRYPOINT ["/sbin/tini", "--"]

# Start the MCP HTTP server
CMD ["node", "dist/server/index.js"]

EXPOSE 3000