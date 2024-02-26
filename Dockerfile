FROM golang:1.20.5-alpine


ARG DOCKER_GIT_CREDENTIALS


ENV SRC_DIR=/go/src/app

WORKDIR $SRC_DIR


#COPY  /app for just binaries
COPY . $SRC_DIR

#RUN git config --global credential.helper store && echo "${DOCKER_GIT_CREDENTIALS}" > ~/.git-credentials


RUN go get ./
RUN go build .

# RUN ["go", "install", "github.com/githubnemo/CompileDaemon"]
RUN go install github.com/githubnemo/CompileDaemon@latest


#RUN go get -d -v ./...
# RUN go install -v ./...

EXPOSE 8080

CMD ["./backend"]


