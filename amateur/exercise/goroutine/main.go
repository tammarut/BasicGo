package main

import (
	"fmt"
	"time"
)

func buyGkassesMinimart() {
	time.Sleep(1 * time.Second)
	fmt.Println("ซื้อแว่นตา ที่มินิมาท")
}
func buyWatchesCentral() {
	time.Sleep(1 * time.Second)
	fmt.Println("ซื้อนาฬิกา ที่เซนทรัล")
}
func buyFruitsParagon() {
	time.Sleep(1 * time.Second)
	fmt.Println("ซื้อผลไม้ ที่พารากอน")
}
func buyCarHonda() {
	time.Sleep(1 * time.Second)
	fmt.Println("ซื้อรถ ที่ฮอนด้า")
}
func sendToA(message chan<- string) {
	time.Sleep(1 * time.Second)
	message <- "กำลังส่งของให้A"
}

func main() {
	start := time.Now()
	go buyGkassesMinimart()
	go buyWatchesCentral()
	ch := make(chan string)
	go sendToA(ch)
	buyFruitsParagon()
	buyCarHonda()
	messageFromB := <-ch
	if messageFromB == "กำลังส่งของให้A" {
		fmt.Println("Aได้รับของแล้ว")
		fmt.Println("total time: ", time.Since(start))
	}
}
func init() {
	fmt.Println("Hi, goroutine")
	fmt.Println("-----------------------------")
}
