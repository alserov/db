package main

import (
	"context"
	"fmt"
	"time"
)

type Model struct {
	ID        int       `json:"id"`
	Text      string    `json:"age"`
	CreatedAt time.Time `json:"createdAt"`
}

func main() {
	db := NewCHClient()

	fmt.Println("DB connected")

	ctx := context.TODO()

	// Create table

	_, err := db.NewTruncateTable().Model((*Model)(nil)).Exec(ctx)
	if err != nil {
		panic("failed to create table: " + err.Error())
	}

	// Insert

	models := []Model{
		{ID: 1, Text: "Foo", CreatedAt: time.Now()},
		{ID: 2, Text: "Bar", CreatedAt: time.Now().Add(time.Hour)},
	}

	_, err = db.NewInsert().Model(&models).Exec(ctx)
	if err != nil {
		panic("failed to insert: " + err.Error())
	}

	// Select

	var selectedModels []Model
	err = db.NewSelect().Model(&selectedModels).Where("id < ?", 2).Scan(ctx)

	fmt.Println(selectedModels)
}
