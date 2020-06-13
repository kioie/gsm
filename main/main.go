package main

import (
	"fmt"
	gcp_secret_manager "github.com/kioie/gcp-secret-manager"
)

func main() {
	gcp_secret_manager.ProjectID = "secret-manager-test"
	//projectID := "secret-manager-test"
	//connect :=&gcp_secret_manager.NewClient{}
	//	if !gcp_secret_manager.SecretExists("projects/1092054168008/secrets/my-secret") {
	//		fmt.Println("Secret does not Exists")
	//	}

	//	fmt.Print(gcp_secret_manager.CreateSecret(projectID, []byte("my secret data")))
	//result := gcp_secret_manager.GetSecret("my-secret", "", projectID)
	//fmt.Println(result)
	//	newResult := gcp_secret_manager.AddNewSecretVersion("my-secret", projectID, []byte("a new test"))
	//	fmt.Println(newResult)
	result2 := gcp_secret_manager.GetSecret("my-secret", "")
	fmt.Println(result2)
	//	result3 := gcp_secret_manager.GetSecretMetadata("my-secret", "1", projectID)
	//	fmt.Println(result3.GetName())

	//fmt.Println(gcp_secret_manager.SecretExists(projectID,"my-secret"))
}
