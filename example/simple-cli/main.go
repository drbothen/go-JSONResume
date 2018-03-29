package main

import (
	"flag"
	"repo.theoremforge.com/me/go-JSONResume"
	"fmt"
)

func main()  {
	// register cli args
	rStingPtr := flag.String("resume", "C:\\jsonresume.json", "A json file using the jsonresume schema")

	// Parse incoming cli args
	flag.Parse()

	resume, err := JSONResume.LoadResumeFile(*rStingPtr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resume)
}
