/*
 * // Licensed to the Apache Software Foundation (ASF) under one
 * // or more contributor license agreements.  See the NOTICE file
 * // distributed with this work for additional information
 * // regarding copyright ownership.  The ASF licenses this file
 * // to you under the Apache License, Version 2.0 (the
 * // "License"); you may not use this file except in compliance
 * // with the License.  You may obtain a copy of the License at
 * //
 * //   http://www.apache.org/licenses/LICENSE-2.0
 * //
 * // Unless required by applicable law or agreed to in writing,
 * // software distributed under the License is distributed on an
 * // "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * // KIND, either express or implied.  See the License for the
 * // specific language governing permissions and limitations
 * // under the License.
 *
 *
 *
 *
 * author: Eddy Kioi
 * project: gcp-secret-manager
 * date: 15/06/2020, 14:17
 */

package gcp_secret_manager

import (
	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"context"
	"errors"
	"fmt"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
	"log"
)

var (
	// Api Secret Client
	client    SecretClient
	// ProjectID that is used as a global var
	ProjectID string
)
var ctx = context.Background()
var c, _ = secretmanager.NewClient(ctx)

func init() {

	client = &secretClientImpl{client: c, ctx: ctx}

}
// CreateEmptySecret function
func CreateEmptySecret(secretName string) (*secretmanagerpb.Secret, error) {
	if SecretExists(secretName) == true {
		log.Println("failed to create secret as secret already exists")
		return nil, errors.New("failed to create secret as secret already exists")
	}
	createSecretReq := &secretmanagerpb.CreateSecretRequest{
		Parent:   fmt.Sprintf("projects/%s", ProjectID),
		SecretId: secretName,
		Secret: &secretmanagerpb.Secret{
			Replication: &secretmanagerpb.Replication{
				Replication: &secretmanagerpb.Replication_Automatic_{
					Automatic: &secretmanagerpb.Replication_Automatic{},
				},
			},
		},
	}
	secret, err := client.CreateSecret(createSecretReq)
	if err != nil {
		log.Printf("failed to create secret: %v", err)
		return nil, err
	}
	return secret, nil
}
// CreateSecretWithData creates secret with data
func CreateSecretWithData(secretName string, payload []byte) (*secretmanagerpb.SecretVersion, error) {
	if SecretExists(secretName) == true {
		log.Println("failed to create secret as secret already exists")
		return nil, errors.New("failed to create secret as secret already exists")
	}
	createSecretReq := &secretmanagerpb.CreateSecretRequest{
		Parent:   fmt.Sprintf("projects/%s", ProjectID),
		SecretId: secretName,
		Secret: &secretmanagerpb.Secret{
			Replication: &secretmanagerpb.Replication{
				Replication: &secretmanagerpb.Replication_Automatic_{
					Automatic: &secretmanagerpb.Replication_Automatic{},
				},
			},
		},
	}
	secret, err := client.CreateSecret(createSecretReq)
	if err != nil {
		log.Printf("failed to create secret: %v\n", err)
		return nil, err
	}
	addSecretVersionReq := &secretmanagerpb.AddSecretVersionRequest{
		Parent: secret.Name,
		Payload: &secretmanagerpb.SecretPayload{
			Data: payload,
		},
	}
	version, err := client.AddSecretVersion(addSecretVersionReq)
	if err != nil {
		log.Printf("failed to add secret version: %v\n", err)
		return nil, err
	}
	return version, err
}
// SecretExists Checks if secret exists
func SecretExists(secretName string) bool {
	accessRequest := &secretmanagerpb.GetSecretRequest{
		Name: fmt.Sprintf("projects/%v/secrets/%v", ProjectID, secretName)}
	_, err := client.GetSecret(accessRequest)
	if err != nil {
		return false
	}
	return true
}
// AddNewSecretVersion Adds a new Version of a secret on a secret name
func AddNewSecretVersion(secretName string, payload []byte) (*secretmanagerpb.SecretVersion, error) {
	addSecretVersionReq := &secretmanagerpb.AddSecretVersionRequest{
		Parent: fmt.Sprintf("projects/%v/secrets/%v", ProjectID, secretName),
		Payload: &secretmanagerpb.SecretPayload{
			Data: payload,
		},
	}
	version, err := client.AddSecretVersion(addSecretVersionReq)
	if err != nil {
		log.Printf("failed to add secret version: %v", err)
		return nil, err
	}
	return version, nil
}
// GetSecret Gets secret data
func GetSecret(secretName string, version string) (*secretmanagerpb.SecretPayload, error) {
	if version == "" {
		version = "latest"
	}
	getSecret := &secretmanagerpb.AccessSecretVersionRequest{
		Name: fmt.Sprintf("projects/%v/secrets/%v/versions/%v", ProjectID, secretName, version),
	}
	result, err := client.AccessSecretVersion(getSecret)
	if err != nil {
		log.Printf("failed to get secret: %v", err)
		return nil, err
	}
	return result.Payload, nil
}
// DeleteSecretAndVersions Deletes secret with all the versions included
func DeleteSecretAndVersions(secretName string) error {
	deleteSecretReq := &secretmanagerpb.DeleteSecretRequest{
		Name: fmt.Sprintf("projects/%v/secrets/%v", ProjectID, secretName),
	}
	err := client.DeleteSecret(deleteSecretReq)
	if err == nil {
		log.Printf("Secret Deleted Successfully")
	}
	return err
}
// DeleteSecretVersion Deletes specific version of a secret
func DeleteSecretVersion(secretName string, version string) (*secretmanagerpb.SecretVersion, error) {
	destroySecretReq := &secretmanagerpb.DestroySecretVersionRequest{
		Name: fmt.Sprintf("projects/%v/secrets/%v/versions/%v", ProjectID, secretName, version),
	}
	result, err := client.DestroySecretVersion(destroySecretReq)
	if err != nil {
		log.Printf("failed to get secret: %v", err)
		return nil, err
	}
	return result, nil
}
// GetSecretMetadata Gets metadata of a secret Name
func GetSecretMetadata(secretName string, version string) (*secretmanagerpb.SecretVersion, error) {
	getSecretReq := &secretmanagerpb.GetSecretVersionRequest{
		Name: fmt.Sprintf("projects/%v/secrets/%v/versions/%v", ProjectID, secretName, version),
	}
	result, err := client.GetSecretVersion(getSecretReq)
	if err != nil {
		log.Printf("failed to get secret: %v", err)
		return nil, err
	}
	return result, nil
}
// DisableSecret Disables secret
func DisableSecret(secretName string, version string) (*secretmanagerpb.SecretVersion, error) {
	disableSecretReq := &secretmanagerpb.DisableSecretVersionRequest{
		Name: fmt.Sprintf("projects/%v/secrets/%v/versions/%v", ProjectID, secretName, version),
	}
	result, err := client.DisableSecretVersion(disableSecretReq)
	if err != nil {
		log.Printf("failed to get secret: %v", err)
		return nil, err
	}
	return result, nil
}
// EnableSecret Enables secret
func EnableSecret(secretName string, version string) (*secretmanagerpb.SecretVersion, error) {
	enableSecretReq := &secretmanagerpb.EnableSecretVersionRequest{
		Name: fmt.Sprintf("projects/%v/secrets/%v/versions/%v", ProjectID, secretName, version),
	}
	result, err := client.EnableSecretVersion(enableSecretReq)
	if err != nil {
		log.Printf("failed to get secret: %v\n", err)
		return nil, err
	}
	return result, nil
}
