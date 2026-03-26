//go:build !windows
// +build !windows

package model

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)
