package orm_test

import (
	"fmt"

	"github.com/phogolabs/orm"
	"github.com/phogolabs/parcello"
	lk "github.com/ulule/loukoum"
)

type User struct {
	ID        int64  `db:"id"`
	FirstName string `db:"last_name"`
	LastName  string `db:"first_name"`
}

func ExampleGateway_SelectOne() {
	gateway, err := orm.Open("sqlite3", "example.db")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func() {
		if dbErr := gateway.Close(); dbErr != nil {
			fmt.Println(dbErr)
		}
	}()

	query := lk.Select("id", "first_name", "last_name").
		From("users").
		Where(lk.Condition("first_name").Equal("John"))

	user := User{}
	if err := gateway.SelectOne(&user, query); err != nil {
		fmt.Println(err)
	}
}

func ExampleGateway_Select() {
	gateway, err := orm.Open("sqlite3", "example.db")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func() {
		if dbErr := gateway.Close(); dbErr != nil {
			fmt.Println(dbErr)
		}
	}()

	query := lk.Select("id", "first_name", "last_name").From("users")
	users := []User{}

	if err := gateway.Select(&users, query); err != nil {
		fmt.Println(err)
	}
}

func ExampleGateway_QueryRow() {
	gateway, err := orm.Open("sqlite3", "example.db")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func() {
		if dbErr := gateway.Close(); dbErr != nil {
			fmt.Println(dbErr)
		}
	}()

	query := lk.Select("id", "first_name", "last_name").
		From("users").
		Where(lk.Condition("first_name").Equal("John"))

	var row *orm.Row

	row, err = gateway.QueryRow(query)
	if err != nil {
		fmt.Println(err)
		return
	}

	user := User{}
	if err := row.StructScan(&user); err != nil {
		fmt.Println(err)
	}
}

func ExampleGateway_Query() {
	gateway, err := orm.Open("sqlite3", "example.db")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func() {
		if dbErr := gateway.Close(); dbErr != nil {
			fmt.Println(dbErr)
		}
	}()

	query := lk.Select("id", "first_name", "last_name").From("users")
	rows, err := gateway.Query(query)

	if err != nil {
		fmt.Println(err)
		return
	}

	user := User{}

	for rows.Next() {
		if err = rows.StructScan(&user); err != nil {
			fmt.Println(err)
			return
		}

		if user.FirstName == "John" {
			fmt.Println(user.LastName)
		}
	}
}

func ExampleGateway_Exec() {
	gateway, err := orm.Open("sqlite3", "example.db")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func() {
		if dbErr := gateway.Close(); dbErr != nil {
			fmt.Println(dbErr)
		}
	}()

	query := lk.Insert("users").
		Set(
			lk.Pair("first_name", "John"),
			lk.Pair("last_name", "Doe"),
		).
		Returning("id")

	if _, err := gateway.Exec(query); err != nil {
		fmt.Println(err)
	}
}

func ExampleRoutine() {
	gateway, err := orm.Open("sqlite3", "example.db")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func() {
		if dbErr := gateway.Close(); dbErr != nil {
			fmt.Println(dbErr)
		}
	}()

	err = gateway.ReadDir(parcello.Dir("./database/command"))
	if err != nil {
		fmt.Println(err)
		return
	}

	if _, err := gateway.Exec(orm.Routine("show-sqlite-master")); err != nil {
		fmt.Println(err)
	}
}

func ExampleSQL() {
	gateway, err := orm.Open("sqlite3", "example.db")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func() {
		if dbErr := gateway.Close(); dbErr != nil {
			fmt.Println(dbErr)
		}
	}()

	rows, err := gateway.Query(orm.SQL("SELECT tbl_name FROM sqlite_master"))
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func() {
		if dbErr := rows.Close(); dbErr != nil {
			fmt.Println(dbErr)
		}
	}()

	var name string

	for rows.Next() {
		if err = rows.Scan(&name); err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(name)
	}
}
