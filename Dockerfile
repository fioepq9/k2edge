From golang:1.19.5
ENV GOOS=linux 
ENV GOARCH=arm64
ENV GOARM=8
ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct
WORKDIR /app
COPY . .

CMD make bin
