package service

import (
	"io"
	"log"
	"os"
)

func setLog() {
	f, err := os.OpenFile("log.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	// defer f.Close()
	multiWriter := io.MultiWriter(os.Stderr, f)
	log.SetOutput(multiWriter)
	//打印日志时间，文件名，行号
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}
