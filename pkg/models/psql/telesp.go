package psql

import (
	"database/sql"
	"fmt"
	"log"
	"telesp/pkg/models"
)

var connection connParams = connParams{
	host:     "localhost",
	password: "1234",
	user:     "postgres",
	dbname:   "postgres",
}

type connParams struct {
	user     string
	password string
	dbname   string
	host     string
}

type TeleSp struct {
	DB *sql.DB
}

// TODO: creation all funcs of work with db
func OpenConn() (*sql.DB, error) {
	connStr := func() (CS string) { //MARK: CS = shorter name of connStr
		// SASL -> "user=username password=password host=localhost dbname=mydb sslmode=disable"
		CS = fmt.Sprintf("user=%s password=%s host=%V dbname=%s sslmode=disable", connection.user, connection.password, connection.host, connection.dbname)
		return CS
	}()
	db, err := sql.Open("postgres", connStr)
	return db, err
}

func ExecuteQuery(act string) {
	db, err := OpenConn()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	testTeleSp := &TeleSp{DB: db}

	// Start call functions
	//switch act {
	//case "GET":
	//	testTeleSp.Get()
	//}
	testTeleSp.Get()
}

// TODO: make a Insert func

// Insert
func (m *TeleSp) Insert() error {
	return nil
}

// Find - select func for looking for based offered params
func (m *TeleSp) Get() (*models.PersonData, error) {
	tx, err := m.DB.Begin()
	if err != nil {
		return nil, err
	}

	someParams := 1
	rows, err := tx.Query("SELECT * FROM person WHERE id = ?", someParams)
	if err != nil {
		log.Fatal("Ошибка выполнения запроса:", err)
	}
	defer rows.Close()

	// Обрабатываем результат запроса
	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal("Ошибка сканирования строки:", err)
		}
		fmt.Printf("id:%d name:%s\n", id, name)
	}

	// Проверяем на ошибки после обработки строк
	if err := rows.Err(); err != nil {
		log.Fatal("Ошибка при итерации по строкам: ", err)
	}

	return nil, nil
}

//
