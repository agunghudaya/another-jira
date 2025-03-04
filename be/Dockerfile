FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /another-jira-be ./cmd/api

EXPOSE 8080

CMD [ "/another-jira-be" ]