package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"text/template"
)

type joke struct {
	category  string
	blacklist string
	word      string
}

var data joke
var count int = 0
var blacklist string = ""
var category string = ""
var word string = ""
var url string

func c(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		err := tp1.ExecuteTemplate(w, "index.gohtml", nil)
		if err != nil {
			log.Println(err)

		}

	case "POST":
		error := r.ParseForm()
		if error != nil {
			log.Println(error)
		} else {
			data = joke{
				category:  r.Form.Get("category"),
				blacklist: r.Form.Get("blacklist"),
				word:      r.Form.Get("word"),
			}
			blacklist = data.blacklist
			category = data.category
			word = data.word
			url = "https://v2.jokeapi.dev/joke/" + category + "?blacklistFlags=" + blacklist + "&contains=" + word
			log.Println(url)
			res, err := http.Get(url)
			if err != nil {
				log.Println(err)
			} else {
				body, err := io.ReadAll(res.Body)
				res.Body.Close()
				if res.StatusCode > 299 {
					log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
				}
				if err != nil {
					log.Println(err)
				} else {
					fmt.Printf("%s", body)
					fmt.Printf("%T", body)
					io.WriteString(w, "The joke can be seen in the console")
					// for i, v := range body {
					// 	fmt.Println(i)
					// 	fmt.Println(v)
					// }
				}

			}
		}
	}
}

func d(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		io.WriteString(w, "You won't get your dog")
	}
}

func e(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		io.WriteString(w, "You won't get your cat")
	}

}

var tp1 *template.Template

func init() {
	tp1 = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {

	http.HandleFunc("/", c)
	http.HandleFunc("/dog", d)
	http.HandleFunc("/cat", e)
	http.ListenAndServe(":3000", nil)
	log.Println(url)

}
