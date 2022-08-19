package handler

import (
	"AqiWithCron/database"
	"AqiWithCron/database/helper"
	"AqiWithCron/models"
	"encoding/json"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"net/http"
)

func InsertAqi() {
	url := "https://api.ambeedata.com/ghg/latest/by-place?place=Mumbai"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("x-api-key", "6f2ba3e4d229d9eeb28e15843d747212b993f47b0f284f400d629b9edb498b7b")
	req.Header.Add("Content-type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		logrus.Error("GetDataFromApi: error in getting data")
		return
	}

	defer res.Body.Close()
	var data models.CityAqiDetails
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		logrus.Error("GetDataFromApi: error in unmarshal data %v", err)
		return
	}

	txErr := database.Tx(func(tx *sqlx.Tx) error {
		err = helper.InsertAqiDetails(models.City, data.Data, tx)
		if err != nil {
			logrus.Error("InsertAqiDetails: Error in adding data %v", err)
			//writer.WriteHeader(http.StatusBadRequest)
			return err
		}
		return err
	})
	if txErr != nil {
		//writer.WriteHeader(http.StatusInternalServerError)
		return
	}

}
