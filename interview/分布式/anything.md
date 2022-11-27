#### 分布式id生成器

```
timestamp 
datacenter_id
worker_id
sequence_id
```

#### 分布式锁

go
// 进程内加锁
package main

import (
    "sync"
)

// 全局变量
var counter int
var wg sync.WaitGroup
var l sync.Mutex
for i := 0; i < 1000; i++ {
    wg.Add(1)
    go func() {
        defer wg.Done()
        l.Lock()
        counter++
        l.Unlock()
    }()
}

wg.Wait()
println(counter)

// trylock
package main

import (
    "sync"
)

type Lock struct {
    c chan struct{}
}

func NewLock() Lock {
    var l Lock
    l.c = make(chan struct{}, 1)
    l.c <- struct{}{}
    return l
}

func (l Lock) Lock() bool {
    lockResult := false
    select {
    case <-l.c:
        lockResult = true
    default:
    }
    return lockResult
}

func (l Lock) Unlock() {
    l.c <- struct{}{}
}

var counter int

func main() {
    var l = NewLock()
    var wg sync.WaitGroup
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            if !l.Lock() {
                println("lock failed")
                return
            }
            counter++
            println("current counter", counter)
            l.Unlock()
        }()
    }
    wg.Wait()
}

// 高并发场景下
// 基于Redis的setnx
package main

import (
    "fmt"
    "sync"
    "time"

    "github.com/go-redis/redis"
)

func incr() {
    client := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })

    var lockKey = "counter_lock"
    var counterKey = "counter"

    // lock
    resp := client.SetNX(lockKey, 1, time.Second*5)
    lockSuccess, err := resp.Result()

    if err != nil || !lockSuccess {
        fmt.Println(err, "lock result: ", lockSuccess)
        return
    }

    // counter ++
    getResp := client.Get(counterKey)
    cntValue, err := getResp.Int64()
    if err == nil || err == redis.Nil {
        cntValue++
        resp := client.Set(counterKey, cntValue, 0)
        _, err := resp.Result()
        if err != nil {
            // log err
            println("set value error!")
        }
    }
    println("current counter is ", cntValue)

    delResp := client.Del(lockKey)
    unlockSuccess, err := delResp.Result()
    if err == nil && unlockSuccess > 0 {
        println("unlock success!")
    } else {
        println("unlock failed", err)
    }
}

func main() {
    var wg sync.WaitGroup
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            incr()
        }()
    }
    wg.Wait()
}

// 分布式任务调度场景
// 基于ZooKeeper
```
package main

import (
    "time"

    "github.com/samuel/go-zookeeper/zk"
)

func main() {
    c, _, err := zk.Connect([]string{"127.0.0.1"}, time.Second)
    if err != nil {
        panic(err)
    }
    l := zk.NewLock(c, "/lock", zk.WorldACL(zk.PermAll))
    err = l.Lock()
    if err != nil {
        panic(err)
    }
    println("lock succ, do your business logic")

    time.Sleep(time.Second * 10)
    // do something
    l.Unlock()
    println("unlock succ, finish business logic")
}
```


// 分布式阻塞锁
// etcd
```
package main

import (
    "log"

    "github.com/zieckey/etcdsync"
)

func main() {
    // 检查/lock路径下是否有值
    m, err := etcdsync.New("/lock", 10, []string{"http://127.0.0.1:2379"})
    if m == nil || err != nil {
        log.Printf("etcdsync.New failed")
        return
    }
    // 加锁
    err = m.Lock()
    if err != nil {
        log.Printf("etcdsync.Lock failed")
        return
    }

    log.Printf("etcdsync.Lock OK")
    log.Printf("Get the lock. Do something here.")

    err = m.Unlock()
    if err != nil {
        log.Printf("etcdsync.Unlock failed")
    } else {
        log.Printf("etcdsync.Unlock OK")
    }
}
```

#### 延时任务系统

+ 分布式定时任务管理系统

+ 定时发送消息的消息队列

  ```
  // 定时器的实现
  // 常见的时间堆一般使用小顶堆实现   四叉堆
  // 时间轮  哈希表  触发时间%时间轮元素大小
  ```

#### 分布式搜索引擎

+ Elasticsearch
+ 倒排列表





#### 负载均衡

+ 基于洗牌算法

```
  func shuffle(slice []int) {
      for i := 0; i < len(slice); i++ {
          a := rand.Intn(len(slice))
          b := rand.Intn(len(slice))
          slice[a], slice[b] = slice[b], slice[a]
      }
  }
  
  // 设置种子
  rand.Seed(time.Now().UnixNano())
  ```

+ 修正洗牌算法

  ```
  func shuffle(indexes []int) {
      for i:=len(indexes); i>0; i-- {
          lastIdx := i - 1
          idx := rand.Int(i)
          indexes[lastIdx], indexes[idx] = indexes[idx], indexes[lastIdx]
      }
  }
  // Go内置
  rand.Perm()
  ```

#### 分布式配置管理

+ 应用场景：报表系统、业务配置

+ 例子

  ```
  package main
  
  import (
      "log"
      "time"
  
      "golang.org/x/net/context"
      "github.com/coreos/etcd/client"
  )
  
  var configPath =  `/configs/remote_config.json`
  var kapi client.KeysAPI
  
  type ConfigStruct struct {
      Addr           string `json:"addr"`
      AesKey         string `json:"aes_key"`
      HTTPS          bool   `json:"https"`
      Secret         string `json:"secret"`
      PrivateKeyPath string `json:"private_key_path"`
      CertFilePath   string `json:"cert_file_path"`
  }
  
  var appConfig ConfigStruct
  
  func init() {
      cfg := client.Config{
          Endpoints:               []string{"http://127.0.0.1:2379"},
          Transport:               client.DefaultTransport,
          HeaderTimeoutPerRequest: time.Second,
      }
  
      c, err := client.New(cfg)
      if err != nil {
          log.Fatal(err)
      }
      kapi = client.NewKeysAPI(c)
      initConfig()
  }
  
  func watchAndUpdate() {
      w := kapi.Watcher(configPath, nil)
      go func() {
          // watch 该节点下的每次变化
          for {
              resp, err := w.Next(context.Background())
              if err != nil {
                  log.Fatal(err)
              }
              log.Println("new values is ", resp.Node.Value)
  
              err = json.Unmarshal([]byte(resp.Node.Value), &appConfig)
              if err != nil {
                  log.Fatal(err)
              }
          }
      }()
  }
  
  func initConfig() {
      resp, err = kapi.Get(context.Background(), configPath, nil)
      if err != nil {
          log.Fatal(err)
      }
  
      err := json.Unmarshal(resp.Node.Value, &appConfig)
      if err != nil {
          log.Fatal(err)
      }
  }
  
  func getConfig() ConfigStruct {
      return appConfig
  }
  
  func main() {
      // init your app
  }
  ```



+ 配置膨胀、配置版本管理、客户端容错





在线交易处理（OLTP, Online transaction processing）

AST抽象语法树



















