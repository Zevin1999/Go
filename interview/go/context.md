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