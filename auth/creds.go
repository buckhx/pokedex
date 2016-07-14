package auth

import (
	"encoding/base64"
	"fmt"

	"golang.org/x/net/context"
	"google.golang.org/grpc/credentials"
)

type security int

const (
	public security = iota
	access
	basic
)

type creds struct {
	payload  string
	security security
}

// AccessCredentials generates grpc credentials based on the access token string
func BasicCredentials(key, secret string) credentials.PerRPCCredentials {
	payload := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s")))
	return creds{payload: payload, security: basic}
}

// AccessCredentials generates grpc credentials based on the access token string
func AccessCredentials(token string) credentials.PerRPCCredentials {
	return creds{payload: token, security: access}
}

// EmptyCredentials generates grpc credentials that can be used to call unsecured methods such as authentication
func PublicCredentials() credentials.PerRPCCredentials {
	return creds{security: public}
}

func (c creds) GetRequestMetadata(ctx context.Context, uri ...string) (md map[string]string, err error) {
	switch c.security {
	case public:
		break //TODO make sure md == nil is ok
	case access:
		md = map[string]string{AUTH_HEADER: BEARER_PREFIX + c.payload}
	case basic:
		md = map[string]string{AUTH_HEADER: BASIC_PREFIX + c.payload}
	default:
		err = fmt.Errorf("Invalid Credentials Security")
	}
	return
}

func (c creds) RequireTransportSecurity() bool {
	return false //TODO change
}