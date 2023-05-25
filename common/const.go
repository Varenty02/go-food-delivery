package common

import "log"

func AppRecover() {
	if err := recover(); err != nil {
		log.Println("Recovery error", err)
	}
}

const (
	DbTypeRestaurant = 1
	DbTypeUser       = 2
)
const (
	CurrentUser = "user"
)

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() string
}
