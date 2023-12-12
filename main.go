package main

import (
	"business-things/ent"
	"business-things/ent/group"
	"context"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	ctx := context.Background()

	a8m := client.User.Create().
		SetName("a8m").
		SetAge(30).
		SaveX(ctx)
	neta := client.User.Create().
		SetName("neta").
		SetAge(28).
		SaveX(ctx)

	client.Car.Create().
		SetModel("Tesla").
		SetRegisteredAt(time.Now()).
		SetOwner(a8m).
		ExecX(ctx)

	client.Car.Create().
		SetModel("Mazda").
		SetRegisteredAt(time.Now()).
		SetOwner(a8m).
		ExecX(ctx)

	client.Car.
		Create().
		SetModel("Ford").
		SetRegisteredAt(time.Now()).
		SetOwner(neta).
		ExecX(ctx)

	client.Group.
		Create().
		SetName("GitLab").
		AddUsers(neta, a8m).
		ExecX(ctx)

	client.Group.
		Create().
		SetName("GitHub").
		AddUsers(a8m).
		ExecX(ctx)

	cars := client.Group.Query().
		Where(
			group.NameNEQ("GitLab"),
		).
		QueryUsers().
		QueryCars().
		AllX(ctx)
	fmt.Println(cars)

}
