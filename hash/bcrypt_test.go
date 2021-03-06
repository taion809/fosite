package hash

import (
	"github.com/pborman/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHash(t *testing.T) {
	h := &BCrypt{
		WorkFactor: 10,
	}
	password := []byte("foo")
	hash, err := h.Hash(password)
	assert.Nil(t, err)
	assert.NotNil(t, hash)
	assert.NotEqual(t, hash, password)
}

func TestCompareEquals(t *testing.T) {
	h := &BCrypt{
		WorkFactor: 10,
	}
	password := []byte("foo")
	hash, err := h.Hash(password)
	assert.Nil(t, err)
	assert.NotNil(t, hash)
	err = h.Compare(hash, password)
	assert.Nil(t, err)
}

func TestCompareDifferent(t *testing.T) {
	h := &BCrypt{
		WorkFactor: 10,
	}
	password := []byte("foo")
	hash, err := h.Hash(password)
	assert.Nil(t, err)
	assert.NotNil(t, hash)
	err = h.Compare(hash, []byte(uuid.NewRandom().String()))
	assert.NotNil(t, err)
}
