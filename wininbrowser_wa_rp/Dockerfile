FROM nginx

  RUN apt-get update -qq && apt-get -y install apache2-utils
  ENV NODE_ROOT /var/www/frontend_web_node
  WORKDIR $NODE_ROOT
  RUN mkdir log
  COPY app.conf /tmp/app.nginx
  COPY /etc/nginx/nginx.crt /etc/nginx/nginx.crt
  COPY /etc/nginx/nginx.key /etc/nginx/nginx.key

  ENV FRONTEND_WEB_NODE=frontend_web_node

  RUN envsubst '$NODE_ROOT $FRONTEND_WEB_NODE' < /tmp/app.nginx > /etc/nginx/conf.d/default.conf

  EXPOSE 443

  CMD [ "nginx", "-g", "daemon off;" ]
