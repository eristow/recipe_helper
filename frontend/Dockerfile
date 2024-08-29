# Use the official Node.js 22 image as the base image
FROM node:22

# Set the working directory inside the container
WORKDIR /frontend

# Copy the package.json and pnpm-lock.yaml files to the working directory
COPY package.json pnpm-lock.yaml ./

# Install dependencies using PNPM
RUN npm install -g pnpm && pnpm install

# Copy the rest of the application code to the working directory
COPY . .

# Build the React app using Vite
RUN pnpm run build

# Expose port 3000 for the React app
EXPOSE 3000

# Start the React app
CMD ["pnpm", "run", "preview"]