package gcp_secret_manager

import (
	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"context"
	"fmt"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
	"log"
)

var ctx = context.Background()
var (
	Client *secretmanager.Client
)

func init() {
	var err error
	Client, err = secretmanager.NewClient(ctx)
	if err != nil {
		log.Fatalf("failed to setup client: %v", err)
	}
}

func Connect() *secretmanager.Client {
	ctx := context.Background()
	newClient, err := secretmanager.NewClient(ctx)
	if err != nil {
		log.Fatalf("failed to setup client: %v", err)
	}
	return newClient
}

func CreateEmptySecret(projectID string, secretName string, ) *secretmanagerpb.Secret {
	if SecretExists(projectID, secretName) == true {
		log.Fatalf("failed to create secret as secret already exists")
	}
	createSecretReq := &secretmanagerpb.CreateSecretRequest{
		Parent:   fmt.Sprintf("projects/%s", projectID),
		SecretId: secretName,
		Secret: &secretmanagerpb.Secret{
			Replication: &secretmanagerpb.Replication{
				Replication: &secretmanagerpb.Replication_Automatic_{
					Automatic: &secretmanagerpb.Replication_Automatic{},
				},
			},
		},
	}
	secret, err := Client.CreateSecret(ctx, createSecretReq)
	if err != nil {
		log.Fatalf("failed to create secret: %v", err)
	}
	return secret
}

func CreateSecretWithData(projectID string, secretName string, payload []byte) (*secretmanagerpb.SecretVersion, error) {
	if SecretExists(projectID, secretName) == true {
		log.Fatalf("failed to create secret as secret already exists")
	}
	createSecretReq := &secretmanagerpb.CreateSecretRequest{
		Parent:   fmt.Sprintf("projects/%s", projectID),
		SecretId: secretName,
		Secret: &secretmanagerpb.Secret{
			Replication: &secretmanagerpb.Replication{
				Replication: &secretmanagerpb.Replication_Automatic_{
					Automatic: &secretmanagerpb.Replication_Automatic{},
				},
			},
		},
	}
	secret, err := Client.CreateSecret(ctx, createSecretReq)
	if err != nil {
		log.Fatalf("failed to create secret: %v", err)
	}
	addSecretVersionReq := &secretmanagerpb.AddSecretVersionRequest{
		Parent: secret.Name,
		Payload: &secretmanagerpb.SecretPayload{
			Data: payload,
		},
	}
	version, err := Client.AddSecretVersion(ctx, addSecretVersionReq)
	if err != nil {
		log.Fatalf("failed to add secret version: %v", err)
	}
	return version, err
}

func SecretExists(projectID string, secretName string, ) bool {
	accessRequest := &secretmanagerpb.GetSecretRequest{
		Name: fmt.Sprintf("projects/%v/secrets/%v", projectID, secretName)}
	_, err := Client.GetSecret(ctx, accessRequest)
	if err != nil {
		//log.Fatalf("failed to check if secret exists: %v", err)
		return false
	}
	return true
}

func ListSecrets(projectID string) *secretmanager.SecretIterator {
	listSecretsReq := &secretmanagerpb.ListSecretsRequest{
		Parent: fmt.Sprintf("projects/%v", projectID),
	}
	results := Client.ListSecrets(ctx, listSecretsReq)
	return results
}

func AddNewSecretVersion(projectID string, secretName string, payload []byte) *secretmanagerpb.SecretVersion {
	addSecretVersionReq := &secretmanagerpb.AddSecretVersionRequest{
		Parent: fmt.Sprintf("projects/%v/secrets/%v", projectID, secretName),
		Payload: &secretmanagerpb.SecretPayload{
			Data: payload,
		},
	}
	version, err := Client.AddSecretVersion(ctx, addSecretVersionReq)
	if err != nil {
		log.Fatalf("failed to add secret version: %v", err)
	}
	return version
}
func GetSecret(projectID string, secretName string, version string) *secretmanagerpb.SecretPayload {
	if version == "" {
		version = "latest"
	}
	getSecret := &secretmanagerpb.AccessSecretVersionRequest{
		Name: fmt.Sprintf("projects/%v/secrets/%v/versions/%v", projectID, secretName, version),
	}
	result, err := Client.AccessSecretVersion(ctx, getSecret)
	if err != nil {
		log.Fatalf("failed to get secret: %v", err)
	}
	return result.Payload
}

func DeleteSecret(projectID string, secretName string) {
	deleteSecretReq := &secretmanagerpb.DeleteSecretRequest{
		Name: fmt.Sprintf("projects/%v/secrets/%v", projectID, secretName),
	}
	err := Client.DeleteSecret(ctx, deleteSecretReq)
	if err != nil {
		log.Fatalf("failed to delete secret: %v", err)
	}
}

func DeleteSecretVersion(projectID string, secretName string, version string) *secretmanagerpb.SecretVersion {
	destroySecretReq := &secretmanagerpb.DestroySecretVersionRequest{
		Name: fmt.Sprintf("projects/%v/secrets/%v/versions/%v", projectID, secretName, version),
	}
	result, err := Client.DestroySecretVersion(ctx, destroySecretReq)
	if err != nil {
		log.Fatalf("failed to get secret: %v", err)
	}
	return result
}

func ListSecretVersions() {}
func GetSecretMetadata(projectID string, secretName string, version string) *secretmanagerpb.SecretVersion {
	getSecretReq := &secretmanagerpb.GetSecretVersionRequest{
		Name: fmt.Sprintf("projects/%v/secrets/%v/versions/%v", projectID, secretName, version),
	}
	result, err := Client.GetSecretVersion(ctx, getSecretReq)
	if err != nil {
		log.Fatalf("failed to get secret: %v", err)
	}
	return result
}

func AccessSecretVersion() {

}

func DisableSecret(projectID string, secretName string, version string) *secretmanagerpb.SecretVersion {
	disableSecretReq := &secretmanagerpb.DisableSecretVersionRequest{
		Name: fmt.Sprintf("projects/%v/secrets/%v/versions/%v", projectID, secretName, version),
	}
	result, err := Client.DisableSecretVersion(ctx, disableSecretReq)
	if err != nil {
		log.Fatalf("failed to get secret: %v", err)
	}
	return result
}

func EnableSecret(projectID string, secretName string, version string) *secretmanagerpb.SecretVersion {
	enableSecretReq := &secretmanagerpb.EnableSecretVersionRequest{
		Name: fmt.Sprintf("projects/%v/secrets/%v/versions/%v", projectID, secretName, version),
	}
	result, err := Client.EnableSecretVersion(ctx, enableSecretReq)
	if err != nil {
		log.Fatalf("failed to get secret: %v", err)
	}
	return result
}

