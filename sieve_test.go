package sieve

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSieve(t *testing.T) {
	s := New(10)

	primes, err := ioutil.ReadAll(s)

	assert.Nil(t, err)
	assert.Equal(t, primes, []byte("2\n3\n5\n7\n"))
}
