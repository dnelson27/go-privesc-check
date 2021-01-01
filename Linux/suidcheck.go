package main

import (
	"os"
	"path/filepath"
)

func FindInterestingFiles(findings []string) ([]string, error){
	var interestingFiles []string
	// TODO Probably should replace this with some kind of external data store
	gtfobinsSuidFiles := []string{
		"FAKE_FILE","aria2c","arp","ash","base64",
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
	for _, finding := range findings {
		for _, gtfobin := range gtfobinsSuidFiles {
			if finding == gtfobin{
				interestingFiles = append(interestingFiles, finding)
			}
		}
	}
	return interestingFiles, nil
}


func CheckDirectoryForSuid(fwriter *FindingsWriter, root string){
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
}
