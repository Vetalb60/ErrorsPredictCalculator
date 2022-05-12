package main

import (
	"database/sql"
	"errors"
	"github.com/cenk/backoff"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
	"os"
)

const migrationsPath = "file://migrations/sql"

const (
	MYSQL_ROOT_PASSWORD = "MYSQL_ROOT_PASSWORD"
	MYSQL_POD_IP        = "MYSQL_POD_IP"
	MYSQL_DATABASE_PORT = "MYSQL_PORT"
	MYSQL_USER          = "MYSQL_USER"
	MYSQL_TABLE_NAME    = "MYSQL_TABLE_NAME"
)

// Data base configuration.
type InitDBInfo struct {
	User_name_     string
	Password_      string
	Ip_address_    string
	Port_          string
	Database_name_ string
	Table_name_    string
}

func main() {
	// Open connection.
	conn, driver := getConnection()

	// Close connection.
	defer func() {
		if conn != nil {
			_ = conn.Close()
		}
	}()

	// Establish connection to a database.
	establishConnectionWithRetry(conn)

	// Starting migration job.
	err := migrateSQL(conn, driver)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("Migration successfully finished.")
	}
}

func getConnection() (*sql.DB, string) {
	info, err := auth()
	if err != nil {
		log.Fatalln(err.Error())
	}

	driver := "mysql"
	address := info.Ip_address_
	username := info.User_name_
	password := info.Password_
	database := "mysql"

	// Open may just validate its arguments without creating a connection to the database.
	sqlconn, err := sql.Open(driver, username+":"+password+"@tcp("+address+")/"+database)
	if err != nil {
		log.Fatalln("Cannot establish connection to a database")
	}

	return sqlconn, driver
}

// This function executes the migration scripts.
func migrateSQL(conn *sql.DB, driverName string) error {
	driver, _ := mysql.WithInstance(conn, &mysql.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		migrationsPath,
		driverName,
		driver,
	)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}

func establishConnectionWithRetry(conn *sql.DB) {
	b := backoff.NewExponentialBackOff()
	// We wait forever until the connection will be established.
	// In practice k8s will kill the pod if it takes too long.
	b.MaxElapsedTime = 0

	_ = backoff.Retry(func() error {
		log.Printf("Connecting to a database ...")
		// Ping verifies a connection to the database is still alive,
		// establishing a connection if necessary.
		if errPing := conn.Ping(); errPing != nil {
			log.Printf("ping failed %s", errPing.Error())
		}
		return nil
	}, b)
}

// auth
//	Function for providing authentication and authorization in the database.
func auth() (InitDBInfo, error) {
	var connection InitDBInfo
	if err := setEnvironmentVar(&connection); err != nil {
		return InitDBInfo{}, err
	}

	return connection, nil
}

func setEnvironmentVar(connection *InitDBInfo) error {
	var ok bool

	if connection.Ip_address_, ok = os.LookupEnv(MYSQL_POD_IP); ok != true {
		return errors.New("error get environment MYSQL_POD_IP")
	}
	if connection.Password_, ok = os.LookupEnv(MYSQL_ROOT_PASSWORD); ok != true {
		return errors.New("error get environment MYSQL_ROOT_PASSWORD")
	}
	if connection.Port_, ok = os.LookupEnv(MYSQL_DATABASE_PORT); ok != true {
		return errors.New("error get environment MYSQL_PORT")
	}
	if connection.User_name_, ok = os.LookupEnv(MYSQL_USER); ok != true {
		return errors.New("error get environment MYSQL_USER")
	}
	if connection.Table_name_, ok = os.LookupEnv(MYSQL_TABLE_NAME); ok != true {
		return errors.New("error get environment MYSQL_TABLE_NAME")
	}

	return nil
}
