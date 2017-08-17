package ec2

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

type aErrFunc func(assert.TestingT, error, ...interface{}) bool

func TestIAMRoleName(t *testing.T) {
	l := []struct{ n, v string; exp IAMRoleName; a aErrFunc; }{
		{ "valid", "abc", "abc", assert.NoError },
		{ "invalid", "!!Z", "", assert.Error },
	}
	for _, v := range l {
		t.Run(v.n, func(t *testing.T) {
			var r IAMRoleName
			if v.a(t, yaml.Unmarshal([]byte(v.v), &r)) {
				assert.Equal(t, v.exp, r)
			}
		})
	}
}

func TestELBName(t *testing.T) {
	l := []struct{ n, v string; exp ELBName; a aErrFunc; }{
		{ "valid", "abc-123", "abc-123", assert.NoError },
		{ "start with hyphen", "-Z", "", assert.Error },
		{ "ends with hyphen", "Z-", "", assert.Error },
		{ "non alpha-anumeric", "Z-()", "", assert.Error },
	}
	for _, v := range l {
		t.Run(v.n, func(t *testing.T) {
			var r ELBName
			t.Log(v.v)
			if v.a(t, yaml.Unmarshal([]byte(v.v), &r)) {
				assert.Equal(t, v.exp, r)
			}
		})
	}
}
