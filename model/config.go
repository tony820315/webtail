package model

import (
	"flag"
	"fmt"
)

type ArgsConfig struct {
	Path string
}

var Args = &ArgsConfig{}

func init() {
	flag.StringVar(&Args.Path, "p", "", "input the file absolute path")
	flag.Parse()

	fmt.Println(Args)
}
