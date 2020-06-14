package gcp_secret_manager

import (
	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"errors"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
	"testing"
)

func TestMockClient_Close(t *testing.T) {
	AddSecretVersionFunc = func(req *secretmanagerpb.AddSecretVersionRequest) (*secretmanagerpb.SecretVersion, error) {
		return secretVersionPositiveReturn, nil
	}
	GetSecretFunc = func(req *secretmanagerpb.GetSecretRequest) (*secretmanagerpb.Secret, error) {
		return nil, errors.New("Secret does not exist")
	}
	GetSecretVersionFunc = func(req *secretmanagerpb.GetSecretVersionRequest) (*secretmanagerpb.SecretVersion, error) {
		return secretVersionPositiveReturn, nil
	}
	CreateSecretFunc = func(req *secretmanagerpb.CreateSecretRequest) (*secretmanagerpb.Secret, error) {
		return secretPositiveReturn, nil
	}
	DestroySecretVersionFunc = func(req *secretmanagerpb.DestroySecretVersionRequest) (*secretmanagerpb.SecretVersion, error) {
		return secretVersionPositiveReturn, nil
	}
	DisableSecretVersionFunc = func(req *secretmanagerpb.DisableSecretVersionRequest) (*secretmanagerpb.SecretVersion, error) {
		return secretVersionPositiveReturn, nil
	}
	EnableSecretVersionFunc = func(req *secretmanagerpb.EnableSecretVersionRequest) (*secretmanagerpb.SecretVersion, error) {
		return secretVersionPositiveReturn, nil
	}
	AccessSecretVersionFunc = func(req *secretmanagerpb.AccessSecretVersionRequest) (*secretmanagerpb.AccessSecretVersionResponse, error) {
		return &secretmanagerpb.AccessSecretVersionResponse{
			Name:    "projects/myProjects/secrets/mySecrets/versions/latest",
			Payload: &secretmanagerpb.SecretPayload{Data: []byte("mySecret")},
		}, nil
	}
	type fields struct {
		GetSecretFunc            func(req *secretmanagerpb.GetSecretRequest) (*secretmanagerpb.Secret, error)
		ListSecretVersionsFunc   func(req *secretmanagerpb.ListSecretVersionsRequest) SecretListIterator
		AccessSecretVersionFunc  func(req *secretmanagerpb.AccessSecretVersionRequest) (*secretmanagerpb.AccessSecretVersionResponse, error)
		DestroySecretVersionFunc func(req *secretmanagerpb.DestroySecretVersionRequest) (*secretmanagerpb.SecretVersion, error)
		CreateSecretFunc         func(req *secretmanagerpb.CreateSecretRequest) (*secretmanagerpb.Secret, error)
		AddSecretVersionFunc     func(req *secretmanagerpb.AddSecretVersionRequest) (*secretmanagerpb.SecretVersion, error)
		DeleteSecretFunc         func(req *secretmanagerpb.DeleteSecretRequest) error
		ListSecretsFunc          func(req *secretmanagerpb.ListSecretsRequest) *secretmanager.SecretIterator
		GetSecretVersionFunc     func(req *secretmanagerpb.GetSecretVersionRequest) (*secretmanagerpb.SecretVersion, error)
		DisableSecretVersionFunc func(req *secretmanagerpb.DisableSecretVersionRequest) (*secretmanagerpb.SecretVersion, error)
		EnableSecretVersionFunc  func(req *secretmanagerpb.EnableSecretVersionRequest) (*secretmanagerpb.SecretVersion, error)
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				GetSecretFunc:            nil,
				ListSecretVersionsFunc:   nil,
				AccessSecretVersionFunc:  nil,
				DestroySecretVersionFunc: nil,
				CreateSecretFunc:         nil,
				AddSecretVersionFunc:     nil,
				DeleteSecretFunc:         nil,
				ListSecretsFunc:          nil,
				GetSecretVersionFunc:     nil,
				DisableSecretVersionFunc: nil,
				EnableSecretVersionFunc:  nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockClient{
				GetSecretFunc:            tt.fields.GetSecretFunc,
				ListSecretVersionsFunc:   tt.fields.ListSecretVersionsFunc,
				AccessSecretVersionFunc:  tt.fields.AccessSecretVersionFunc,
				DestroySecretVersionFunc: tt.fields.DestroySecretVersionFunc,
				CreateSecretFunc:         tt.fields.CreateSecretFunc,
				AddSecretVersionFunc:     tt.fields.AddSecretVersionFunc,
				DeleteSecretFunc:         tt.fields.DeleteSecretFunc,
				ListSecretsFunc:          tt.fields.ListSecretsFunc,
				GetSecretVersionFunc:     tt.fields.GetSecretVersionFunc,
				DisableSecretVersionFunc: tt.fields.DisableSecretVersionFunc,
				EnableSecretVersionFunc:  tt.fields.EnableSecretVersionFunc,
			}
			if err := m.Close(); (err != nil) != tt.wantErr {
				t.Errorf("Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
