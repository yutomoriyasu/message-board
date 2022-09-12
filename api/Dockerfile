FROM golang:1.18.2-buster as gobuilder
WORKDIR /app
COPY . /app
RUN go mod tidy
RUN go build -buildvcs=false -o /go/bin/app
RUN go install github.com/google/wire/cmd/wire@latest && \
  go install github.com/kyoh86/richgo@latest && \
  go install github.com/cosmtrek/air@latest && \
  go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
RUN go generate
CMD ["air"]

FROM gcr.io/distroless/base
COPY --from=gobuilder /go/bin/app /
CMD ["/app"]