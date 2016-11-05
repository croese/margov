package margov

import (
	"errors"
	"strings"
)

type followersMap map[string]map[string]int

// Generator represents a first-order
// Markov chain generator
type Generator struct {
	followers followersMap
}

// NewGenerator creates and initializes a
// Markov generator, seeded with the given string
func NewGenerator(seed string) (*Generator, error) {
	if seed == "" {
		return nil, errors.New("input string can't be empty")
	}

	gen := &Generator{followers: createFollowers(seed)}
	return gen, nil
}

func createFollowers(seed string) followersMap {
	followers := make(map[string]map[string]int)

	words := strings.Fields(seed)
	for i, word := range words[:len(words)-1] {
		if _, ok := followers[word]; !ok {
			followers[word] = make(map[string]int)
		}

		followers[word][words[i+1]]++
	}

	return followers
}

// Len returns the number of unique words
// that this Generator can pull from
func (g *Generator) Len() int {
	return len(g.followers)
}
