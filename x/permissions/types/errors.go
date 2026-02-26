package types

// DONTCOVER

import (
	"cosmossdk.io/errors"
)

// x/permissions module sentinel errors
var (
	ErrInvalidSigner = errors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrUnauthorized  = errors.Register(ModuleName, 1101, "unauthorized")
	ErrNotFound     = errors.Register(ModuleName, 1102, "not found")
	ErrAlreadyExists = errors.Register(ModuleName, 1103, "already exists")
)
