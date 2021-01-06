package main

import (
	"os"
	"path/filepath"
	"time"
	"fmt"
	"strings"
)

func FindInterestingFiles(paths []string) ([]string, error){
	var interestingFiles []string
	// TODO Probably should replace this with some kind of external data store
	gtfobinsSuidFiles := []string{
		"aria2c","arp","ash","base64",
		"basenc","busybox","capsh","chmod","chown",
		"column","comm","csh","csplit","cut","dash",
		"dd","dialog","dmsetup","docker","env","eqn",
		"expect","file","flock","fmt","gdb","gimp",
		"gtester","hd","hexdump","highlight","iconv",
		"install","ip","jjs","jq","jrunscript","ksshell",
		"ld.so","logsave","look","lwp-request","make","mv",
		"nano","nl","node","od","openssl","perl","pg","pico",
		"pr","readelf","restic","rlwrap","rpm","rsync",
		"run-parts","rvim","sed","shuf","soelim","ss","ssh-keyscan",
		"stdbuf","strace","sysctl","systemctl","tail","taskset","tclsh",
		"tee","time","timeout","ul","unexpand","unshare",
		"update-alternatives","uuencode","view","watch","wget","xmodmap",
		"xxd","zsh","zsoelim"}
	for _, path := range paths {
		for _, gtfobin := range gtfobinsSuidFiles {
			pathsSplit := strings.Split(path, "/")
			filename := pathsSplit[len(pathsSplit) -1]
			if filename == gtfobin{
				interestingFiles = append(interestingFiles, path)
			}
		}
	}
	return interestingFiles, nil
}

func CheckDirectoryForSuid(fwriter *FindingsWriter, root string){
	fmt.Printf(ColorFmt("- - - - - - - - - - - -\nStarting SUID Scan In %s Directory\n- - - - - - - - - - - -\n", "Purple"), root)
	start := time.Now()
	var findings []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error{
		if info != nil {
			if info.Mode()&os.ModeSetuid != 0 == true {
				findings = append(findings, path)
			}
		}
		return nil
	})
	check(err)
	interestingFiles, err := FindInterestingFiles(findings)
	if err != nil{
		panic(err)
	}
	fwriter.Output(findings, "All Files With SUID Set", "Green")
	fwriter.Output(interestingFiles, "Interesting Files With SUID Set", "Yellow")
	searchTime := []string{time.Since(start).String()}
	fwriter.Output(searchTime, "Search Time", "Red")
}
