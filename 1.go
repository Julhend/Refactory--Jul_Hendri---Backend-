package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Product struct {
	Name     string `json:"name"`
	Category string `json:"category"`
	Price    string `json:"price"`
}

// type ByPrice []Product
// func (a ByPrice) Len() int           { return len(a) }
// func (a ByPrice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
// func (a ByPrice) Less(i, j int) bool { return a[i].Price < a[j].Price }

func main() {

	csvFile, _ := os.Open("1.csv")
	reader := csv.NewReader(csvFile)
	var gadget []Product
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		gadget = append(gadget, Product{
			Name:     line[0],
			Category: line[1],
			Price:    line[2],
		})
	}
	gadgetJson, _ := json.Marshal(gadget)
	fmt.Println(string(gadgetJson))
	// sort.Sort(ByPrice(gadget))
}
