FROM mysql:8
COPY ./conf.d/my.cnf /etc/mysql/conf.d/my.cnf
COPY ./initdb.d /docker-entrypoint-initdb.d
EXPOSE 3306