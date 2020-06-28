package main

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	root := "D:/SmartsFiles/DefectsData/"
	files := GetFiles(root)

	defectSli := make([]int, 0, 10)

	for _, file := range files {
		defectSli = append(defectSli, ReadData(file)...)
	}

	fmt.Println("Found " + strconv.Itoa(len(defectSli)) + " defects")
	fmt.Println(defectSli)
}

/** Example https://www.bogotobogo.com/GoLang/GoLang_SQLite.php */
func ReadData(file string) []int {
	fmt.Println("Reading Data", file)
	defectSli := make([]int, 0, 100)
	database, _ := sql.Open("sqlite3", "./0a-swath_0.db")
	rows, _ := database.Query("SELECT defectId FROM attribTable")
	var defectId int
	for rows.Next() {
		rows.Scan(&defectId)
		defectSli = append(defectSli, defectId)
	}

	return defectSli
}

func GetFiles(root string) []string {
	var files []string
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	return files
}
