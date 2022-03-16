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

func GetVideoObjectHandler(W http.ResponseWriter, R *http.Request) {
	log.Info("Fetching video...please wait")
	
	// mapper := (paramMap)(mux.Vars(R))
	// VideoId := mapper.get("videoId")
	// ChannelId := mapper.get("channelid")
	VideoId, ok := vars["videoId"]
    if !ok {
        log.Errorf("Video ID is missing in parameters")
    }
	ChannelId, ok1 := vars["channelId"]
	if !ok1 {
        log.Errorf("Channel ID is missing in parameters")
    }
	
	GetVideoObject(W, R, VideoId, ChannelId)
}

func initServer() *http.Server {
	//Creating the routers
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", myFunc)
	myRouter.HandleFunc("/{channelId:[0-9]+}/{videoId:[0-9]+}/video.ts", GetVideoObjectHandler).Methods("GET")
	http.ListenAndServe(":8080", myRouter)
}

func main() {
	server := initServer()
	shutdown(server)
	log.Infof("Application closed")
}