package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
)

func main() {
	file, err := os.Open("example-sheet.csv")

	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	dataFrame := dataframe.ReadCSV(file)
	fmt.Println(dataFrame)

	// select row
	// row := dataFrame.Subset([]int{0, 2})
	// fmt.Println(row)

	// select column
	column := dataFrame.Select([]string{"kecamatan", "provinsi"})
	fmt.Println(column)
}
