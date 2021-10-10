package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Computer struct {
	MemoLoad float64 `json:"memory_load"`
	CPULoad  float64 `json:"cpu_load"`
}

type JsonLogger struct {
	Time         string     `json:"time"`
	MainComputer Computer   `json:"main"`
	SubComputer  []Computer `json:"sub"`
}

func (j JsonLogger) Log() {
	j.Time = time.Now().Format("2006/01/02 15:04:05.99")
	json, err := json.Marshal(&j)
	if err != nil {
		panic(err)
	}
	j.writeJsonFormat("output.json", json)
}

func (j JsonLogger) writeJsonFormat(file string, obj []byte) {
	if len(obj) == 0 {
		return
	}
	f, _ := os.OpenFile(file, os.O_CREATE|os.O_RDWR, 0755)
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()
	fi, _ := f.Stat()
	length := fi.Size()
	fmt.Printf("位置=%d\n", length)
	if length == 0 {
		f.Write([]byte(fmt.Sprintf("[%s]", obj)))
	} else {
		f.WriteAt([]byte(fmt.Sprintf(",%s]", obj)), length-1)
	}
}

func main() {
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)

		main := Computer{
			MemoLoad: float64(i) * 0.9,
			CPULoad:  float64(i) * 0.7,
		}
		sub := []Computer{
			{

				MemoLoad: float64(i) * 0.9,
				CPULoad:  float64(i) * 0.7,
			},
			{

				MemoLoad: float64(i) * 0.9,
				CPULoad:  float64(i) * 0.7,
			},
			{

				MemoLoad: float64(i) * 0.9,
				CPULoad:  float64(i) * 0.7,
			},
			{

				MemoLoad: float64(i) * 0.9,
				CPULoad:  float64(i) * 0.7,
			},
			{

				MemoLoad: float64(i) * 0.9,
				CPULoad:  float64(i) * 0.7,
			},
		}
		jl := JsonLogger{MainComputer: main, SubComputer: sub}
		jl.Log()
	}
}
