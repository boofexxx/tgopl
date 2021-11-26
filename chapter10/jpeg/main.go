package main

import (
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	flag.Parse()

	var format = flag.String("f", "png", "output format")
	if len(flag.Args()) == 0 {
		os.Exit(1)
	}

	for _, file := range flag.Args() {
		fin, err := os.Open(file)
		if err != nil {
			log.Println(err)
			continue
		}

		// what if the name of the file is something like .go ???
		fout, err := os.Create(file[:len(file)-len(filepath.Ext(file))] + "_changed." + *format)
		if err != nil {
			log.Println(err)
			continue
		}

		switch *format {
		case "jpg", "jpeg":
			if err := toJPEG(fin, fout); err != nil {
				fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
				os.Exit(1)
			}
		case "png":
			if err := toPNG(fin, fout); err != nil {
				fmt.Fprintf(os.Stderr, "png: %v\n", err)
				os.Exit(1)
			}
		}
		fin.Close()
		fout.Close()
	}

}

func toPNG(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	return png.Encode(out, img)
}

func toJPEG(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
}
