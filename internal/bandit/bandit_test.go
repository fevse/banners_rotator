package bandit_test

import (
	"testing"

	"github.com/fevse/banners_rotator/internal/bandit"
	"github.com/stretchr/testify/require"
)

func TestNewBandit(t *testing.T) {
	_, err := bandit.New(3)
	require.NoError(t, err)
}

func TestSelectArm(t *testing.T) {
	ban, err := bandit.New(3)
	require.NoError(t, err)
	ban.Clicks[1] = 0
	ban.Views[1] = 0
	ban.Rewards[1] = 1
	ban.Clicks[2] = 0
	ban.Views[2] = 0
	ban.Rewards[2] = 1
	ban.Clicks[3] = 0
	ban.Views[3] = 0
	ban.Rewards[3] = 1

	for i := 0; i < 3; i++ {
		_ = ban.SelectArm()
	}

	require.Equal(t, 1, ban.Views[1])
	require.Equal(t, 1, ban.Views[2])
	require.Equal(t, 1, ban.Views[3])

	ban.Update(1, 1.1)

	require.Equal(t, 1, ban.SelectArm())
	require.Equal(t, 2, ban.Views[1])
}

func TestUpdate(t *testing.T) {
	ban, err := bandit.New(3)
	require.NoError(t, err)
	ban.Clicks[1] = 0
	ban.Views[1] = 0
	ban.Rewards[1] = 1
	ban.Clicks[2] = 0
	ban.Views[2] = 0
	ban.Rewards[2] = 1
	ban.Clicks[3] = 0
	ban.Views[3] = 0
	ban.Rewards[3] = 1

	for i := 0; i < 3; i++ {
		_ = ban.SelectArm()
	}

	ban.Update(1, 1.1)
	ban.Update(1, 1.1)
	ban.Update(1, 1.1)
	ban.Update(1, 1.1)
	require.GreaterOrEqual(t, ban.Rewards[1], 1.1)
}
