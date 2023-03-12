rm {logic,database}/env.list
docker rm -f wininbrowser_clock_{ms,db}
docker rmi wininbrowser_clock_{ms,db}
