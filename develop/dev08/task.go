package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

/*
=== Взаимодействие с ОС ===

# Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
func exec(c string) error {
	c = strings.Replace(c, "\r", "", 1)
	c = strings.Replace(c, "\n", "", 1)
	allc := strings.Split(c, " ")
	l := len(allc)
	if l < 1 {
		return errors.New("Неизвестная команда")
	}
	switch allc[0] {
	case "exit":
		os.Exit(0)
	case "cd":
		if l != 2 {
			return errors.New("Нужен лишь один аргумент - путь")
		}
		move(allc[1])
		break
	case "pwd":
		if l != 1 {
			return errors.New("Не должно быть аргументов")
		}
		fmt.Println(urPath())
	case "echo":
		if l != 2 {
			return errors.New("Нужен лишь один аргумент - данные")
		}
		fmt.Println(allc[1])
	case "kill":
		if l != 2 {
			return errors.New("Нужен лишь один аргумент - pid")
		}
		pid, err := strconv.Atoi(allc[1])
		if err != nil {
			return errors.New("Невалидный аргумент")
		}
		err = kill(pid)
		if err != nil {
			return err
		}
	case "ps":
		if l != 1 {
			return errors.New("Не должно быть аргументов")
		}
		list, err := allproc()
		if err != nil {
			return err
		}
		for _, v := range list {
			fmt.Print(v)
		}
	default:
		return fmt.Errorf("Неизвестная команда: %s", allc[0])
	}
	return nil
}

// only Linux
func allproc() ([]string, error) {
	matches, err := filepath.Glob("proc/*/exe")
	if err != nil {
		return nil, err
	}
	var allp []string
	for _, file := range matches {
		target, err := os.Readlink(file)
		if err != nil {
			return allp, err
		}
		if len(target) > 0 {
			allp = append(allp, target)
		}
	}
	return allp, nil
}

func kill(pid int) error {
	p, err := os.FindProcess(pid)
	if err != nil {
		return err
	}
	err = p.Kill()
	if err != nil {
		return err
	}
	return nil
}
func urPath() string {
	path, err := os.Getwd()
	if err != nil {
		return ""
	}
	return path
}
func move(path string) {
	err := os.Chdir(path)
	if err != nil {
		return
	}
}

func main() {
	r := bufio.NewReader(os.Stdin)
	for {
		c, err := r.ReadString('\r')
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}
		err = exec(c)
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}
	}
}
