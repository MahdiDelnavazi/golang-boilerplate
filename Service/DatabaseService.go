package Service

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"net/url"
	"strings"
)

// DatabaseConfig Config is the required properties to use the database.
type DatabaseConfig struct {
	User         string
	Password     string
	Host         string
	Name         string
	MaxIdleConns int
	MaxOpenConns int
	DisableTLS   bool
}

// DatabaseOpen Open knows how to open a database connection based on the configuration.
func DatabaseOpen(cfg DatabaseConfig) (*sqlx.DB, error) {
	sslMode := "disable"
	if cfg.DisableTLS {
		sslMode = "disable"
	}

	q := make(url.Values)
	q.Set("sslmode", sslMode)
	q.Set("timezone", "utc")
	q.Set("user", cfg.User)
	q.Set("password", cfg.Password)
	q.Set("dbname", cfg.Name)

	databaseConfig := strings.Replace(q.Encode(), "&", " ", -1)
	fmt.Println("TESTSSTSTST: ", databaseConfig)
	db, err := sqlx.Open("postgres", databaseConfig)
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetMaxOpenConns(cfg.MaxOpenConns)

	return db, nil
}

// queryString provides a pretty print version of the query and parameters.
func queryString(query string, args ...interface{}) string {
	query, params, err := sqlx.Named(query, args)
	if err != nil {
		return err.Error()
	}

	for _, param := range params {
		var value string
		switch v := param.(type) {
		case string:
			value = fmt.Sprintf("%q", v)
		case []byte:
			value = fmt.Sprintf("%q", string(v))
		default:
			value = fmt.Sprintf("%v", v)
		}
		query = strings.Replace(query, "?", value, 1)
	}

	query = strings.ReplaceAll(query, "\t", "")
	query = strings.ReplaceAll(query, "\n", " ")

	return strings.Trim(query, " ")
}
