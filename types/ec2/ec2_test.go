package ec2

import (
	"testing"
	"encoding/json"
	"github.com/stretchr/testify/assert"
)

type aErrFunc func(*testing.T, error, ...interface{})

func TestIAMRoleName(t *testing.T) {
	l := []struct{ n, v string; exp IAMRoleName; a aErrFunc; }{
		{ "valid", "abc", "abc", assert.NoError },
		{ "invalid", "+Z", "", assert.Error },
	}
	for _, v := range l {
		t.Run(v.n, func(t *testing.T) {
			var r IAMRoleName
			v.a(t, json.Unmarshal([]byte(v.v, &r)))
			assert.Equal(v.exp, r)
		})
	}
}

func TestELBName(t *testing.T) {
	l := []struct{ n, v string; exp ELBName; a aErrFunc; }{
		{ "valid", "abc-123", "abc-123", assert.NoError },
		{ "start with hyphen", "-Z", "", assert.Error },
		{ "ends with hyphen", "Z-", "", assert.Error },
		{ "non alph-anumeric", "Z-()", "", assert.Error },
	}
	for _, v := range l {
		t.Run(v.n, func(t *testing.T) {
			var r ELBName
			v.a(t, json.Unmarshal([]byte(v.v, &r)))
			assert.Equal(t, v.exp, r)
		})
	}
}
