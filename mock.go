package gcp_secret_manager

import (
	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

// MockClient is the mock client
type MockClient struct {
	GetSecretFunc func(req *secretmanagerpb.GetSecretRequest) (*secretmanagerpb.Secret, error)
}

func (m *MockClient) ListSecretVersions(req *secretmanagerpb.ListSecretVersionsRequest) SecretListIterator {
	panic("implement me")
}

func (m *MockClient) AccessSecretVersion(req *secretmanagerpb.AccessSecretVersionRequest) (*secretmanagerpb.AccessSecretVersionResponse, error) {
	panic("implement me")
}

func (m *MockClient) DestroySecretVersion(req *secretmanagerpb.DestroySecretVersionRequest) (*secretmanagerpb.SecretVersion, error) {
	panic("implement me")
}

func (m *MockClient) CreateSecret(req *secretmanagerpb.CreateSecretRequest) (*secretmanagerpb.Secret, error) {
	panic("implement me")
}

func (m *MockClient) AddSecretVersion(req *secretmanagerpb.AddSecretVersionRequest) (*secretmanagerpb.SecretVersion, error) {
	panic("implement me")
}

func (m *MockClient) DeleteSecret(req *secretmanagerpb.DeleteSecretRequest) error {
	panic("implement me")
}

func (m *MockClient) ListSecrets(req *secretmanagerpb.ListSecretsRequest) *secretmanager.SecretIterator {
	panic("implement me")
}

func (m *MockClient) GetSecretVersion(req *secretmanagerpb.GetSecretVersionRequest) (*secretmanagerpb.SecretVersion, error) {
	panic("implement me")
}

func (m *MockClient) DisableSecretVersion(req *secretmanagerpb.DisableSecretVersionRequest) (*secretmanagerpb.SecretVersion, error) {
	panic("implement me")
}

func (m *MockClient) EnableSecretVersion(req *secretmanagerpb.EnableSecretVersionRequest) (*secretmanagerpb.SecretVersion, error) {
	panic("implement me")
}

func (m *MockClient) Close() error {
	panic("implement me")
}

var (
	// GetSecretFunc fetches the mock client's `GetSecret` func
	GetSecretFunc func(req *secretmanagerpb.GetSecretRequest) (*secretmanagerpb.Secret, error)
)

func (m *MockClient) GetSecret(req *secretmanagerpb.GetSecretRequest) (*secretmanagerpb.Secret, error) {
	return GetSecretFunc(req)
}
