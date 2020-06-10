package main

import (
	"fmt"
	gcp_secret_manager "github.com/gcp-secret-manager"
)

func main() {
	projectID := "secret-manager-test"
	//connect :=&gcp_secret_manager.NewClient{}
	if !gcp_secret_manager.SecretExists("projects/1092054168008/secrets/my-secret") {
		fmt.Println("Secret does not Exists")
	}

	//	fmt.Print(gcp_secret_manager.CreateSecret(projectID, []byte("my secret data")))
	result := gcp_secret_manager.GetSecret("my-secret", "", projectID)
	fmt.Println(result.Payload)
	newResult := gcp_secret_manager.AddNewSecretVersion("my-secret", projectID, []byte("a new test"))
	fmt.Println(newResult)
	result2 := gcp_secret_manager.GetSecret("my-secret", "1", projectID)
	fmt.Println(result2.Payload)
}
