package main

import (
	"fmt",
	"log",
	"net/http",
	"github.com/gorilla/mux"
)

type paramMap map[string]string

funv (pm paramMap) fetchValue (key string) string {
	return pm[key]
}

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

func GetVideoObjectHandler(W http.ResponseWriter, R *http.Request) {
	log.Info("Fetching video...please wait")
	
	mapper := (paramMap)(mux.Vars(R))
	VideoId := mapper.get("videoId")
	ChannelId := mapper.get("channelid")
	GetVideoObject(W, R, VideoId, ChannelId)
}

func initServer() *http.Server {
	//Creating the routers
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", myFunc)
	myRouter.HandleFunc("/{channelId:[0-9]+}/{videoId:[0-9]+}/video.ts", GetVideoObjectHandler).Methods("GET")
}

func main() {
	server := initServer()
	shutdown(server)
	log.Infof("Application closed")
}