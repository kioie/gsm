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
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

var (
	// Declare Vars
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

// MockClient is the mock client
type MockClient struct {
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
// GetSecret Mock Get Secret
func (m *MockClient) GetSecret(req *secretmanagerpb.GetSecretRequest) (*secretmanagerpb.Secret, error) {
	return GetSecretFunc(req)
}
// AccessSecretVersion Mock Access SecretVersion
func (m *MockClient) AccessSecretVersion(req *secretmanagerpb.AccessSecretVersionRequest) (*secretmanagerpb.AccessSecretVersionResponse, error) {
	return AccessSecretVersionFunc(req)
}
// DestroySecretVersion Mock Destroy Secret Version
func (m *MockClient) DestroySecretVersion(req *secretmanagerpb.DestroySecretVersionRequest) (*secretmanagerpb.SecretVersion, error) {
	return DestroySecretVersionFunc(req)
}
// CreateSecret Mock Create Secret Version
func (m *MockClient) CreateSecret(req *secretmanagerpb.CreateSecretRequest) (*secretmanagerpb.Secret, error) {
	return CreateSecretFunc(req)
}
// AddSecretVersion Mock Add Secret Version
func (m *MockClient) AddSecretVersion(req *secretmanagerpb.AddSecretVersionRequest) (*secretmanagerpb.SecretVersion, error) {
	return AddSecretVersionFunc(req)
}
// DeleteSecret Mock Delete Secret
func (m *MockClient) DeleteSecret(req *secretmanagerpb.DeleteSecretRequest) error {
	return DeleteSecretFunc(req)
}
// GetSecretVersion Mock Get Secret Version
func (m *MockClient) GetSecretVersion(req *secretmanagerpb.GetSecretVersionRequest) (*secretmanagerpb.SecretVersion, error) {
	return GetSecretVersionFunc(req)
}
// DisableSecretVersion Mock Disable Secret Version
func (m *MockClient) DisableSecretVersion(req *secretmanagerpb.DisableSecretVersionRequest) (*secretmanagerpb.SecretVersion, error) {
	return DisableSecretVersionFunc(req)
}
// EnableSecretVersion Mock Enable Secret Version
func (m *MockClient) EnableSecretVersion(req *secretmanagerpb.EnableSecretVersionRequest) (*secretmanagerpb.SecretVersion, error) {
	return EnableSecretVersionFunc(req)
}
// Close Mock Close Client
func (m *MockClient) Close() error {
	return nil
}
