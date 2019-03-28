package main

import (
	"copynew/copyfile"
	"copynew/template"
	"flag"
	"fmt"
	"os"
)

func main() {
	var pd string
	flag.Parse()
	var SRC = flag.Arg(0)
	var DST = flag.Arg(1)
	Srclx := template.ChecksrcexisT(SRC)
	DstExist, Dstlx := template.CheckdstexisT(DST)
	DSTFILEE := template.CheckdstfileexisT(SRC, DST)
	if Srclx && Dstlx {
		fmt.Println("it is dir")
	} else if Srclx == false && Dstlx == false {
		if DstExist {
			fmt.Printf(SRC + " is oready exist,oevr it Y/N:")
			fmt.Scanln(&pd)
			if pd == "y" || pd == "Y" {
				copyfile.CopyF(SRC, DST)
			} else {
				os.Exit(0)
			}
		} else {
			copyfile.CopyF(SRC, DST)
		}

	} else if Srclx == false && Dstlx {
		if DSTFILEE {
			fmt.Printf(SRC + " is oready exist,oevr it Y/N:")
			fmt.Scanln(&pd)
			if pd == "y" || pd == "Y" {
				copyfile.CopyF(SRC, DST+"\\"+SRC)
			} else {
				os.Exit(0)
			}
		} else {
			copyfile.CopyF(SRC, DST+"\\"+SRC)
		}
	}

	// _, err := copyfile.CopyF(SRC, DST)
	// if err != nil {
	// 	panic(err)
	// }
	// Dirtest:="h:\\test1233"
	// errr := os.Mkdir(Dirtest, 0755)
	// if errr != nil {
	// 	fmt.Printf(": ")
	// }
	// fmt.Println("copy bak ,del dstdir ,copy src ,del bak ")

}
file > file ok
file > dir ok
file > dir\file no