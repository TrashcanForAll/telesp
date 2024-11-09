package psql

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
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
func (m *TeleSp) Insert() error {
	return nil
}

// Find - select func for looking for based offered params
// TODO: add [id] param into Get func
func (m *TeleSp) Get(storage *models.TestPerson) {

	//someParams := 1
	//row := m.DB.QueryRow("SELECT id, name FROM main WHERE id = $1", someParams)
	//
	////err := row.Scan(&storage.Id, &storage.Name)
	//if err != nil {
	//	log.Fatal("telesp.go; Error of scan string: ", err)
	//}

}
