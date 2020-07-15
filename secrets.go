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

package gsm

import (
	"context"
	"fmt"
	"log"
	
	sm "cloud.google.com/go/secretmanager/apiv1"
	pb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

// Client is a global exported Client struct
type Client struct {
	smc SecretClient
}

// NewClient is a global exported function that creates a new client
func NewClient(ctx context.Context) (*sm.Client, error) {
	client, err := sm.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// CreateEmptySecret function
func (c *Client) CreateEmptySecret(ctx context.Context, secretName string, projectId string) (*pb.Secret, error) {
	createSecretReq := pb.CreateSecretRequest{
		Parent:   fmt.Sprintf("projects/%s", projectId),
		SecretId: secretName,
		Secret: &pb.Secret{
			Replication: &pb.Replication{
				Replication: &pb.Replication_Automatic_{
					Automatic: &pb.Replication_Automatic{},
				},
			},
		},
	}
	
	secret, err := c.smc.CreateSecret(ctx, &createSecretReq)
	if err != nil {
		log.Printf("failed to create secret: %v", err)
		return nil, err
	}
	
	return secret, nil
}

// CreateSecretWithData creates secret with data
func (c *Client) CreateSecretWithData(ctx context.Context, secretName string, payload []byte, projectId string) (*pb.SecretVersion, error) {
	createSecretReq := pb.CreateSecretRequest{
		Parent:   fmt.Sprintf("projects/%s", projectId),
		SecretId: secretName,
		Secret: &pb.Secret{
			Replication: &pb.Replication{
				Replication: &pb.Replication_Automatic_{
					Automatic: &pb.Replication_Automatic{},
				},
			},
		},
	}
	
	secret, err := c.smc.CreateSecret(ctx, &createSecretReq)
	if err != nil {
		log.Printf("failed to create secret: %v\n", err)
		return nil, err
	}
	
	addSecretVersionReq := pb.AddSecretVersionRequest{
		Parent: secret.Name,
		Payload: &pb.SecretPayload{
			Data: payload,
		},
	}
	
	version, err := c.smc.AddSecretVersion(ctx, &addSecretVersionReq)
	if err != nil {
		log.Printf("failed to add secret version: %v\n", err)
		return nil, err
	}
	
	return version, err
}

// SecretExists Checks if secret exists
func (c *Client) SecretExists(ctx context.Context, secretName string, projectId string) bool {
	accessRequest := pb.GetSecretRequest{
		Name: fmt.Sprintf("projects/%v/secrets/%v", projectId, secretName)}
	
	_, err := c.smc.GetSecret(ctx, &accessRequest)
	if err != nil {
		return false
	}
	return true
}

// AddNewSecretVersion Adds a new Version of a secret on a secret name
func (c *Client) AddNewSecretVersion(ctx context.Context, secretName string, projectId string, payload []byte) (*pb.SecretVersion, error) {
	addSecretVersionReq := pb.AddSecretVersionRequest{
		Parent: fmt.Sprintf("projects/%v/secrets/%v", projectId, secretName),
		Payload: &pb.SecretPayload{
			Data: payload,
		},
	}
	version, err := c.smc.AddSecretVersion(ctx, &addSecretVersionReq)
	if err != nil {
		log.Printf("failed to add secret version: %v", err)
		return nil, err
	}
	
	return version, nil
}

// GetSecret Gets secret data
func (c *Client) GetSecret(ctx context.Context, secretName string, projectId string, version string) (*pb.SecretPayload, error) {
	if version == "" {
		version = "latest"
	}
	
	getSecret := pb.AccessSecretVersionRequest{
		Name: fmt.Sprintf("projects/%v/secrets/%v/versions/%v", projectId, secretName, version),
	}
	
	result, err := c.smc.AccessSecretVersion(ctx, &getSecret)
	if err != nil {
		log.Printf("failed to get secret: %v", err)
		return nil, err
	}
	
	return result.Payload, nil
}

// DeleteSecretAndVersions Deletes secret with all the versions included
func (c *Client) DeleteSecretAndVersions(ctx context.Context, secretName string, projectId string) error {
	deleteSecretReq := pb.DeleteSecretRequest{
		Name: fmt.Sprintf("projects/%v/secrets/%v", projectId, secretName),
	}
	
	err := c.smc.DeleteSecret(ctx, &deleteSecretReq)
	if err == nil {
		log.Printf("Secret Deleted Successfully")
	}
	
	return err
}

// DeleteSecretVersion Deletes specific version of a secret
func (c *Client) DeleteSecretVersion(ctx context.Context, secretName string, projectId string, version string) (*pb.SecretVersion, error) {
	
	destroySecretReq := pb.DestroySecretVersionRequest{
		Name: fmt.Sprintf("projects/%v/secrets/%v/versions/%v", projectId, secretName, version),
	}
	result, err := c.smc.DestroySecretVersion(ctx, &destroySecretReq)
	if err != nil {
		log.Printf("failed to get secret: %v", err)
		return nil, err
	}
	return result, nil
}

// GetSecretMetadata Gets metadata of a secret Name
func (c *Client) GetSecretMetadata(ctx context.Context, secretName string, projectId string, version string) (*pb.SecretVersion, error) {
	getSecretReq := pb.GetSecretVersionRequest{
		Name: fmt.Sprintf("projects/%v/secrets/%v/versions/%v", projectId, secretName, version),
	}
	
	result, err := c.smc.GetSecretVersion(ctx, &getSecretReq)
	if err != nil {
		log.Printf("failed to get secret: %v", err)
		return nil, err
	}
	
	return result, nil
}

// DisableSecret Disables secret
func (c *Client) DisableSecret(ctx context.Context, secretName string, projectId string, version string) (*pb.SecretVersion, error) {
	disableSecretReq := pb.DisableSecretVersionRequest{
		Name: fmt.Sprintf("projects/%v/secrets/%v/versions/%v", projectId, secretName, version),
	}
	
	result, err := c.smc.DisableSecretVersion(ctx, &disableSecretReq)
	if err != nil {
		log.Printf("failed to get secret: %v", err)
		return nil, err
	}
	
	return result, nil
}

// EnableSecret Enables secret
func (c *Client) EnableSecret(ctx context.Context, secretName string, projectId string, version string) (*pb.SecretVersion, error) {
	enableSecretReq := pb.EnableSecretVersionRequest{
		Name: fmt.Sprintf("projects/%v/secrets/%v/versions/%v", projectId, secretName, version),
	}
	
	result, err := c.smc.EnableSecretVersion(ctx, &enableSecretReq)
	if err != nil {
		log.Printf("failed to get secret: %v\n", err)
		return nil, err
	}
	
	return result, nil
}
