package main

import (
	"flag"
	"fmt"
	"os"
	"windowscmdtools/copynew/copyfile"
	"windowscmdtools/copynew/template"
)

func main() {
	var pd string
	flag.Parse()
	var SRC = flag.Arg(0)
	var DST = flag.Arg(1)
	Srclx := template.ChecksrcexisT(SRC)
	DstExist, Dstlx := template.CheckdstexisT(DST)
	DSTFILEE := template.CheckdstfileexisT(SRC, DST)
	//src is dir and dst is dir
	if Srclx {
		DIrfile := copyfile.MKFDIR(SRC, DST)
		copyfile.CPall(SRC, DIrfile)
	}

	//src is file and dst is not exist file
	if Srclx == false && DstExist {
		copyfile.CopyF(SRC, DST)
	}
	//src is file and dst is exist file
	if Srclx == false && DstExist == false {
		if Dstlx == false {
			fmt.Printf(DST + " is oready exist,oevr it Y/N:")
			fmt.Scanln(&pd)
			if pd == "y" || pd == "Y" {
				if SRC == DST {
					// src and dst is in the same dir
					copyfile.CopyF(SRC, DST+".bak")
				}
				copyfile.CopyF(SRC, DST)

			} else {
				os.Exit(0)
			}
		} else {
			// src is file and dst is dir not exist file
			copyfile.CopyF(SRC, DST+"\\"+SRC)
		}

	}

	//src is file and dst is dir and exist file
	if Srclx == false && Dstlx {
		if DSTFILEE {
			fmt.Printf(DST + "\\" + SRC + " is oready exist,oevr it Y/N:")
			fmt.Scanln(&pd)
			if pd == "y" || pd == "Y" {
				copyfile.CopyF(SRC, DST)
			} else {
				os.Exit(0)
			}
		} else {
			copyfile.CopyF(SRC, DST)
		}
	}

}
