package fileutils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

func ReadFileSync(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		//print(scanner.Text())
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file: ", err)
		return nil
	}

	return lines
}

func ReadFile(filePath string, wg *sync.WaitGroup, dataChan chan<- []string) {
	defer wg.Done()

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}

	dataChan <- lines
}

func GetFiles(gameFolder string) []string {
	files := []string{}
	//gameFolder := ".\\data"

	err := filepath.Walk(gameFolder, func(yearpath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && yearpath != gameFolder {
			//fmt.Println("Directory:", yearpath)
			err := filepath.Walk(yearpath, func(monthfile string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}

				if !info.IsDir() && monthfile != yearpath {
					//fmt.Println("Filepath:", monthfile)
					files = append(files, monthfile)
				}

				return nil
			})

			if err != nil {
				fmt.Printf("Year error")
			}
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Walk error")
	}

	return files

}

func GetFilesYear(gameFolder string, year int) []string {
	files := []string{}
	yearFolder := gameFolder + "\\" + strconv.Itoa(year)
	//gameFolder := ".\\data"

	err := filepath.Walk(yearFolder, func(yearpath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && yearpath != yearFolder {
			//fmt.Println("Directory:", yearpath)
			files = append(files, yearpath)
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Walk error")
	}

	return files
}
