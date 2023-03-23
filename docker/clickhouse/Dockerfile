FROM yandex/clickhouse-server:21.3.20.1

CMD [ "/bin/sh -c apt-get update" ] 

# Add data
COPY init-db.sh /docker-entrypoint-initdb.d
COPY users.xml /etc/clickhouse-server/users.xml

EXPOSE 8123
EXPOSE 9000