package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"log"
	"os"
	"time"
)

// GetNTPTime получает текущее время с NTP сервера.
func GetNTPTime(server string) (time.Time, error) {
	ntpTime, err := ntp.Time(server)
	if err != nil {
		return time.Time{}, err
	}
	return ntpTime, nil
}

// ConvertToLocalTime преобразует время в локальное время.
func ConvertToLocalTime(ntpTime time.Time) time.Time {
	return ntpTime.In(time.Local)
}

func main() {

	ntpTime, err := GetNTPTime("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		log.Printf("error while receiving time info: %v ", err)
		os.Exit(1)
	}

	localTime := ConvertToLocalTime(ntpTime)

	fmt.Println(localTime)
}
