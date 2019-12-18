FROM golang:1.12-alpine AS build-stage

RUN apk add --update make git
RUN mkdir /migrations
WORKDIR /migrations

COPY . /migrations

RUN make get-migrations-deps build-migrations
RUN ["chmod", "+x", "/migrations/migrate.sh"]

# Final Stage
FROM alpine

COPY --from=build-stage /migrations/* /migrations/

ENTRYPOINT ["/migrations/migrate.sh"]