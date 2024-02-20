FROM golang:1.22-alpine AS firststage

WORKDIR /user_management

COPY ./ .

RUN go mod download

RUN go build -o ./build ./cmd/api/main.go

FROM scratch

WORKDIR /user_management_system

COPY --from=firststage /user_management/build ./
COPY --from=firststage /user_management/.env .
COPY --from=firststage /user_management/template ./template/

CMD [ "./build" ]

