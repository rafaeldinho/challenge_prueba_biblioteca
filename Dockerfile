# Go image for building the project
FROM golang:alpine as builder

ENV GOBIN=$GOPATH/bin
ENV GO111MODULE="on"

RUN mkdir -p $GOPATH/src/github/meli
WORKDIR $GOPATH/src/github/meli

COPY . .

RUN cd src
RUN ls -l


RUN CGO_ENABLED=0 GOOS=linux go build -installsuffix cgo -ldflags '-extldflags "-static"' -o $GOBIN/main ./main.go

# Runtime image with scratch container
FROM scratch
ARG VERSION
ENV VERSION_APP=$VERSION

COPY --from=builder /go/bin/ /src/
COPY . .

EXPOSE 8080
ENTRYPOINT ["/src/main"]