package common

import "fmt"

type PatchOperation string

// Validate validates the PatchOperation string.
func (p *PatchOperation) Validate() error {
	if p == nil {
		return fmt.Errorf("patchOperation must not be nil")
	}
	switch *p {
	case "add":
	case "remove":
	case "replace":
	case "move":
	case "copy":
	case "test":
	default:
		return fmt.Errorf("invalid patch operation: %s", p)
	}
	return nil
}

const (
	PatchOperationAdd     PatchOperation = "add"     // Add operation as defined in IETF RFC 6902.
	PatchOperationCopy    PatchOperation = "copy"    // Copy operation as defined in IETF RFC 6902.
	PatchOperationMove    PatchOperation = "move"    // Move operation as defined in IETF RFC 6902.
	PatchOperationRemove  PatchOperation = "remove"  // Remove operation as defined in IETF RFC 6902.
	PatchOperationReplace PatchOperation = "replace" // Replace operation as defined in IETF RFC 6902.
	PatchOperationTest    PatchOperation = "test"    // Test operation as defined in IETF RFC 6902.
)

type UriScheme string

// Validate validates the UriScheme string.
func (u *UriScheme) Validate() error {
	if u == nil {
		return fmt.Errorf("uriScheme must not be nil")
	}
	switch *u {
	case "http":
	case "https":
	default:
		return fmt.Errorf("invalid uri scheme: %s", u)
	}
	return nil
}

const (
	UriSchemeHttp  UriScheme = "http"  // HTTP URI scheme.
	UriSchemeHttps UriScheme = "https" // HTTPS URI scheme.
)

type ChangeType string

// Validate validates the ChangeType string.
func (c *ChangeType) Validate() error {
	if c == nil {
		return fmt.Errorf("changeType must not be nil")
	}
	switch *c {
	case "ADD":
	case "MOVE":
	case "REMOVE":
	case "REPLACE":
	default:
		return fmt.Errorf("invalid change type: %s", c)
	}
	return nil
}

const (
	ChangeTypeAdd     ChangeType = "ADD"     // This value indicates new attribute has been added to the resource.
	ChangeTypeMove    ChangeType = "MOVE"    // This value indicates existing attribute has been moved to a different path in the resource.
	ChangeTypeRemove  ChangeType = "REMOVE"  // This value indicates existing attribute has been deleted from the resource.
	ChangeTypeReplace ChangeType = "REPLACE" // This value indicates existing attribute has been updated with new value.
)
