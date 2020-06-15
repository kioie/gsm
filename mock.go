package gcp_secret_manager

import (
	"context"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

// MockClient is the mock client
type MockClient struct {
	NewClientFactoryFunc     func(ctx context.Context) (SecretClient, error)
	GetSecretFunc            func(req *secretmanagerpb.GetSecretRequest) (*secretmanagerpb.Secret, error)
	AccessSecretVersionFunc  func(req *secretmanagerpb.AccessSecretVersionRequest) (*secretmanagerpb.AccessSecretVersionResponse, error)
	DestroySecretVersionFunc func(req *secretmanagerpb.DestroySecretVersionRequest) (*secretmanagerpb.SecretVersion, error)
	CreateSecretFunc         func(req *secretmanagerpb.CreateSecretRequest) (*secretmanagerpb.Secret, error)
	AddSecretVersionFunc     func(req *secretmanagerpb.AddSecretVersionRequest) (*secretmanagerpb.SecretVersion, error)
	DeleteSecretFunc         func(req *secretmanagerpb.DeleteSecretRequest) error
	GetSecretVersionFunc     func(req *secretmanagerpb.GetSecretVersionRequest) (*secretmanagerpb.SecretVersion, error)
	DisableSecretVersionFunc func(req *secretmanagerpb.DisableSecretVersionRequest) (*secretmanagerpb.SecretVersion, error)
	EnableSecretVersionFunc  func(req *secretmanagerpb.EnableSecretVersionRequest) (*secretmanagerpb.SecretVersion, error)
}

func (m *MockClient) NewClientFactory(ctx context.Context) (SecretClient, error) {
	return NewClientFactoryFunc(ctx)
}
func (m *MockClient) AccessSecretVersion(req *secretmanagerpb.AccessSecretVersionRequest) (*secretmanagerpb.AccessSecretVersionResponse, error) {
	return AccessSecretVersionFunc(req)
}

func (m *MockClient) DestroySecretVersion(req *secretmanagerpb.DestroySecretVersionRequest) (*secretmanagerpb.SecretVersion, error) {
	return DestroySecretVersionFunc(req)
}

func (m *MockClient) CreateSecret(req *secretmanagerpb.CreateSecretRequest) (*secretmanagerpb.Secret, error) {
	return CreateSecretFunc(req)
}

func (m *MockClient) AddSecretVersion(req *secretmanagerpb.AddSecretVersionRequest) (*secretmanagerpb.SecretVersion, error) {
	return AddSecretVersionFunc(req)
}

func (m *MockClient) DeleteSecret(req *secretmanagerpb.DeleteSecretRequest) error {
	return DeleteSecretFunc(req)
}

func (m *MockClient) GetSecretVersion(req *secretmanagerpb.GetSecretVersionRequest) (*secretmanagerpb.SecretVersion, error) {
	return GetSecretVersionFunc(req)
}

func (m *MockClient) DisableSecretVersion(req *secretmanagerpb.DisableSecretVersionRequest) (*secretmanagerpb.SecretVersion, error) {
	return DisableSecretVersionFunc(req)
}

func (m *MockClient) EnableSecretVersion(req *secretmanagerpb.EnableSecretVersionRequest) (*secretmanagerpb.SecretVersion, error) {
	return EnableSecretVersionFunc(req)
}

func (m *MockClient) Close() error {
	return nil
}

var (
	NewClientFactoryFunc     func(ctx context.Context) (SecretClient, error)
	GetSecretFunc            func(req *secretmanagerpb.GetSecretRequest) (*secretmanagerpb.Secret, error)
	AccessSecretVersionFunc  func(req *secretmanagerpb.AccessSecretVersionRequest) (*secretmanagerpb.AccessSecretVersionResponse, error)
	DestroySecretVersionFunc func(req *secretmanagerpb.DestroySecretVersionRequest) (*secretmanagerpb.SecretVersion, error)
	CreateSecretFunc         func(req *secretmanagerpb.CreateSecretRequest) (*secretmanagerpb.Secret, error)
	AddSecretVersionFunc     func(req *secretmanagerpb.AddSecretVersionRequest) (*secretmanagerpb.SecretVersion, error)
	DeleteSecretFunc         func(req *secretmanagerpb.DeleteSecretRequest) error
	GetSecretVersionFunc     func(req *secretmanagerpb.GetSecretVersionRequest) (*secretmanagerpb.SecretVersion, error)
	DisableSecretVersionFunc func(req *secretmanagerpb.DisableSecretVersionRequest) (*secretmanagerpb.SecretVersion, error)
	EnableSecretVersionFunc  func(req *secretmanagerpb.EnableSecretVersionRequest) (*secretmanagerpb.SecretVersion, error)
)

func (m *MockClient) GetSecret(req *secretmanagerpb.GetSecretRequest) (*secretmanagerpb.Secret, error) {
	return GetSecretFunc(req)
}
