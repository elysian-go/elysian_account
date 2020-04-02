FROM postgres:latest

WORKDIR /app
COPY . .

ADD add_extension.sh /docker-entrypoint-initdb.d/createExtension.sh
RUN chmod 755 /docker-entrypoint-initdb.d/createExtension.sh