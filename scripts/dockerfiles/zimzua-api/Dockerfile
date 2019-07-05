FROM ubuntu
RUN apt-get update -y && apt-get install -y vim locales

# 언어 설정
RUN sed -i -e 's/# ko_KR.UTF-8 UTF-8/ko_KR.UTF-8 UTF-8/' /etc/locale.gen && \
    echo 'LANG="ko_KR.UTF-8 UTF-8"'>/etc/default/locale && \
    dpkg-reconfigure --frontend=noninteractive locales && \
    echo 'export LANG=ko_KR.utf-8' >> /etc/bash.bashrc

COPY ./bin/zimzua-api /opt/
COPY ./bin/config.yml /opt/

CMD ["/opt/zimzua-api --config /opt/config.yml"]