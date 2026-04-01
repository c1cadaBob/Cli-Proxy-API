package gemini

import (
	"os"
	"strings"
)

func GetClientID() string {
	return lookupOAuthEnv("GEMINI_OAUTH_CLIENT_ID", "GEMINI_CLIENT_ID")
}

func GetClientSecret() string {
	return lookupOAuthEnv("GEMINI_OAUTH_CLIENT_SECRET", "GEMINI_CLIENT_SECRET")
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
