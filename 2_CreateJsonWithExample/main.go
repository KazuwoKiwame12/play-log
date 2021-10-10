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

type LogContent struct {
	Time         string     `json:"time"`
	MainComputer Computer   `json:"main"`
	SubComputer  []Computer `json:"sub"`
}

type JsonLogger struct {
	Content LogContent
	Output  *os.File
}

func (j *JsonLogger) Log() {
	j.Content.Time = time.Now().Format("2006/01/02 15:04:05.99")
	json, err := json.Marshal(&j.Content)
	if err != nil {
		panic(err)
	}
	j.writeJsonFormat(json)
}

func (j *JsonLogger) writeJsonFormat(obj []byte) {
	if len(obj) == 0 {
		return
	}

	fi, _ := j.Output.Stat()
	length := fi.Size()
	fmt.Printf("位置=%d\n", length)
	if length == 0 {
		j.Output.Write([]byte(fmt.Sprintf("[%s]", obj)))
	} else {
		j.Output.WriteAt([]byte(fmt.Sprintf(",%s]", obj)), length-1)
	}
}

func main() {
	f, _ := os.OpenFile("output.json", os.O_CREATE|os.O_RDWR, 0755)
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()
	jsonLogger := &JsonLogger{Output: f}

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
		jsonLogger.Content = LogContent{MainComputer: main, SubComputer: sub}
		jsonLogger.Log()
	}
}
