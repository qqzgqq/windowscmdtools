package template

import (
	"fmt"
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
