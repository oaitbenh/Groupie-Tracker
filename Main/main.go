package main

import (
	"fmt"
	"net/http"

	Root "Api/Root"
)

func main() {
	ApiFetchAllData("https://groupietrackers.herokuapp.com/api")
	http.HandleFunc("/", Root.Home)
	http.HandleFunc("/artists/{id}", Root.Artist)
	http.Handle("/Media/", http.StripPrefix("/Media/", http.FileServer(http.Dir("../Media"))))
	fmt.Println("Listen to Localhost at Port 4000 http://localhost:4000")
	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	}
}

func ApiFetchAllData(Url string) {
	Root.ApiData = Root.ApiData.GetAllData(Url)
	Root.Artists = Root.Artists.GetAllData(Root.ApiData.ArtistsUrl)
	Root.Dates = Root.Dates.GetAllData(Root.ApiData.DatesUrl)
	Root.Locations = Root.Locations.GetAllData(Root.ApiData.LocationsUrl)
	Root.Relations = Root.Relations.GetAllData(Root.ApiData.RelationUrl)
}
