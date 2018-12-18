package main

import (
	"fmt"
	"time"

	randomdata "github.com/Pallinder/go-randomdata"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/phogolabs/orm/example/database"
	"github.com/phogolabs/parcello"
	validator "gopkg.in/go-playground/validator.v9"

	"github.com/apex/log"
	"github.com/phogolabs/orm"
	"github.com/phogolabs/orm/example/database/model"
	lk "github.com/ulule/loukoum"
)

func main() {
	gateway, err := orm.Connect("sqlite3://prana.db")
	if err != nil {
		log.WithError(err).Fatal("Failed to open database connection")
	}
	defer gateway.Close()

	if err = gateway.Migrate(parcello.ManagerAt("migration")); err != nil {
		log.WithError(err).Fatal("Failed to setup OAK")
	}

	if err = create(gateway); err != nil {
		log.WithError(err).Fatal("Failed to generate users")
	}

	if err = gateway.ReadDir(parcello.ManagerAt("routine")); err != nil {
		log.WithError(err).Fatal("Failed to load routine")
	}

	users := []model.User{}
	if err = gateway.Select(&users, orm.Routine("select-all-users")); err != nil {
		log.WithError(err).Fatal("Failed to select all users")
	}

	show(users)
}

func create(gateway *orm.Gateway) error {
	for i := 0; i < 10; i++ {
		var lastName interface{}

		if i%2 == 0 {
			lastName = randomdata.LastName()
		}

		query := lk.Insert("users").
			Set(
				lk.Pair("id", time.Now().UnixNano()),
				lk.Pair("first_name", randomdata.FirstName(randomdata.Male)),
				lk.Pair("last_name", lastName),
			)

		if _, err := gateway.Exec(query); err != nil {
			return err
		}
	}

	return nil
}

func show(users []model.User) {
	validate := validator.New()

	for _, user := range users {
		if err := validate.Struct(user); err != nil {
			log.WithError(err).Error("Failed to validate user")
			continue
		}

		fmt.Printf("User ID: %v\n", user.ID)
		fmt.Printf("First Name: %v\n", user.FirstName)

		if user.LastName.Valid {
			fmt.Printf("Last Name: %v\n", user.LastName.String)
		} else {
			fmt.Println("Last Name: null")
		}

		fmt.Println("---")
	}
}
