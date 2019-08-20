package cpall

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// CopyF cy file
func CopyF(src, dst string) (int64, error) {

	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

// MKFDIR mkdir first dir
func MKFDIR(src, dst string) {
	LL, serr := os.Stat(src)
	if serr != nil {
		fmt.Print(src + " is not exist\n")
		os.Exit(0)
	}

	aerr := os.MkdirAll(dst, LL.Mode())
	if aerr != nil {
		// fmt.Println("create  " + dst + " bad\n")
		os.Exit(0)
	}
	// fmt.Println("copy  " + dst + " sucess\n")
}

// CPall cp all to dst
func CPall(src, dst string) {
	var FileeSrcdir, FileDstdir string
	List, _ := ioutil.ReadDir(src)
	for _, v := range List {
		fmt.Println("copy  " + v.Name() + " sucess\n")
		if v.IsDir() {
			FileeSrcdir = src + "\\" + v.Name()
			FileDstdir = dst + "\\" + v.Name()
			err := os.Mkdir(dst+"\\"+v.Name(), v.Mode())
			if err != nil {
				fmt.Println("create  " + dst + "\\" + v.Name() + " bad\n")
				os.Exit(0)
			}
			CPall(FileeSrcdir, FileDstdir)
			// fmt.Println("copy  " + FileDstdir + " sucess\n")
		}
		CopyF(src+"\\"+v.Name(), dst+"\\"+v.Name())
	}
}

// ChecksrcexisT check file exist
func ChecksrcexisT(s string) bool {

	f, err := os.Stat(s)
	if err != nil {
		fmt.Println(s + " is not exist !!")
		os.Exit(0)
	}
	if f.IsDir() {
		return true
	}
	return false
}
