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

	sm "cloud.google.com/go/secretmanager/apiv1"
	smpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

// ClientFactory is used to create SecretClient, which is the GRPC Secret Client
// in normal use, but can be mocked for tests.
type ClientFactory interface {
	NewSecretClient(ctx context.Context) (SecretClient, error)
}

/*// SecretClient is a wrapper around the secretmanager APIs that are used by smcache.
// It is entirely for the purpose of being able to mock these for testing.
type SecretClient interface {
	AccessSecretVersion(ctx context.Context, req *smpb.AccessSecretVersionRequest) (*smpb.AccessSecretVersionResponse, error)
	DestroySecretVersion(ctx context.Context, req *smpb.DestroySecretVersionRequest) (*smpb.SecretVersion, error)
	CreateSecret(ctx context.Context, req *smpb.CreateSecretRequest) (*smpb.Secret, error)
	AddSecretVersion(ctx context.Context, req *smpb.AddSecretVersionRequest) (*smpb.SecretVersion, error)
	DeleteSecret(ctx context.Context, req *smpb.DeleteSecretRequest) error
	GetSecret(ctx context.Context, req *smpb.GetSecretRequest) (*smpb.Secret, error)
	GetSecretVersion(ctx context.Context, req *smpb.GetSecretVersionRequest) (*smpb.SecretVersion, error)
	DisableSecretVersion(ctx context.Context, req *smpb.DisableSecretVersionRequest) (*smpb.SecretVersion, error)
	EnableSecretVersion(ctx context.Context, req *smpb.EnableSecretVersionRequest) (*smpb.SecretVersion, error)
	Close() error
}*/

// SecretListIterator is an interface for the GRPC secret manager response from ListSecretVersions.
type SecretListIterator interface {
	Next() (*smpb.SecretVersion, error)
}

type secretClientImpl struct {
	client *sm.Client
	//ctx    context.Context
}

// SecretClientFactoryImpl implements ClientFactory for the real GRPC client.
type SecretClientFactoryImpl struct{}

// NewSecretClient creates a GRPC NewClient for secretmanager.
func (*SecretClientFactoryImpl) NewSecretClient(ctx context.Context) (SecretClient, error) {
	c, err := sm.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to setup client: %w", err)
	}

	return &secretClientImpl{client: c}, nil
}

func (sc *secretClientImpl) AccessSecretVersion(ctx context.Context, req *smpb.AccessSecretVersionRequest) (*smpb.AccessSecretVersionResponse, error) {
	return sc.client.AccessSecretVersion(ctx, req)
}

func (sc *secretClientImpl) DestroySecretVersion(ctx context.Context, req *smpb.DestroySecretVersionRequest) (*smpb.SecretVersion, error) {
	return sc.client.DestroySecretVersion(ctx, req)
}
func (sc *secretClientImpl) CreateSecret(ctx context.Context, req *smpb.CreateSecretRequest) (*smpb.Secret, error) {
	return sc.client.CreateSecret(ctx, req)
}
func (sc *secretClientImpl) AddSecretVersion(ctx context.Context, req *smpb.AddSecretVersionRequest) (*smpb.SecretVersion, error) {
	return sc.client.AddSecretVersion(ctx, req)
}
func (sc *secretClientImpl) DeleteSecret(ctx context.Context, req *smpb.DeleteSecretRequest) error {
	return sc.client.DeleteSecret(ctx, req)
}

func (sc *secretClientImpl) GetSecret(ctx context.Context, req *smpb.GetSecretRequest) (*smpb.Secret, error) {
	return sc.client.GetSecret(ctx, req)
}

func (sc *secretClientImpl) GetSecretVersion(ctx context.Context, req *smpb.GetSecretVersionRequest) (*smpb.SecretVersion, error) {
	return sc.client.GetSecretVersion(ctx, req)
}

func (sc *secretClientImpl) DisableSecretVersion(ctx context.Context, req *smpb.DisableSecretVersionRequest) (*smpb.SecretVersion, error) {
	return sc.client.DisableSecretVersion(ctx, req)
}

func (sc *secretClientImpl) EnableSecretVersion(ctx context.Context, req *smpb.EnableSecretVersionRequest) (*smpb.SecretVersion, error) {
	return sc.client.EnableSecretVersion(ctx, req)
}

func (sc *secretClientImpl) Close() error {
	return sc.client.Close()
}
