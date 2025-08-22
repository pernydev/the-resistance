package utils

import (
	"math/rand"
	"time"
)

func Shuffle[T any](cards []T) []T {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})
	return cards
}
