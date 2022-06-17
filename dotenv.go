package dotconfig

import (
	"github.com/joho/godotenv"
)

// Load env from .env with os env fallback value
// into a Struct pointer
func Load(ptr any, filenames ...string) error {
	if len(filenames) > 0 {
		godotenv.Load(filenames...)
	}
	return Process("", ptr)
}
