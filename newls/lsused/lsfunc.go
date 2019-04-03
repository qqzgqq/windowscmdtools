package lsused

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
	"windowscmdtools/newls/template"
)

// DirSIZESUM cale the dir sum size
func DirSIZESUM(DIR string) int64 {
	var DIRINtoNAme string
	var Dirsizesum int64
	DirList, _ := ioutil.ReadDir(DIR)
	for _, v := range DirList {
		DIRINtoNAme = DIR + "\\" + v.Name()
		if v.IsDir() == true {
			Dirsizesum = Dirsizesum + DirSIZESUM(DIRINtoNAme)
		}

		Dirsizesum = Dirsizesum + v.Size()
	}

	return Dirsizesum
}

// DISPLAYcolor select color
func DISPLAYcolor(s, ss string) {
	switch s {
	case "sh", "SH":
		template.GreeN(ss)
	case "bat", "BAT":
		template.GreeN(ss)
	case "gz", "GZ":
		template.YelloW(ss)
	case "zip", "ZIP":
		template.YelloW(ss)
	case "rar", "RAR":
		template.YelloW(ss)
	case "exe", "EXE":
		template.YelloW(ss)
	default:
		fmt.Println(ss)
	}
}

// CHECKFILEORDIr ls one file
func CHECKFILEORDIr(DIr string) {
	f, err := os.Stat(DIr)
	if err != nil {
		template.ReD(DIr + "is not exsit")
		os.Exit(0)
	}
	if f != nil {
		splitstring := strings.Split(DIr, "\\")
		FIlename := splitstring[len(splitstring)-1]
		Splitstrings := strings.Split(FIlename, ".")
		DISPLAYcolor(Splitstrings[len(Splitstrings)-1], FIlename)
	}
}

// ONEFILEDISplay return NewDir, FIlename
func ONEFILEDISplay(s string) (string, string) {
	var NewDir string
	var FIlename string
	f, err := os.Stat(s)
	if err != nil {
		template.ReD(s + "is not exsit")
		os.Exit(0)

	}
	if f != nil {

		SPlitstring := strings.Split(s, "\\")
		FIlename = SPlitstring[len(SPlitstring)-1]
		Spnum := len(SPlitstring)
		for i := 0; i < Spnum-1; i++ {
			NewDir = NewDir + SPlitstring[i] + "\\"
		}
	}
	return NewDir, FIlename
}

// CHECKFILEORDIR2 use ls -l one file info
func CHECKFILEORDIR2(DIr string) {
	f, err := os.Stat(DIr)
	if err != nil {
		template.ReD(DIr + "is not exsit")
		os.Exit(0)
	}
	if f != nil {
		if strings.Contains(DIr, "\\") {
			NewDir, FIlename := ONEFILEDISplay(DIr)
			DirList, _ := ioutil.ReadDir(NewDir)
			for _, v := range DirList {

				if v.Name() == FIlename {

					Splitstrings := strings.Split(FIlename, ".")
					fmt.Print(v.Mode())
					BZPRINT(v.ModTime().Format("2006-01-02 15:04:05"), strconv.FormatInt(v.Size(), 10))
					DISPLAYcolor(Splitstrings[len(Splitstrings)-1], FIlename)

				}

			}

		} else {
			LLONEFILECOLOR(DIr)
		}
	}
}

// LLONEFILECOLOR ls -l one file color and list display
func LLONEFILECOLOR(DIr string) {
	f, _ := os.Stat(".\\" + DIr)
	fmt.Print(f.Mode())
	BZPRINT(f.ModTime().Format("2006-01-02 15:04:05"), strconv.FormatInt(f.Size(), 10))
	Splitstrings := strings.Split(DIr, ".")
	DISPLAYcolor(Splitstrings[len(Splitstrings)-1], DIr)
}

// HLONEFILECOLOR ls -hl one file color and list display
func HLONEFILECOLOR(DIr string) {
	f, _ := os.Stat(".\\" + DIr)
	fmt.Print(f.Mode())
	BZPRINT(f.ModTime().Format("2006-01-02 15:04:05"), JSSIZE(f.Size()))
	Splitstrings := strings.Split(DIr, ".")
	DISPLAYcolor(Splitstrings[len(Splitstrings)-1], DIr)
}

// PPCO use ls
func PPCO(DIr string) {
	//checkout DIr exist

	DirList, e := ioutil.ReadDir(DIr)
	if e != nil {
		CHECKFILEORDIr(DIr)
		return
	}

	for _, v := range DirList {
		if v.IsDir() == true {
			template.BluE(v.Name() + "    ")
		} else if strings.Contains(v.Name(), ".sh") == true {
			template.GreeN(v.Name() + "    ")
		} else if strings.Contains(v.Name(), ".bat") == true {
			template.GreeN(v.Name() + "    ")
		} else if strings.Contains(v.Name(), ".tar.") == true {
			template.YelloW(v.Name() + "    ")
		} else if strings.Contains(v.Name(), ".rar") == true {
			template.YelloW(v.Name() + "    ")
		} else if strings.Contains(v.Name(), ".zip") == true {
			template.YelloW(v.Name() + "    ")
		} else if strings.Contains(v.Name(), ".exe") == true {
			template.GreeN(v.Name() + "    ")
		} else {
			fmt.Printf(v.Name() + "    ")

		}

	}
}

// BZPRINT print formatting
func BZPRINT(s1, s2 string) {
	fmt.Printf("%25s", s1)
	fmt.Printf("%17s", s2)
	fmt.Printf("%7s", "")
}

// PPCOLL ls -l
func PPCOLL(DIr string) {
	//checkout DIr exist
	DirList, e := ioutil.ReadDir(DIr)
	if e != nil {
		CHECKFILEORDIR2(DIr)
		return
	}
	for _, v := range DirList {

		if v.IsDir() == true {
			DIRSUMNAME := DIr + "\\" + v.Name()
			fmt.Print(v.Mode())
			BZPRINT(v.ModTime().Format("2006-01-02 15:04:05"), strconv.FormatInt(DirSIZESUM(DIRSUMNAME), 10))
			template.BluE(v.Name() + "\n")
		} else if strings.Contains(v.Name(), ".sh") == true {
			fmt.Print(v.Mode())
			BZPRINT(v.ModTime().Format("2006-01-02 15:04:05"), strconv.FormatInt(v.Size(), 10))
			template.GreeN(v.Name() + "\n")
		} else if strings.Contains(v.Name(), ".bat") == true {
			fmt.Print(v.Mode())
			BZPRINT(v.ModTime().Format("2006-01-02 15:04:05"), strconv.FormatInt(v.Size(), 10))
			template.GreeN(v.Name() + "\n")
		} else if strings.Contains(v.Name(), ".tar.") == true {
			fmt.Print(v.Mode())
			BZPRINT(v.ModTime().Format("2006-01-02 15:04:05"), strconv.FormatInt(v.Size(), 10))
			template.YelloW(v.Name() + "\n")
		} else if strings.Contains(v.Name(), ".rar") == true {
			fmt.Print(v.Mode())
			BZPRINT(v.ModTime().Format("2006-01-02 15:04:05"), strconv.FormatInt(v.Size(), 10))
			template.YelloW(v.Name() + "\n")
		} else if strings.Contains(v.Name(), ".zip") == true {
			fmt.Print(v.Mode())
			BZPRINT(v.ModTime().Format("2006-01-02 15:04:05"), strconv.FormatInt(v.Size(), 10))
			template.YelloW(v.Name() + "\n")
		} else if strings.Contains(v.Name(), ".exe") == true {
			fmt.Print(v.Mode())
			BZPRINT(v.ModTime().Format("2006-01-02 15:04:05"), strconv.FormatInt(v.Size(), 10))
			template.GreeN(v.Name() + "\n")
		} else {
			fmt.Print(v.Mode())
			BZPRINT(v.ModTime().Format("2006-01-02 15:04:05"), strconv.FormatInt(v.Size(), 10))
			fmt.Printf(v.Name() + "\n")

		}
	}
}

// Decimal make the float64
func Decimal(value float64) float64 {
	return math.Trunc(value*1e2+0.5) * 1e-2
}

// JSSIZE calc the size
func JSSIZE(a int64) string {
	var b float64 = 1
	var filesize string
	if b < 1024 {
		b = float64(a) / 1024
		if b < 1024 {
			filesize = fmt.Sprintf("%.2f", Decimal(b)) + " KB"
		} else {
			b = b / 1024
			if b < 1024 {
				filesize = fmt.Sprintf("%.2f", Decimal(b)) + " MB"
			} else {
				b = b / 1024
				if b < 1024 {
					filesize = fmt.Sprintf("%.2f", Decimal(b)) + " GB"
				} else {
					b = b / 1024
					if b < 1024 {
						filesize = fmt.Sprintf("%.2f", Decimal(b)) + " TB"
					}
				}
			}
		}
	}
	return filesize
}

// PPCOHL ls -hl
func PPCOHL(DIr string) {
	//checkout DIr exist
	DirList, e := ioutil.ReadDir(DIr)
	if e != nil {
		if strings.Contains(DIr, "\\") {
			NewDir, FIlename := ONEFILEDISplay(DIr)
			DirList, _ := ioutil.ReadDir(NewDir)
			for _, v := range DirList {

				if v.Name() == FIlename {

					Splitstrings := strings.Split(FIlename, ".")
					fmt.Print(v.Mode())
					BZPRINT(v.ModTime().Format("2006-01-02 15:04:05"), JSSIZE(v.Size()))
					DISPLAYcolor(Splitstrings[len(Splitstrings)-1], FIlename)

				}

			}

		} else {
			HLONEFILECOLOR(DIr)
		}
		return
	}

	for _, v := range DirList {

		if v.IsDir() == true {
			DIRSUMNAMEH := DIr + "\\" + v.Name()
			fmt.Print(v.Mode())
			BZPRINT(v.ModTime().Format("2006-01-02 15:04:05"), JSSIZE(DirSIZESUM(DIRSUMNAMEH)))
			template.BluE(v.Name() + "\n")
		} else if strings.Contains(v.Name(), ".sh") == true {
			fmt.Print(v.Mode())
			BZPRINT(v.ModTime().Format("2006-01-02 15:04:05"), JSSIZE(v.Size()))
			template.GreeN(v.Name() + "\n")
		} else if strings.Contains(v.Name(), ".bat") == true {
			fmt.Print(v.Mode())
			BZPRINT(v.ModTime().Format("2006-01-02 15:04:05"), JSSIZE(v.Size()))
			template.GreeN(v.Name() + "\n")
		} else if strings.Contains(v.Name(), ".tar.") == true {
			fmt.Print(v.Mode())
			BZPRINT(v.ModTime().Format("2006-01-02 15:04:05"), JSSIZE(v.Size()))
			template.YelloW(v.Name() + "\n")
		} else if strings.Contains(v.Name(), ".rar") == true {
			fmt.Print(v.Mode())
			BZPRINT(v.ModTime().Format("2006-01-02 15:04:05"), JSSIZE(v.Size()))
			template.YelloW(v.Name() + "\n")
		} else if strings.Contains(v.Name(), ".zip") == true {
			fmt.Print(v.Mode())
			BZPRINT(v.ModTime().Format("2006-01-02 15:04:05"), JSSIZE(v.Size()))
			template.YelloW(v.Name() + "\n")
		} else if strings.Contains(v.Name(), ".exe") == true {
			fmt.Print(v.Mode())
			BZPRINT(v.ModTime().Format("2006-01-02 15:04:05"), JSSIZE(v.Size()))
			template.GreeN(v.Name() + "\n")
		} else {
			fmt.Print(v.Mode())
			BZPRINT(v.ModTime().Format("2006-01-02 15:04:05"), JSSIZE(v.Size()))
			fmt.Printf(v.Name() + "\n")

		}
	}

}
