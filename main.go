package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var xi []map[string][]string

// Scans the file and searches for imports
func scanFile(f string) {
	content, err := os.Open(f)
	if err != nil {
		log.Fatal(err)
	}

	defer content.Close()

	scanner := bufio.NewScanner(content)

	//rgx := regexp.MustCompile(`\((import(.*?)\)`)
	importPresent := false

	var imp []string
	for scanner.Scan() {

		if strings.Contains(scanner.Text(), ")") {
			importPresent = false
		}

		if importPresent {
			s := strings.TrimSpace(scanner.Text())
			s = strings.Trim(s, `""`)
			imp = append(imp, s)
		}

		if strings.Contains(scanner.Text(), "import") {
			importPresent = true

			if strings.Contains(scanner.Text(), "(") {
				continue
			}

			s := strings.Split(scanner.Text(), "import")
			str := strings.Join(s, "")
			s = strings.Split(str, "(")
			str = strings.Join(s, "")
			str = strings.TrimSpace(str)
			str = strings.Trim(str, `""`)
			imp = append(imp, str)
			importPresent = false
		}
	}

	toJson(filepath.Base(f), imp)
}

func toJson(f string, imp []string) {
	m := map[string][]string{
		f: imp,
	}
	xi = append(xi, m)
}

func listFiles(f string) {
	err := filepath.Walk(f,
		func(path string, info os.FileInfo, err error) error {
			matched, err := regexp.MatchString(`.go$`, filepath.Base(path))
			if matched {
				scanFile(path)
				//fmt.Println(path)
				//fmt.Println(filepath.Base(path))
			}
			if err != nil {
				return err
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}
}

func writeToFile() error {
	js, err := json.MarshalIndent(xi, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Create("imports.json")
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(js)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func main() {
	listFiles(os.Args[1])
	err := writeToFile()
	if err != nil {
		log.Fatal(err)
	}
	//"./paperspace-project"
}
