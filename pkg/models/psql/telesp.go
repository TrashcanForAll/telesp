package psql

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"telesp/pkg/models"
)

var connection connParams = connParams{
	host:     "localhost",
	password: "1234",
	user:     "postgres",
	dbname:   "telesp",
	port:     "5432",
}

type connParams struct {
	user     string
	password string
	dbname   string
	host     string
	port     string
}

type TeleSp struct {
	DB *sql.DB
}

// TODO: creation all funcs of work with db
func OpenConn() (TeleSp, error) {

	connStr := func() (CS string) { //MARK: CS = shorter name of connStr
		// SASL -> "user=username password=password host=localhost dbname=mydb sslmode=disable"
		CS = fmt.Sprintf("user=%s port=%s password=%s host=%s dbname=%s sslmode=disable", connection.user, connection.port, connection.password, connection.host, connection.dbname)
		return CS
	}()

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("not connect")
		fmt.Println(err)
	}

	return TeleSp{DB: db}, err
}

func (m *TeleSp) SecondaryTablesQuery(storage *models.PersonData) []int {
	params := make([]int, 8)
	if storage.FirstName != "" {
		row := m.DB.QueryRow("SELECT id FROM firstname WHERE firstname_val = $1", storage.FirstName)
		err := row.Scan(&params[0])
		if err != nil {
			log.Fatal("SecondaryTablesQuery; FirstName ", err)
		}
	}
	if storage.LastName != "" {
		row := m.DB.QueryRow("SELECT id FROM lastname WHERE lastname_val = $1", storage.LastName)
		err := row.Scan(&params[1])
		if err != nil {
			log.Fatal("SecondaryTablesQuery; LastName ", err)
		}
	}
	if storage.MiddleName != "" {
		row := m.DB.QueryRow("SELECT id FROM middlename WHERE middlename_val = $1", storage.MiddleName)
		err := row.Scan(&params[2])
		if err != nil {
			log.Fatal("SecondaryTablesQuery; MiddleName ", err)
		}
	}
	if storage.Street != "" {
		row := m.DB.QueryRow("SELECT id FROM street WHERE street_val = $1", storage.Street)
		err := row.Scan(&params[3])
		if err != nil {
			log.Fatal("SecondaryTablesQuery; Street ", err)
		}
	}
	if storage.House != "" {
		row := m.DB.QueryRow("SELECT id FROM house WHERE house_val = $1", storage.House)
		err := row.Scan(&params[4])
		if err != nil {
			log.Fatal("SecondaryTablesQuery; House ", err)
		}
	}
	if storage.Building != "" {
		row := m.DB.QueryRow("SELECT id FROM building WHERE building_val = $1", storage.Building)
		err := row.Scan(&params[5])
		if err != nil {
			log.Fatal("SecondaryTablesQuery; Building ", err)
		}
	}
	if storage.Apartment != "" {
		row := m.DB.QueryRow("SELECT id FROM apartment WHERE apartment_val = $1", storage.Apartment)
		err := row.Scan(&params[6])
		if err != nil {
			log.Fatal("SecondaryTablesQuery; Apartment ", err)
		}
	}
	if storage.PhoneNumber != "" {
		row := m.DB.QueryRow("SELECT id FROM phonenumber WHERE building_val = $1", storage.PhoneNumber)
		err := row.Scan(&params[7])
		if err != nil {
			log.Fatal("SecondaryTablesQuery; PhoneNumber ", err)
		}
	}
	return params
}

// TODO: make a Insert func

// Insert
func (m *TeleSp) Insert() error {
	return nil
}

// Find - select func for looking for based offered params
// TODO: add [id] param into Get func
func (m *TeleSp) Get(storage *models.PersonData) []models.PersonData {
	//IdParams := m.SecondaryTablesQuery(storage)
	Query := CreateSqlQuery(storage)
	response := []models.PersonData{}

	rows, err := m.DB.Query(Query)
	if err != nil {
		log.Fatal("telesp.go; Error of scan string: ", err)
	}
	//
	for rows.Next() {
		p := models.PersonData{}
		rows.Scan(&p.FirstName, &p.LastName, &p.MiddleName, &p.Street, &p.House, &p.Building, &p.Apartment, &p.PhoneNumber)
		response = append(response, p)
	}

	return response
}
