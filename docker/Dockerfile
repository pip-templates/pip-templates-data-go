FROM golang:1.13

# set working directory
WORKDIR /app

# copy package files
COPY go.mod go.sum ./

# install ALL go_modules
RUN go mod download

# copy all project
COPY . .

ENV HTTP_PORT=8080
ENV MONGO_SERVICE_HOST=
ENV MONGO_SERVICE_PORT=27017

EXPOSE "8080:8080"

CMD [ "./obj/app" ]