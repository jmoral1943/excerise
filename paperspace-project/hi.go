package paperspace_project

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

import "how"

func main() {
	err := filepath.Walk(".",
		func(path string, info os.FileInfo, err error) error {
			matched, err := regexp.MatchString(`.go$`, filepath.Base(path))
			if matched {
				fmt.Println(filepath.Base(path))
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


