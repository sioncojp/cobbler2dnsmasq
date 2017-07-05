package main

import (
	"flag"
	"log"
	"os"

	"github.com/sioncojp/cobbler2dnsmasq"
)

const (
	app = "c2d"
)

// init ...flagを初期化する
func init() {
	flag.StringVar(&c2d.ConfigFile, "c", "", "configファイル読み込み")
	flag.Parse()
}

func main() {
	log.SetOutput(os.Stderr)
	log.SetPrefix(app + ": ")

	// dnsmasq.template作成
	err := c2d.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
