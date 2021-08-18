package main

import (
	"fmt"
	"image/color"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"

	dat "github.com/OpenDiablo2/dat_palette/pkg"
	gpl "github.com/gravestench/gpl/pkg"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	srcPath := os.Args[1]
	baseName := path.Base(srcPath)
	fileName := fileNameWithoutExt(baseName)

	fileContents, err := ioutil.ReadFile(srcPath)
	if err != nil {
		const fmtErr = "could not read file, %w"

		fmt.Print(fmt.Errorf(fmtErr, err))

		return
	}

	p, err := dat.Decode(fileContents)
	if err != nil {
		fmt.Print(err)
		return
	}

	f, err := os.Create(fmt.Sprintf("%s.gpl", fileName))
	if err != nil {
		log.Fatal(err)
	}

	if err := gpl.FromPalette(color.Palette(p)).Encode(baseName, f); err != nil {
		_ = f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func printUsage() {
	fmt.Printf("Usage:\r\n\t%s path/to/file.dat", os.Args[0])
}

func fileNameWithoutExt(fileName string) string {
	return fileName[:len(fileName)-len(filepath.Ext(fileName))]
}
