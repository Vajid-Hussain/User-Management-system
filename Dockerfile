FROM golang:1.22

WORKDIR /user_management/

COPY ./ .

RUN go mod download

RUN go build -o /user_management/cmd/api/api ./cmd/api/main.go

EXPOSE 8000

CMD [ "./cmd/api/api" ]