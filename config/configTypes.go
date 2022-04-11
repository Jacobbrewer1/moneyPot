package config

type (
	StructConfig struct {
		RemoteConfig *RemoteConfigStruct
		LocalConfig  *LocalConfigStruct
	}

	LocalConfigStruct struct {
		ConnectionStrings *Endpoints `json:"ConnectionStrings,omitempty"`
		UseDatabase       *bool      `json:"UseDatabase,omitempty"`
	}

	Endpoints struct {
		ConfigIpAddress *string `json:"ConfigIpAddress,omitempty"`
	}

	RemoteConfigStruct struct {
		Secrets          *ApiSecrets `json:"Secrets,omitempty"`
		SheetId          *string     `json:"SheetId,omitempty"`
		ConnectionString *string     `json:"ConnectionString"`
	}

	ApiSecrets struct {
		GoogleSheetCredentials *GoogleSheetCredentials `json:"GoogleSheetCredentials,omitempty"`
	}

	GoogleSheetCredentials struct {
		Type                    *string `json:"type,omitempty"`
		ProjectId               *string `json:"project_id,omitempty"`
		PrivateKeyId            *string `json:"private_key_id,omitempty"`
		PrivateKey              *string `json:"private_key,omitempty"`
		ClientEmail             *string `json:"client_email,omitempty"`
		ClientId                *string `json:"client_id,omitempty"`
		AuthUri                 *string `json:"auth_uri,omitempty"`
		TokenUri                *string `json:"token_uri,omitempty"`
		AuthProviderX509CertUrl *string `json:"auth_provider_x509_cert_url,omitempty"`
		ClientX509CertUrl       *string `json:"client_x509_cert_url,omitempty"`
	}
)

func (c StructConfig) UseDB() bool {
	return *c.LocalConfig.UseDatabase
}
