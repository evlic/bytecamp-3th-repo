package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

const (
	maxNum = 100
	retry  = 3
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func getInputInt(maxRetry int) (res int, err error) {
	// fmt.Println()
	for maxRetry > 0 {
		fmt.Printf("\r请您输入一个整数( <= %d): ", maxNum)
		_, err = fmt.Scanf("%d\r\n", &res)
		if err != nil {
			log.Println(err)
			maxRetry--
			continue
		}
		return
	}
	return
}

func main() {
	secretNumber := rand.Intn(maxNum)
	for {
		input, err := getInputInt(retry)
		if err != nil {
			fmt.Println("错误的输入", err)
			continue
		}
		if input == secretNumber {
			fmt.Println("猜中了! >> ", input)
			break
		}

		if input > secretNumber {
			fmt.Println("猜大了")
		} else {
			fmt.Println("猜小了")
		}
		continue
	}

}
