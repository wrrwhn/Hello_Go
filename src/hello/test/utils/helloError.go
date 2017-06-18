package utils

import (
	"log"
)

func HelloError() {

	defer func() {
		log.Println("defer.init")
		if err := recover(); nil != err {
			log.Println("defer-recover-error:", err)
		}
		log.Println("defer.end")
	}()
	caller()
}

func caller() {
	log.Println("caller-1")
	panic("caller-error")
	log.Println("caller-2")
	log.Println("caller-3")
}
