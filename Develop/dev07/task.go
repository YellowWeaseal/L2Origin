package main

import (
	"fmt"
	"time"
)

func or(channels ...<-chan interface{}) <-chan interface{} {
	// Создаем канал для объединенного результата
	result := make(chan interface{})

	// Запускаем горутину, которая будет слушать закрытие каналов
	go func() {
		defer close(result)

		// Создаем канал для синхронизации
		done := make(chan struct{})

		// Запускаем отдельную горутину для каждого канала
		for _, ch := range channels {
			go func(ch <-chan interface{}) {
				select {
				case <-ch:
					// Канал закрыт, сигнализируем через done
					done <- struct{}{}
				case <-done:
					// Другой канал уже был закрыт, игнорируем
				}
			}(ch)
		}

		// Ожидаем закрытия одного из каналов
		<-done
	}()

	return result
}

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(5*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf("Done after %v", time.Since(start))
}
