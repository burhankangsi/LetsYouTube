package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/burhankangsi/LetsYouTube/flash_api"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type File struct {
	id         string
	name       string
	channelId  string
	length     string
	uploadDate string
}

var allFiles []File

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

	item, err := flash_api.GetVideoObject(W, R, VideoId, ChannelId)
	if err != nil {
		log.Fatal("Error in getting the video. Please try again")
		return
	}
	json.NewEncoder(W).Encode(item)
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
	http.ListenAndServe("https://vocal-starship-53117c.netlify.app/endpoint", myRouter)
}

func main() {
	server := initServer()
	shutdown(server)
	log.Infof("Application closed")
}
