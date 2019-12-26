FROM golang:alpine as build
RUN mkdir /app
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY  . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix -cgo -o server server.go

FROM scratch
COPY --from=build /app/server .
COPY --from=build /app/data /data
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
ENTRYPOINT [ "/server" ]
EXPOSE 3000