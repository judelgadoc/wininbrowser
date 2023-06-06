FROM nginx:latest

# Remove the default Nginx configuration
RUN rm /etc/nginx/conf.d/default.conf

# Copy your app.conf file to the container
COPY app.conf /etc/nginx/conf.d/

# Expose port 55690
EXPOSE 55690

# Start Nginx
CMD ["nginx", "-g", "daemon off;"]
