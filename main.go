package main

import (
	"flag"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"golang.org/x/image/bmp"
	"image"
	"image/color"
	"os"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	i := flag.Uint("size", 3, "Set the map size in multiples of 2 (total will be 2^n + 1)")
	colorValue := flag.Int("color", 32725, "Color to make outputted gif")
	path := flag.String("path", "test.bmp", "Filename to write to")
	dry := flag.Bool("dry", false, "Dry run (don't create image, just print size)")
	verbose := flag.Bool("v", false, "verbose")
	veryverbose := flag.Bool("vv", false, "very verbose")
	flag.Parse()

	if *colorValue > 65535 {
		fmt.Println("Color must be less than or equal to 65535")
		return
	}

	switch {
	case *veryverbose:
		log.SetLevel(log.DebugLevel)
	case *verbose:
		log.SetLevel(log.InfoLevel)
	default:
		log.SetLevel(log.WarnLevel)
	}

	t := InitializeTerrain(uint16(*i), 65535)

	fmt.Println("Generating terrain of size", t.max)
	t.Generate(0.3, 32767, 32767, 32767, 32767)

	if *dry {
		return
	}

	intMax := int(t.max)

	r := image.Rect(0, 0, intMax, intMax)
	m := image.NewGray16(r)

	for x := 0; x < intMax; x++ {
		for y := 0; y < intMax; y++ {
			n, _ := t.GetHeight(uint16(x), uint16(y))
			m.SetGray16(x, y, color.Gray16{Y: n})
		}
	}

	fi, err := os.Create(*path)
	defer fi.Close()
	if err != nil {
		panic(err)
	}
	bmp.Encode(fi, m)
}
