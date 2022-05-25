package src

import (
	"os"

	"github.com/fauna/faunadb-go/v4/faunadb"
)

func ConnectDB() (*faunadb.FaunaClient, error) {
	client := faunadb.NewFaunaClient(
		os.Getenv("FAUNA_KEY"),
		faunadb.Endpoint("https://db.fauna.com"),
	)

	return client, nil
}
