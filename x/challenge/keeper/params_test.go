package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"topchain/x/challenge/types"
)

func TestGetParams(t *testing.T) {
	k, ctx, _ := MockChallengeKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
