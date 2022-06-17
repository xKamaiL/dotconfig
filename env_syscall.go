//go:build !appengine && !go1.5
// +build !appengine,!go1.5

package dotconfig

import "syscall"

var lookupEnv = syscall.Getenv
