package main

import (
	"fmt"
	"html/template"
	"math"
	"net/http"
	"regexp"
	"strconv"

	"challenge.haraj.com.sa/kraicklist/datautils"
)

const STATIC_DIR = "./static/"

func initHandler(searcher *datautils.Searcher) {
	fs := http.FileServer(http.Dir(STATIC_DIR))
	http.Handle("/", fs)
	http.HandleFunc("/home", makeHandler(renderHome))
	http.HandleFunc("/search", handleSearch(searcher))
}

var validPath = regexp.MustCompile("^/(home|result)$")

func makeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		m := validPath.FindStringSubmatch(req.URL.Path)
		if m == nil {
			fmt.Println("NOT FOUND: 404")
			http.NotFound(res, req)
			return
		}
		fn(res, req)
	}
}

func renderHome(res http.ResponseWriter, req *http.Request) {
	pageCount := []uint64{0}
	responseData := datautils.Page{Title: "home", Records: []datautils.Record{}, Pagecount: pageCount}
	renderTemplate(res, "index", &responseData)
}

func handleSearch(s *datautils.Searcher) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// fetch query string from query params
			q := r.URL.Query().Get("q")
			if len(q) == 0 {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("missing search query in query params"))
				return
			}
			// search relevant records
			records, err := s.Search(q)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}
			pageNum, _ := strconv.ParseUint(r.URL.Query().Get("p"), 10, 8)
			recordData, pageCount := paginateRecord(records, pageNum)

			responseData := datautils.Page{Title: "search", Records: recordData, Pagecount: pageCount}
			renderTemplate(w, "result", &responseData)
		},
	)
}

var templates = template.Must(template.ParseFiles(STATIC_DIR+"result.html", STATIC_DIR+"index.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *datautils.Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func paginateRecord(records []datautils.Record, pageNum uint64) ([]datautils.Record, []uint64) {
	numOfShow := uint64(5)
	if len(records) < 5 {
		numOfShow = uint64(len(records))
	}

	startNum := numOfShow * pageNum
	floatPage := float64(len(records)) / float64(numOfShow)
	PageCount := uint64(math.Ceil(floatPage))
	pages := []uint64{}
	for i := 1; i <= int(PageCount); i++ {
		pages = append(pages, uint64(i))
	}
	return records[startNum : startNum+numOfShow], pages
}
