1. channel怎么理解的?channel写数据和读数据的过程是怎样的?
    channel是Golang在语言层面提供的goroutine间的通信方式，比Unix管道更易用也更轻便。
    channel主要用于进程内各goroutine间通信。
    Channel 分为两种：带缓冲、不带缓冲。对不带缓冲的 channel 进行的操作实际上可以看作“同步模式”，带缓冲的则称为“异步模式”。
    同步模式下，发送方和接收方要同步就绪，只有在两者都 ready 的情况下，数据才能在两者间传输（实际上就是内存拷贝）。否则，任意一方先行进行发送或接收操作，都会被挂起，等待另一方的出现才能被唤醒 
    异步模式下，在缓冲槽可用的情况下（有剩余容量），发送和接收操作都可以顺利进行。否则，操作的一方（如写入）同样会被挂起，直到出现相反操作（如接收）才会被唤醒。
2. channel组成 
   1. buf是有缓冲的channel所特有的结构，用来存储缓存数据，是个循环链表 
   2. sendx和recvx用于记录buf这个循环链表中的发送或者接收的index
   3. lock是个互斥锁 
   4. recvq和sendq分别是接收(<-channel)或者发送(channel <- xxx)的goroutine抽象出来的结构体(sudog)的队列，是个双向链表
3. channel是线程安全的
4. channel死锁场景
   1. 当一个channel中没有数据，而直接读取时，会发生死锁
   2. 当channel数据满了，再尝试写数据会造成死锁
   3. 向一个关闭的channel写数据
5. 对已经关闭的chan进行读写会怎么样？
   1. 读已经关闭的chan能一直读到东西，但是读到的内容根据通道内关闭前是否有元素而不同
      1. 如果chan关闭前，buffer内有元素还未读,会正确读到chan内的值，且返回的第二个bool值（是否读成功）为true
      2. 如果chan关闭前，buffer内有元素已经被读完，chan内无值，接下来所有接收的值都会非阻塞直接成功，返回 channel 元素的零值，但是第二个bool值一直为false
   2. 写已经关闭的chan会panic
6. 
向channel写数据 向一个channel中写数据简单过程如下：
如果等待接收队列recvq不为空，说明缓冲区中没有数据或者没有缓冲区，此时直接从recvq取出G,并把数据写入，最后把该G唤醒，结束发送过程； 如果缓冲区中有空余位置，将数据写入缓冲区，结束发送过程； 如果缓冲区中没有空余位置，将待发送数据写入G，将当前G加入sendq，进入睡眠，等待被读goroutine唤醒；

从channel读数据 从一个channel读数据简单过程如下：
如果等待发送队列sendq不为空，且没有缓冲区，直接从sendq中取出G，把G中数据读出，最后把G唤醒，结束读取过程； 如果等待发送队列sendq不为空，此时说明缓冲区已满，从缓冲区中首部读出数据，把G中数据写入缓冲区尾部，把G唤醒，结束读取过程； 如果缓冲区中有数据，则从缓冲区取出数据，结束读取过程； 将当前goroutine加入recvq，进入睡眠，等待被写goroutine唤醒；



Channel

+ 概述

    + 用于goroutine之间消息的传递
    + 通信顺序进程并发模式（CSP）

+ 类型

    + 缓冲/非缓冲

      ```
      // 缓冲通道声明长度
      // 缓冲通道在容量为空时，读端goroutine会阻塞；容量未满时，读写两端都不会阻塞；容量满了之后，写端goroutine会阻塞
      ch := make(chan int, 1024)  
      // 非缓冲通道未声明长度
      // 非缓冲通道对于读写两端的goroutine都会阻塞
      ch := make(chan int)
      ```

    + 获取（读）/发送（写）

      ```
      chan     读写
      <-chan   只读
      chan<-   只写
      ```

    + 状态（空/满/关闭）

      ```
      ch := make(chan interface{})
      close(ch)
      ```

+ 使用

    + 1

      ```
      package main
      
      import "fmt"
      
      func main(){
          ch := make(chan int, 1)
      
          go func() {
              ch <- 999
          }()
      
          value := <- ch
          fmt.Println("value:", value)
      }
      ```

    + close

      ```
      func main() {
          ch := make(chan int, 5)
          sign := make(chan string, 2)
      
          go func() {
              for i := 1; i <= 5; i++ {
                  ch <- i
      
                  time.Sleep(time.Second)
              }
      
              close(ch)
      
              fmt.Println("the channel is closed")
      
              sign <- "func1"
      
          }()
      
          go func() {
              for {
                  i, ok := <-ch
                  fmt.Printf("%d, %v \n", i, ok)
      
                  if !ok {
                      break
                  }
      
                  time.Sleep(time.Second * 2 )
              }
      
              sign <- "func2"
      
          }()
      
          <-sign
          <-sign
      }
      ```

    + for select

      ```
      func main() {
          put := make(chan int)
      
          go func() {
              for i := 0; i < 10; i++ {
                  put <- i
      
                  time.Sleep(time.Millisecond * 100)
              }
          }()
      
          go func() {
              for {
                  select {
                  case value := <-put:
                      fmt.Println("输出：", value)
                  }
              }
          }()
      
          time.Sleep(time.Second * 2)
          fmt.Println("退出")
      }
      ```

    + 生产者消费者

      ```
      func main() {
      
          cook := make(chan int)
          quit := make(chan bool)
          quit2 := make(chan bool)
      
          go func() {
              for i := 0; i < 5; i++ {
                  cook <- i + 1
                  fmt.Println("cook:", i+1)
                  time.Sleep(time.Second)
              }
          quit <- true
          }()
          go func() {
              for {
                  select {
                  case v := <-cook:
                      fmt.Println("eat:", v)
                  case <-quit:
                      quit2 <- true
                  }
              }
          }()
        
        <-quit2
          fmt.Println("done")
      }
      ```


参考：https://zhuanlan.zhihu.com/p/338985480

