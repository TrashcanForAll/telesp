package psql

import (
	"fmt"
	"telesp/pkg/models"
)

func AllRows() string {
	Query := "SELECT firstname.firstname_val, lastname.lastname_val, middlename.middlename_val, street.street_val," +
		"house.house_val, apartment.apartment_val, building.building_val, phonenumber.phonenumber_val " +
		"FROM main " +
		"JOIN firstname ON firstname.id = main.first_name " +
		"JOIN lastname ON lastname.id = main.last_name " +
		"JOIN middlename ON middlename.id = main.middle_name " +
		"JOIN street ON street.id = main.street " +
		"JOIN house ON house.id = main.house " +
		"JOIN building ON building.id = main.building " +
		"JOIN apartment ON apartment.id = main.apartment " +
		"JOIN phonenumber ON phonenumber.id = main.phone_number;"
	return Query
}

func CreateSqlQuery(storage *models.PersonData) string {
	Query := "SELECT firstname.firstname_val, lastname.lastname_val, middlename.middlename_val, street.street_val," +
		"house.house_val, apartment.apartment_val, building.building_val, phonenumber.phonenumber_val " +
		"FROM main " +
		"JOIN firstname ON firstname.id = main.first_name " +
		"JOIN lastname ON lastname.id = main.last_name " +
		"JOIN middlename ON middlename.id = main.middle_name " +
		"JOIN street ON street.id = main.street " +
		"JOIN house ON house.id = main.house " +
		"JOIN building ON building.id = main.building " +
		"JOIN apartment ON apartment.id = main.apartment " +
		"JOIN phonenumber ON phonenumber.id = main.phone_number " +
		"WHERE "
	if storage.FirstName != "" {
		Query += "firstname.firstname_val = " + fmt.Sprintf("'%s'", storage.FirstName) + " and "
	}
	if storage.LastName != "" {
		Query += "lastname.lastname_val = " + fmt.Sprintf("'%s'", storage.LastName) + " and "
	}
	if storage.MiddleName != "" {
		Query += "middlename.middlename_val = " + fmt.Sprintf("'%s'", storage.MiddleName) + " and "
	}
	if storage.Street != "" {
		Query += "street.street_val = " + fmt.Sprintf("'%s'", storage.Street) + " and "
	}
	if storage.House != "" {
		Query += "house.house_val = " + fmt.Sprintf("'%s'", storage.House) + " and "
	}
	if storage.Building != "" {
		Query += "building.building_val = " + fmt.Sprintf("'%s'", storage.Building) + " and "
	}
	if storage.Apartment != "" {
		Query += "apartment.apartment_val = " + fmt.Sprintf("'%s'", storage.Apartment) + " and "
	}
	if storage.PhoneNumber != "" {
		Query += "phonenumber.phonenumber_val = " + fmt.Sprintf("'%s'", storage.PhoneNumber)
	}
	if Query[len(Query)-4:] == "and " {
		Query = Query[:len(Query)-4]
	}
	Query += ";"
	//fmt.Println("Query:", Query)
	return Query
}
