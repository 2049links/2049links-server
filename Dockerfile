
FROM sunmi-docker-images-registry.cn-hangzhou.cr.aliyuncs.com/public/golang:1.18.10 As builder

ENV GOPROXY https://goproxy.cn,direct
ENV GO111MODULE on

#step 1 build go cache
WORKDIR /go/cache
ADD app/go.mod .
ADD app/go.sum .
RUN go mod download

#step 2 build binary project
WORKDIR /project
ADD ./app/ .
RUN ls
RUN go build main.go

FROM sunmi-docker-images-registry.cn-hangzhou.cr.aliyuncs.com/public/debian:stable-slim
#run binary project
WORKDIR /app
COPY --from=builder /project/main .

# your project shell [project] [arg1] [arg2] ...
CMD [ "/app/main","api","start"]