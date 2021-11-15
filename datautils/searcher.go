package datautils

import (
	"bufio"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Searcher struct {
	records []Record
}

func (s *Searcher) Load(filepath string) error {
	// open file
	file, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("unable to open source file due: %v", err)
	}
	defer file.Close()
	// read as gzip
	reader, err := gzip.NewReader(file)
	if err != nil {
		return fmt.Errorf("unable to initialize gzip reader due: %v", err)
	}
	// read the reader using scanner to contstruct records
	var records []Record
	cs := bufio.NewScanner(reader)
	for cs.Scan() {
		var r Record
		err = json.Unmarshal(cs.Bytes(), &r)
		if err != nil {
			continue
		}
		records = append(records, r)
	}
	s.records = records

	return nil
}

func (s *Searcher) Search(query string) ([]Record, error) {
	// var result []Record
	// for _, record := range s.records {
	// 	if strings.Contains(record.Title, query) || strings.Contains(record.Content, query) {
	// 		result = append(result, record)
	// 	}
	// }
	ch := make(chan []Record)
	go searchContains(ch, s, "title", query)
	go searchContains(ch, s, "content", query)
	go searchContains(ch, s, "thumbURL", query)

	res1 := <-ch
	res2 := <-ch
	res3 := <-ch
	// result := res1
	result := append(res1, res2...)
	result = append(result, res3...)
	return result, nil
}

func searchContains(c chan []Record, s *Searcher, source string, query string) {
	var result []Record
	for _, record := range s.records {
		if strings.Contains(strings.ToLower(record.getItem(source)), strings.ToLower(query)) {
			result = append(result, record)
		}
	}
	c <- result
}
