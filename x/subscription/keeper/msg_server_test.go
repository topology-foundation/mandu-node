package keeper_test

import (
	"testing"

	"mandu/x/subscription/keeper"
	subscription "mandu/x/subscription/module"
	"mandu/x/subscription/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func setupMsgServer(t testing.TB) (keeper.Keeper, types.MsgServer, sdk.Context, subscription.AppModule) {
	k, ctx, am := MockSubscriptionKeeper(t)
	return k, keeper.NewMsgServerImpl(k), ctx, am
}

func TestMsgServer(t *testing.T) {
	k, ms, ctx, am := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
	require.NotNil(t, am)
	require.NotEmpty(t, k)
}
