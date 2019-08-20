package main

import (
	"flag"
	"os"
	"strings"
	"windowscmdtools/rmall/cpall"
)

func main() {
	flag.Parse()
	var FiNe = flag.Arg(0)
	// get the file path
	var Dir1 string
	if strings.Contains(FiNe, ":") == false {
		Dir1, _ = os.Getwd()
		FiNe = Dir1 + "\\" + FiNe
	}

	var FileDST = strings.Split(Dir1, ":")[0] + ":\\deletefiles\\"
	_, err1 := os.Stat(FileDST)
	if err1 != nil {
		err := os.Mkdir(FileDST, 775)
		if err != nil {
			panic(err)
		}
	}

	FiNeshuzu := strings.Split(FiNe, "\\")
	FiNenumber := len(FiNeshuzu) - 1
	FileDST = FileDST + FiNeshuzu[FiNenumber]
	if cpall.ChecksrcexisT(FiNe) {
		cpall.MKFDIR(FiNe, FileDST)
		cpall.CPall(FiNe, FileDST)
	} else {

		cpall.CopyF(FiNe, FileDST)
	}

	_, error := os.Stat(FileDST)
	if error == nil {
		os.RemoveAll(FiNe)
	}

}
