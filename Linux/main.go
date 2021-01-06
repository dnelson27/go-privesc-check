package main

import (
	"os"
	"fmt"
	"bufio"
)

func check(e error) {
	if e != nil {
		fmt.Println(ColorFmt("ERROR!", "Red"))
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
		// If file exists, truncate
		if _, err := os.Stat(output_file); !os.IsNotExist(err) {
			existingFile, err := os.Open(output_file)
			check(err)
			err = os.Truncate(output_file, 0)
			check(err)
			existingFile.Close()
		}
		fp, err := os.OpenFile(output_file, os.O_WRONLY|os.O_CREATE , 0700)
		check(err)
		// Create os.File and bufio.Writer to close when main returns
		defer fp.Close()
		w := bufio.NewWriter(fp)
		defer w.Flush()
		// Create writer object for functions to use for output
		fwriter := FindingsWriter{output_type: "file", buffered_writer: w}
		CheckDirectoryForSuid(&fwriter, check_dir)

	} else if output_type == "stdout" {
		fwriter := FindingsWriter{output_type: "stdout"}
		CheckDirectoryForSuid(&fwriter, check_dir)

	} else {
		printhelp("No Valid Output Method")
		os.Exit(0)
	}
	fmt.Println(ColorFmt("\u2713 Done! \u2713", "Green"))
}