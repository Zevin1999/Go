1. 一个完整的Docker由哪些部分组成? 
   1. DockerClient 客户端 
   2. Docker Daemon 守护进程 
   3. Docker Image 镜像 
   4. Docker Container 容器
2. Docker的本质：进程
3. Docker常用命令
4. Docker容器的生命周期
   1. 创建容器
   2. 运行容器
   3. 暂停容器、取消暂停容器（可选）
   4. 启动容器
   5. 停止容器
   6. 重启容器
   7. 杀死容器
   8. 销毁容器
5. DockerFile是什么？ 
   Dockerfile 是一个文本文件，其中包含我们需要运行已构建 Docker 镜像的所有命令，每一条指令构建一层，因此每一条指令的内容，就是描述该层应当如何构建。
   Docker 使用 Dockerfile 中的指令自动构建镜像。我们可以 docker build 用来创建按顺序执行多个命令行指令的自动构建。


参考：https://developer.aliyun.com/article/976339