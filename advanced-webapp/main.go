package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
)

var validPath = regexp.MustCompile(`^/(ask|qa)/([a-zA-Z0-9\s-]+)$`)

func makeHndler(fn func(http.ResponseWriter, *http.Request, string, int) string) http.HandlerFunc {
	var i int
	return func(rw http.ResponseWriter, r *http.Request) {
		i++
		fmt.Println("this is run first")
		m := validPath.FindStringSubmatch(r.URL.Path)
		if len(m) < 3 || m == nil {
			http.NotFound(rw, r)
			return
		}
		ask := m[2]
		answer := fn(rw, r, ask, i)

		fmt.Println("this is run last")
		fmt.Println(answer, i)

	}
}

func viewHandler(w http.ResponseWriter, r *http.Request, q string, i int) string {
	if q == "can-i-go" {
		q = "yes!"
	} else {
		q = "where is your question!"
	}
	fmt.Fprintf(w, "Hi someone Answer yor question!: %s %d", q, i)
	return q

}

func main() {
	http.HandleFunc("/", makeHndler(viewHandler))

	log.Fatal(http.ListenAndServe(":8888", nil))

}
