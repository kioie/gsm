package gcp_secret_manager

import (
	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"context"
	"fmt"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
	"log"
)

var ctx = context.Background()
var client = Connect()

func Connect() *secretmanager.Client {
	ctx := context.Background()
	newClient, err := secretmanager.NewClient(ctx)
	if err != nil {
		log.Fatalf("failed to setup client: %v", err)
	}
	return newClient
}

func CreateSecret(projectID string, payload []byte) (*secretmanagerpb.SecretVersion, error) {

	// Create the request to create the secret.
	createSecretReq := &secretmanagerpb.CreateSecretRequest{
		Parent:   fmt.Sprintf("projects/%s", projectID),
		SecretId: "my-secret",
		Secret: &secretmanagerpb.Secret{
			Replication: &secretmanagerpb.Replication{
				Replication: &secretmanagerpb.Replication_Automatic_{
					Automatic: &secretmanagerpb.Replication_Automatic{},
				},
			},
		},
	}
	secret, err := client.CreateSecret(ctx, createSecretReq)
	if err != nil {
		log.Fatalf("failed to create secret: %v", err)
	}
	addSecretVersionReq := &secretmanagerpb.AddSecretVersionRequest{
		Parent: secret.Name,
		Payload: &secretmanagerpb.SecretPayload{
			Data: payload,
		},
	}
	version, err := client.AddSecretVersion(ctx, addSecretVersionReq)
	if err != nil {
		log.Fatalf("failed to add secret version: %v", err)
	}
	return version, err
}

func CreateNewSecretVersion(projectID string, payload []byte) (*secretmanagerpb.SecretVersion, error) {
	// Create the client.
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		log.Fatalf("failed to setup client: %v", err)
	}

	// Create the request to create the secret.
	createSecretReq := &secretmanagerpb.CreateSecretRequest{
		Parent:   fmt.Sprintf("projects/%s", projectID),
		SecretId: "my-secret",
		Secret: &secretmanagerpb.Secret{
			Replication: &secretmanagerpb.Replication{
				Replication: &secretmanagerpb.Replication_Automatic_{
					Automatic: &secretmanagerpb.Replication_Automatic{},
				},
			},
		},
	}
	secret, err := client.CreateSecret(ctx, createSecretReq)
	if err != nil {
		log.Fatalf("failed to create secret: %v", err)
	}
	// Declare the payload to store.
	//payload := []byte("my super secret data")

	// Build the request.
	addSecretVersionReq := &secretmanagerpb.AddSecretVersionRequest{
		Parent: secret.Name,
		Payload: &secretmanagerpb.SecretPayload{
			Data: payload,
		},
	}
	// Call the API.
	version, err := client.AddSecretVersion(ctx, addSecretVersionReq)
	if err != nil {
		log.Fatalf("failed to add secret version: %v", err)
	}
	return version, err
}

func SecretExists(secretId string) bool {
	accessRequest := &secretmanagerpb.GetSecretRequest{Name: secretId}
	_, err := client.GetSecret(ctx, accessRequest)
	if err != nil {
		return false
	}
	return true
}

func ListSecrets(projectID string) *secretmanager.SecretIterator {
	listSecretsReq := &secretmanagerpb.ListSecretsRequest{
		Parent: fmt.Sprintf("projects/%v", projectID),
	}
	results := client.ListSecrets(ctx, listSecretsReq)
	return results
}

func AddNewSecretVersion(secretName string, projectID string, payload []byte) *secretmanagerpb.SecretVersion {
	addSecretVersionReq := &secretmanagerpb.AddSecretVersionRequest{
		Parent: fmt.Sprintf("projects/%v/secrets/%v", projectID, secretName),
		Payload: &secretmanagerpb.SecretPayload{
			Data: payload,
		},
	}
	version, err := client.AddSecretVersion(ctx, addSecretVersionReq)
	if err != nil {
		log.Fatalf("failed to add secret version: %v", err)
	}
	return version
}
func GetSecret(secretName string, version string, projectID string) *secretmanagerpb.AccessSecretVersionResponse {
	if version == "" {
		version = "latest"
	}
	getSecret := &secretmanagerpb.AccessSecretVersionRequest{
		Name: fmt.Sprintf("projects/%v/secrets/%v/versions/%v", projectID, secretName, version),
	}
	result, err := client.AccessSecretVersion(ctx, getSecret)
	if err != nil {
		log.Fatalf("failed to get secret: %v", err)
	}
	return result
}
func UpdateSecret() {
}
func DeleteSecret(secretName string, projectID string) {
	deleteSecretReq := &secretmanagerpb.DeleteSecretRequest{
		Name: fmt.Sprintf("projects/%v/secrets/%v", projectID, secretName),
	}
	err := client.DeleteSecret(ctx, deleteSecretReq)
	if err != nil {
		log.Fatalf("failed to delete secret: %v", err)
	}
}
func ListSecretVersions()   {}
func GetSecretVersion()     {}
func AccessSecretVersion()  {}
func DisableSecretVersion() {}
func EnableSecretVersion()  {}
func DestroySecretVersion() {}
