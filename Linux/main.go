package main

import (
	"os"
	"fmt"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func printhelp(err string){
	var Red    = "\033[31m"
	var Reset  = "\033[0m"
	var Green  = "\033[32m"
	var Yellow = "\033[33m"
	var Purple = "\033[35m"
	var Gray   = "\033[37m"
	fmt.Println(Red + err + Reset)
	fmt.Println(
		Green + "Linux PrivEsc Scanner "  + Purple + "- In Development - \n" + Reset +
		Yellow + "USAGE: " + Reset + "./suidcheck target-directory {stdout | file} [output file]\n\n" +
		Gray + "This is a work in progress that aims to achieve some of the basic functionality found in tools\nlike https://github.com/carlospolop/privilege-escalation-awesome-scripts-suite in binary form\n")
}

func main(){
	if len(os.Args) < 3{
		printhelp("Invalid Arguments")
		os.Exit(0)
	}

	check_dir, output_type := os.Args[1], os.Args[2]
	if (output_type == "file"){
		output_file := os.Args[3]
		fwriter := FindingsWriter{output_file:  output_file, output_type: "file"}
		checkDirectoryForSuid(&fwriter, check_dir)
	} else if output_type == "stdout" {
		fwriter := FindingsWriter{output_type: "stdout"}
		checkDirectoryForSuid(&fwriter, check_dir)
	} else {
		printhelp("No Output Method")
	}
}