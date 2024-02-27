package utils

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type PageData struct {
	ID           int
	Image        string
	Name         string
	Members      []string
	CreationDate int
	FirstAlbum   string
	Locations    []string
	ConcertDates []string
	Relations    map[string][]string
}
type SearchData struct {
	ID           int                 `json:"id"`
	Image        string              `json:"image"`
	Name         string              `json:"name"`
	Members      []string            `json:"members"`
	CreationDate int                 `json:"creationDate"`
	FirstAlbum   string              `json:"firstAlbum"`
	Locations    []string            `json:"locations"`
	ConcertDates []string            `json:"concertDates"`
	Relations    map[string][]string `json:"relations"`
}

type ArtistFullData struct {
	ID           int                 `json:"id"`
	Image        string              `json:"image"`
	Name         string              `json:"name"`
	Members      []string            `json:"members"`
	CreationDate int                 `json:"creationDate"`
	FirstAlbum   string              `json:"firstAlbum"`
	Locations    []string            `json:"locations"`
	ConcertDates []string            `json:"concertDates"`
	Relations    map[string][]string `json:"relations"`
}

type MyArtist struct {
	ID           int                 `json:"id"`
	Image        string              `json:"image"`
	Name         string              `json:"name"`
	Members      []string            `json:"members"`
	CreationDate int                 `json:"creationDate"`
	FirstAlbum   string              `json:"firstAlbum"`
	Locations    string              `json:"locations"`
	ConcertDates string              `json:"concertDates"`
	Relations    map[string][]string `json:"relations"`
}

type MyLocation struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}
type LocationData struct {
	Index []MyLocation `json:"index"`
}

type MyConcertDate struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}
type ConcertDateData struct {
	Index []MyConcertDate `json:"index"`
}

type MyRelationDate struct {
	ID            int                 `json:"id"`
	DatesLocation map[string][]string `json:"datesLocations"`
}
type RelationData struct {
	Index []MyRelationDate `json:"index"`
}

func GetArtistsData() error {

	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return errors.New("error by get")
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New("error by reading")
	}
	json.Unmarshal(bytes, &Artists)
	return nil
}

func GetLocations() error {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		return errors.New("error by get")
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New("error by reading")
	}
	json.Unmarshal(bytes, &LocationsData)
	return nil
}

func GetDates() error {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/dates")
	if err != nil {
		return errors.New("error by get")
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New("error by reading")
	}
	json.Unmarshal(bytes, &ConcertDatesData)
	return nil
}

func GetRelations() error {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		return errors.New("error by get")
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New("error by reading")
	}
	json.Unmarshal(bytes, &RelationsData)
	return nil
}
