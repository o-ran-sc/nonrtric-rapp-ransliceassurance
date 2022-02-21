FROM golang:1.15.2-alpine3.12 as build
RUN apk add git
RUN mkdir /build
ADD . /build
WORKDIR /build
RUN go build -o simulator .
 
FROM alpine:latest
RUN mkdir /app
WORKDIR /app/
 
# Copy the Pre-built binary file from the previous stage
COPY --from=build /build .
 
# Expose port 8080
EXPOSE 8080
 
# Run Executable
ENTRYPOINT  ["/stub"]