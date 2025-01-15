FROM freelabspace/adoptopenjdk8:alpine-slim-cn

WORKDIR /home

ARG JAR_FILE
COPY ${JAR_FILE} app.jar

ENTRYPOINT [ \
    "java", \
    "-Djava.security.egd=file:/dev/./urandom", \
    "-jar", \
    "app.jar" \
]