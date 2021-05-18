# Stage 1: Compile and Build angular codebase

# Use official node image as the base image
FROM node:alpine as BUILD_IMAGE

# Set the working directory
WORKDIR /usr/local/app

# Add the source code to app
COPY package.json package-lock.json ./
# Install all the dependencies
RUN npm install

COPY . /usr/local/app/
# Generate the build of the application
RUN npm run build


# Stage 2: Serve app with nginx server

# Use official nginx image as the base image
FROM nginx:alpine

# Copy the build output to replace the default nginx contents.
COPY nginx.conf /etc/nginx/conf.d/
RUN rm -rf /usr/share/nginx/html/*
COPY --from=BUILD_IMAGE /usr/local/app/dist/hmn-frontend-app /usr/share/nginx/html


# Expose port 80
EXPOSE 80