# 第一阶段：构建应用程序
FROM golang:1.19 AS builder

# 设置工作目录
WORKDIR /source

ENV GOPROXY=https://goproxy.cn,direct

# 复制源码并构建应用程序
COPY . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/user-api api/user.go


# 第二阶段：创建最终镜像
FROM alpine:3.18.4

RUN mkdir /user
# 复制构建阶段的应用程序到最终镜像
COPY --from=builder /source /user

RUN chmod +x /user/bin/user-api
WORKDIR /user

ENTRYPOINT ["bin/user-api", "-f", "api/etc/user.yaml"]


#GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/user-api ./api/user.go
#RUN mkdir /go && cd /go \
#    && wget --no-check-certificate https://studygolang.com/dl/golang/go1.21.0.linux-amd64.tar.gz \
#    && tar -C /usr/local -zxvf go1.21.0.linux-amd64.tar.gz \
#    && rm -rf /go/go.1.21.0.linux-amd64.tar.gz \
#    && mkdir /lib64 \
#    && ln -s /lib/libc.musl-x86_64_so.1 /lib64/ld-linux-x86_64_so.2
#
#ENV GOPATH /go
#ENV PATH /usr/local/go/bin:$GOPATH/bin:$PATH
#
#CMD ["ping", "baidu.com"]



#docker run --name user-api -p 8888:8888 --net user --ip 168.10.0.50 --link etcd -d user-api


