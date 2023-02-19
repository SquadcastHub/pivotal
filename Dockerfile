FROM alpine:3.14
LABEL app="pivotal" description="pivotal app"
WORKDIR /app
COPY pivotal /app/
EXPOSE 80
CMD /app/pivotal serve /config/*.json
