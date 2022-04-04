package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
	"math/rand"
	"github.com/burhankangsi/LetsYouTube/bucket_api"
)

type File struct {
	id         string
	name       string
	channelId  string
	length     string
	uploadDate string
}

var allFiles []File
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

func UploadVideoHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("Fetching video...please wait")

	params := mux.Vars(r)
	chanId := params["channelId"]
	var item File
	obj := json.NewDecoder(r.Body).Decode(&item)
	item.id = strconv.Itoa(rand.Intn(1000000))

	allFiles = append(allFiles, item)
	json.NewEncoder(w).Encode(item)

}

func GetVideoObjectHandler(W http.ResponseWriter, R *http.Request) {
	log.Info("Fetching video...please wait")
	W.Header().Set("Content-Type", "application/json")
	params := mux.Vars(R)
	VideoId, ok := params["videoId"]
    if !ok {
        log.Errorf("Video ID is missing in parameters")
    }
	ChannelId, ok1 := params["channelId"]
	if !ok1 {
        log.Errorf("Channel ID is missing in parameters")
    }
	
	item, err:= GetVideoObject(W, R, VideoId, ChannelId)
	if err != nil {
		log.Fatal("Error in getting the video. Please try again")
		return
	}
	json.NewEncoder(w).Encode(item)
}

func initUpload() {
	go func() {
		saveToBucket()
	}
}

func Demo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Endpoint created")
}

func initServer() *http.Server {
	//Creating the routers
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", Demo).Methods("GET")
	myRouter.HandleFunc("/{channelId}/{videoId}/video.ts", GetVideoObjectHandler).Methods("GET")
	//myRouter.HandleFunc("/{channelId:[0-9]+}/{videoId:[0-9]+}/video.ts", GetVideoObjectHandler).Methods("GET")
	myRouter.HandleFunc("/video/{channelId}", UploadVideoHandler).Methods("POST")
	initUpload()
	http.ListenAndServe("https://vocal-starship-53117c.netlify.app/endpoint", myRouter)
}

func main() {
	server := initServer()
	shutdown(server)
	log.Infof("Application closed")
}