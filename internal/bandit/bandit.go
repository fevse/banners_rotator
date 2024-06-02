package bandit

import (
	"errors"
	"fmt"
	"math"
)

var ErrorInvalidNumberArms = errors.New("wrong number of arms")

type Bandit struct {
	Clicks  map[int]int
	Views   map[int]int
	Rewards map[int]float64
}

func New(n int) (Bandit, error) {
	if n >= 1 {
		return Bandit{
			Clicks:  make(map[int]int),
			Views:   make(map[int]int),
			Rewards: make(map[int]float64),
		}, nil
	}
	return Bandit{}, ErrorInvalidNumberArms
}

func (b *Bandit) Update(arm int, reward float64) error {
	b.Clicks[arm]++
	if b.Views[arm] == 0 {
		return fmt.Errorf("views = 0, should to choose banner to view")
	}
	b.Rewards[arm] = (b.Rewards[arm] * ((float64(b.Clicks[arm]) - 1) + reward)) / float64(b.Views[arm])
	return nil
}

func (b *Bandit) SelectArm() int {
	for i, v := range b.Views {
		if v == 0 {
			b.Views[i]++
			return i
		}
	}

	values := make(map[int]float64)

	for i := range b.Views {
		values[i] = b.Rewards[i] + math.Sqrt((1*math.Log(1+float64(b.Clicks[i])))/(1+float64(b.Views[i])))
	}

	arm := 0
	var m float64
	for i, v := range values {
		if v >= m {
			arm = i
			m = v
		}
	}
	b.Views[arm]++
	return arm
}
