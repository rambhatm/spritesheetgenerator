package main

import (
	"fmt"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"strings"

	gim "github.com/ozankasikci/go-image-merge"
)

func main() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	var grids []*gim.Grid
	fmt.Println("Creating sprite sheet from following files...")

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".png") {

			grid := gim.Grid{ImageFilePath: file.Name()}
			grids = append(grids, &grid)
			fmt.Println(file.Name())
		}
	}

	rgba, err := gim.New(grids, len(grids), 1).Merge()

	// save the output to jpg or png
	file, err := os.Create("sheet.png")
	err = png.Encode(file, rgba)
}
