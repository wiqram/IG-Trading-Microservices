FROM golang:1.15 AS builder

# Set GOPATH to build Go app
ENV GOPATH=/go

# Set apps source directory
ENV SRC_DIR=${GOPATH}/github.com/wiqram/IG-Trading-Microservices/notificationService

# Define current working directory
WORKDIR ${SRC_DIR}

# Copy apps scource code to the image
#COPY . ${SRC_DIR}
ADD . ${SRC_DIR}
# Build App

RUN ./build.sh
#RUN go get notificationservice
#RUN go install
EXPOSE 50054
ENTRYPOINT [ "notificationService" ]
