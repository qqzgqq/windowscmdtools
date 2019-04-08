package copyfile

import (
	"io/ioutil"
	"os"
	"windowscmdtools/copynew/template"
)

// MKFDIR mkdir first dir
func MKFDIR(src, dst string) string {
	LL, serr := os.Stat(src)
	if serr != nil {
		template.ReD(src + " is not exist\n")
		os.Exit(0)
	}

	aerr := os.MkdirAll(dst+"\\"+LL.Name(), LL.Mode())
	if aerr != nil {
		template.ReD("create  " + dst + "\\" + LL.Name() + " bad\n")
		os.Exit(0)
	}
	template.GreeN("copy  " + dst + "\\" + LL.Name() + " sucess\n")
	return dst + "\\" + LL.Name()

}

// CPall cp all to dst
func CPall(src, dst string) {
	var FileeSrcdir, FileDstdir string
	List, _ := ioutil.ReadDir(src)
	for _, v := range List {
		template.YelloW("copy  " + v.Name() + " sucess\n")
		if v.IsDir() {
			FileeSrcdir = src + "\\" + v.Name()
			FileDstdir = dst + "\\" + v.Name()
			err := os.Mkdir(dst+"\\"+v.Name(), v.Mode())
			if err != nil {
				template.ReD("create  " + dst + "\\" + v.Name() + " bad\n")
				os.Exit(0)
			}
			CPall(FileeSrcdir, FileDstdir)
			template.GreeN("copy  " + FileDstdir + " sucess\n")
		}
		CopyF(src+"\\"+v.Name(), dst+"\\"+v.Name())
	}
}
