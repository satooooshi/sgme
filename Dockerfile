FROM golang:1.18

COPY go.mod /code/go.mod
COPY go.sum /code/go.sum
WORKDIR /code
RUN ls
#RUN go mod download

COPY . /code

ENV KUBECONFIG '/code/config'
ENV NAMESPACE 'default'

RUN go get .

CMD go run main.go

# https://qiita.com/Nossa/items/f70c561020b87532d933
# https://t8.dev/ja/blog/docker_dockerfile_env/
# docker build -t superheatedboy/sgme:latest .
# cp /Users/satoshiaikawa/.kube/config ./

# this cannot setup the server
# docker run --rm -it -p 8080:8080 superheatedboy/sgme:latest sh
# go run main.go

# this sets up the server
# docker container run -p 8080:8080 superheatedboy/sgme:latest
# docker ps
# docker stop 

