package petnames

import (
	"math/rand"
	"time"
)

type Petnames struct {
	dictionary []string
}

// Opt overrides a default value in the client config.
type Opt func(c *Petnames)

func New(opts ...Opt) Petnames {
	return new(opts...)
}

func new(opts ...Opt) Petnames {
	rand.Seed(time.Now().UTC().Unix())

	pn := Petnames{
		dictionary: dictionary,
	}

	for _, opt := range opts {
		opt(&pn)
	}

	return pn
}

// UseDictionary overrides the default dictionary with
// the provided list.
func UseDictionary(dictionary []string) Opt {
	return func(pn *Petnames) {
		pn.dictionary = dictionary
	}
}

// Name returns a random petname from the underlying dictionary.
func (pn Petnames) Name() string {
	return pn.dictionary[rand.Intn(len(pn.dictionary))]
}
