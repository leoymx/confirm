package oauth2

import (
	"testing"
)

var (
	oa *OAGithub
)

func init() {
	oa = NewOAGithub("8ba2991113e68b4805c1", "secret", "user")
}

func TestAuth(t *testing.T) {
	t.Log(oa.AuthURL())
}
