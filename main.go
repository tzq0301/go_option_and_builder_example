package main

import (
	"log"
	"option_and_builder/example"
)

func main() {
	handler := example.NewHandler()

	resp := handler.HasPermission(example.HasPermissionRequest{
		User:        example.User{Name: "123"},
		Data:        example.Data{CreatorName: "123"},
		UserIsSuper: true,
		UserIsBoss:  true,
	})

	if !resp.HasPermission {
		log.Fatal("OH NO")
	}
}
