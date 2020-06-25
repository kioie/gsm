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

package gcpSecretManager

import (
	"errors"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
	"reflect"
	"testing"
)

func init() {
	client = &MockClient{}
}

var secretVersionPositiveReturn = &secretmanagerpb.SecretVersion{
	Name:        "projects/myProject/secrets/mySecrets/versions/1",
	CreateTime:  nil,
	DestroyTime: nil,
	State:       0,
}

var secretPositiveReturn = &secretmanagerpb.Secret{
	Name:        "",
	Replication: nil,
	CreateTime:  nil,
	Labels:      nil,
}

func TestAddNewSecretVersion(t *testing.T) {
	AddSecretVersionFunc = func(req *secretmanagerpb.AddSecretVersionRequest) (*secretmanagerpb.SecretVersion, error) {
		return secretVersionPositiveReturn, nil
	}

	type args struct {
		secretName string
		payload    []byte
	}
	tests := []struct {
		name    string
		args    args
		want    *secretmanagerpb.SecretVersion
		wantErr bool
	}{
		{name: "Success",
			args:    args{secretName: "mysecret", payload: []byte("a new test")},
			want:    secretVersionPositiveReturn,
			wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := AddNewSecretVersion(tt.args.secretName, tt.args.payload); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddNewSecretVersion() = %v, want %v", got, tt.want)
			}
		})
	}
	AddSecretVersionFunc = func(req *secretmanagerpb.AddSecretVersionRequest) (*secretmanagerpb.SecretVersion, error) {
		return nil, errors.New("failed to add secret version")
	}

	tests = []struct {
		name    string
		args    args
		want    *secretmanagerpb.SecretVersion
		wantErr bool
	}{
		{name: "Failure",
			args: args{
				secretName: "mysecret",
				payload:    []byte("a new test")},
			want:    nil,
			wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := AddNewSecretVersion(tt.args.secretName, tt.args.payload); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddNewSecretVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateEmptySecret(t *testing.T) {
	GetSecretFunc = func(req *secretmanagerpb.GetSecretRequest) (*secretmanagerpb.Secret, error) {
		return nil, errors.New("secret does not exist")
	}
	CreateSecretFunc = func(req *secretmanagerpb.CreateSecretRequest) (*secretmanagerpb.Secret, error) {
		return secretPositiveReturn, nil
	}
	type args struct {
		projectID  string
		secretName string
	}
	tests := []struct {
		name    string
		args    args
		want    *secretmanagerpb.Secret
		wantErr bool
	}{
		{name: "Success",
			args: args{
				projectID:  "myProject",
				secretName: "mySecret",
			},
			want:    secretPositiveReturn,
			wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := CreateEmptySecret(tt.args.secretName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateEmptySecret() = %v, want %v", got, tt.want)
			}
		})
	}
	CreateSecretFunc = func(req *secretmanagerpb.CreateSecretRequest) (*secretmanagerpb.Secret, error) {
		return nil, errors.New("failed to create Secret")
	}
	tests = []struct {
		name    string
		args    args
		want    *secretmanagerpb.Secret
		wantErr bool
	}{
		{name: "Failure",
			args: args{
				projectID:  "myProject",
				secretName: "mySecret",
			},
			want:    nil,
			wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := CreateEmptySecret(tt.args.secretName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateEmptySecret() = %v, want %v", got, tt.want)
			}
		})
	}
	GetSecretFunc = func(req *secretmanagerpb.GetSecretRequest) (*secretmanagerpb.Secret, error) {
		return secretPositiveReturn, nil
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := CreateEmptySecret(tt.args.secretName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateEmptySecret() = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestCreateSecretWithData(t *testing.T) {
	GetSecretFunc = func(req *secretmanagerpb.GetSecretRequest) (*secretmanagerpb.Secret, error) {
		return nil, errors.New("secret does not exist")
	}
	CreateSecretFunc = func(req *secretmanagerpb.CreateSecretRequest) (*secretmanagerpb.Secret, error) {
		return secretPositiveReturn, nil
	}
	AddSecretVersionFunc = func(req *secretmanagerpb.AddSecretVersionRequest) (*secretmanagerpb.SecretVersion, error) {
		return secretVersionPositiveReturn, nil
	}
	type args struct {
		projectID  string
		secretName string
		payload    []byte
	}
	tests := []struct {
		name    string
		args    args
		want    *secretmanagerpb.SecretVersion
		wantErr bool
	}{
		{name: "Success",
			args: args{
				projectID:  "myProject",
				secretName: "mySecret",
				payload:    []byte("a new test"),
			}, want: secretVersionPositiveReturn},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateSecretWithData(tt.args.secretName, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateSecretWithData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateSecretWithData() got = %v, want %v", got, tt.want)
			}
		})
	}
	AddSecretVersionFunc = func(req *secretmanagerpb.AddSecretVersionRequest) (*secretmanagerpb.SecretVersion, error) {
		return nil, errors.New("failed to add secret version")
	}
	tests1 := []struct {
		name    string
		args    args
		want    *secretmanagerpb.SecretVersion
		wantErr bool
	}{
		{name: "Success",
			args: args{
				projectID:  "myProject",
				secretName: "mySecret",
				payload:    []byte("a new test"),
			},
			want:    nil,
			wantErr: true},
	}
	for _, tt := range tests1 {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateSecretWithData(tt.args.secretName, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateSecretWithData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateSecretWithData() got = %v, want %v", got, tt.want)
			}
		})
	}

	GetSecretFunc = func(req *secretmanagerpb.GetSecretRequest) (*secretmanagerpb.Secret, error) {
		return secretPositiveReturn, nil
	}
	tests2 := []struct {
		name    string
		args    args
		want    *secretmanagerpb.SecretVersion
		wantErr bool
	}{
		{name: "FailSecretExists",
			args: args{
				projectID:  "myProject",
				secretName: "mySecret",
				payload:    []byte("a new test"),
			},
			want:    nil,
			wantErr: true},
	}
	for _, tt := range tests2 {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateSecretWithData(tt.args.secretName, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateSecretWithData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateSecretWithData() got = %v, want %v", got, tt.want)
			}
		})
	}
	GetSecretFunc = func(req *secretmanagerpb.GetSecretRequest) (*secretmanagerpb.Secret, error) {
		return nil, errors.New("secret does not exist")
	}
	CreateSecretFunc = func(req *secretmanagerpb.CreateSecretRequest) (*secretmanagerpb.Secret, error) {
		return nil, errors.New("failed to create secret")
	}
	tests3 := []struct {
		name    string
		args    args
		want    *secretmanagerpb.SecretVersion
		wantErr bool
	}{
		{name: "FailCreateSecret",
			args: args{
				projectID:  "myProject",
				secretName: "mySecret",
				payload:    []byte("a new test"),
			},
			want:    nil,
			wantErr: true},
	}
	for _, tt := range tests3 {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateSecretWithData(tt.args.secretName, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateSecretWithData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateSecretWithData() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteSecretVersion(t *testing.T) {
	DestroySecretVersionFunc = func(req *secretmanagerpb.DestroySecretVersionRequest) (*secretmanagerpb.SecretVersion, error) {
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
		want    *secretmanagerpb.SecretVersion
		wantErr bool
	}{
		{name: "Success",
			args: args{
				projectID:  "myProjects",
				secretName: "mySecrets",
				version:    "",
			},
			want:    secretVersionPositiveReturn,
			wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := DeleteSecretVersion(tt.args.secretName, tt.args.version); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteSecretVersion() = %v, want %v", got, tt.want)
			}
		})
	}
	DestroySecretVersionFunc = func(req *secretmanagerpb.DestroySecretVersionRequest) (*secretmanagerpb.SecretVersion, error) {
		return nil, errors.New("failed to delete secret version")
	}
	tests = []struct {
		name    string
		args    args
		want    *secretmanagerpb.SecretVersion
		wantErr bool
	}{
		{name: "Failure",
			args: args{
				projectID:  "myProjects",
				secretName: "mySecrets",
				version:    "",
			},
			want:    nil,
			wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := DeleteSecretVersion(tt.args.secretName, tt.args.version); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteSecretVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDisableSecret(t *testing.T) {
	DisableSecretVersionFunc = func(req *secretmanagerpb.DisableSecretVersionRequest) (*secretmanagerpb.SecretVersion, error) {
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
		want    *secretmanagerpb.SecretVersion
		wantErr bool
	}{
		{name: "Success",
			args: args{
				projectID:  "myProjects",
				secretName: "mySecrets",
				version:    "1",
			},
			want:    secretVersionPositiveReturn,
			wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := DisableSecret(tt.args.secretName, tt.args.version); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DisableSecret() = %v, want %v", got, tt.want)
			}
		})
	}
	DisableSecretVersionFunc = func(req *secretmanagerpb.DisableSecretVersionRequest) (*secretmanagerpb.SecretVersion, error) {
		return nil, errors.New("failed to Disable Secret")
	}
	tests = []struct {
		name    string
		args    args
		want    *secretmanagerpb.SecretVersion
		wantErr bool
	}{
		{name: "Failure",
			args: args{
				projectID:  "myProjects",
				secretName: "mySecrets",
				version:    "1",
			},
			want:    nil,
			wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := DisableSecret(tt.args.secretName, tt.args.version); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DisableSecret() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnableSecret(t *testing.T) {
	EnableSecretVersionFunc = func(req *secretmanagerpb.EnableSecretVersionRequest) (*secretmanagerpb.SecretVersion, error) {
		return secretVersionPositiveReturn, nil
	}
	type args struct {
		projectID  string
		secretName string
		version    string
	}
	tests := []struct {
		name string
		args args
		want *secretmanagerpb.SecretVersion
	}{
		{name: "Success",
			args: args{
				projectID:  "",
				secretName: "",
				version:    "",
			},
			want: secretVersionPositiveReturn},
		{name: "Failure",
			args: args{
				projectID:  "myProjects",
				secretName: "mySecrets",
				version:    "1",
			},
			want: secretVersionPositiveReturn},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := EnableSecret(tt.args.secretName, tt.args.version); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EnableSecret() = %v, want %v", got, tt.want)
			}
		})
	}
	EnableSecretVersionFunc = func(req *secretmanagerpb.EnableSecretVersionRequest) (*secretmanagerpb.SecretVersion, error) {
		return nil, errors.New("failed to enable Secret")
	}
	tests2 := []struct {
		name    string
		args    args
		want    *secretmanagerpb.SecretVersion
		wantErr bool
	}{
		{name: "Failure",
			args: args{
				projectID:  "myProjects",
				secretName: "mySecrets",
				version:    "1",
			},
			want:    nil,
			wantErr: true},
	}
	for _, tt := range tests2 {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := EnableSecret(tt.args.secretName, tt.args.version); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EnableSecret() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSecret(t *testing.T) {
	AccessSecretVersionFunc = func(req *secretmanagerpb.AccessSecretVersionRequest) (*secretmanagerpb.AccessSecretVersionResponse, error) {
		return &secretmanagerpb.AccessSecretVersionResponse{
			Name:    "projects/myProjects/secrets/mySecrets/versions/latest",
			Payload: &secretmanagerpb.SecretPayload{Data: []byte("mySecret")},
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
		want    *secretmanagerpb.SecretPayload
		wantErr bool
	}{
		{name: "Success",
			args: args{
				projectID:  "myProject",
				secretName: "mySecret",
				version:    "",
			},
			want:    &secretmanagerpb.SecretPayload{Data: []byte("mySecret")},
			wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := GetSecret(tt.args.secretName, tt.args.version); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSecret() = %v, want %v", got, tt.want)
			}
		})
	}
	AccessSecretVersionFunc = func(req *secretmanagerpb.AccessSecretVersionRequest) (*secretmanagerpb.AccessSecretVersionResponse, error) {
		return nil, errors.New("failed to get secret")
	}
	tests = []struct {
		name    string
		args    args
		want    *secretmanagerpb.SecretPayload
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
	GetSecretVersionFunc = func(req *secretmanagerpb.GetSecretVersionRequest) (*secretmanagerpb.SecretVersion, error) {
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
		want    *secretmanagerpb.SecretVersion
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
	GetSecretVersionFunc = func(req *secretmanagerpb.GetSecretVersionRequest) (*secretmanagerpb.SecretVersion, error) {
		return nil, errors.New("failed to Get Secret Version")
	}
	tests = []struct {
		name    string
		args    args
		want    *secretmanagerpb.SecretVersion
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
	GetSecretFunc = func(req *secretmanagerpb.GetSecretRequest) (*secretmanagerpb.Secret, error) {
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

	GetSecretFunc = func(req *secretmanagerpb.GetSecretRequest) (*secretmanagerpb.Secret, error) {
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
	DeleteSecretFunc = func(req *secretmanagerpb.DeleteSecretRequest) error {
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
	DeleteSecretFunc = func(req *secretmanagerpb.DeleteSecretRequest) error {
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
