package sieve

import (
	"io"
	"strconv"
)

// Sieve is an io.Reader that ouputs a stream of newline seperated prime numbers
// prime numbers are generated using the sieve of eratosthenes.
type Sieve struct {
	nums    []bool
	current int
}

// New generates a new Sieve with the max value assigned
func New(max int) *Sieve {
	return &Sieve{nums: make([]bool, max/2), current: 2}
}

func (s *Sieve) Read(p []byte) (n int, err error) {
	if s.current == 2 {
		// special case 2 to cut down nums by half
		n := copy(p, []byte("2\n"))
		s.current = 3
		return n, nil
	}

	for s.current < len(s.nums)*2 {
		if s.nums[s.current/2] == false {
			n := copy(p, []byte(strconv.Itoa(s.current)+"\n"))

			for j := s.current; j < len(s.nums)*2; j += 2 * s.current {
				s.nums[j/2] = true
			}

			s.current += 2
			return n, nil
		}
		s.current += 2
	}

	return 0, io.EOF
}
