package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

/*
=== Утилита wget ===

# Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
func wget(file, url string) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	data, err := os.Create(file)
	if err != nil {
		return err
	}
	defer func() {
		if data.Close(); err != nil {
			log.Fatal(err.Error())
		}
	}()
	w := bufio.NewWriter(data)
	if _, err := io.Copy(w, res.Body); err != nil {
		return err
	}
	return nil
}

func main() {
	var file, url string
	fmt.Scan(&file, &url)
	if err := wget(file, url); err != nil {
		log.Fatal(err.Error())
	}
}
