package gcp_secret_manager

import (
	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"errors"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
	"reflect"
	"testing"
)

func init() {
	Client = &MockClient{}
}

func TestAddNewSecretVersion(t *testing.T) {
	AddSecretVersionFunc = func(req *secretmanagerpb.AddSecretVersionRequest) (*secretmanagerpb.SecretVersion, error) {
		return &secretmanagerpb.SecretVersion{
			Name:        "projects/myProject/secrets/mySecrets/versions/1",
			CreateTime:  nil,
			DestroyTime: nil,
			State:       0,
		}, nil
	}

	type args struct {
		secretName string
		payload    []byte
	}
	tests := []struct {
		name string
		args args
		want *secretmanagerpb.SecretVersion
	}{
		{name: "Success",
			args: args{secretName: "mysecret", payload: []byte("a new test")},
			want: &secretmanagerpb.SecretVersion{
				Name:        "projects/myProject/secrets/mySecrets/versions/1",
				CreateTime:  nil,
				DestroyTime: nil,
				State:       0,
			}},
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
	GetSecretFunc = func(req *secretmanagerpb.GetSecretRequest) (*secretmanagerpb.Secret, error) {
		return &secretmanagerpb.Secret{
			Name:        "nil",
			Replication: nil,
			CreateTime:  nil,
			Labels:      nil,
		}, errors.New("Secret does not exist")
	}
	CreateSecretFunc = func(req *secretmanagerpb.CreateSecretRequest) (*secretmanagerpb.Secret, error) {
		return &secretmanagerpb.Secret{
			Name:        "projects/myProject/secrets/mySecrets",
			Replication: nil,
			CreateTime:  nil,
			Labels:      nil,
		}, nil
	}
	type args struct {
		projectID  string
		secretName string
	}
	tests := []struct {
		name string
		args args
		want *secretmanagerpb.Secret
	}{
		{name: "Success",
			args: args{
				projectID:  "myProject",
				secretName: "mySecret",
			},
			want: &secretmanagerpb.Secret{
				Name:        "projects/myProject/secrets/mySecrets",
				Replication: nil,
				CreateTime:  nil,
				Labels:      nil,
			}},
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
	GetSecretFunc = func(req *secretmanagerpb.GetSecretRequest) (*secretmanagerpb.Secret, error) {
		return &secretmanagerpb.Secret{
			Name:        "nil",
			Replication: nil,
			CreateTime:  nil,
			Labels:      nil,
		}, errors.New("Secret does not exist")
	}
	CreateSecretFunc = func(req *secretmanagerpb.CreateSecretRequest) (*secretmanagerpb.Secret, error) {
		return &secretmanagerpb.Secret{
			Name:        "projects/myProject/secrets/mySecrets",
			Replication: nil,
			CreateTime:  nil,
			Labels:      nil,
		}, nil
	}
	AddSecretVersionFunc = func(req *secretmanagerpb.AddSecretVersionRequest) (*secretmanagerpb.SecretVersion, error) {
		return &secretmanagerpb.SecretVersion{
			Name:        "projects/myProject/secrets/secrets/versions/1",
			CreateTime:  nil,
			DestroyTime: nil,
			State:       0,
		}, nil
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
			}, want: &secretmanagerpb.SecretVersion{
			Name:        "projects/myProject/secrets/secrets/versions/1",
			CreateTime:  nil,
			DestroyTime: nil,
			State:       0,
		}},
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



func TestDeleteSecretVersion(t *testing.T) {
	DestroySecretVersionFunc = func(req *secretmanagerpb.DestroySecretVersionRequest) (*secretmanagerpb.SecretVersion, error) {
		return &secretmanagerpb.SecretVersion{
			Name:        "projects/myProjects/secrets/mySecrets/versions/1",
			CreateTime:  nil,
			DestroyTime: nil,
			State:       0,
		}, nil
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
				projectID:  "myProjects",
				secretName: "mySecrets",
				version:    "",
			},
			want: &secretmanagerpb.SecretVersion{
				Name:        "projects/myProjects/secrets/mySecrets/versions/1",
				CreateTime:  nil,
				DestroyTime: nil,
				State:       0,
			}},
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
	DisableSecretVersionFunc = func(req *secretmanagerpb.DisableSecretVersionRequest) (*secretmanagerpb.SecretVersion, error) {
		return &secretmanagerpb.SecretVersion{
			Name:        "projects/myProjects/secrets/mySecrets/versions/1",
			CreateTime:  nil,
			DestroyTime: nil,
			State:       0,
		}, nil
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
				projectID:  "myProjects",
				secretName: "mySecrets",
				version:    "1",
			},
			want: &secretmanagerpb.SecretVersion{
				Name:        "projects/myProjects/secrets/mySecrets/versions/1",
				CreateTime:  nil,
				DestroyTime: nil,
				State:       0,
			}},
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
	EnableSecretVersionFunc = func(req *secretmanagerpb.EnableSecretVersionRequest) (*secretmanagerpb.SecretVersion, error) {
		return &secretmanagerpb.SecretVersion{
			Name:        "projects/myProjects/secrets/mySecrets/versions/1",
			CreateTime:  nil,
			DestroyTime: nil,
			State:       0,
		}, nil
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
			want: &secretmanagerpb.SecretVersion{
				Name:        "projects/myProjects/secrets/mySecrets/versions/1",
				CreateTime:  nil,
				DestroyTime: nil,
				State:       0,
			}},
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
		name string
		args args
		want *secretmanagerpb.SecretPayload
	}{
		{name: "Success",
			args: args{
				projectID:  "myProject",
				secretName: "mySecret",
				version:    "",
			},
			want: &secretmanagerpb.SecretPayload{Data: []byte("mySecret")}},
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
	GetSecretVersionFunc = func(req *secretmanagerpb.GetSecretVersionRequest) (*secretmanagerpb.SecretVersion, error) {
		return &secretmanagerpb.SecretVersion{
			Name:        "projects/myProject/secrets/mySecrets/versions/1",
			CreateTime:  nil,
			DestroyTime: nil,
			State:       0,
		}, nil
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
				projectID:  "myProject",
				secretName: "mySecrets",
				version:    "1",
			},
			want: &secretmanagerpb.SecretVersion{
				Name:        "projects/myProject/secrets/mySecrets/versions/1",
				CreateTime:  nil,
				DestroyTime: nil,
				State:       0,
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetSecretMetadata(tt.args.secretName, tt.args.version); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSecretMetadata() = %v, want %v", got, tt.want)
			}
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

func TestSecretExists(t *testing.T) {
	type args struct {
		secretName string
	}
	GetSecretFunc = func(req *secretmanagerpb.GetSecretRequest) (*secretmanagerpb.Secret, error) {
		return &secretmanagerpb.Secret{
			Name:        "projects/myProject/secrets/my-secret",
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
		return errors.New("Delete failed")
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
}