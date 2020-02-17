FROM golang:1.13

# set working directory
WORKDIR /app

# install ALL go_modules
COPY go.mod go.sum ./
RUN go mod download

# copy all project
COPY . .

# build project
RUN go build -o app ./bin

ENV HTTP_PORT=8080
ENV MONGO_SERVICE_HOST=
ENV MONGO_SERVICE_PORT=27017

EXPOSE "8080:8080"

CMD [ "./app" ]