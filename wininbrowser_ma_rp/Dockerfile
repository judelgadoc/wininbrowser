FROM nginx:latest

# Remove the default Nginx configuration
RUN rm /etc/nginx/conf.d/default.conf

# Copy your app.conf file to the container
COPY app.conf /etc/nginx/conf.d/
COPY nginx.crt /etc/nginx/nginx.crt
COPY nginx.key /etc/nginx/nginx.key

# Expose port 31695
EXPOSE 31695

# Start Nginx
CMD ["nginx", "-g", "daemon off;"]
