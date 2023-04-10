# First stage: Build the application
FROM node:alpine AS build

# Install pnpm globally
RUN npm install -g pnpm

# Set the working directory
WORKDIR /app

# Copy package.json and pnpm-lock.yaml
COPY ./web/package.json ./web/pnpm-lock.yaml ./

# Install dependencies
RUN pnpm install --frozen-lockfile

# Copy the rest of the application files
COPY ./web .

# Build the application
RUN pnpm build

# Second stage: Serve the application with NGINX
FROM nginx:stable-alpine

# Copy built files from the build stage to the NGINX html directory
COPY --from=build /app/dist /usr/share/nginx/html

# Copy an optional custom nginx configuration file if needed
# COPY nginx.conf /etc/nginx/conf.d/default.conf

# Expose port 80
EXPOSE 80

# Start NGINX
CMD ["nginx", "-g", "daemon off;"]
