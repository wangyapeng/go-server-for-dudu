package models

import (
	"fmt"
	// "errors"
	// "strconv"
	// "time"
)

type Animal interface {
	eat()
}

type User struct {
	ID        int    `orm:"column(id)"`
	Name      string `orm:"column(name)"`
	Telephone string `orm:"column(telephone)"`
}

func Test() int {
	fmt.Println("test func execute")
	return 1
}

func init() {
	fmt.Println("user model init")
}
