package main

import (
	"fmt"
	"bufio"
)

type FindingsWriter struct{
	output_type string
	buffered_writer *bufio.Writer
}


func (fw *FindingsWriter) Output(findings []string, header string, headerColor string){
	output_type := fw.output_type
	buffered_writer := fw.buffered_writer
	switch output_type {
	case "file":
		buffered_writer.WriteString("- - - - - - - - - - - -\n")
		buffered_writer.WriteString(header + "\n")
		for _, i := range findings{
			buffered_writer.WriteString(i + "\n")
		}
	case "stdout":
		fmt.Println(ColorFmt(header, headerColor))
		for _, finding := range findings {
			fmt.Println(finding)
		}
	default: // Should never fire, input validation in main
		fmt.Println(ColorFmt(header, headerColor))
		for _, finding := range findings {
			fmt.Println(finding)
		}
	}
}

func ColorFmt(text, color string) string{
	colors := make(map[string]string)
	colors["Red"]    = "\033[31m"
	colors["Reset"]  = "\033[0m"
	colors["Green"]  = "\033[32m"
	colors["Yellow"] = "\033[33m"
	colors["Purple"] = "\033[35m"
	colors["Gray"]   = "\033[37m"
	return string(colors[color] + text + colors["Reset"])
}


func FileWrite(w *bufio.Writer, data []string){
	for _, i := range data {
		_, err := w.WriteString(i)
		check(err)
	}
}