package main

import (
	"flag"
	log "github.com/Sirupsen/logrus"
	"golang.org/x/image/bmp"
	"image"
	"image/color"
	"os"
)

func main() {
	i := flag.Uint("size", 3, "Set the map size in multiples of 2 (total will be 2^n + 1)")
	path := flag.String("path", "test.bmp", "Filename to write to")
	dry := flag.Bool("dry", false, "Dry run (don't create image, just print size)")
	loglevel := flag.String("loglevel", "Info", "Log level. Can be Debug, Info, Warn, Error, Fatal, or Panic")
	flag.Parse()

	switch *loglevel {
	case "Debug":
		log.SetLevel(log.DebugLevel)
	case "Info":
		log.SetLevel(log.InfoLevel)
	case "Warn":
		log.SetLevel(log.WarnLevel)
	case "Error":
		log.SetLevel(log.ErrorLevel)
	case "Fatal":
		log.SetLevel(log.FatalLevel)
	case "Panic":
		log.SetLevel(log.PanicLevel)
	default:
		log.Fatal("Unknown log level. Valid values are Debug, Info, Warn, Error, Fatal, or Panic. Supplied value: ", *loglevel)
		return // Unnecessary, but do so just in case the former one didn't
	}

	t := InitializeTerrain(uint16(*i), 65535)

	log.Info("Generating terrain of size ", t.max)
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
