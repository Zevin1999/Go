context包的理解和用途
 1. golang无协程ID,如何关闭创建的子孙协程呢? 
    方式: chan ; context 
 2. context接口提供方法及作用? 
    type Context interface {
         Done() <-chan struct{} // 当Context 被 canceled 或是 times out 的时候，Done 返回一个被 closed 的channel
         Err() error // 在 Done 的 channel被closed 后， Err 代表被关闭的原因
         Deadline() (deadline time.Time, ok bool) // 如果存在，Deadline 返回Context将要关闭的时间
         Value(key interface{}) interface{} // 如果存在，Value 返回与 key 相关了的值，不存在返回 nil 
    } 
 3. context有那些后代?
   func WithCancel(parent Context) (ctx Context, cancel CancelFunc) //WithCancel对应的是cancelCtx，返回一个 cancelCtx和一个CancelFunc，CancelFunc是 context包中定义的一个函数类型：type CancelFunc func()。调用这个CancelFunc 时，关闭对应的c.done，也就是让他的后代goroutine退出。 2.func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc) 3.func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)//WithDeadline和WithTimeout对应的是timerCtx ，WithDeadline和WithTimeout是相似的，WithDeadline 是设置具体的deadline时间，到达deadline的时候，后代 goroutine退出，而WithTimeout简单粗暴，直接return WithDeadline(parent, time.Now().Add(timeout))。 4.func WithValue(parent Context, key interface{}, val interface{}) Context // WithValue对应valueCtx，WithValue是在Context 中设置一个map，拿到这个Context以及它的后代的goroutine 都可以拿到map里的值。 
 4. 如何实现一个cancelCtx? 
    父代保存子代的cancel函数组成的map(或者list) 当WithCancel生成子代时将子代cancel函数增加到父代map(或者list)中 当父代调用cancel时候遍历map(或者list)调用所有子代的cancel函数


### context

+ context作用：在 Goroutine 构成的树形结构中对信号进行同步以减少计算资源的浪费

  （goroutine 之间传递上下文信息，包括：取消信号、超时时间、截止时间、k-v 等）

+ context.go整体梳理
```
  |      名称       |  类型  |                           功能                           |
    | :-------------: | :----: | :------------------------------------------------------: |
  |     Context     |  接口  |       定义Deadline() Done() Err() Value()四个方法        |
  |    canceler     |  接口  |         context取消接口，定义cancel() Done()方法         |
  |    emptyCtx     |  int   |              实现Context接口, 实际上是空的               |
  |    cancelCtx    | 结构体 |                        可以被取消                        |
  |    timerCtx     | 结构体 |                         超时取消                         |
  |    valueCtx     | 结构体 |                存储k-v对、协程间传递数据                 |
  |   WithCancel    |  函数  |         基于父context，生成一个可以取消的context         |
  |  WithDeadline   |  函数  |               创建一个有deadline的context                |
  |   WithTimeout   |  函数  |                创建一个有timeout的context                |
  |    WithValue    |  函数  |                创建一个存储k-v对的context                |
  |   Background    |  函数  |            返回一个空context，常作为根context            |
  |      TODO       |  函数  | 返回一个空context，常用于重构时期，没有合适的context可用 |
  |   CancelFunc    |  函数  |                         取消函数                         |
  |  newCancelCtx   |  函数  |                 创建一个可取消的context                  |
  | propagateCancel |  函数  |             向下传递context节点间的取消关系              |
  | parentCancelCtx |  函数  |                 找到第一个可取消的父节点                 |
  |   removeChild   |  函数  |                   移除父节点的子上下文                   |
  |      init       |  函数  |                                                          |

```

+



参考：

https://zhuanlan.zhihu.com/p/68792989

https://www.codenong.com/cs106949500/

https://www.cnblogs.com/yjf512/p/10399190.html

https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-context/

