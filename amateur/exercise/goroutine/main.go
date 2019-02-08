package main

import (
	"fmt"
	"time"
)

func buyGkassesMinimart() { //=> ซื้อแว่นตา
	time.Sleep(1 * time.Second)
	fmt.Println("ซื้อแว่นตา ที่มินิมาท")
}
func buyWatchesCentral() { //=> ซื้อนาฬิกา
	time.Sleep(1 * time.Second)
	fmt.Println("ซื้อนาฬิกา ที่เซนทรัล")
}
func buyFruitsParagon() { //=> ซื้อผลไม้
	time.Sleep(1 * time.Second)
	fmt.Println("ซื้อผลไม้ ที่พารากอน")
}
func buyCarHonda() { //=> ซื้อรถ
	time.Sleep(1 * time.Second)
	fmt.Println("ซื้อรถ ที่ฮอนด้า")
}
func sendToA(message chan<- string) { //=> ส่งของให้A
	time.Sleep(1 * time.Second)
	message <- "กำลังส่งของให้A"
}

func main() {
	start := time.Now() //=> Start จับเวลา
	go buyGkassesMinimart()
	go buyWatchesCentral()
	ch := make(chan string)
	go sendToA(ch) //=> อะไรซื้อแล้วส่งให้A
	buyFruitsParagon()
	buyCarHonda()
	messageFromB := <-ch //=> รับของจากB
	if messageFromB == "กำลังส่งของให้A" {
		fmt.Println("Aได้รับของแล้ว")
		fmt.Println("total time: ", time.Since(start)) //=> ตัดเวลา
	}
}
func init() {
	fmt.Println("Hi, goroutine")
	fmt.Println("-----------------------------")
}
