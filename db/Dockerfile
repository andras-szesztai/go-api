FROM mysql:latest

ENV MYSQL_ROOT_PASSWORD=password

COPY ./db.sql /docker-entrypoint-initdb.d/

EXPOSE 3306