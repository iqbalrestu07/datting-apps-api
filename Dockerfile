# Builder
FROM golang:1.20.12-alpine as builder

# RUN apk update && apk upgrade && \
#     apk --update add git make bash build-base

WORKDIR /date-app

COPY . .
# RUN go install github.com/securego/gosec/v2/cmd/gosec@latest
# RUN gosec ./...
 
RUN go build -trimpath  -o datting-apps-api ./app/
# Distribution
FROM alpine AS runner
# ADD https://github.com/golang/go/raw/master/lib/time/zoneinfo.zip /zoneinfo.zip
# ENV ZONEINFO /zoneinfo.zip

WORKDIR /date-app

COPY --from=builder /date-app/datting-apps-api .
COPY --from=builder /date-app/.env ./.env
