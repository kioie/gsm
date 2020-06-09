package main

import (
	"fmt"
	gcp_secret_manager "github.com/gcp-secret-manager"
)

func main() {
	projectID := "secret-manager-test"
	fmt.Print(gcp_secret_manager.CreateSecret(projectID, []byte("my secret data")))
}
