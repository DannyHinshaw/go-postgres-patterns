package env

import (
	"fmt"
	"net/url"
)

// PostgresConfig holds the most commonly used postgres environment vars.
type PostgresConfig struct {
	Port    string `envconfig:"POSTGRES_PORT" default:"5432"`
	Host    string `envconfig:"POSTGRES_HOST" default:"localhost"`
	User    string `envconfig:"POSTGRES_USER" default:"postgres"`
	Pass    string `envconfig:"POSTGRES_PASS" default:"postgres"`
	Name    string `envconfig:"DATABASE_NAME" default:"postgres"`
	TZone   string `envconfig:"POSTGRES_TZONE" default:"UTC"`
	SSLMode string `envconfig:"POSTGRES_SSL" default:"false"`
}

// ConnectionURI builds a valid postgres database connection URI from the PostgresConfig field values.
// Spec: https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-CONNSTRING
func (c *PostgresConfig) ConnectionURI() (string, error) {
	u, err := url.Parse("postgres:///")
	if err != nil {
		return "", fmt.Errorf("error parsing postgres base url: %w", err)
	}

	// An empty host string will cause the url builder
	// to remove the third "/" from "postgres:///" base.
	// We assume that if someone overrode the default
	// value for this field (localhost) to an empty
	// string they want the "bind-all" address.
	if c.Host == "" {
		c.Host = "0.0.0.0"
	}

	u.Host = url.QueryEscape(c.Host)
	u.Path = url.QueryEscape(c.Name)

	q := u.Query()
	if c.Port != "" {
		q.Set("port", c.Port)
	}
	if c.User != "" {
		q.Set("username", c.User)
	}
	if c.Pass != "" {
		q.Set("password", c.Pass)
	}
	if c.SSLMode != "" {
		q.Set("sslmode", c.SSLMode)
	}
	if c.TZone != "" {
		q.Set("timezone", c.TZone)
	}

	u.RawQuery = q.Encode()

	return u.String(), nil
}
