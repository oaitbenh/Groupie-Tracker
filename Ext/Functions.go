package Api

import (
	"encoding/json"
	"io"
	"net/http"
)

var ApiError bool

func (Api Api) GetAllData(Url string) Api {
	Json_Data := GetJson(Url)
	json.Unmarshal(Json_Data, &Api)
	return Api
}

func (Art Artists) GetAllData(Url string) Artists {
	Json_Data := GetJson(Url)
	json.Unmarshal(Json_Data, &Art)
	return Art
}

func (Loc AllLocations) GetAllData(Url string) AllLocations {
	Json_Data := GetJson(Url)
	json.Unmarshal(Json_Data, &Loc)
	return Loc
}

func (Dates AllDates) GetAllData(Url string) AllDates {
	Json_Data := GetJson(Url)
	json.Unmarshal(Json_Data, &Dates)
	return Dates
}

func (Relations AllRelations) GetAllData(Url string) AllRelations {
	Json_Data := GetJson(Url)
	json.Unmarshal(Json_Data, &Relations)
	return Relations
}

func GetJson(Url string) []byte {
	Response_Data, err := http.Get(Url)
	if err != nil {
		ApiError = true
		return nil
	}
	Json_Data, err := io.ReadAll(Response_Data.Body)
	if err != nil {
		ApiError = true
		return nil
	}
	return Json_Data
}
