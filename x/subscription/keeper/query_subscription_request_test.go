package keeper_test

import (
	"testing"

	"mandu/x/subscription/types"

	qtypes "github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
)

func TestQuerySubscriptionRequest(t *testing.T) {
	k, ms, ctx, _ := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
	require.NotEmpty(t, k)

	// Create a subReq
	createSubscriptionRequest := types.MsgCreateSubscriptionRequest{Requester: Alice, CroId: "alicecro", Amount: 1000, StartBlock: 10, EndBlock: 20}
	createResponse, err := ms.CreateSubscriptionRequest(ctx, &createSubscriptionRequest)
	require.Nil(t, err)

	// Query the subReq
	queryResponse, err := k.SubscriptionRequest(ctx, &types.QuerySubscriptionRequestRequest{Id: createResponse.SubscriptionRequestId})
	require.Nil(t, err)

	require.EqualValues(t, createSubscriptionRequest, types.MsgCreateSubscriptionRequest{Requester: queryResponse.SubscriptionRequest.Requester, CroId: queryResponse.SubscriptionRequest.CroId, Amount: queryResponse.SubscriptionRequest.TotalAmount, StartBlock: queryResponse.SubscriptionRequest.StartBlock, EndBlock: queryResponse.SubscriptionRequest.EndBlock})
}

func TestQuerySubscriptionRequestStatus(t *testing.T) {
	k, ms, ctx, _ := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
	require.NotEmpty(t, k)

	// Create a subReq
	createSubscriptionRequest := types.MsgCreateSubscriptionRequest{Requester: Alice, CroId: "alicecro", Amount: 1000, StartBlock: 10, EndBlock: 20}
	createResponse, err := ms.CreateSubscriptionRequest(ctx, &createSubscriptionRequest)
	require.Nil(t, err)

	// Query the subReq
	queryStatusResponse, err := k.SubscriptionRequestStatus(ctx, &types.QuerySubscriptionRequestStatusRequest{Id: createResponse.SubscriptionRequestId})
	require.Nil(t, err)

	require.EqualValues(t, queryStatusResponse.Status, types.SubscriptionRequest_SCHEDULED)
}

func TestQuerySubscriptionRequests(t *testing.T) {
	k, ms, ctx, _ := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
	require.NotEmpty(t, k)

	// Create a first subReq
	createSubscriptionRequest1 := types.MsgCreateSubscriptionRequest{Requester: Alice, CroId: "alicecro1", Amount: 1000, StartBlock: 10, EndBlock: 20}
	createResponse1, err := ms.CreateSubscriptionRequest(ctx, &createSubscriptionRequest1)
	require.Nil(t, err)

	subReq1, found := k.GetSubscriptionRequest(ctx, createResponse1.SubscriptionRequestId)
	require.True(t, found)

	// Create a second subReq
	createSubscriptionRequest2 := types.MsgCreateSubscriptionRequest{Requester: Alice, CroId: "alicecro2", Amount: 1000, StartBlock: 20, EndBlock: 30}
	createResponse2, err := ms.CreateSubscriptionRequest(ctx, &createSubscriptionRequest2)
	require.Nil(t, err)

	subReq2, found := k.GetSubscriptionRequest(ctx, createResponse2.SubscriptionRequestId)
	require.True(t, found)

	// Create a third subReq
	createSubscriptionRequest3 := types.MsgCreateSubscriptionRequest{Requester: Alice, CroId: "alicecro", Amount: 1000, StartBlock: 10, EndBlock: 20}
	createResponse3, err := ms.CreateSubscriptionRequest(ctx, &createSubscriptionRequest3)
	require.Nil(t, err)

	subReq3, found := k.GetSubscriptionRequest(ctx, createResponse3.SubscriptionRequestId)
	require.True(t, found)

	// Query for all the subReqs by alice
	querySubscriptionRequestsResponse, err := k.SubscriptionRequests(ctx, &types.QuerySubscriptionRequestsRequest{Requester: Alice})
	require.Nil(t, err)

	require.Contains(t, querySubscriptionRequestsResponse.SubscriptionRequests, subReq1)
	require.Contains(t, querySubscriptionRequestsResponse.SubscriptionRequests, subReq2)
	require.Contains(t, querySubscriptionRequestsResponse.SubscriptionRequests, subReq3)
}

func TestQuerySubscriptionRequestsWithPagination(t *testing.T) {
	k, ms, ctx, _ := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
	require.NotEmpty(t, k)

	// Create a first subReq
	createSubscriptionRequest1 := types.MsgCreateSubscriptionRequest{Requester: Alice, CroId: "alicecro1", Amount: 1000, StartBlock: 10, EndBlock: 20}
	createResponse1, err := ms.CreateSubscriptionRequest(ctx, &createSubscriptionRequest1)
	require.Nil(t, err)

	subReq1, found := k.GetSubscriptionRequest(ctx, createResponse1.SubscriptionRequestId)
	require.True(t, found)

	// Create a second subReq
	createSubscriptionRequest2 := types.MsgCreateSubscriptionRequest{Requester: Alice, CroId: "alicecro2", Amount: 1000, StartBlock: 20, EndBlock: 30}
	createResponse2, err := ms.CreateSubscriptionRequest(ctx, &createSubscriptionRequest2)
	require.Nil(t, err)

	subReq2, found := k.GetSubscriptionRequest(ctx, createResponse2.SubscriptionRequestId)
	require.True(t, found)

	// Create a third subReq
	createSubscriptionRequest3 := types.MsgCreateSubscriptionRequest{Requester: Alice, CroId: "alicecro", Amount: 1000, StartBlock: 10, EndBlock: 20}
	createResponse3, err := ms.CreateSubscriptionRequest(ctx, &createSubscriptionRequest3)
	require.Nil(t, err)

	subReq3, found := k.GetSubscriptionRequest(ctx, createResponse3.SubscriptionRequestId)
	require.True(t, found)

	// Query for all the subReqs by alice
	querySubscriptionRequestsResponse, err := k.SubscriptionRequests(ctx, &types.QuerySubscriptionRequestsRequest{Requester: Alice, Pagination: &qtypes.PageRequest{Limit: 2}})
	require.Nil(t, err)

	require.Equal(t, len(querySubscriptionRequestsResponse.SubscriptionRequests), 2)
	require.Contains(t, []types.SubscriptionRequest{subReq1, subReq2, subReq3}, querySubscriptionRequestsResponse.SubscriptionRequests[0])
	require.Contains(t, []types.SubscriptionRequest{subReq1, subReq2, subReq3}, querySubscriptionRequestsResponse.SubscriptionRequests[1])
}
