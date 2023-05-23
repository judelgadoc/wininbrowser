FROM mysql:latest

COPY setup_with_docker.sql /docker-entrypoint-initdb.d/
ENV MYSQL_ROOT_PASSWORD=12345
ENV MYSQL_USER=fredy
ENV MYSQL_PASSWORD=12345
ENV MYSQL_DATABASE=wininbrowser_calendar_db

EXPOSE 65001
