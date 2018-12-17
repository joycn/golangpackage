package main

import (
	"vrouter/server/api"
)

func main() {
	client := api.NewClient("127.0.0.1:8000")
	err := client.CreateAndStart("user1234")
}
