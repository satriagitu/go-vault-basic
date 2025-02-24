package main

import (
	"fmt"
	"log"

	"github.com/hashicorp/vault/api"
)

func main() {
	client, err := api.NewClient(&api.Config{Address: "http://127.0.0.1:8200"})
	if err != nil {
		log.Fatal(err)
	}

	client.SetToken("-") // Ganti dengan root token

	secret, err := client.Logical().Read("secret/data/myapp")
	if err != nil {
		log.Fatal(err)
	}

	if secret != nil {
		data := secret.Data["data"].(map[string]interface{})
		fmt.Println("Username:", data["username"])
		fmt.Println("Password:", data["password"])
	} else {
		fmt.Println("Secret not found")
	}
}

