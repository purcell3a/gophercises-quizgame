package main

import (
	"encoding/csv"
	// "flag"
	"fmt"
	"os"
	// "strings"
)

// func NewReader(r io.Reader) *Reader why is this different than the docs??
func getProblemFile() {
	// var file string 
	// fmt.Scan(&file)
	fmt.Println("what's broken")


}

func parseProblemFile(file string){

	f, err := os.Open(file)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s", file))
	}
	reader := csv.NewReader(f)
	records, _ := reader.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}
	fmt.Println(records)
}

func main () {
	parseProblemFile("problems.csv")
	getProblemFile()
	// parseProblemFile()



	// we need a filename variable so user can customize problems filename

}
//  https://stackoverflow.com/questions/12907653/what-is-the-difference-between-string-and-string-in-golang
//  Link on why the brackets are there......could use more explanation here 
func parseLines (lines [][]string) []problem {
	//  make() creates a slice....so this is making a slice out of the output and the ...length??
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q:line[0],
			a : strings.TrimSpace(line[1])
			// need to look up TrimSpace
		}
	}
	return ret
}

//  so structs are like objects....and this is creating an object "outline" for....formatting??
type problem struct{
	q string
	a string
}
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}