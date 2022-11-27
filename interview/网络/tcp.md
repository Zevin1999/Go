1. TCP三次握手？ 
   开始客户端和服务器都处于CLOSED状态，然后服务端开始监听某个端口，进入LISTEN状态 
   第一次握手(SYN=1, seq=x)，发送完毕后，客户端进入 SYN_SEND 状态 【syn】
   第二次握手(SYN=1, ACK=1, seq=y, ACKnum=x+1)， 发送完毕后，服务器端进入 SYN_RCV 状态。 【ack+syn】
   第三次握手(ACK=1，ACKnum=y+1)，发送完毕后，客户端进入 ESTABLISHED 状态，当服务器端接收到这个包时,也进入 ESTABLISHED 状态，TCP 握手，即可以开始数据传输。[ack]
2. TCP为什么是三次握手？而不是两次或四次？
   两次握手服务器无法确认客户端是否收到消息；
   四次握手无必要，因为三次握手可以满足通信双方开始数据传输条件。
3. TCP四次挥手？ 
   第一次挥手：Client将FIN置为1，发送一个序列号seq给Server；进入FIN_WAIT_1状态；(FIN=1，seq=u)
   第二次挥手：Server收到FIN之后，发送一个ACK=1，ack number=收到的序列号+1；进入CLOSE_WAIT状态。此时客户端已经没有要发送的数据了，但仍可以接受服务器发来的数据。(ACK=1，ack=u+1,seq =v)
   第三次挥手：Server将FIN置1，发送一个序列号给Client；进入LAST_ACK状态；(FIN=1，ACK=1,seq=w)
   第四次挥手：Client收到服务器的FIN后，进入TIME_WAIT状态；接着将ACK置1，发送一个ack number=序列号+1给服务器；服务器收到后，确认acknowledge number后，变为CLOSED状态，不再向客户端发送数据。客户端等待2MSL（报文段最长寿命）时间后，也进入CLOSED状态。完成四次挥手。 \
4. TCP为什么是四次挥手？而不是两次或三次？
   TCP是全双工通信，服务端和客服端都能发送和接收数据。在断开连接时，需要服务端和客服端都确定对方将不再发送数据。 
   为什么不是3次挥手？ 在客户端第1次挥手时，服务端可能还在发送数据。 所以第2次挥手和第3次挥手不能合并。
5. TIME-WAIT 状态为什么需要等待2MSL（两个最大段生命周期）？
   1个MSL保证四次挥手中主动关闭方最后的ACK报文最终能到达对端
   1个MSL保证对端没有收到ACK报文时，进行重传的FIN报文能够到达
6. 

参考链接：https://juejin.cn/post/6983639186146328607