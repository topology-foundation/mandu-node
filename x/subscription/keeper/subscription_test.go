package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"mandu/x/subscription/types"

	query "github.com/cosmos/cosmos-sdk/types/query"
)

func TestGetSetSubscriptionRequest(t *testing.T) {
	k, ms, ctx, _ := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
	require.NotEmpty(t, k)

	// Create a subReq
	subReq := types.SubscriptionRequest{Id: "12345", CroId: "alicecro", Requester: Alice, Status: types.SubscriptionRequest_SCHEDULED, AvailableAmount: 1000, TotalAmount: 100, StartBlock: 10, EndBlock: 20}
	k.SetSubscriptionRequest(ctx, subReq)

	subReqResponse, found := k.GetSubscriptionRequest(ctx, subReq.Id)
	require.True(t, found)
	require.EqualValues(t, subReq, subReqResponse)
}

func TestSubscriptionRequestActive(t *testing.T) {
	k, ms, ctx, am := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
	require.NotEmpty(t, k)

	// create a subscription
	sub := types.Subscription{Id: "123", SubscriptionRequestId: "12345", Provider: "provider1", StartBlock: 10, EndBlock: 15}
	k.SetSubscription(ctx, sub)
	// Create a subReq
	subReq := types.SubscriptionRequest{Id: "12345", CroId: "alicecro", Requester: Alice, Status: types.SubscriptionRequest_UNDEFINED, AvailableAmount: 1000, TotalAmount: 100, StartBlock: 10, EndBlock: 20, SubscriptionIds: []string{"123"}}
	k.SetSubscriptionRequest(ctx, subReq)

	// The subReq must be inactive at block number 0
	isActive := k.IsSubscriptionRequestActive(ctx, subReq)
	require.False(t, isActive)

	// Jump to block 12
	ctx = MockBlockHeight(ctx, am, 12)

	// The subReq must be active at block number 12
	isActive = k.IsSubscriptionRequestActive(ctx, subReq)
	require.True(t, isActive)

	// Jump to block 18
	ctx = MockBlockHeight(ctx, am, 18)
	// The subReq must be inactive at block number 18
	isActive = k.IsSubscriptionRequestActive(ctx, subReq)
	require.False(t, isActive)
}

func TestGetAllActiveProviders(t *testing.T) {
	k, ms, ctx, am := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
	require.NotEmpty(t, k)

	// create a subscription
	sub := types.Subscription{Id: "123", SubscriptionRequestId: "12345", Provider: "provider1", StartBlock: 10, EndBlock: 15}
	k.SetSubscription(ctx, sub)
	// Create a subReq
	subReq := types.SubscriptionRequest{Id: "12345", CroId: "alicecro", Requester: Alice, Status: types.SubscriptionRequest_UNDEFINED, AvailableAmount: 1000, TotalAmount: 100, StartBlock: 10, EndBlock: 20, SubscriptionIds: []string{"123"}}
	k.SetSubscriptionRequest(ctx, subReq)

	activeSubs := k.GetAllActiveSubscriptions(ctx, subReq)
	// there shouldn't be any active subs at block 0
	require.True(t, len(activeSubs) == 0)

	// Jump to block 12
	ctx = MockBlockHeight(ctx, am, 12)
	activeSubs = k.GetAllActiveSubscriptions(ctx, subReq)
	// there should be an active subs at block 12
	require.True(t, len(activeSubs) == 1)
	_, ok := activeSubs[sub.Id]
	require.True(t, ok)

	// Jump to block 18
	ctx = MockBlockHeight(ctx, am, 18)
	activeSubs = k.GetAllActiveSubscriptions(ctx, subReq)
	// there shouldn't be an active subs at block 18
	require.True(t, len(activeSubs) == 0)
}

func TestIsSubscriptionRequestUnavailable(t *testing.T) {
	k, ms, ctx, _ := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
	require.NotEmpty(t, k)

	require.False(t, k.IsSubscriptionRequestUnavailable(types.SubscriptionRequest_ACTIVE))
	require.False(t, k.IsSubscriptionRequestUnavailable(types.SubscriptionRequest_SCHEDULED))
	require.False(t, k.IsSubscriptionRequestUnavailable(types.SubscriptionRequest_INITIALIZED))
	require.True(t, k.IsSubscriptionRequestUnavailable(types.SubscriptionRequest_CANCELLED))
	require.True(t, k.IsSubscriptionRequestUnavailable(types.SubscriptionRequest_EXPIRED))
}

func TestSubscriptionRequestHasProvider(t *testing.T) {
	k, ms, ctx, am := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
	require.NotEmpty(t, k)

	// create a subscription
	sub := types.Subscription{Id: "123", SubscriptionRequestId: "12345", Provider: "provider1", StartBlock: 10, EndBlock: 15}
	k.SetSubscription(ctx, sub)
	// Create a subReq
	subReq := types.SubscriptionRequest{Id: "12345", CroId: "alicecro", Requester: Alice, Status: types.SubscriptionRequest_UNDEFINED, AvailableAmount: 1000, TotalAmount: 100, StartBlock: 10, EndBlock: 20, SubscriptionIds: []string{"123"}}
	k.SetSubscriptionRequest(ctx, subReq)

	hasProvider := k.SubscriptionRequestHasProvider(ctx, subReq, "provider1")
	require.True(t, hasProvider)

	hasProvider = k.SubscriptionRequestHasProvider(ctx, subReq, "provider2")
	require.False(t, hasProvider)

	// Jump to block 18
	ctx = MockBlockHeight(ctx, am, 18)
	hasProvider = k.SubscriptionRequestHasProvider(ctx, subReq, "provider1")
	require.False(t, hasProvider)
}

func TestSubscription(t *testing.T) {
	keeper, ctx, _ := MockSubscriptionKeeper(t)
	subscription := types.Subscription{
		Id:       "sub1",
		Provider: "provider1",
	}

	keeper.SetSubscription(ctx, subscription)

	req := &types.QuerySubscriptionRequest{Id: "sub1"}

	retrievedSubscription, err := keeper.Subscription(ctx, req)
	require.NoError(t, err)
	require.Equal(t, subscription, retrievedSubscription.Subscription)
}

func TestSubscriptions(t *testing.T) {
	keeper, ctx, _ := MockSubscriptionKeeper(t)
	subscription1 := types.Subscription{
		Id:       "sub1",
		Provider: "provider1",
	}
	subscription2 := types.Subscription{
		Id:       "sub2",
		Provider: "provider1",
	}
	subscription3 := types.Subscription{
		Id:       "sub3",
		Provider: "provider2",
	}

	keeper.SetSubscription(ctx, subscription1)
	keeper.SetSubscription(ctx, subscription2)
	keeper.SetSubscription(ctx, subscription3)

	req := &types.QuerySubscriptionsRequest{Provider: "provider1"}
	res, err := keeper.Subscriptions(ctx, req)
	require.NoError(t, err)
	require.Len(t, res.Subscriptions, 2)
	require.Contains(t, res.Subscriptions, subscription1)
	require.Contains(t, res.Subscriptions, subscription2)
}

func TestSubscriptionsWithPaginationOne(t *testing.T) {
	keeper, ctx, _ := MockSubscriptionKeeper(t)
	subscription1 := types.Subscription{
		Id:       "sub1",
		Provider: "provider1",
	}
	subscription2 := types.Subscription{
		Id:       "sub2",
		Provider: "provider1",
	}
	subscription3 := types.Subscription{
		Id:       "sub3",
		Provider: "provider2",
	}

	keeper.SetSubscription(ctx, subscription1)
	keeper.SetSubscription(ctx, subscription2)
	keeper.SetSubscription(ctx, subscription3)

	req := &types.QuerySubscriptionsRequest{Provider: "provider1", Pagination: &query.PageRequest{Limit: 1}}
	res, err := keeper.Subscriptions(ctx, req)
	require.NoError(t, err)
	require.Len(t, res.Subscriptions, 1)
	require.Contains(t, res.Subscriptions, subscription1)
	req = &types.QuerySubscriptionsRequest{Provider: "provider1", Pagination: &query.PageRequest{Key: res.Pagination.NextKey, Limit: 1}}
	res, err = keeper.Subscriptions(ctx, req)
	require.NoError(t, err)

	require.Len(t, res.Subscriptions, 1)
	require.Contains(t, res.Subscriptions, subscription2)
}
