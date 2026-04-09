FROM golang:1.26.1-alpine AS build
WORKDIR /src

RUN apk add --no-cache ca-certificates

COPY go.mod go.sum ./
RUN go mod download

COPY cmd ./cmd
COPY internal ./internal
COPY content ./content
COPY static ./static
COPY templates ./templates

RUN CGO_ENABLED=0 GOOS=linux go build -o /out/blog ./cmd

FROM gcr.io/distroless/static-debian12:nonroot
WORKDIR /app

COPY --from=build /out/blog /app/blog
COPY --from=build /src/content /app/content
COPY --from=build /src/static /app/static
COPY --from=build /src/templates /app/templates

ENV PORT=8080
EXPOSE 8080

ENTRYPOINT ["/app/blog"]
