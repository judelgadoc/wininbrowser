FROM mariadb

COPY setup.sql /docker-entrypoint-initdb.d/
ENV MARIADB_USER=clock_ms
ENV MARIADB_PASSWORD=clock_ms12#$
ENV MARIADB_ROOT_PASSWORD=clock_ms12#$
ENV MARIADB_DATABASE=clock_db

EXPOSE 3306

