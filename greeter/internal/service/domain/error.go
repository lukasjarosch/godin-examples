package domain

import (
	"github.com/pkg/errors"
)

var EmptyNameError = errors.New("greeting name cannot be empty")
