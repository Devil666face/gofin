FROM ubuntu:20.04

ENV TOKEN ${TOKEN}
ENV SUPERUSERID 446545799
ENV DB_GOFINABOT db.sqlite3
ENV LOG true
ENV DEPS "wget ca-certificates locales libsqlite3-0"
ENV APP_NAME gofinabot

RUN DEBIAN_FRONTEND=noninteractive \
    apt-get update --quiet --quiet && \
    apt-get upgrade --quiet --quiet && \
    apt-get install --quiet --quiet --yes \
    --no-install-recommends --no-install-suggests \
    ${DEPS} \
    && apt-get --quiet --quiet clean \
    && rm --recursive --force /var/lib/apt/lists/* /tmp/* /var/tmp/*

RUN sed -i '/en_US.UTF-8/s/^# //g' /etc/locale.gen && \
    sed -i '/ru_RU.UTF-8/s/^# //g' /etc/locale.gen && \
    locale-gen
ENV LANG en_US.UTF-8  
ENV LANGUAGE en_US:en  
ENV LC_ALL en_US.UTF-8   

WORKDIR ${APP_NAME}

RUN wget --no-check-certificate https://github.com/Devil666face/${APP_NAME}/releases/latest/download/${APP_NAME}.tgz && \
    tar -xf ${APP_NAME}.tgz && \
    rm -rf ${APP_NAME}.tgz

CMD ["/bin/bash","-c","./$APP_NAME"]
