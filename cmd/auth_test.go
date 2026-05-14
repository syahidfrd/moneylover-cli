package cmd

import (
	"testing"
)

func TestParseCallbackURL(t *testing.T) {
	tests := []struct {
		name        string
		url         string
		wantAccess  string
		wantRefresh string
		wantExpire  int64
		wantErr     bool
	}{
		{
			name:        "valid URL",
			url:         "https://web.moneylover.me/login?access_token=abc123&refresh_token=def456&expire=1779365792&status=true",
			wantAccess:  "abc123",
			wantRefresh: "def456",
			wantExpire:  1779365792,
			wantErr:     false,
		},
		{
			name:    "missing query params",
			url:     "https://web.moneylover.me/login",
			wantErr: true,
		},
		{
			name:    "missing access_token",
			url:     "https://web.moneylover.me/login?refresh_token=def456&expire=123",
			wantErr: true,
		},
		{
			name:    "missing refresh_token",
			url:     "https://web.moneylover.me/login?access_token=abc123&expire=123",
			wantErr: true,
		},
		{
			name:        "URL-encoded values",
			url:         "https://web.moneylover.me/login?access_token=eyJhbGciOiJIUzI1NiJ9.eyJ0ZXN0IjoxfQ.sig&refresh_token=eyJhbGciOiJIUzI1NiJ9.eyJrZXkiOiJ2YWwifQ.sig&expire=2094120992&status=true",
			wantAccess:  "eyJhbGciOiJIUzI1NiJ9.eyJ0ZXN0IjoxfQ.sig",
			wantRefresh: "eyJhbGciOiJIUzI1NiJ9.eyJrZXkiOiJ2YWwifQ.sig",
			wantExpire:  2094120992,
			wantErr:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, err := parseCallbackURL(tt.url)
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if token.AccessToken != tt.wantAccess {
				t.Errorf("AccessToken = %q, want %q", token.AccessToken, tt.wantAccess)
			}
			if token.RefreshToken != tt.wantRefresh {
				t.Errorf("RefreshToken = %q, want %q", token.RefreshToken, tt.wantRefresh)
			}
			if token.Expire != tt.wantExpire {
				t.Errorf("Expire = %d, want %d", token.Expire, tt.wantExpire)
			}
		})
	}
}

func TestDecodeJWTClaims(t *testing.T) {
	// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0eXBlIjoiYWNjZXNzLXRva2VuIiwidXNlcklkIjoiMTIzIiwiY2xpZW50Ijoia0hpWmJGUU93NUxWIn0.sig
	// payload: {"type":"access-token","userId":"123","client":"kHiZbFQOw5LV"}
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0eXBlIjoiYWNjZXNzLXRva2VuIiwidXNlcklkIjoiMTIzIiwiY2xpZW50Ijoia0hpWmJGUU93NUxWIn0.fake_sig"

	claims, err := decodeJWTClaims(token)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if claims["type"] != "access-token" {
		t.Errorf("type = %v, want access-token", claims["type"])
	}
	if claims["userId"] != "123" {
		t.Errorf("userId = %v, want 123", claims["userId"])
	}
	if claims["client"] != "kHiZbFQOw5LV" {
		t.Errorf("client = %v, want kHiZbFQOw5LV", claims["client"])
	}
}

func TestDecodeJWTClaims_Invalid(t *testing.T) {
	tests := []struct {
		name  string
		token string
	}{
		{"empty", ""},
		{"no dots", "nodots"},
		{"one dot", "one.dot"},
		{"invalid base64", "header.!!!invalid!!!.sig"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := decodeJWTClaims(tt.token)
			if err == nil {
				t.Fatal("expected error, got nil")
			}
		})
	}
}
