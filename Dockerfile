# 使用alpine开启
FROM alpine
LABEL maintainer="18158899797@163.com"
# 拷贝编译程序
COPY . /app
WORKDIR /app
# 打开8080端口
EXPOSE 8000
# 运行!
CMD ["./main"]