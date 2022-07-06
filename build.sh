


docker build -t registry.cn-beijing.aliyuncs.com/ning1875_haiwai_image/go-mem-allocate:v1  .

docker push registry.cn-beijing.aliyuncs.com/ning1875_haiwai_image/go-mem-allocate:v1



# for subprocess
CGO_ENABLED=0  GOOS=linux GOARCH=amd64 go build -o go-mem-allocate main.go

docker build -t registry.cn-beijing.aliyuncs.com/ning1875_haiwai_image/go-mem-allocate-subprocess:v1  -f subprocess.Dockerfile .

docker push registry.cn-beijing.aliyuncs.com/ning1875_haiwai_image/go-mem-allocate-subprocess:v1
