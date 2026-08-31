package daemon

import (
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestValidateURLTestURL(t *testing.T) {
	testCases := []struct {
		name string
		url  string
		code codes.Code
	}{
		{"default", "", codes.OK},
		{"https", "https://example.com/generate_204", codes.OK},
		{"https with port and query", "https://example.com:8443/generate_204?source=test", codes.OK},
		{"http", "http://example.com/generate_204", codes.InvalidArgument},
		{"other scheme", "ftp://example.com/resource", codes.InvalidArgument},
		{"missing scheme", "example.com/generate_204", codes.InvalidArgument},
		{"missing host", "https://", codes.InvalidArgument},
		{"garbage", "abc", codes.InvalidArgument},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			if code := status.Code(validateURLTestURL(testCase.url)); code != testCase.code {
				t.Fatalf("unexpected status code: got %s, want %s", code, testCase.code)
			}
		})
	}
}
