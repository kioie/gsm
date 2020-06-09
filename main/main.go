package main

import (
	"fmt"
	gcp_secret_manager "github.com/gcp-secret-manager"
)

func main() {
	projectID := "secret-manager-test"

	if !gcp_secret_manager.SecretExists("projects/1092054168008/secrets/my-secr9et", projectID) {
		fmt.Println("Secret does not Exists")
	}
	fmt.Print(gcp_secret_manager.CreateSecret(projectID, []byte("my secret data")))

}
