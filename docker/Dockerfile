FROM golang:1.13

# set working directory
WORKDIR /app

# install ALL go_modules
RUN go get -u

# copy all project
COPY . .

ENV HTTP_PORT=8080
ENV MONGO_SERVICE_HOST=
ENV MONGO_SERVICE_PORT=27017

EXPOSE "8080:8080"

RUN go install ./bin/run.go

CMD [ "app" ]