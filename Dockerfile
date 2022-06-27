##
## Build
##
FROM golang:1.17-buster AS build

WORKDIR /ShortLinksService

COPY . ./
RUN go mod tidy

RUN go build -o /links-service ./cmd/main.go

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /links-service /links-service

EXPOSE 8080

USER nonroot:nonroot

ENV DB_NAME=d7r6s29qsbs7ah
ENV DB_HOST=ec2-52-212-228-71.eu-west-1.compute.amazonaws.com
ENV DB_PORT=5432
ENV DB_USER=fzcamrgntritxl
ENV DB_PASSWORD=7945968eed261e91271a72e87735b1f85715d098b45847238aaeae55b2d9d535

ENTRYPOINT ["/links-service"]