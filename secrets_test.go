package gcp_secret_manager

import (
	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"errors"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
	"reflect"
	"testing"
)

func TestAddNewSecretVersion(t *testing.T) {

	type args struct {
		projectID  string
		secretName string
		payload    []byte
	}
	tests := []struct {
		name string
		args args
		want *secretmanagerpb.SecretVersion
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddNewSecretVersion(tt.args.secretName, tt.args.payload); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddNewSecretVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateEmptySecret(t *testing.T) {
	type args struct {
		projectID  string
		secretName string
	}
	tests := []struct {
		name string
		args args
		want *secretmanagerpb.Secret
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateEmptySecret(tt.args.secretName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateEmptySecret() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateSecretWithData(t *testing.T) {
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
		// TODO: Add test cases.
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
}

func TestDeleteSecret(t *testing.T) {
	type args struct {
		projectID  string
		secretName string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestDeleteSecretVersion(t *testing.T) {
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteSecretVersion(tt.args.secretName, tt.args.version); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteSecretVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDisableSecret(t *testing.T) {
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DisableSecret(tt.args.secretName, tt.args.version); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DisableSecret() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnableSecret(t *testing.T) {
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EnableSecret(tt.args.secretName, tt.args.version); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EnableSecret() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSecret(t *testing.T) {
	type args struct {
		projectID  string
		secretName string
		version    string
	}
	tests := []struct {
		name string
		args args
		want *secretmanagerpb.SecretPayload
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetSecret(tt.args.secretName, tt.args.version); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSecret() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSecretMetadata(t *testing.T) {
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetSecretMetadata(tt.args.secretName, tt.args.version); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSecretMetadata() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListSecretVersions(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestListSecrets(t *testing.T) {
	type args struct {
		projectID string
	}
	tests := []struct {
		name string
		args args
		want *secretmanager.SecretIterator
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ListSecrets(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListSecrets() = %v, want %v", got, tt.want)
			}
		})
	}
}
func init() {
	Client = &MockClient{}
}

func TestSecretExists(t *testing.T) {
	type args struct {
		projectID  string
		secretName string
	}
	GetSecretFunc = func(req *secretmanagerpb.GetSecretRequest) (*secretmanagerpb.Secret, error) {
		return &secretmanagerpb.Secret{
			Name:        "projects/secret-manager-test/secrets/my-secret",
			Replication: nil,
			CreateTime:  nil,
			Labels:      nil,
		}, nil
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "success",
			args: args{projectID: "secret-manager-test", secretName: "my-secret"},
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
		return &secretmanagerpb.Secret{
			Name:        "nil",
			Replication: nil,
			CreateTime:  nil,
			Labels:      nil,
		}, errors.New("Secret not found")
	}
	tests2 := []struct {
		name string
		args args
		want bool
	}{
		{name: "failure",
			args: args{projectID: "secret-manager-test", secretName: "mysecret"},
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
