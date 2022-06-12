FROM ubuntu
LABEL name="httpserver"
LABEL version="1.1"
COPY ./bin/httpserver /bin/
EXPOSE 8080
CMD nohup /bin/httpserver >my.out 2>&1 &
