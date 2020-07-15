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

/*
func TestGetSecret(t *testing.T) {
	AccessSecretVersionFunc = func(ctx context.Context, req *pb.AccessSecretVersionRequest) (*pb.AccessSecretVersionResponse, error) {
		return &pb.AccessSecretVersionResponse{
			Name:    "projects/myProjects/secrets/mySecrets/versions/latest",
			Payload: &pb.SecretPayload{Data: []byte("mySecret")},
		}, nil
	}
	type args struct {
		projectID  string
		secretName string
		version    string
	}
	tests := []struct {
		name    string
		args    args
		want    *pb.SecretPayload
		wantErr bool
	}{
		{name: "Success",
			args: args{
				projectID:  "myProject",
				secretName: "mySecret",
				version:    "",
			},
			want:    &pb.SecretPayload{Data: []byte("mySecret")},
			wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := GetSecret(tt.args.secretName, tt.args.version); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSecret() = %v, want %v", got, tt.want)
			}
		})
	}
	AccessSecretVersionFunc = func(ctx context.Context, req *pb.AccessSecretVersionRequest) (*pb.AccessSecretVersionResponse, error) {
		return nil, errors.New("failed to get secret")
	}
	tests = []struct {
		name    string
		args    args
		want    *pb.SecretPayload
		wantErr bool
	}{
		{name: "Failure",
			args: args{
				projectID:  "myProject",
				secretName: "mySecret",
				version:    "",
			},
			want:    nil,
			wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := GetSecret(tt.args.secretName, tt.args.version); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSecret() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSecretMetadata(t *testing.T) {
	GetSecretVersionFunc = func(ctx context.Context, req *pb.GetSecretVersionRequest) (*pb.SecretVersion, error) {
		return secretVersionPositiveReturn, nil
	}
	type args struct {
		projectID  string
		secretName string
		version    string
	}
	tests := []struct {
		name    string
		args    args
		want    *pb.SecretVersion
		wantErr bool
	}{
		{name: "Success",
			args: args{
				projectID:  "myProject",
				secretName: "mySecrets",
				version:    "1",
			},
			want:    secretVersionPositiveReturn,
			wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := GetSecretMetadata(tt.args.secretName, tt.args.version); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSecretMetadata() = %v, want %v", got, tt.want)
			}
		})
	}
	GetSecretVersionFunc = func(ctx context.Context, req *pb.GetSecretVersionRequest) (*pb.SecretVersion, error) {
		return nil, errors.New("failed to Get Secret Version")
	}
	tests = []struct {
		name    string
		args    args
		want    *pb.SecretVersion
		wantErr bool
	}{
		{name: "Failure",
			args: args{
				projectID:  "myProject",
				secretName: "mySecrets",
				version:    "1",
			},
			want:    nil,
			wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := GetSecretMetadata(tt.args.secretName, tt.args.version); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSecretMetadata() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSecretExists(t *testing.T) {
	type args struct {
		secretName string
	}
	GetSecretFunc = func(ctx context.Context, req *pb.GetSecretRequest) (*pb.Secret, error) {
		return secretPositiveReturn, nil
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Success",
			args: args{secretName: "my-secret"},
			want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SecretExists(tt.args.secretName); got != tt.want {
				t.Errorf("SecretExists() = %v, want %v", got, tt.want)
			}
		})
	}

	GetSecretFunc = func(ctx context.Context, req *pb.GetSecretRequest) (*pb.Secret, error) {
		return nil, errors.New("secret not found")
	}
	tests2 := []struct {
		name string
		args args
		want bool
	}{
		{name: "Failure",
			args: args{secretName: "mysecret"},
			want: false},
	}

	for _, tt := range tests2 {
		t.Run(tt.name, func(t *testing.T) {

			if got := SecretExists(tt.args.secretName); got != tt.want {
				t.Errorf("SecretExists() = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestDeleteSecretAndVersions(t *testing.T) {
	DeleteSecretFunc = func(ctx context.Context, req *pb.DeleteSecretRequest) error {
		return errors.New("delete failed")
	}
	type args struct {
		secretName string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Failure",
			args:    args{secretName: "mySecret"},
			wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteSecretAndVersions(tt.args.secretName); (err != nil) != tt.wantErr {
				t.Errorf("DeleteSecretAndVersions() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	DeleteSecretFunc = func(ctx context.Context, req *pb.DeleteSecretRequest) error {
		return nil
	}
	tests = []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Success",
			args:    args{secretName: "mySecret"},
			wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteSecretAndVersions(tt.args.secretName); (err != nil) != tt.wantErr {
				t.Errorf("DeleteSecretAndVersions() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
*/

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
