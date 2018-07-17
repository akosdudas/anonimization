package anondb

import (
	"anonmodel"
	"fmt"
	"math"
	"regexp"
	"strings"

	"github.com/globalsign/mgo/bson"
)

//GPSBoundary same
type GPSBoundary struct {
	*anonmodel.GPSBoundary
}

//CreateCoordHelpers interprets each records to coordinate pairs
func CreateCoordHelpers(fieldName string) error {
	LongitudeField := "__coords_longitude_" + fieldName
	LatitudeField := "__coords_latitude_" + fieldName
	var err error
	/*value := bson......[fieldname] <-- string
	LongitudeField, LatitudeField, err = readCoords(value, findFormat(value))
	*/
	return err
}

//FindCoordDif ...
func FindCoordDif(name string, LL string, idx uint64) (first float64, second float64) {
	fieldName := "__coords_" + LL + "_" + name
	return
}

//FindFormat ... (coord string) (format string)
func FindFormat(coord string) (format string) {
	format = "ERR"
	DDFormat := "/([-+]?[1-9]\\d*([,. ]\\d\\d)*)°?[ ,]*[-+]?[1-9]\\d*([,. ]\\d\\d*)*°?/"
	DMSFormat := "/[NS][ ]*[1-9]\\d*°[ ]*[1-9]\\d*['][ ]*[1-9]\\d*([,.]\\d*)?(\"|'')[ ,]*[WE][ ]*[1-9]\\d*°[ ]*[1-9]\\d*['][ ]*[1-9]\\d*([,.]\\d*)?(\"|'')[ \n,.]*/"
	match, _ := regexp.MatchString(DDFormat, coord)
	if match {
		format = "DD"
	}
	match, _ = regexp.MatchString(DMSFormat, coord)
	if match {
		format = "DMS"
	}
	return
}

func readDMS(coords string) (Latitude float64, Longitude float64, err error) {
	var Lat, Lon uint8
	var dLat, mLat, sLat, dLon, mLon, sLon int
	coords = strings.ToUpper(coords)
	strings.Replace(coords, "''", "\"", -1)
	strings.Replace(coords, " ", "", -1)
	strings.Replace(coords, ",", "", -1)
	fmt.Sscanf(coords, "%c%d°%d'%d\"%c%d°%d'%d\"", &Lat, &dLat, &mLat, &sLat, &Lon, &dLon, &mLon, &sLon)

	Latitude = float64(dLat) + float64(mLat)/60 + float64(sLat)/360
	if Lat == 'W' {
		Latitude = -Latitude
	} else if Lat != 'E' {
		err = fmt.Errorf("Bad coord string")
		return
	}

	Longitude = float64(dLon) + float64(mLon)/60 + float64(sLon)/360
	if Lat == 'S' {
		Longitude = -Longitude
	} else if Lat != 'N' {
		err = fmt.Errorf("Bad coord string")
		return
	}
	err = nil
	return
}

func readDD(coords string) (Latitude float64, Longitude float64, err error) {
	strings.Replace(coords, " ", "", -1)
	fmt.Sscanf(coords, "%f,%f", &Latitude, &Longitude)
	return
}

func readERROR(coords string) (Latitude float64, Longitude float64, err error) {
	Latitude = math.NaN()
	Longitude = math.NaN()
	err = fmt.Errorf("Unrecognised coord format")
	return
}

//ReadCordsValue a
func ReadCordsValue(coords string, format string) (Latitude, Longitude float64, err error) {
	if format == "DD" {
		return readDD(coords)
	} else if format == "DMS" {
		return readDMS(coords)
	} else {
		return readERROR(coords)
	}
}

func (b *GPSBoundary) setMatch(fieldName string, match *[]bson.M) {
	b.setMatch("__coords_longitude_"+fieldName, match)
	b.setMatch("__coords_latitude_"+fieldName, match)
}

func (b *GPSBoundary) getResult(fieldName string, mainGroupResult bson.M, queryResult bson.M) interface{} {
	return nil
}
