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
	"errors"
	"reflect"
	"testing"

	pb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

var client = &MockClient{}

var secretVersionPositiveReturn = &pb.SecretVersion{
	Name:        "projects/myProject/secrets/mySecrets/versions/1",
	CreateTime:  nil,
	DestroyTime: nil,
	State:       0,
}

var secretPositiveReturn = &pb.Secret{
	Name:        "",
	Replication: nil,
	CreateTime:  nil,
	Labels:      nil,
}

func TestClient_AddNewSecretVersion(t *testing.T) {
	addNewSecretTest := func(ctx context.Context, secretName string, projectId string, payload []byte, want *pb.SecretVersion, wantErr bool) func(t *testing.T) {
		return func(t *testing.T) {
			c := &Client{
				smc: client,
			}
			if got, _ := c.AddNewSecretVersion(ctx, secretName, projectId, payload); !reflect.DeepEqual(got, want) {
				t.Errorf("AddNewSecretVersion() = %v, want %v", got, want)
			}
		}
	}

	AddSecretVersionFunc = func(ctx context.Context, req *pb.AddSecretVersionRequest) (*pb.SecretVersion, error) {
		return secretVersionPositiveReturn, nil
	}
	t.Run("Success", addNewSecretTest(nil, "mysecret", "myproject", []byte("a new test"), secretVersionPositiveReturn, false))

	AddSecretVersionFunc = func(ctx context.Context, req *pb.AddSecretVersionRequest) (*pb.SecretVersion, error) {
		return nil, errors.New("failed to add secret version")
	}
	t.Run("Failure", addNewSecretTest(nil, "mysecret", "myproject", []byte("a new test"), nil, true))
}

func TestClient_CreateEmptySecret(t *testing.T) {

	createEmptySecretTest := func(ctx context.Context, secretName string, projectId string, want *pb.Secret, wantErr bool) func(t *testing.T) {
		return func(t *testing.T) {
			c := &Client{
				smc: client,
			}
			got, err := c.CreateEmptySecret(ctx, secretName, projectId)
			if (err != nil) != wantErr {
				t.Errorf("CreateEmptySecret() error = %v, wantErr %v", err, wantErr)
				return
			}
			if !reflect.DeepEqual(got, want) {
				t.Errorf("CreateEmptySecret() got = %v, want %v", got, want)
			}
		}
	}
	GetSecretFunc = func(ctx context.Context, req *pb.GetSecretRequest) (*pb.Secret, error) {
		return nil, errors.New("secret does not exist")
	}
	CreateSecretFunc = func(ctx context.Context, req *pb.CreateSecretRequest) (*pb.Secret, error) {
		return secretPositiveReturn, nil
	}
	t.Run("Success", createEmptySecretTest(nil, "mySecret", "myProject", secretPositiveReturn, false))

	CreateSecretFunc = func(ctx context.Context, req *pb.CreateSecretRequest) (*pb.Secret, error) {
		return nil, errors.New("failed to create Secret")
	}
	t.Run("Failure", createEmptySecretTest(nil, "mySecret", "myProject", nil, true))

	GetSecretFunc = func(ctx context.Context, req *pb.GetSecretRequest) (*pb.Secret, error) {
		return secretPositiveReturn, nil
	}
	t.Run("Failure", createEmptySecretTest(nil, "mySecret", "myProject", nil, true))

}

func TestClient_CreateSecretWithData(t *testing.T) {
	createSecretWithDataTest := func(ctx context.Context, secretName string, payload []byte, projectId string, want *pb.SecretVersion, wantErr bool) func(t *testing.T) {
		return func(t *testing.T) {
			c := &Client{
				smc: client,
			}
			got, err := c.CreateSecretWithData(ctx, secretName, payload, projectId)
			if (err != nil) != wantErr {
				t.Errorf("CreateSecretWithData() error = %v, wantErr %v", err, wantErr)
				return
			}
			if !reflect.DeepEqual(got, want) {
				t.Errorf("CreateSecretWithData() got = %v, want %v", got, want)
			}
		}
	}
	GetSecretFunc = func(ctx context.Context, req *pb.GetSecretRequest) (*pb.Secret, error) {
		return nil, errors.New("secret does not exist")
	}
	CreateSecretFunc = func(ctx context.Context, req *pb.CreateSecretRequest) (*pb.Secret, error) {
		return secretPositiveReturn, nil
	}
	AddSecretVersionFunc = func(ctx context.Context, req *pb.AddSecretVersionRequest) (*pb.SecretVersion, error) {
		return secretVersionPositiveReturn, nil
	}
	t.Run("Success", createSecretWithDataTest(nil, "mySecret", []byte("a new test"), "myProject", secretVersionPositiveReturn, false))

	AddSecretVersionFunc = func(ctx context.Context, req *pb.AddSecretVersionRequest) (*pb.SecretVersion, error) {
		return nil, errors.New("failed to add secret version")
	}
	t.Run("Success", createSecretWithDataTest(nil, "mySecret", []byte("a new test"), "myProject", nil, true))

	GetSecretFunc = func(ctx context.Context, req *pb.GetSecretRequest) (*pb.Secret, error) {
		return secretPositiveReturn, nil
	}
	t.Run("FailSecretExists", createSecretWithDataTest(nil, "mySecret", []byte("a new test"), "myProject", nil, true))

	GetSecretFunc = func(ctx context.Context, req *pb.GetSecretRequest) (*pb.Secret, error) {
		return nil, errors.New("secret does not exist")
	}
	CreateSecretFunc = func(ctx context.Context, req *pb.CreateSecretRequest) (*pb.Secret, error) {
		return nil, errors.New("failed to create secret")
	}
	t.Run("FailCreateSecret", createSecretWithDataTest(nil, "mySecret", []byte("a new test"), "myProject", nil, true))

}

func TestClient_DeleteSecretVersion(t *testing.T) {
	deleteSecretVersionTest := func(ctx context.Context, secretName string, projectId string, version string, want *pb.SecretVersion, wantErr bool) func(t *testing.T) {
		return func(t *testing.T) {
			c := &Client{
				smc: client,
			}
			got, err := c.DeleteSecretVersion(ctx, secretName, projectId, version)
			if (err != nil) != wantErr {
				t.Errorf("DeleteSecretVersion() error = %v, wantErr %v", err, wantErr)
				return
			}
			if !reflect.DeepEqual(got, want) {
				t.Errorf("DeleteSecretVersion() got = %v, want %v", got, want)
			}
		}
	}

	DestroySecretVersionFunc = func(ctx context.Context, req *pb.DestroySecretVersionRequest) (*pb.SecretVersion, error) {
		return secretVersionPositiveReturn, nil
	}
	t.Run("Success", deleteSecretVersionTest(nil, "mySecrets", "myProjects", "", secretVersionPositiveReturn, false))

	DestroySecretVersionFunc = func(ctx context.Context, req *pb.DestroySecretVersionRequest) (*pb.SecretVersion, error) {
		return nil, errors.New("failed to delete secret version")
	}
	t.Run("Failure", deleteSecretVersionTest(nil, "mySecrets", "myProjects", "", nil, true))

}

func TestClient_DisableSecret(t *testing.T) {
	disableSecretTest := func(ctx context.Context, secretName string, projectId string, version string, want *pb.SecretVersion, wantErr bool) func(t *testing.T) {
		return func(t *testing.T) {
			c := &Client{
				smc: client,
			}
			got, err := c.DisableSecret(ctx, secretName, projectId, version)
			if (err != nil) != wantErr {
				t.Errorf("DisableSecret() error = %v, wantErr %v", err, wantErr)
				return
			}
			if !reflect.DeepEqual(got, want) {
				t.Errorf("DisableSecret() got = %v, want %v", got, want)
			}
		}
	}

	DisableSecretVersionFunc = func(ctx context.Context, req *pb.DisableSecretVersionRequest) (*pb.SecretVersion, error) {
		return secretVersionPositiveReturn, nil
	}
	t.Run("Success", disableSecretTest(nil, "mySecrets", "myProjects", "1", secretVersionPositiveReturn, false))

	DisableSecretVersionFunc = func(ctx context.Context, req *pb.DisableSecretVersionRequest) (*pb.SecretVersion, error) {
		return nil, errors.New("failed to Disable Secret")
	}
	t.Run("Failure", disableSecretTest(nil, "mySecrets", "myProjects", "1", nil, true))

}

func TestClient_EnableSecret(t *testing.T) {
	enableSecretTest := func(ctx context.Context, secretName string, projectId string, version string, want *pb.SecretVersion, wantErr bool) func(t *testing.T) {
		return func(t *testing.T) {
			c := &Client{
				smc: client,
			}
			got, err := c.EnableSecret(ctx, secretName, projectId, version)
			if (err != nil) != wantErr {
				t.Errorf("EnableSecret() error = %v, wantErr %v", err, wantErr)
				return
			}
			if !reflect.DeepEqual(got, want) {
				t.Errorf("EnableSecret() got = %v, want %v", got, want)
			}
		}
	}

	EnableSecretVersionFunc = func(ctx context.Context, req *pb.EnableSecretVersionRequest) (*pb.SecretVersion, error) {
		return secretVersionPositiveReturn, nil
	}
	t.Run("Success", enableSecretTest(nil, "", "", "", secretVersionPositiveReturn, false))
	t.Run("Failure", enableSecretTest(nil, "mySecrets", "myProjects", "1", secretVersionPositiveReturn, false))

	EnableSecretVersionFunc = func(ctx context.Context, req *pb.EnableSecretVersionRequest) (*pb.SecretVersion, error) {
		return nil, errors.New("failed to enable Secret")
	}
	t.Run("Failure", enableSecretTest(nil, "mySecrets", "myProjects", "1", nil, true))

}

func TestClient_GetSecret(t *testing.T) {
	getSecretTest := func(ctx context.Context, secretName string, projectId string, version string, want *pb.SecretPayload, wantErr bool) func(t *testing.T) {
		return func(t *testing.T) {
			c := &Client{
				smc: client,
			}
			got, err := c.GetSecret(ctx, secretName, projectId, version)
			if (err != nil) != wantErr {
				t.Errorf("GetSecret() error = %v, wantErr %v", err, wantErr)
				return
			}
			if !reflect.DeepEqual(got, want) {
				t.Errorf("GetSecret() got = %v, want %v", got, want)
			}
		}
	}

	AccessSecretVersionFunc = func(ctx context.Context, req *pb.AccessSecretVersionRequest) (*pb.AccessSecretVersionResponse, error) {
		return &pb.AccessSecretVersionResponse{
			Name:    "projects/myProjects/secrets/mySecrets/versions/latest",
			Payload: &pb.SecretPayload{Data: []byte("mySecret")},
		}, nil
	}
	t.Run("Success", getSecretTest(nil, "mySecret", "myProject", "", &pb.SecretPayload{Data: []byte("mySecret")}, false))

	AccessSecretVersionFunc = func(ctx context.Context, req *pb.AccessSecretVersionRequest) (*pb.AccessSecretVersionResponse, error) {
		return nil, errors.New("failed to get secret")
	}
	t.Run("Failure", getSecretTest(nil, "mySecret", "myProject", "", nil, true))

}

func TestClient_GetSecretMetadata(t *testing.T) {

	getSecretMetadataTest := func(ctx context.Context, secretName string, projectId string, version string, want *pb.SecretVersion, wantErr bool) func(t *testing.T) {
		return func(t *testing.T) {
			c := &Client{
				smc: client,
			}
			got, err := c.GetSecretMetadata(ctx, secretName, projectId, version)
			if (err != nil) != wantErr {
				t.Errorf("GetSecretMetadata() error = %v, wantErr %v", err, wantErr)
				return
			}
			if !reflect.DeepEqual(got, want) {
				t.Errorf("GetSecretMetadata() got = %v, want %v", got, want)
			}
		}
	}
	GetSecretVersionFunc = func(ctx context.Context, req *pb.GetSecretVersionRequest) (*pb.SecretVersion, error) {
		return secretVersionPositiveReturn, nil
	}
	t.Run("Success", getSecretMetadataTest(nil, "mySecrets", "myProject", "1", secretVersionPositiveReturn, false))

	GetSecretVersionFunc = func(ctx context.Context, req *pb.GetSecretVersionRequest) (*pb.SecretVersion, error) {
		return nil, errors.New("failed to Get Secret Version")
	}
	t.Run("Failure", getSecretMetadataTest(nil, "mySecrets", "myProject", "1", nil, true))

}

func TestClient_DeleteSecretAndVersions(t *testing.T) {
	deleteSecretAndVersionsTest := func(ctx context.Context, secretName string, projectId string, wantErr bool) func(t *testing.T) {
		return func(t *testing.T) {
			c := &Client{
				smc: client,
			}
			if err := c.DeleteSecretAndVersions(ctx, secretName, projectId); (err != nil) != wantErr {
				t.Errorf("DeleteSecretAndVersions() error = %v, wantErr %v", err, wantErr)
			}
		}
	}

	DeleteSecretFunc = func(ctx context.Context, req *pb.DeleteSecretRequest) error {
		return errors.New("delete failed")
	}
	t.Run("Failure", deleteSecretAndVersionsTest(nil, "mySecret", "myProject", true))

	DeleteSecretFunc = func(ctx context.Context, req *pb.DeleteSecretRequest) error {
		return nil
	}
	t.Run("Success", deleteSecretAndVersionsTest(nil, "mySecret", "myProject", false))

}

func TestClient_SecretExists(t *testing.T) {
	secretExists := func(ctx context.Context, secretName string, projectId string, want bool) func(t *testing.T) {
		return func(t *testing.T) {

			client := &MockClient{}
			c := &Client{
				smc: client,
			}
			if got := c.SecretExists(ctx, secretName, projectId); got != want {
				t.Errorf("SecretExists() = %v, want %v", got, want)
			}

		}
	}
	GetSecretFunc = func(ctx context.Context, req *pb.GetSecretRequest) (*pb.Secret, error) {
		return secretPositiveReturn, nil
	}
	t.Run("Success", secretExists(nil, "my-secret", "my-project", true))
	GetSecretFunc = func(ctx context.Context, req *pb.GetSecretRequest) (*pb.Secret, error) {
		return nil, errors.New("secret not found")
	}
	t.Run("Failure", secretExists(nil, "mysecret", "my-project", false))
}
