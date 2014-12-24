package eccore

import (
	"flag"
)

var psqlHost string
var psqlUser string
var psqlPass string
var psqlDb string

func init() {
	flag.StringVar(&psqlHost, "eccore.psql.host", "localhost", "Database host")
	flag.StringVar(&psqlUser, "eccore.psql.user", "evec", "Database user")
	flag.StringVar(&psqlPass, "eccore.psql.pass", "evec", "Database password")
	flag.StringVar(&psqlDb, "eccore.psql.db", "evec", "Database password")
}
