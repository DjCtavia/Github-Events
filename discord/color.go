package discord

import (
	"log"
	"strconv"
)

func RGBToInt(r int, g int, b int) int {
	return r + g<<8 + b<<16
}

func HexRGBToInt(hexCode string) int {
	if len(hexCode) != 6 {
		log.Fatal("Hex code given is invalid, it must be 6 (FFFFFF)")
	}
	value, err := strconv.ParseInt(hexCode, 16, 24)
	if err != nil {
		log.Fatal(err)
	}
	return int(value)
}
