FROM mysql

RUN apt-get update && apt-get install -y vim

COPY sql/init.sql /opt/
COPY sql/point_example.sql /opt/