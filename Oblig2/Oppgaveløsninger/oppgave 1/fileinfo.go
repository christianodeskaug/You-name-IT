package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"log"
)


//main funskjonen
func main() {


fmt.Printf("Information about <fileinfo.go>: \n")

	file, err := os.Open("fileinfo.go")
	if err != nil {
		fmt.Printf("couldn't open file: %s\n", err)
		os.Exit(1)
	}





 	//Størrelse på filen

	size, err := ioutil.ReadFile("fileinfo.go")
	if err != nil{
		log.Panicf("failed at reading data from file: %s", err)
	}
	fmt.Printf("Size: %d bytes\n", len(size))
	i := len(size)
	f := float64(i)
	fmt.Printf("Size: %f Kilobytes\n", f/1e3)
	fmt.Printf("Size: %f Megabytes\n", f/1e6)
	fmt.Printf("Size: %f Gigabytes\n", f/1e9)




	//Ulike "stats"
	stats, err := file.Stat()
	if err != nil {
		fmt.Printf("couldn't get stats: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Is dir: %t\n", stats.IsDir())
	fmt.Printf("Is regular: %t\n", stats.Mode().IsRegular())
	fmt.Printf("Unix permissions: %s\n", stats.Mode().String())
	fmt.Printf("Append only: %t\n", stats.Mode()&os.ModeAppend != 0)
	fmt.Printf("Is device: %t\n", stats.Mode()&os.ModeDevice != 0)
	fmt.Printf("Is symlink: %t\n", stats.Mode()&os.ModeSymlink != 0)

}
