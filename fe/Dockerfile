FROM node:18-alpine

WORKDIR /app

# Copy package.json and package-lock.json first (to optimize caching)
COPY package.json package-lock.json ./

# Ensure all dependencies are installed
RUN npm install

# Copy the rest of the application files
COPY . ./

EXPOSE 3000

CMD ["npm", "start"]
