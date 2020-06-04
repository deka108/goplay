package sqlite

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"

	"github.com/deka108/goplay/pkg/testutil"
	"github.com/spf13/viper"
	"log"
	"os"
	"testing"
)

type TestConfig struct{
	dbPath1 string
	dbPath2 string
}

var testConfig TestConfig

func TestMain(m *testing.M) {
	log.Println("Do stuff BEFORE the tests!")
	testutil.LoadConfig()
	testConfig = TestConfig{
		dbPath1: os.ExpandEnv(viper.GetString("db.TOKEN_DB")),
		dbPath2: os.ExpandEnv(viper.GetString("db.CREDENTIALS_DB")),
	}
	exitVal := m.Run()
	log.Println("Do stuff AFTER the tests!")

	os.Exit(exitVal)
}

func initDb(dbPath string) *sql.DB{
	fmt.Println(dbPath)

	// Create connection
	conn, err := sql.Open("sqlite3", dbPath)
	logError(err)

	return conn
}

func TestSqliteAllTables(t *testing.T) {
	conn := initDb(testConfig.dbPath1)
	defer conn.Close()

	// Prepare Query
	rows, err := conn.Query("SELECT name FROM sqlite_master where type='table'")
	logError(err)
	defer rows.Close()

	// Get Rows
	fmt.Println("Results")
	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		logError(err)
		fmt.Println(name)
	}
}

func TestSqliteGetColumns(t *testing.T){
	testCases := []struct{
		db_path string
		tbl_name string
	}{
		{db_path: testConfig.dbPath1, tbl_name: "access_tokens"},
		{db_path: testConfig.dbPath2, tbl_name: "credentials"},
	}

	for _, tc := range testCases{
		conn := initDb(tc.db_path)

		// Prepare Query
		rows, err := conn.Query(fmt.Sprintf("SELECT name, type FROM PRAGMA_TABLE_INFO('%s')", tc.tbl_name))
		logError(err)

		// Get Rows
		fmt.Println(fmt.Sprintf("Columns for Table %s", tc.tbl_name))
		for rows.Next(){
			var name, colType string
			err = rows.Scan(&name, &colType)
			logError(err)
			fmt.Println(name)
		}

		rows.Close()
		conn.Close()
	}
}

func TestSqliteGetCredentials(t *testing.T){
	conn := initDb(testConfig.dbPath2)
	defer conn.Close()

	// Prepare Query
	rows, err := conn.Query("SELECT account_id, value FROM credentials;")
	logError(err)
	defer rows.Close()

	// Get Rows
	fmt.Println("Results")
	for rows.Next(){
		var account_id string
		var value []byte
		err = rows.Scan(&account_id, &value)
		logError(err)
		fmt.Println(account_id, string(value))
	}
}

func logError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}