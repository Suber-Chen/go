package main

import (
	"log"
	"os"
)

func init() {
	// fmt.Printf("log.Flags(): %v\n", log.Flags())
	// log.Print("111")
	// log.SetFlags(1)
	// fmt.Printf("log.Flags(): %v\n", log.Flags())
	// log.Print("222")
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	// fmt.Printf("log.Flags(): %v\n", log.Flags())
	// log.Print("333")
	f, err := os.OpenFile("test.log", os.O_RDWR|os.O_APPEND|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Panic("日志文件异常")
	}
	log.SetOutput(f)
}
func main() {

	log.Print("111111111111")
}
