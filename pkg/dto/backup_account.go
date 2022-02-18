package dto

import "github.com/KubeOperator/KubeOperator/pkg/model"

type BackupAccount struct {
	model.BackupAccount
	CredentialVars map[string]interface{} `json:"credentialVars"`
}

type BackupAccountOp struct {
	Operation string          `json:"operation" validate:"required"`
	Items     []BackupAccount `json:"items" validate:"required"`
}

type BackupAccountRequest struct {
	ID             string                 `json:"id" validate:"required"`
	Name           string                 `json:"name" validate:"required"`
	CredentialVars map[string]interface{} `json:"credentialVars" validate:"required"`
	Bucket         string                 `json:"bucket" validate:"required"`
	Type           string                 `json:"type" validate:"oneof=OSS AZURE SFTP"`
}

type CloudStorageRequest struct {
	CredentialVars map[string]interface{} `json:"credentialVars" validate:"required"`
	Type           string                 `json:"type" validate:"required"`
}
