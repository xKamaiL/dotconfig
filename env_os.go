//go:build appengine || go1.5
// +build appengine go1.5

package dotconfig

import "os"

var lookupEnv = os.LookupEnv
