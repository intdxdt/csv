package main

import (
	"fmt"
	"github.com/intdxdt/csv"
)

func main(){
	fname := "example/coords.csv"
	data := csv.ParseCSV(fname,',', true)
	for _, o := range data {
		for k := range o {
			fmt.Println([]byte(k))
		}
	}
	fmt.Println(len(data))
}