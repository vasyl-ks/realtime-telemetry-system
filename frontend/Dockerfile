# --- Stage 1: Build React App ---
FROM node:20-alpine AS builder

# Set working directory inside container
WORKDIR /app

# Copy package files first to leverage Docker cache for dependencies
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy all source code
COPY . .

# Build the production-ready React app
RUN npm run build

# --- Stage 2: Serve with Nginx ---
FROM nginx:alpine

# Remove default nginx HTML content
RUN rm -rf /usr/share/nginx/html/*

# Copy React build output from builder stage
COPY --from=builder /app/build /usr/share/nginx/html

# Expose HTTP port
EXPOSE 80

# Run Nginx in foreground
CMD ["nginx", "-g", "daemon off;"]