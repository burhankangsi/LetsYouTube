package main

import (
	"fmt",
	"log",
	"net/http",
	"github.com/gorilla/mux"
)

func myFunc(w http.ResponseWriter, r *http.Request) {
	log.info("My Handler")
}

// func createNewArticle(w http.ResponseWriter, r *http.Request) {
//     // get the body of our POST request
//     // unmarshal this into a new Article struct
//     // append this to our Articles array.    
//     reqBody, _ := ioutil.ReadAll(r.Body)
//     var article Article 
//     json.Unmarshal(reqBody, &article)
//     // update our global Articles array to include
//     // our new Article
//     Articles = append(Articles, article)

//     json.NewEncoder(w).Encode(article)
// }

func handleRequests() {

	// myRouter := mux.NewRouter().StrictSlash(true)
	// myRouter.HandleFunc("/", myFunc)
    // myRouter.HandleFunc("/all", returnAllArticles)
	//myRouter.HandleFunc("/article/{id}", returnSingleArticle)
	//myRouter.HandleFunc("/article", createNewArticle).Methods("POST")

	http.HandleFunc("/", myFunc)
	log.info("Handler called")
}

func main() {
	handleRequests()
}