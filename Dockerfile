FROM golang:latest

COPY timerbot /bin/timerbot

CMD ["/bin/timerbot"]

# vim: ft=dockerfile
