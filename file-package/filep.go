package filepackage

import (
	"fmt"
	"os"
	"io"
	"bufio"
	"strings"

)


func ReadText() {
	for _, fname := range os.Args[1:] {
		file, err := os.Open(fname)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		if  _, err := io.Copy(os.Stdout, file); err != nil {
			fmt.Fprint(os.Stderr, err)
			continue
		}
		file.Close()
	}
}


func ReadTextUtil() {
	for _, fname := range os.Args[1:] {
		file, err := os.Open(fname)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		data, err := io.ReadAll(file)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			continue
		}

		fmt.Println("the file has", len(data), "bytes")
		file.Close()
	}
}


func CountFileUnits() {
	total_count := map[string]uint{
		"total_words": 0,
		"total_chars": 0,
		"total_lines": 0,
	}
	for _, fname := range os.Args[1:] {

		var line_count, word_count, char_count uint

		file, err := os.Open(fname)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		scan := bufio.NewScanner(file)

		for scan.Scan() {
			s := scan.Text()

			word_count += uint(len(strings.Fields(s)))
			total_count["total_words"]+= word_count

			char_count += uint(len(s))
			total_count["total_chars"]+= char_count

			line_count++
			total_count["total_lines"]+= line_count
			fmt.Printf(" %7d %7d %7d %s\n", word_count, char_count, line_count, fname)
		}


		file.Close()

		}
		if len(os.Args) > 2 {
			fmt.Printf(" %7d %7d %7d\n", 
			total_count["total_words"],
			total_count["total_chars"], 
			total_count["total_lines"] )	
	}
}
