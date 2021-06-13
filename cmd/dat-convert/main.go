package main

import (
	"fmt"
	"image/color"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"

	dat "github.com/gravestench/dat_palette/pkg"
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

	p, err := dat.Unmarshal(fileContents)
	if err != nil {
		fmt.Print(err)
		return
	}

	f, err := os.Create(fmt.Sprintf("%s.gpl", fileName))
	if err != nil {
		log.Fatal(err)
	}

	if err := writeGimpPalette(baseName, p, f); err != nil {
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

func writeGimpPalette(name string, p color.Palette, w io.Writer) error {
	// GIMP Palette
	// Name: Bears
	// #
	// 8   8   8	grey3
	// 68  44  44
	// 80   8  12
	// 72  56  56
	// 104  84  68
	// 116  96  80
	// 84  56  44
	// 140 104  88
	const (
		line1        = "GIMP Palette\r\n"
		line2        = "Name: %s\r\n"
		line3        = "#\r\n"
		fmtComponent = "  %v"
		fmtLine      = "%s %s %s\r\n"
		fmtErr = "error encoding DAT to gpl format, %v"
	)

	numHeaderLines := 3
	numColors := len(p)
	numLines := numColors + numHeaderLines
	lines := make([]string, numLines)

	lines[0] = line1
	lines[1] = fmt.Sprintf(line2, fileNameWithoutExt(name))
	lines[2] = line3

	strComponent := func(n int) string {
		s := fmt.Sprintf(fmtComponent, n)
		return s[len(s)-3:]
	}

	for idx := range p {
		r, g, b, _ := p[idx].RGBA()
		rs := strComponent(int(uint8(r)))
		gs := strComponent(int(uint8(g)))
		bs := strComponent(int(uint8(b)))

		lines[numHeaderLines+idx] = fmt.Sprintf(fmtLine, rs, gs, bs)
	}

	for idx := range lines {
		if _, err := w.Write([]byte(lines[idx])); err != nil {
			return fmt.Errorf(fmtErr, err)
		}
	}

	return nil
}
