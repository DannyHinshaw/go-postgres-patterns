package env_test

import (
	"testing"

	"github.com/kelseyhightower/envconfig"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/dannyhinshaw/pg-patterns/internal/env"
)

// pointStr returns a pointer to a string.
func pointStr(s string) *string {
	return &s
}

func TestPostgresConfig_ConnectionURI(t *testing.T) {
	emptyStr := pointStr("")

	testCases := map[string]struct {
		host *string
		port *string
		name *string
		user *string
		pass *string
		sslm *string
		zone *string
		want string
	}{
		"envconfig defaults produce URI": {
			want: "postgres://localhost/postgres?password=postgres&port=5432&sslmode=false&timezone=UTC&username=postgres",
		},
		"all overridden to empty strings still sets the host to 0.0.0.0": {
			host: emptyStr,
			port: emptyStr,
			name: emptyStr,
			user: emptyStr,
			pass: emptyStr,
			sslm: emptyStr,
			zone: emptyStr,
			want: "postgres://0.0.0.0",
		},
		"only database name still creates valid connection URI": {
			host: emptyStr,
			port: emptyStr,
			name: pointStr("somedb"),
			user: emptyStr,
			pass: emptyStr,
			sslm: emptyStr,
			zone: emptyStr,
			want: "postgres://0.0.0.0/somedb",
		},
		"only host provided creates valid connection URI": {
			host: pointStr("127.0.0.1"),
			port: emptyStr,
			name: emptyStr,
			user: emptyStr,
			pass: emptyStr,
			sslm: emptyStr,
			zone: emptyStr,
			want: "postgres://127.0.0.1",
		},
		"database connection URI is fully query escaped": {
			host: pointStr("h//h"),
			port: pointStr("p::p"),
			name: pointStr("n@@n"),
			user: pointStr("u::u"),
			pass: pointStr("p..p"),
			sslm: pointStr("s\\s"),
			zone: pointStr("z&&z"),
			want: "postgres://h%252F%252Fh/n%2540%2540n?password=p..p&port=p%3A%3Ap&sslmode=s%5Cs&timezone=z%26%26z&username=u%3A%3Au",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			a := assert.New(t)
			r := require.New(t)

			pgConfig := env.PostgresConfig{}
			err := envconfig.Process("go-postgres-patterns-test", &pgConfig)
			r.NoError(err)

			if tc.host != nil {
				pgConfig.Host = *tc.host
			}
			if tc.port != nil {
				pgConfig.Port = *tc.port
			}
			if tc.name != nil {
				pgConfig.Name = *tc.name
			}
			if tc.user != nil {
				pgConfig.User = *tc.user
			}
			if tc.pass != nil {
				pgConfig.Pass = *tc.pass
			}
			if tc.sslm != nil {
				pgConfig.SSLMode = *tc.sslm
			}
			if tc.zone != nil {
				pgConfig.TZone = *tc.zone
			}

			cs, err := pgConfig.ConnectionURI()
			r.NoError(err)
			a.Equal(tc.want, cs)
		})
	}
}
