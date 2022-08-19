package helper

import (
	"AqiWithCron/models"
	"github.com/elgris/sqrl"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strconv"
)

func InsertAqiDetails(cityName string, aqiDetails []models.Combination, tx *sqlx.Tx) error {
	psql := sqrl.StatementBuilder.PlaceholderFormat(sqrl.Dollar)
	insertQuery := psql.Insert("aqi_details").Columns("city_name", "ozone_value", "ozone_last_updated", "ozone_units", "co2_value", "co2_last_updated", "co2_units", "ch4_value", "ch4_last_updated", "ch4_unit")
	for _, details := range aqiDetails {
		ozoneValue, _ := strconv.ParseFloat(details.Ozone.Value, 8)
		co2Value, _ := strconv.ParseFloat(details.Co2.Value, 8)
		ch4Value, _ := strconv.ParseFloat(details.Ch4.Value, 8)
		insertQuery.Values(cityName, ozoneValue, details.Ozone.LastUpdated, details.Ozone.Units,
			co2Value, details.Co2.LastUpdated, details.Co2.Units, ch4Value, details.Ch4.LastUpdated, details.Ch4.Units)
	}

	sql, args, err := insertQuery.ToSql()
	if err != nil {
		logrus.Error("InsertAqiDetails: Error in making the query")
		return err
	}

	_, err = tx.Exec(sql, args...)
	if err != nil {
		logrus.Error("InsertAqiDetails: Error in Adding product")
		return err
	}

	return nil
}
