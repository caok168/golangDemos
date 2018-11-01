package main

import (
	"flag"
	"fmt"
)

func main() {
	username := flag.String("name", "", "Input your username")
	flag.Parse()
	fmt.Println("Hello, ", *username)

	data_path := flag.String("D","/home/manu/sample","DB data path")
	log_file := flag.String("l","/home/manu/sample.log","log file")
	nowait_flag := flag.Bool("W",false,"do not wait until operation completes")
	flag.Parse()

	var cmd string = flag.Arg(0)
	fmt.Println("action   : %s\n",cmd)
	fmt.Println("data path: %s\n",*data_path)
	fmt.Println("log file : %s\n",*log_file)
	fmt.Println("nowait  :%s\n",*nowait_flag)

	fmt.Println("------------------------------------------------------------------")

	fmt.Println("there are %d non-flag input param\n",flag.NArg())
	for i,param := range flag.Args(){
		fmt.Printf("#%d    :%s\n",i,param)
	}
}
