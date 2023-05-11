# Specify a base image
FROM node:14.17.0-alpine

# Set the working directory
WORKDIR /app

# Copy the package.json and package-lock.json files to the container
COPY package*.json ./

# Install the dependencies
RUN npm install

# Copy the rest of the application files to the container
COPY . .

# Expose the port that the server will run on
EXPOSE 4000

# Start the server
CMD ["node", "index.js"]

