package main

import (
	"os"
	"path/filepath"
	"fmt"
)

func checkDirectoryForSuid(fwriter *FindingsWriter, root string){
	var findings []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error{
		if info != nil {
			if info.Mode()&os.ModeSetuid != 0 == true {
				findings = append(findings, fmt.Sprintf("%s has SUID set", path))
			}
		}
		return nil
	})
	check(err)
	fwriter.Output(findings)
}