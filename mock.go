package gcp_secret_manager

import (
	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

// MockClient is the mock client
type MockClient struct {
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

func (m *MockClient) ListSecretVersions(req *secretmanagerpb.ListSecretVersionsRequest) SecretListIterator {
	return m.ListSecretVersionsFunc(req)
}

func (m *MockClient) AccessSecretVersion(req *secretmanagerpb.AccessSecretVersionRequest) (*secretmanagerpb.AccessSecretVersionResponse, error) {
	return m.AccessSecretVersionFunc(req)
}

func (m *MockClient) DestroySecretVersion(req *secretmanagerpb.DestroySecretVersionRequest) (*secretmanagerpb.SecretVersion, error) {
	return m.DestroySecretVersionFunc(req)
}

func (m *MockClient) CreateSecret(req *secretmanagerpb.CreateSecretRequest) (*secretmanagerpb.Secret, error) {
	return m.CreateSecretFunc(req)
}

func (m *MockClient) AddSecretVersion(req *secretmanagerpb.AddSecretVersionRequest) (*secretmanagerpb.SecretVersion, error) {
	return m.AddSecretVersionFunc(req)
}

func (m *MockClient) DeleteSecret(req *secretmanagerpb.DeleteSecretRequest) error {
	return m.DeleteSecretFunc(req)
}

func (m *MockClient) ListSecrets(req *secretmanagerpb.ListSecretsRequest) *secretmanager.SecretIterator {
	return m.ListSecretsFunc(req)
}

func (m *MockClient) GetSecretVersion(req *secretmanagerpb.GetSecretVersionRequest) (*secretmanagerpb.SecretVersion, error) {
	return m.GetSecretVersionFunc(req)
}

func (m *MockClient) DisableSecretVersion(req *secretmanagerpb.DisableSecretVersionRequest) (*secretmanagerpb.SecretVersion, error) {
	return m.DisableSecretVersionFunc(req)
}

func (m *MockClient) EnableSecretVersion(req *secretmanagerpb.EnableSecretVersionRequest) (*secretmanagerpb.SecretVersion, error) {
	return m.EnableSecretVersionFunc(req)
}

func (m *MockClient) Close() error {
	panic("implement me")
}

var (
	// GetSecretFunc fetches the mock client's `GetSecret` func
	GetSecretFunc        func(req *secretmanagerpb.GetSecretRequest) (*secretmanagerpb.Secret, error)
	ListSecretVersions   func(req *secretmanagerpb.ListSecretVersionsRequest) SecretListIterator
	AccessSecretVersion  func(req *secretmanagerpb.AccessSecretVersionRequest) (*secretmanagerpb.AccessSecretVersionResponse, error)
	DestroySecretVersion func(req *secretmanagerpb.DestroySecretVersionRequest) (*secretmanagerpb.SecretVersion, error)
	CreateSecret         func(req *secretmanagerpb.CreateSecretRequest) (*secretmanagerpb.Secret, error)
	AddSecretVersion     func(req *secretmanagerpb.AddSecretVersionRequest) (*secretmanagerpb.SecretVersion, error)
	DeleteSecret         func(req *secretmanagerpb.DeleteSecretRequest) error
	ListSecrets          func(req *secretmanagerpb.ListSecretsRequest) *secretmanager.SecretIterator
	GetSecretVersion     func(req *secretmanagerpb.GetSecretVersionRequest) (*secretmanagerpb.SecretVersion, error)
	DisableSecretVersion func(req *secretmanagerpb.DisableSecretVersionRequest) (*secretmanagerpb.SecretVersion, error)
	EnableSecretVersion  func(req *secretmanagerpb.EnableSecretVersionRequest) (*secretmanagerpb.SecretVersion, error)
)

func (m *MockClient) GetSecret(req *secretmanagerpb.GetSecretRequest) (*secretmanagerpb.Secret, error) {
	return m.GetSecretFunc(req)
}
