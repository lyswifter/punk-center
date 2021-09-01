package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

const repoPath = "~/.punkcenter"

var AllIps = []string{}

var lock sync.Mutex
var received []string = []string{}

func main() {
	log.Print("This is the punk center server")

	DataStores()

	go func() {
		AllIps = readline("ips")

		ticker := time.NewTicker(5 * time.Second)

		for {
			select {
			case <-ticker.C:
				var notReceived []string
				for _, ip := range AllIps {
					isHas := false
					for _, r := range received {
						if ip == r {
							isHas = true
							break
						}
					}

					if isHas {
						continue
					}

					notReceived = append(notReceived, ip)
				}
				log.Printf("notReceived(%d) ips: %v", len(notReceived), notReceived)
			}
		}
	}()

	SetupServers()
}

func readline(path string) []string {
	f, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer f.Close()

	rd := bufio.NewReader(f)

	var ret = []string{}
	for {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行

		if err != nil || io.EOF == err {
			break
		}

		line = strings.Replace(line, "\n", "", -1)

		ret = append(ret, line)
	}

	return ret
}
