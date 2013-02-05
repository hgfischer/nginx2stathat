package main

import (
	"./loghit"
	"fmt"
	"github.com/ActiveState/tail"
)

func main() {
	t, err := tail.TailFile("/Users/herbert/Rocket/logs/www.braprint.dev.access.log", tail.Config{Follow: true})
	if err != nil {
		panic(err)
	}

	for line := range t.Lines {
		logHit, err := loghit.New(line.Text)
		if err != nil {
			panic(err)
		}
		fmt.Println(logHit)
	}
}
