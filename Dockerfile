FROM ubuntu:22.04

ENV TOKEN ${TOKEN}
ENV SUPERUSERID 446545799
ENV DB ./db/db.sqlite3
ENV LOG true
ENV DEPS "wget ca-certificates"
ENV APP_NAME gofinabot

RUN DEBIAN_FRONTEND=noninteractive \
    apt-get update --quiet --quiet && \
    apt-get upgrade --quiet --quiet && \
    apt-get install --quiet --quiet --yes \
    --no-install-recommends --no-install-suggests \
    ${DEPS} \
    && apt-get --quiet --quiet clean \
    && rm --recursive --force /var/lib/apt/lists/* /tmp/* /var/tmp/*

WORKDIR /var/www/${APP_NAME}

RUN wget --no-check-certificate https://github.com/Devil666face/${APP_NAME}/releases/latest/download/${APP_NAME}.tgz && \
    tar -xf ${APP_NAME}.tgz && \
    rm -rf ${APP_NAME}.tgz

CMD ["/bin/bash","-c","./$APP_NAME"]
