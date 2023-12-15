# registry.numberly.in/team-infrastructure/3ans:1.0.0
FROM golang:1.21.3-alpine3.18 AS build

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /3ans

FROM gcr.io/distroless/static:nonroot
WORKDIR /

COPY --from=build /3ans /3ans

USER 65534
EXPOSE 8443 8080

ENTRYPOINT ["/3ans"]