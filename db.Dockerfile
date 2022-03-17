# start with base image
FROM mariadb:latest

COPY ./resources/database/*.sql /docker-entrypoint-initdb.d/