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

	fmt.Println(ColorFmt(err, "Red"))
	fmt.Println(
		ColorFmt("Linux PrivEsc Scanner ", "Green") + ColorFmt("- In Development - \n", "Purple") +
		ColorFmt("USAGE: ", "Yellow") + "./suidcheck target-directory {stdout | file} [output file]\n\n" +
		"This is a work in progress that aims to achieve some of the basic functionality found in tools\nlike https://github.com/carlospolop/privilege-escalation-awesome-scripts-suite in binary form\n")
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
		CheckDirectoryForSuid(&fwriter, check_dir)
	} else if output_type == "stdout" {
		fwriter := FindingsWriter{output_type: "stdout"}
		CheckDirectoryForSuid(&fwriter, check_dir)
	} else {
		printhelp("No Output Method")
	}
}