package Api

import (
	"html/template"
	"net/http"
	"strconv"

	Ext "Api/Ext"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if Ext.ApiError {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	} else if r.URL.Path != "/" {
		ErrorHandler(w, http.StatusNotFound)
		return
	} else if r.Method != "GET" {
		ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	template, err := template.ParseFiles("../Pages/home.html")
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	err = template.Execute(w, Artists)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
	}
}

func Artist(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	template, err := template.ParseFiles("../Pages/artists.html")
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 || id > len(Artists) {
		ErrorHandler(w, http.StatusBadRequest)
		return
	}
	data := Ext.DATA{
		Artist:    Artists[id-1],
		Locations: Locations.Index[id-1].Locations,
		Dates:     Dates.Index[id-1].Dates,
		Relations: Relations.Index[id-1].DatesLocations,
	}
	err = template.Execute(w, data)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
}

func ErrorHandler(w http.ResponseWriter, n int) {
	w.WriteHeader(n)
	template, err := template.ParseFiles("../Pages/error.html")
	if err != nil {
		http.Error(w, Errors[n].Tittle, n)
		return
	}
	err = template.Execute(w, Errors[n])
	if err != nil {
		http.Error(w, Errors[n].Tittle, n)
		return
	}
}
