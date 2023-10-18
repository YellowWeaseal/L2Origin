package main

import (
	"testing"
	"time"
)

func TestGetNTPTime(t *testing.T) {
	ntpTime, err := GetNTPTime("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		t.Errorf("Ошибка при получении времени NTP: %v", err)
	}
	if ntpTime.IsZero() {
		t.Error("Получено нулевое время NTP")
	}
}

func TestConvertToLocalTime(t *testing.T) {
	ntpTime := time.Date(2023, 9, 22, 12, 0, 0, 0, time.UTC)
	localTime := ConvertToLocalTime(ntpTime)
	currentTime := time.Now()
	timeDiff := localTime.Sub(currentTime)
	if timeDiff.Seconds() > 1 {
		t.Errorf("Ошибка в преобразовании времени, разница больше 1 секунды: %v", timeDiff)
	}
}
