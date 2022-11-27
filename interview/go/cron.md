### cron库

+ 定时任务库，主要用于定时备份系统数据、周期性清理缓存、定时重启服务等

+ linux 中可以通过 crontab -e 来配置定时任务。不过，linux 中的 cron 只能精确到分钟。而Go 实现的 cron 可以精确到秒，除了这点比较大的区别外，cron 表达式的基本语法是类似的。

+ 示例1：不同时间格式的使用

  ```
  package main
  
  import (
  	"fmt"
  	"github.com/robfig/cron"
  	"time"
  )
  
  func main() {
  	// 创建cron对象，此对象用于管理定时任务
  	c := cron.New()
  
  	// 向管理器中添加定时任务
  	// 每个小时的30分触发
  	c.AddFunc("30 * * * *", func() {
  		fmt.Println("Every hour on the half hour")
  	})
  	// 6、7、8、18、19、20时的30分触发
  	c.AddFunc("30 6-8,18-20 * * *", func() {
  		fmt.Println("On the half hour of 6-8am, 6-8pm")
  	})
  	// 1月2日3时4分触发
  	c.AddFunc("4 3 2 1 *", func() {
  		fmt.Println(" 3:4 Jun 2 every year")
  	})
  
  	c.AddFunc("@hourly", func() {
  		fmt.Println("Every hour")
  	})
  
  	c.AddFunc("@daily", func() {
  		fmt.Println("Every day on midnight")
  	})
  
  	c.AddFunc("@weekly", func() {
  		fmt.Println("Every week")
  	})
  
  	// 每一时两分三秒触发
  	c.AddFunc("@every 1h2m3s", func() {
  		fmt.Println("every 1 hour 2 minute 3 second")
  	})
  	// 东京时区每两秒触发
  	c.AddFunc("CRON_TZ=Asia/Tokyo @every 2s", func() {
  		fmt.Println("every 2 second")
  	})
  
  	// 启动新的goroutinue做循环检测
  	c.Start()
  	for {
  		time.Sleep(time.Second)
  	}
  }
  ```

+ 示例2：特殊字符/以及cron.WithSeconds()的使用

  ```
  package main
  
  import (
  	"github.com/robfig/cron"
  	"log"
  	"time"
  )
  
  func main() {
  	i := 0
  	// 每两秒执行一次
  	c := cron.New(cron.WithSeconds())
  	spec := "*/2 * * * * ?"
  	c.AddFunc(spec, func() {
  		i++
  		log.Println("cron running:", i)
  	})
  	c.Start()
  
  	for {
  		time.Sleep(time.Second)
  	}
  }
  ```

+ 示例3：自实现Job接口

  ```
  package main
  
  import (
  	"fmt"
  	"github.com/robfig/cron"
  )
  
  type TestJob struct {
  	id   int
  	name string
  }
  
  func (t TestJob) Run() {
  	fmt.Printf("TestJob: %d is %s\n", t.id, t.name)
  }
  
  func main() {
  	c := cron.New()
  	// AddJob方法
  	c.AddJob("@every 2s", TestJob{
  		id:   999,
  		name: "感冒灵",
  	})
  	// 启动任务
  	c.Start()
  	select {}
  }
  ```

+ 示例4：cron.WithLogger()的使用

  ```
  package main
  
  import (
  	"fmt"
  	"github.com/robfig/cron"
  	"log"
  	"os"
  )
  
  func main() {
  	c := cron.New(cron.WithLogger(
  		cron.VerbosePrintfLogger(
  			log.New(os.Stdout, "log: ", log.LUTC))))
  	c.AddFunc("@every 2s", func() {
  		fmt.Println("every 2 second")
  	})
  	c.Start()
  	select {}
  }
  ```

​	参考：https://jishuin.proginn.com/p/763bfbd2c3a3

​               https://segmentfault.com/a/1190000023029219

