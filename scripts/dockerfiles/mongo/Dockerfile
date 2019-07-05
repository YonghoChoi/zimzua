FROM mongo

RUN apt-get update -y && apt-get install -y curl vim apt-transport-https
ADD ./sample /opt/
COPY ./entrypoint.sh /docker-entrypoint-initdb.d/
RUN chmod +x /docker-entrypoint-initdb.d/entrypoint.sh