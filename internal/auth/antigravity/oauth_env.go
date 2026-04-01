package antigravity

import (
	"os"
	"strings"
)

func GetClientID() string {
	return lookupOAuthEnv("ANTIGRAVITY_OAUTH_CLIENT_ID", "ANTIGRAVITY_CLIENT_ID")
}

func GetClientSecret() string {
	return lookupOAuthEnv("ANTIGRAVITY_OAUTH_CLIENT_SECRET", "ANTIGRAVITY_CLIENT_SECRET")
}

func lookupOAuthEnv(keys ...string) string {
	for _, key := range keys {
		if value, ok := os.LookupEnv(key); ok {
			if trimmed := strings.TrimSpace(value); trimmed != "" {
				return trimmed
			}
		}
	}
	return ""
}
