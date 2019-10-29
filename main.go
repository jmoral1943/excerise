package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func scanFile(f string) {
	content, err := os.Open(f)
	if err != nil {
		log.Fatal(err)
	}

	defer content.Close()

	scanner := bufio.NewScanner(content)

	//rgx := regexp.MustCompile(`\((import(.*?)\)`)
	importPresent := false

	for scanner.Scan() {

		if strings.Contains(scanner.Text(), ")") {
			importPresent = false
		}

		if importPresent == true {
			fmt.Println(scanner.Text())
		}

		if strings.Contains(scanner.Text(), "import") {
			importPresent = true;
			s := strings.Split(scanner.Text(), "import")
			str := strings.Join(s, "")
			s = strings.Split(str, "(")

			fmt.Println(strings.Join(s, ""))
		}

		//fmt.Println("text per line",scanner.Text())

	}

}


func main() {

	scanFile("./hi.go")
	//fmt.Println(os.Args[1:])
}
