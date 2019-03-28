package template

import (
	"fmt"
	"io/ioutil"
	"os"
	"syscall"
)

// ColorPrint windows cmd color
func ColorPrint(s string, i int) {
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	proc := kernel32.NewProc("SetConsoleTextAttribute")
	handle, _, _ := proc.Call(uintptr(syscall.Stdout), uintptr(i))
	fmt.Print(s)
	handle, _, _ = proc.Call(uintptr(syscall.Stdout), uintptr(7))
	CloseHandle := kernel32.NewProc("CloseHandle")
	CloseHandle.Call(handle)
}

// ChecksrcexisT check file exist
func ChecksrcexisT(s string) bool {
	f, err := os.Stat(s)
	if err != nil {
		ReD(s + " is not exist !")
		os.Exit(0)
	}
	if f.IsDir() {
		return true
	}
	return false
}

// CheckdstexisT check file exist
func CheckdstexisT(s string) (bool, bool) {
	// var pd string
	var dstexist = false
	var dstlx = false
	f, err := os.Stat(s)
	if f.IsDir() {
		dstlx = true
	}
	if err == nil {
		dstexist = true
	}
	return dstexist, dstlx
}

// CheckdstfileexisT check file exist
func CheckdstfileexisT(s, d string) bool {
	// var pd string
	var dstfileexist = false
	List, _ := ioutil.ReadDir(d)
	for _, v := range List {
		if v.Name() == s {
			dstfileexist = true
		}
	}

	return dstfileexist
}

// GreeN color
func GreeN(s string) {
	ColorPrint(s, 2|8)
}

// ReD color
func ReD(s string) {
	ColorPrint(s, 4|4)
}

// YelloW color
func YelloW(s string) {
	ColorPrint(s, 12|12)
}

// BluE color
func BluE(s string) {
	ColorPrint(s, 2|1)
}
