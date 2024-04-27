package main

import (
	"context"
	"github.com/uptrace/go-clickhouse/ch"
)

func NewCHClient() *ch.DB {
	db := ch.Connect(
		ch.WithAddr("localhost:9000"),
		ch.WithUser("user"),
		ch.WithPassword("123"),
		ch.WithDatabase("test"),
	)

	if err := db.Ping(context.TODO()); err != nil {
		panic("failed to ping: " + err.Error())
	}

	return db
}
