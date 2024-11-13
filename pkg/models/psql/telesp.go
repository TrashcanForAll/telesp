package psql

import (
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	//_ "github.com/lib/pq"
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

// TODO: make a Insert func

// Insert
func (m *TeleSp) Insert(storage *models.PersonData) error {
	data := []string{
		storage.FirstName,
		storage.LastName,
		storage.MiddleName,
		storage.Street,
		storage.House,
		storage.Building,
		storage.Apartment,
		storage.PhoneNumber,
	}
	_, err := m.DB.Exec("CALL insert_to_main($1);", pq.Array(data))
	if err != nil {
		return err
	}
	return nil
}

// Find - select func for looking for based offered params
// TODO: add [id] param into Get func
func (m *TeleSp) Get(storage *models.PersonData) []models.PersonData {
	//IdParams := m.SecondaryTablesQuery(storage)
	var Query string
	if storage.FirstName+storage.LastName+storage.MiddleName+storage.Street+
		storage.House+storage.Building+storage.Apartment+storage.PhoneNumber == "" {
		Query = AllRows()
	} else {
		Query = CreateSqlQuery(storage)
	}
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
