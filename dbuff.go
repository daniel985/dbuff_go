package main

import (
	"os"
	"time"
	"fmt"
	"io"
	"bufio"
	"strings"
)

type Dbuff struct {
	Path string
	Stamp int64
	Data map[string]string
}

func (df *Dbuff) Load() {
	df.Data = make(map[string]string)
	f, _ := os.Stat(df.Path)
	df.Stamp = f.ModTime().Unix()
	rf, _ := os.Open(df.Path)
	defer rf.Close()
	buff := bufio.NewReader(rf)
	for {
		line, err := buff.ReadString('\n')
		if err != nil || err == io.EOF {
			break
		}
		line = strings.TrimSpace(line)
		splits := strings.Split(line, "\t")
		df.Data[splits[0]] = splits[1]
	}
}

func (df *Dbuff) Reload() {
	for {
		f, _ := os.Stat(df.Path)
		newstamp := f.ModTime().Unix()
		if newstamp > df.Stamp {
			fmt.Printf("Reload data!!!\n")
			newdata := make(map[string]string)
			rf, _ := os.Open(df.Path)
			defer rf.Close()
			buff := bufio.NewReader(rf)
			for {
				line, err := buff.ReadString('\n')
				if err != nil || err == io.EOF {
					break
				}
				line = strings.TrimSpace(line)
				splits := strings.Split(line, "\t")
				newdata[splits[0]] = splits[1]
			}
			df.Data = newdata
			df.Stamp = newstamp
			df.Print()
		} else {
			time.Sleep(100000000)
		}
	}
}

func (df *Dbuff) Print() {
	for k,v := range df.Data {
		fmt.Printf("k=%s, v=%s\n", k, v)
	}
}

func server() {
	for {
		fmt.Printf("Server is working\n")
		time.Sleep(1000000000)
	}
}

func main() {
	dbf := new(Dbuff)
	dbf.Path = "./kv.txt"
	dbf.Load()
	go dbf.Reload()
	server()
}
