package practice

import (
	"fmt"
	"sync"
)

func PrintNumberAndLetter() {
	number, letter := make(chan bool), make(chan bool)
	wait := sync.WaitGroup{}

	go func() {
		num := 1
		for {
			select {
			case <-number:
				fmt.Println(num)
				num++
				fmt.Println(num)
				num++
				letter <- true
			}
		}
	}()

	wait.Add(1)
	go func(wait *sync.WaitGroup) {
		let := 'A'
		for {
			select {
			case <-letter:
				if let >= 'Z' {
					wait.Done()
					return
				}
				fmt.Println(let)
				let++
				fmt.Println(let)
				let++
				number <- true
			}
		}
	}(&wait)
	// 启动数字打印
	number <- true
	// 等待字母打印完成
	wait.Wait()
}
