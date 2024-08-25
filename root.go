package gofaker

import (
	"math/rand"
	"time"
)

type Faker struct {
	r *rand.Rand
}

func New() *Faker {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return &Faker{r: r}
}
