package db

import (
	"log"

	"challenge.haraj.com.sa/kraicklist/datautils"
)

type Db struct{}

func (mdb *Db) InitDB() (*datautils.Searcher, error) {
	searcher := &datautils.Searcher{}
	err := searcher.Load("data.gz")
	if err != nil {
		log.Fatalf("unable to load search data due: %v", err)
		return searcher, err
	}
	return searcher, nil
}

func (mdb Db) closeDB() {

}
