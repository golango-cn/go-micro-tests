FROM golang:1.14

WORKDIR /go/src/app
COPY . .

RUN go env -w GOPROXY=https://goproxy.cn,direct
ENV GO111MODULE on

RUN go get -d -v ./...
## RUN go install -v ./...
RUN go build -o app
RUN chmod +x ./app

EXPOSE 5310/tcp

CMD ./app

#
#FROM alpine
#ADD ccserver-srv /ccserver-srv
#ENTRYPOINT [ "/ccserver-srv" ]
#
#ENV TZ=Asia/Shanghai
#ENV ASPNETCORE_ENVIRONMENT=Development
#
#WORKDIR /app
#COPY ccserver /app
#
#
#CMD dotnet MemberCore.Manage.dll
#
#FROM golang:1.8
#
#WORKDIR /go/src/app
#
#COPY . .
#
#RUN chmod 777 jsj.message
#CMD ./jsj.message