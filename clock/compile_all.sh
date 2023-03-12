#!/bin/bash

cat << EOF > ./database/env.list
MARIADB_USER=$1
MARIADB_PASSWORD=$2
MARIADB_ROOT_PASSWORD=$2
EOF

docker build -t wininbrowser_clock_db ./database --rm

docker run --name wininbrowser_clock_db\
	-p 3306:3306\
	--env-file ./database/env.list\
	-d wininbrowser_clock_db


docker build -t wininbrowser_clock_ms ./logic --rm

cat << EOF > ./logic/env.list
DBHOST=$(docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' wininbrowser_clock_db)
DBUSER=$1
DBPASS=$2
EOF

docker run --name wininbrowser_clock_ms\
	-p 9090:9090\
	--env-file ./logic/env.list\
	-d wininbrowser_clock_ms

