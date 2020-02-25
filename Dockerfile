FROM ubuntu

WORKDIR /app

COPY . /app

EXPOSE 9090

CMD ["/app/main", "-conf=/app/conf.json"]
