package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"log"
	"time"
)

func main() {
	time, err := GetTime()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(time.Format("02-01-2006 15:04:05"))
}

//Функция предоставляет текущее время с удаленного сервера через NTP
func GetTime()(time.Time, error){
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	return time, err
}