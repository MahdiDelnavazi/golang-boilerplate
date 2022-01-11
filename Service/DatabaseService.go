package Service

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"net/url"
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
	//sslMode := "require"
	//if cfg.DisableTLS {
	//	sslMode = "disable"
	//}

	q := make(url.Values)
	q.Set("sslmode", "disable")
	q.Set("timezone", "utc")

	u := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(cfg.User, cfg.Password),
		Host:     cfg.Host,
		Path:     cfg.Name,
		RawQuery: q.Encode(),
	}

	db, err := sqlx.Open("postgres", u.String())
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetMaxOpenConns(cfg.MaxOpenConns)

	return db, nil
}
