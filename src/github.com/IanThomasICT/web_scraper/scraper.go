package main

import (
	"log"
	"net/http"
	"encoding/json"
	"github.com/gocolly/colly"
)

func ping(wri http.ResponseWriter, req *http.Request){
	log.Println("Ping")
	wri.Write([]byte("ping"))
}

func main() {
	addr := ":7171"

	http.HandleFunc("/search", getData)
	http.HandleFunc("/ping", ping)

	log.Println("listening on", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func getData(w http.ResponseWriter, r *http.Request) {
	URL := r.URL.Query().Get("url")
	log.Println("visiting", URL)

	c := colly.NewCollector()

	var response []string

	//onHTML function allows the collector to use a callback function when the specific HTML tag is reached 
	//in this case whenever our collector finds an
	//anchor tag with href it will call the anonymous function
	// specified below which will get the info from the href and append it to our slice
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Request.AbsoluteURL(e.Attr("href"))
		if link != "" {
			response = append(response, link)
		}
	})

	//Command to visit the website
	c.Visit(URL)

	// parse our response slice into JSON format
	b, err := json.Marshal(response)
	if err != nil {
		log.Println("failed to serialize response:", err)
		return
	}
	// Add some header and write the body for our endpoint
	w.Header().Add("Content-Type", "application/json")
	w.Write(b)
}