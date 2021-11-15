package main

import (
	"testing"

	"challenge.haraj.com.sa/kraicklist/datautils"
)

func TestPaginateRecord(t *testing.T) {
	record1 := datautils.Record{1, "tes1", "tes content", "google.com", []string{"Tag 1"}, 1636990324755, []string{"google.com"}}
	record2 := datautils.Record{2, "tes2", "tes content 2", "google.com", []string{"Tag 2"}, 1636990324755, []string{"google.com"}}
	records := []datautils.Record{record1, record2}
	recordData, pageCount := paginateRecord(records, 0)
	if len(recordData) != 2 && len(pageCount) != 1 {
		t.Fatalf("Should paginate result record")
	}
}

func TestSearch(t *testing.T) {
	record1 := datautils.Record{1, "tes1", "tes content", "google.com", []string{"Tag 1"}, 1636990324755, []string{"google.com"}}
	record2 := datautils.Record{2, "tes2", "tes content 2", "google.com", []string{"Tag 2"}, 1636990324755, []string{"google.com"}}
	records := []datautils.Record{record1, record2}
	searcher := &datautils.Searcher{}
	searcher.SetRecords(records)

	result, err := searcher.Search("tes")
	if err != nil {
		t.Fatalf("Should not throw error")
	}
	if result[0].ID != 1 {
		t.Fatalf("Should give first record as first result")
	}
}
