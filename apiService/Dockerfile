FROM golang:1.16 AS builder

# Set GOPATH to build Go app
ENV GOPATH=/go

# Set apps source directory
ENV SRC_DIR=${GOPATH}/src/github.com/wiqram/IG-Trading-Microservices/

# Define current working directory
WORKDIR ${SRC_DIR}

# Copy apps scource code to the image
COPY . ${SRC_DIR}

#private github repo hence need to add it properly so that latest commits can be fetched

# Build App
RUN ./build.sh

EXPOSE 50059
ENTRYPOINT [ "apiService" ]
