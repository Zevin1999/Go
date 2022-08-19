package practice

import (
	"fmt"
	"sync"
)

// PrintNumberAndLetter 打印 12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728
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
