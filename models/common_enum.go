package models

type PatchOperation string

const (
	PatchOperationAdd     PatchOperation = "add"
	PatchOperationCopy    PatchOperation = "copy"
	PatchOperationMove    PatchOperation = "move"
	PatchOperationRemove  PatchOperation = "remove"
	PatchOperationReplace PatchOperation = "replace"
	PatchOperationTest    PatchOperation = "test"
)

type UriScheme string

const (
	UriSchemeHttp  UriScheme = "http"
	UriSchemeHttps UriScheme = "https"
)

type ChangeType string

const (
	ChangeTypeAdd     ChangeType = "ADD"
	ChangeTypeMove    ChangeType = "MOVE"
	ChangeTypeRemove  ChangeType = "REMOVE"
	ChangeTypeReplace ChangeType = "REPLACE"
)
