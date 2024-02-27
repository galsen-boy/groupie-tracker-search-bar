package utils

import (
	// "fmt"
	"html/template"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

var Artists []MyArtist
var ArtistData []ArtistFullData
var LocationsData LocationData
var ConcertDatesData ConcertDateData
var RelationsData RelationData

func GetData() {
	GetArtistsData()
	GetLocations()
	GetDates()
	GetRelations()
	var template ArtistFullData
	for i := range Artists {
		template.ID = i + 1
		template.Image = Artists[i].Image
		template.Name = Artists[i].Name
		template.Members = Artists[i].Members
		template.CreationDate = Artists[i].CreationDate
		template.FirstAlbum = Artists[i].FirstAlbum
		template.Locations = LocationsData.Index[i].Locations
		template.ConcertDates = ConcertDatesData.Index[i].Dates
		template.Relations = RelationsData.Index[i].DatesLocation

		ArtistData = append(ArtistData, template)
	}

}
func SearchHandler(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/search" {
		error404(res)
		return
	}
	if req.Method != "POST" {
		error405(res)
		return
	}
	search := req.FormValue("search")
	SearchDatas, found := Search(search)
	if !found {
		template, err := template.ParseFiles("./templates/notfound.html")
		if err != nil {
			error500(res)
			return
		}
		template.Execute(res, nil)
	} else {
		template, err := template.ParseFiles("./templates/search.html")
		if err != nil {
			error500(res)
			return
		}
		template.Execute(res, SearchDatas)

	}

}
func MainHandler(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		error404(res)
		return
	}
	if req.Method != "GET" {
		error405(res)
		return
	}
	template, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		error500(res)
		return
	}
	template.Execute(res, ArtistData)
}
func ArtistHandler(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/artist/" {
		error404(res)
		return
	}
	if req.Method != "GET" {
		error405(res)
		return
	}
	template, err := template.ParseFiles("./templates/ArtistPage.html")
	if err != nil {
		error500(res)
		return
	}
	queryParams := req.URL.Query() // Obtient les paramètres de la requête dans un map

	id := queryParams.Get("id")
	data := PageData{}

	a, err := strconv.Atoi(id)
	for _, v := range ArtistData {
		if v.ID == a {
			data.Name = v.Name
			data.Members = v.Members
			data.Image = v.Image
			data.CreationDate = v.CreationDate
			data.ConcertDates = v.ConcertDates
			data.FirstAlbum = v.FirstAlbum
			data.Locations = v.Locations
			data.Relations = v.Relations
		}
	}
	if err != nil || a < 0 || a == 0 {
		error400(res)
		return
	}

	template.Execute(res, data)
}

func getDatabyId(id int) ArtistFullData {
	var data ArtistFullData

	for i := range Artists {
		if i == id {
			data.ID = Artists[i].ID
			data.Image = Artists[i].Image
			data.Name = Artists[i].Name
			data.Members = Artists[i].Members
			data.CreationDate = Artists[i].CreationDate
			data.FirstAlbum = Artists[i].FirstAlbum
			data.Locations = LocationsData.Index[i].Locations
			data.ConcertDates = ConcertDatesData.Index[i].Dates
			data.Relations = RelationsData.Index[i].DatesLocation
			break
		}
	}
	return data
}

func Search(search string) ([]ArtistFullData, bool) {
	if search == "" {
		return ArtistData, true
	}
	var resultSearch []ArtistFullData
	search = strings.ToLower(search)
	reg := regexp.MustCompile(`^` + search)
	found := false
	for i := range ArtistData {
		temp := strings.ToLower(ArtistData[i].Name)
		temp2 := strings.ToLower(ArtistData[i].FirstAlbum)
		temp3 := strings.ToLower(strconv.Itoa(ArtistData[i].CreationDate))
		if reg.Match([]byte(temp)) || reg.Match([]byte(temp2)) || reg.Match([]byte(temp3)) {
			resultSearch = append(resultSearch, getDatabyId(i))
			found = true
		}
		dar := ArtistData[i].Locations
		for _, v := range dar {
			temp = strings.ToLower(v)
			if reg.Match([]byte(temp)) {
				resultSearch = append(resultSearch, getDatabyId(i))
				found = true
				break
			}
		}
		dar1 := ArtistData[i].Members
		for _, v := range dar1 {
			temp = strings.ToLower(v)
			if reg.Match([]byte(temp)) {
				resultSearch = append(resultSearch, getDatabyId(i))
				found = true
				break
			}
		}
	}
	// fmt.Println(resultSearch)
	return resultSearch, found
}
