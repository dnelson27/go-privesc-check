package main

import (
	"os"
	"fmt"
)

type FindingsWriter struct{
	output_file string
	output_type string
}

// 	fwriter.Output(findings: interestingFiles, header: "Interesting Files With SUID Set". headerColor: "Yellow")
func (fw *FindingsWriter) Output(findings []string, header string, headerColor string){
	filepath := fw.output_file
	output_type := fw.output_type

	switch output_type {
	case "file":
		fp, err := os.Create(filepath)
		check(err)
		defer fp.Close()
		fp.WriteString(ColorFmt(header, headerColor))
		for _, finding := range findings {
			fp.WriteString(finding + "\n")
		}
		fp.Sync()
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
	return string(colors[color] + text + colors["reset"])
}