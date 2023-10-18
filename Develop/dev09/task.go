package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	// Получение URL и имени файла из аргументов командной строки
	urlFlag := flag.String("url", "", "URL для скачивания")
	outputFlag := flag.String("output", "output.html", "Имя файла для сохранения")
	flag.Parse()

	// Проверка наличия URL
	if *urlFlag == "" {
		fmt.Println("Пожалуйста, укажите URL для скачивания.")
		return
	}

	// Проверка корректности URL
	parsedURL, err := url.Parse(*urlFlag)
	if err != nil || (parsedURL.Scheme != "http" && parsedURL.Scheme != "https") {
		fmt.Println("Некорректный URL. Пожалуйста, укажите URL с префиксом http:// или https://.")
		return
	}

	// Выполнение HTTP-запроса
	resp, err := http.Get(*urlFlag)
	if err != nil {
		fmt.Println("Ошибка при выполнении HTTP-запроса:", err)
		return
	}
	defer resp.Body.Close()

	// Чтение данных из ответа
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении данных:", err)
		return
	}

	// Сохранение данных в файл
	err = ioutil.WriteFile(*outputFlag, body, 0644)
	if err != nil {
		fmt.Println("Ошибка при сохранении данных в файл:", err)
		return
	}

	fmt.Printf("Сайт успешно скачан и сохранен в файл %s\n", *outputFlag)
}
