FROM cenots:7
COPY . /app
WORKDIR /app

EXPOSE 8081 8088
CMD ["/app/changeGo"]
