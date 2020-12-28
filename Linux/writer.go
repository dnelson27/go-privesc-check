package main

import (
	"os"
	"fmt"
)

type FindingsWriter struct{
	output_file string
	output_type string
}

func (fw *FindingsWriter) Output(findings []string){
	filepath := fw.output_file
	output_type := fw.output_type

	switch output_type {
	case "file":
		fp, err := os.Create(filepath)
		check(err)
		defer fp.Close()
		for _, finding := range findings {
			fp.WriteString(finding + "\n")
		}
		fp.Sync()
	case "stdout":
		for _, finding := range findings {
			fmt.Println(finding)
		}
	default: // Should never fire, input validation in main
		for _, finding := range findings {
			fmt.Println(finding)
		}
	}

}
