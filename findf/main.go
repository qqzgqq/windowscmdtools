package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"syscall"
)

// Dirnamefile the file name
var Dirnamefile string
var receive string

// SearchString DIR names to []string
var SearchString []string

// SEARCHNAME search string
var SEARCHNAME *string

// SEARCHDIR search dir
var SEARCHDIR *string

// HeLP1 -h
var HeLP1 *bool

// HeLP1 -help
var HeLP2 *bool

func init() {

	HeLP1 = flag.Bool("h", false, "a bool")
	HeLP2 = flag.Bool("help", false, "a bool")
	SEARCHNAME = flag.String("name", receive, "search string")
	SEARCHDIR = flag.String("dir", receive, "search dir")

}

// Helpinfo help info
func Helpinfo(h, help *bool) {
	var receive2 = flag.Arg(0)
	if receive2 != "" {
		fmt.Println("parameter error,please check `findnew -h`")
		os.Exit(0)
	}
	if *h || *help {
		YelloW("usage:   findnew -name  [string] -dir [string]\n")
		fmt.Printf("   -h      The tool findnew help info  \n")
		fmt.Printf("   -help   The tool findnew help info \n")
		fmt.Printf("   -name   the string will be search\n")
		fmt.Printf("   -dir    the dir for search\n")
		YelloW("eg:     findnew -name  a.txt -dir d:\\  \n")
		os.Exit(0)
	}
}

// INTODIR for checkout the dir to get the string
func INTODIR(DIR string) {

	Dirlist, _ := ioutil.ReadDir(DIR)

	for _, v := range Dirlist {

		Dirnamefile = DIR + "\\" + v.Name()
		SearchString = append(SearchString, Dirnamefile)
		if v.IsDir() == true {
			INTODIR(Dirnamefile)
		}
	}
}

// CHEckouT *string to string
func CHEckouT(s *string) string {
	return *s
}

// ColorPrint windows cmd color
func ColorPrint(s string, i int) {
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	proc := kernel32.NewProc("SetConsoleTextAttribute")
	handle, _, _ := proc.Call(uintptr(syscall.Stdout), uintptr(i))
	fmt.Println(s)
	handle, _, _ = proc.Call(uintptr(syscall.Stdout), uintptr(7))
	CloseHandle := kernel32.NewProc("CloseHandle")
	CloseHandle.Call(handle)
}

// GreeN color
func GreeN(s string) {
	ColorPrint(s, 2|8)
}

// YelloW color
func YelloW(s string) {
	ColorPrint(s, 12|12)
}
func main() {

	flag.Parse()
	Helpinfo(HeLP1, HeLP2)
	SEARCHNAME := CHEckouT(SEARCHNAME)
	SEARCHDIR := CHEckouT(SEARCHDIR)
	//if DIr is in a-z or A-Z then set DIr:\\
	if m, _ := regexp.MatchString("^[a-z]$|^[A-Z]$", SEARCHDIR); m {
		SEARCHDIR = SEARCHDIR + ":\\"
	}
	if SEARCHDIR == "" {
		SEARCHDIR = "."
	}

	INTODIR(SEARCHDIR)

	SSNum := len(SearchString)

	for i := 0; i < SSNum; i++ {
		if m, _ := regexp.MatchString(SEARCHNAME, SearchString[i]); m {
			GreeN(SearchString[i])
		}

	}

}
