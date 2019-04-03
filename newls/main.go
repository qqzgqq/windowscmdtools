package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"windowscmdtools/newls/lsused"
	"windowscmdtools/newls/template"
)

// LL check ls -l
var LL bool

// HL check ls -hl or ls -lh
var HL bool

func init() {
	var hl, lh, l, h, help *bool
	hl = flag.Bool("hl", false, "a bool")
	lh = flag.Bool("lh", false, "a bool")
	l = flag.Bool("l", false, "a bool")
	h = flag.Bool("h", false, "a bool")
	help = flag.Bool("help", false, "a bool")
	flag.Parse()
	if *hl || *lh {
		HL = true
	}
	if *l {
		LL = true
	}
	if *h || *help {
		template.YelloW("usage:   ls [options]... [dir]\n")
		fmt.Printf("   -l    Displays the files as a list and the files size in bytes\n")
		fmt.Printf("   -lh   Displays the files as a list and the files size in KB MB GB TB\n")
		fmt.Printf("   -hl   Displays the files as a list and the files size in KB MB GB TB\n")
		os.Exit(0)
	}
}
func main() {

	var DIr = flag.Arg(0)
	//if DIr is in a-z or A-Z then set DIr:\\
	if m, _ := regexp.MatchString("^[a-z]$|^[A-Z]$", DIr); m {
		DIr = DIr + ":\\"
	}
	if DIr == "" {
		DIr = "."
	}
	if LL {
		lsused.PPCOLL(DIr)
		os.Exit(0)
	}
	if HL {
		lsused.PPCOHL(DIr)
		os.Exit(0)
	}
	lsused.PPCO(DIr)

}
