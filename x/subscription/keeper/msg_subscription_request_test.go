package keeper_test

import (
	"testing"

	"mandu/x/subscription/types"

	"github.com/stretchr/testify/require"
)

func TestMsgServerCreateSubscriptionRequestMsg(t *testing.T) {
	k, ms, ctx, _ := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
	require.NotEmpty(t, k)

	response, err := ms.CreateSubscriptionRequest(ctx, &types.MsgCreateSubscriptionRequest{Requester: Alice, DrpId: "alicedrp", Amount: 10000, StartBlock: 10, EndBlock: 20})
	require.NoError(t, err)

	require.NotEmpty(t, response)
	require.NotEmpty(t, response.SubscriptionRequestId)
}

func TestMsgServerCreateSubscriptionRequestScheduled(t *testing.T) {
	k, ms, ctx, _ := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
	require.NotEmpty(t, k)

	createSubscriptionRequest := types.MsgCreateSubscriptionRequest{Requester: Alice, DrpId: "alicedrp", Amount: 10000, StartBlock: 10, EndBlock: 20}

	response, err := ms.CreateSubscriptionRequest(ctx, &createSubscriptionRequest)
	require.NoError(t, err)

	subReq, found := k.GetSubscriptionRequest(ctx, response.SubscriptionRequestId)
	require.True(t, found)

	require.EqualValues(t, response.SubscriptionRequestId, subReq.Id)

	require.EqualValues(t, createSubscriptionRequest, types.MsgCreateSubscriptionRequest{Requester: subReq.Requester, DrpId: subReq.DrpId, Amount: subReq.TotalAmount, StartBlock: subReq.StartBlock, EndBlock: subReq.EndBlock})

	require.Equal(t, subReq.Status, types.SubscriptionRequest_SCHEDULED)
}

func TestMsgServerCreateSubscriptionRequestInitializedStatus(t *testing.T) {
	k, ms, ctx, am := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
	require.NotEmpty(t, k)

	createSubscriptionRequest := types.MsgCreateSubscriptionRequest{Requester: Alice, DrpId: "alicedrp", Amount: 10000, StartBlock: 10, EndBlock: 20}
	response, err := ms.CreateSubscriptionRequest(ctx, &createSubscriptionRequest)
	require.NoError(t, err)

	// Get the subReq from the storage
	subReq, found := k.GetSubscriptionRequest(ctx, response.SubscriptionRequestId)
	require.True(t, found)

	// Jump to block number 11
	ctx = MockBlockHeight(ctx, am, 10)

	// The subReq must be initialized after entering block 10
	subReq, _ = k.GetSubscriptionRequest(ctx, response.SubscriptionRequestId)

	require.Equal(t, subReq.Status, types.SubscriptionRequest_INITIALIZED)
}

func TestMsgServerCancelSubscriptionRequestCorrectRequester(t *testing.T) {
	k, ms, ctx, _ := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
	require.NotEmpty(t, k)

	createSubscriptionRequest := types.MsgCreateSubscriptionRequest{Requester: Alice, DrpId: "alicedrp", Amount: 10000, StartBlock: 10, EndBlock: 20}
	createResponse, err := ms.CreateSubscriptionRequest(ctx, &createSubscriptionRequest)
	require.NoError(t, err)

	// Get the subReq from the storage
	_, found := k.GetSubscriptionRequest(ctx, createResponse.SubscriptionRequestId)
	require.True(t, found)

	// Now send a cancel message
	cancelSubscriptionRequest := types.MsgCancelSubscriptionRequest{Requester: Alice, SubscriptionRequestId: createResponse.SubscriptionRequestId}
	_, err = ms.CancelSubscriptionRequest(ctx, &cancelSubscriptionRequest)
	require.NoError(t, err)

	// Get the subReq from the storage
	subReq, found := k.GetSubscriptionRequest(ctx, createResponse.SubscriptionRequestId)
	if !found {
		t.Fatalf("SubscriptionRequest not found")
	}
	require.EqualValues(t, subReq.Status, types.SubscriptionRequest_CANCELLED)
}

func TestMsgServerCancelSubscriptionRequestIncorrectRequester(t *testing.T) {
	k, ms, ctx, _ := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
	require.NotEmpty(t, k)

	createSubscriptionRequest := types.MsgCreateSubscriptionRequest{Requester: Alice, DrpId: "alicedrp", Amount: 10000, StartBlock: 10, EndBlock: 20}
	createResponse, err := ms.CreateSubscriptionRequest(ctx, &createSubscriptionRequest)
	require.NoError(t, err)

	// Get the subReq from the storage
	_, found := k.GetSubscriptionRequest(ctx, createResponse.SubscriptionRequestId)
	require.True(t, found)

	// Now send a cancel message
	cancelSubscriptionRequest := types.MsgCancelSubscriptionRequest{Requester: Bob, SubscriptionRequestId: createResponse.SubscriptionRequestId}
	_, err = ms.CancelSubscriptionRequest(ctx, &cancelSubscriptionRequest)

	// The error should not be nil because the incorrect requester sends the CancelSubscriptionRequest message
	require.NotNil(t, err)
}

func TestMsgServerUpdateSubscriptionRequestIncorrectRequesterMsg(t *testing.T) {}
func TestMsgServerUpdateScheduledSubscriptionRequestCorrectStartBlockMsg(t *testing.T) {
	k, ms, ctx, _ := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
	require.NotEmpty(t, k)

	createResponse, err := ms.CreateSubscriptionRequest(ctx, &types.MsgCreateSubscriptionRequest{Requester: Alice, DrpId: "alicedrp", Amount: 10000, StartBlock: 10, EndBlock: 20})
	require.NoError(t, err)

	updateSubscriptionRequest := types.MsgUpdateSubscriptionRequest{Requester: Alice, SubscriptionRequestId: createResponse.SubscriptionRequestId, StartBlock: 11}
	_, err = ms.UpdateSubscriptionRequest(ctx, &updateSubscriptionRequest)
	require.NoError(t, err)

	subReq, found := k.GetSubscriptionRequest(ctx, createResponse.SubscriptionRequestId)

	require.True(t, found)
	require.EqualValues(t, subReq.StartBlock, 11)
}

func TestMsgServerUpdateScheduledSubscriptionRequestIncorrectStartBlockMsg(t *testing.T) {
	k, ms, ctx, am := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
	require.NotEmpty(t, k)

	createResponse, err := ms.CreateSubscriptionRequest(ctx, &types.MsgCreateSubscriptionRequest{Requester: Alice, DrpId: "alicedrp", Amount: 10000, StartBlock: 10, EndBlock: 20})
	require.NoError(t, err)

	// Jump to block 9
	ctx = MockBlockHeight(ctx, am, 9)

	updateSubscriptionRequest := types.MsgUpdateSubscriptionRequest{Requester: Alice, SubscriptionRequestId: createResponse.SubscriptionRequestId, StartBlock: 7}
	_, err = ms.UpdateSubscriptionRequest(ctx, &updateSubscriptionRequest)

	require.NotNil(t, err)
}

func TestMsgServerUpdateInitiatedSubscriptionRequestIncorrectStartBlockMsg(t *testing.T) {
	k, ms, ctx, am := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
	require.NotEmpty(t, k)

	createResponse, err := ms.CreateSubscriptionRequest(ctx, &types.MsgCreateSubscriptionRequest{Requester: Alice, DrpId: "alicedrp", Amount: 10000, StartBlock: 10, EndBlock: 20})
	require.NoError(t, err)

	// Jump to block 12
	ctx = MockBlockHeight(ctx, am, 12)

	updateSubscriptionRequest := types.MsgUpdateSubscriptionRequest{Requester: Alice, SubscriptionRequestId: createResponse.SubscriptionRequestId, StartBlock: 7}
	_, err = ms.UpdateSubscriptionRequest(ctx, &updateSubscriptionRequest)

	// It should return an error because the StartBlock can't be updated once the subReq is initiated.
	require.NotNil(t, err)
}

func TestMsgServerUpdateScheduledSubscriptionRequestIncrementAmountMsg(t *testing.T) {
	k, ms, ctx, _ := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
	require.NotEmpty(t, k)

	createResponse, err := ms.CreateSubscriptionRequest(ctx, &types.MsgCreateSubscriptionRequest{Requester: Alice, DrpId: "alicedrp", Amount: 10000, StartBlock: 10, EndBlock: 20})
	require.NoError(t, err)

	updateSubscriptionRequest := types.MsgUpdateSubscriptionRequest{Requester: Alice, SubscriptionRequestId: createResponse.SubscriptionRequestId, Amount: 12000}
	_, err = ms.UpdateSubscriptionRequest(ctx, &updateSubscriptionRequest)
	require.NoError(t, err)

	subReq, found := k.GetSubscriptionRequest(ctx, createResponse.SubscriptionRequestId)
	require.True(t, found)
	require.EqualValues(t, subReq.TotalAmount, 12000)
}

func TestMsgServerUpdateScheduledSubscriptionRequestDecrementAmountMsg(t *testing.T) {
	k, ms, ctx, _ := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
	require.NotEmpty(t, k)

	createResponse, err := ms.CreateSubscriptionRequest(ctx, &types.MsgCreateSubscriptionRequest{Requester: Alice, DrpId: "alicedrp", Amount: 10000, StartBlock: 10, EndBlock: 20})
	require.NoError(t, err)

	_, err = ms.UpdateSubscriptionRequest(ctx, &types.MsgUpdateSubscriptionRequest{Requester: Alice, SubscriptionRequestId: createResponse.SubscriptionRequestId, Amount: 5000})
	require.NoError(t, err)

	subReq, found := k.GetSubscriptionRequest(ctx, createResponse.SubscriptionRequestId)
	require.True(t, found)
	require.EqualValues(t, subReq.TotalAmount, 5000)
}

func TestMsgServerUpdateScheduledSubscriptionRequestDecrementTotalAmountMsg(t *testing.T) {
	k, ms, ctx, _ := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
	require.NotEmpty(t, k)

	createResponse, err := ms.CreateSubscriptionRequest(ctx, &types.MsgCreateSubscriptionRequest{Requester: Alice, DrpId: "alicedrp", Amount: 10000, StartBlock: 10, EndBlock: 20})
	require.NoError(t, err)

	_, err = ms.UpdateSubscriptionRequest(ctx, &types.MsgUpdateSubscriptionRequest{Requester: Alice, SubscriptionRequestId: createResponse.SubscriptionRequestId, Amount: 0})
	require.NoError(t, err)

	subReq, found := k.GetSubscriptionRequest(ctx, createResponse.SubscriptionRequestId)
	require.True(t, found)
	// Amount should be unchanged because you cannot withdraw full amount while the subReq is still active.
	require.EqualValues(t, subReq.TotalAmount, 10000)
}

func TestMsgServerUpdateInitiatedSubscriptionRequestIncrementAmountMsg(t *testing.T) {
	k, ms, ctx, am := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
	require.NotEmpty(t, k)

	createResponse, err := ms.CreateSubscriptionRequest(ctx, &types.MsgCreateSubscriptionRequest{Requester: Alice, DrpId: "alicedrp", Amount: 10000, StartBlock: 10, EndBlock: 20})
	require.NoError(t, err)

	// Jump to block 12 to initiate the subReq
	ctx = MockBlockHeight(ctx, am, 12)

	_, err = ms.UpdateSubscriptionRequest(ctx, &types.MsgUpdateSubscriptionRequest{Requester: Alice, SubscriptionRequestId: createResponse.SubscriptionRequestId, Amount: 15000})
	require.NoError(t, err)

	subReq, found := k.GetSubscriptionRequest(ctx, createResponse.SubscriptionRequestId)
	require.True(t, found)
	require.EqualValues(t, subReq.TotalAmount, 15000)
}

func TestMsgServerUpdateInitiatedSubscriptionRequestDecrementAmountMsg(t *testing.T) {
	k, ms, ctx, am := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
	require.NotEmpty(t, k)

	createResponse, err := ms.CreateSubscriptionRequest(ctx, &types.MsgCreateSubscriptionRequest{Requester: Alice, DrpId: "alicedrp", Amount: 10000, StartBlock: 10, EndBlock: 20})
	require.NoError(t, err)

	// Jump to block 12 to initiate the subReq
	ctx = MockBlockHeight(ctx, am, 12)

	_, err = ms.UpdateSubscriptionRequest(ctx, &types.MsgUpdateSubscriptionRequest{Requester: Alice, SubscriptionRequestId: createResponse.SubscriptionRequestId, Amount: 9000})
	// It should return an error because you're not allowed to decrease the amount after subReq initiation
	require.NotNil(t, err)
}

func TestMsgServerJoinSubscriptionRequestBeforeInitiationMsg(t *testing.T) {
	k, ms, ctx, _ := setupMsgServer(t)

	require.NotNil(t, ms)
	require.NotNil(t, ctx)
	require.NotEmpty(t, k)

	// Create a new subReq
	createSubscriptionRequest := types.MsgCreateSubscriptionRequest{Requester: Alice, DrpId: "alicedrp", Amount: 1000, StartBlock: 10, EndBlock: 20}
	createResponse, err := ms.CreateSubscriptionRequest(ctx, &createSubscriptionRequest)
	require.NoError(t, err)

	subReqId := createResponse.SubscriptionRequestId

	// Get the subReq from the storage
	subReq, found := k.GetSubscriptionRequest(ctx, subReqId)

	require.True(t, found)
	// Assert the status of the subReq to be "SCHEDULED"
	require.EqualValues(t, subReq.Status, types.SubscriptionRequest_SCHEDULED)

	// Subscriber joins the subReq before it is initiated
	joinSubscriptionRequest := types.MsgJoinSubscriptionRequest{Subscriber: Bob, SubscriptionRequestId: subReqId}
	joinResponse, err := ms.JoinSubscriptionRequest(ctx, &joinSubscriptionRequest)
	require.NoError(t, err)

	// Check if the subscription exists
	sub, found := k.GetSubscription(ctx, joinResponse.SubscriptionId)

	require.True(t, found)
	require.EqualValues(t, sub.Subscriber, Bob)

	// Check if the subscription exists in the subReq's subscriptionIds
	subReq, _ = k.GetSubscriptionRequest(ctx, subReqId)

	// Assert that the last id in subReq's subscriptionIds' is sub's id
	require.EqualValues(t, subReq.SubscriptionIds[len(subReq.SubscriptionIds)-1], sub.Id)
	require.EqualValues(t, subReqId, sub.SubscriptionRequestId)
}

func TestMsgServerJoinInitiatedSubscriptionRequestMsg(t *testing.T) {
	k, ms, ctx, am := setupMsgServer(t)

	require.NotNil(t, ms)
	require.NotNil(t, ctx)
	require.NotEmpty(t, k)

	// Create a new subReq
	createSubscriptionRequest := types.MsgCreateSubscriptionRequest{Requester: Alice, DrpId: "alicedrp", Amount: 1000, StartBlock: 10, EndBlock: 20}
	createResponse, err := ms.CreateSubscriptionRequest(ctx, &createSubscriptionRequest)
	require.NoError(t, err)

	subReqId := createResponse.SubscriptionRequestId

	// Get the subReq from the storage
	subReq, found := k.GetSubscriptionRequest(ctx, subReqId)

	require.True(t, found)
	// Assert the status of the subReq to be "SCHEDULED"
	require.EqualValues(t, subReq.Status, types.SubscriptionRequest_SCHEDULED)

	// Jump to block 12
	ctx = MockBlockHeight(ctx, am, 12)
	// Subscriber joins the subReq after it is initiated
	joinSubscriptionRequest := types.MsgJoinSubscriptionRequest{Subscriber: Bob, SubscriptionRequestId: subReqId}
	joinResponse, err := ms.JoinSubscriptionRequest(ctx, &joinSubscriptionRequest)
	require.NoError(t, err)

	// Check if the subscription exists
	sub, found := k.GetSubscription(ctx, joinResponse.SubscriptionId)

	require.True(t, found)
	require.EqualValues(t, sub.Subscriber, Bob)

	// Check if the subscription exists in the subReq's subscriptionIds
	subReq, _ = k.GetSubscriptionRequest(ctx, subReqId)

	// Assert that the last id in subReq's subscriptionIds' is sub's id
	require.EqualValues(t, subReq.SubscriptionIds[len(subReq.SubscriptionIds)-1], sub.Id)
	require.EqualValues(t, subReqId, sub.SubscriptionRequestId)

	// Check if the subReq's status has changed to ACTIVE
	require.EqualValues(t, subReq.Status, types.SubscriptionRequest_ACTIVE)
}

func TestMsgServerJoinCancelledSubscriptionRequestMsg(t *testing.T) {
	k, ms, ctx, _ := setupMsgServer(t)

	require.NotNil(t, ms)
	require.NotNil(t, ctx)
	require.NotEmpty(t, k)

	// Create a new subReq
	createSubscriptionRequest := types.MsgCreateSubscriptionRequest{Requester: Alice, DrpId: "alicedrp", Amount: 1000, StartBlock: 10, EndBlock: 20}
	createResponse, err := ms.CreateSubscriptionRequest(ctx, &createSubscriptionRequest)
	require.NoError(t, err)

	subReqId := createResponse.SubscriptionRequestId

	// Get the subReq from the storage
	_, found := k.GetSubscriptionRequest(ctx, subReqId)
	require.True(t, found)

	// Cancel the subReq
	cancelSubscriptionRequest := types.MsgCancelSubscriptionRequest{Requester: Alice, SubscriptionRequestId: subReqId}
	_, err = ms.CancelSubscriptionRequest(ctx, &cancelSubscriptionRequest)
	require.NoError(t, err)

	// Check if the status is changed to CANCELLED
	subReq, _ := k.GetSubscriptionRequest(ctx, subReqId)
	require.EqualValues(t, subReq.Status, types.SubscriptionRequest_CANCELLED)

	// Subscriber joins the subReq before it is initiated
	joinSubscriptionRequest := types.MsgJoinSubscriptionRequest{Subscriber: Bob, SubscriptionRequestId: subReqId}
	_, err = ms.JoinSubscriptionRequest(ctx, &joinSubscriptionRequest)

	require.NotNil(t, err)
}

func TestMsgServerJoinSameSubscriptionRequestMoreThanOnceMsg(t *testing.T) {
	k, ms, ctx, _ := setupMsgServer(t)

	require.NotNil(t, ms)
	require.NotNil(t, ctx)
	require.NotEmpty(t, k)

	// Create a new subReq
	createSubscriptionRequest := types.MsgCreateSubscriptionRequest{Requester: Alice, DrpId: "alicedrp", Amount: 1000, StartBlock: 10, EndBlock: 20}
	createResponse, err := ms.CreateSubscriptionRequest(ctx, &createSubscriptionRequest)
	require.NoError(t, err)

	subReqId := createResponse.SubscriptionRequestId

	// Subscriber joins the subReq
	joinSubscriptionRequest := types.MsgJoinSubscriptionRequest{Subscriber: Bob, SubscriptionRequestId: subReqId}
	_, err = ms.JoinSubscriptionRequest(ctx, &joinSubscriptionRequest)
	require.NoError(t, err)

	// Subscriber tries to join the same subReq again
	_, err = ms.JoinSubscriptionRequest(ctx, &joinSubscriptionRequest)

	// It is disallowed to join a subReq already subscribed to
	require.NotNil(t, err)
}

func TestMsgServerIncrementSubscriptionRequestAmount(t *testing.T) {
	k, ms, ctx, _ := setupMsgServer(t)

	require.NotNil(t, ms)
	require.NotNil(t, ctx)
	require.NotEmpty(t, k)

	// Create a new subReq
	createSubscriptionRequest := types.MsgCreateSubscriptionRequest{Requester: Alice, DrpId: "alicedrp", Amount: 1000, StartBlock: 10, EndBlock: 20}
	createResponse, err := ms.CreateSubscriptionRequest(ctx, &createSubscriptionRequest)
	require.NoError(t, err)

	subReqId := createResponse.SubscriptionRequestId

	// topup the subReq amount
	incrementSubscriptionRequest := types.MsgIncrementSubscriptionRequestAmount{Requester: Alice, SubscriptionRequestId: subReqId, Amount: 1000}
	_, err = ms.IncrementSubscriptionRequestAmount(ctx, &incrementSubscriptionRequest)

	require.NoError(t, err)
}

func TestMsgServerIncrementSubscriptionRequestAmountIncorrectRequester(t *testing.T) {
	k, ms, ctx, _ := setupMsgServer(t)

	require.NotNil(t, ms)
	require.NotNil(t, ctx)
	require.NotEmpty(t, k)

	// Create a new subReq
	createSubscriptionRequest := types.MsgCreateSubscriptionRequest{Requester: Alice, DrpId: "alicedrp", Amount: 1000, StartBlock: 10, EndBlock: 20}
	createResponse, err := ms.CreateSubscriptionRequest(ctx, &createSubscriptionRequest)
	require.NoError(t, err)

	subReqId := createResponse.SubscriptionRequestId

	// topup the subReq amount
	incrementSubscriptionRequest := types.MsgIncrementSubscriptionRequestAmount{Requester: Bob, SubscriptionRequestId: subReqId, Amount: 1000}
	_, err = ms.IncrementSubscriptionRequestAmount(ctx, &incrementSubscriptionRequest)

	require.NotNil(t, err)
}

func TestMsgServerIncrementCancelledSubscriptionRequestAmount(t *testing.T) {
	k, ms, ctx, _ := setupMsgServer(t)

	require.NotNil(t, ms)
	require.NotNil(t, ctx)
	require.NotEmpty(t, k)

	// Create a new subReq
	createSubscriptionRequest := types.MsgCreateSubscriptionRequest{Requester: Alice, DrpId: "alicedrp", Amount: 1000, StartBlock: 10, EndBlock: 20}
	createResponse, err := ms.CreateSubscriptionRequest(ctx, &createSubscriptionRequest)
	require.NoError(t, err)

	subReqId := createResponse.SubscriptionRequestId

	// cancel the subReq
	cancelSubscriptionRequest := types.MsgCancelSubscriptionRequest{Requester: Alice, SubscriptionRequestId: subReqId}
	_, err = ms.CancelSubscriptionRequest(ctx, &cancelSubscriptionRequest)
	require.NoError(t, err)

	// topup the subReq amount
	incrementSubscriptionRequest := types.MsgIncrementSubscriptionRequestAmount{Requester: Alice, SubscriptionRequestId: subReqId, Amount: 1000}
	_, err = ms.IncrementSubscriptionRequestAmount(ctx, &incrementSubscriptionRequest)

	require.NotNil(t, err)
}

func TestMsgServerIncrementExpiredSubscriptionRequestAmount(t *testing.T) {
	k, ms, ctx, am := setupMsgServer(t)

	require.NotNil(t, ms)
	require.NotNil(t, ctx)
	require.NotEmpty(t, k)

	// Create a new subReq
	createSubscriptionRequest := types.MsgCreateSubscriptionRequest{Requester: Alice, DrpId: "alicedrp", Amount: 1000, StartBlock: 10, EndBlock: 20}
	createResponse, err := ms.CreateSubscriptionRequest(ctx, &createSubscriptionRequest)
	require.NoError(t, err)

	subReqId := createResponse.SubscriptionRequestId

	// Jump to block 12 to initialize the subReq
	ctx = MockBlockHeight(ctx, am, 12)
	// Jump to block 25 to expire the subReq
	ctx = MockBlockHeight(ctx, am, 25)

	// topup the subReq amount
	incrementSubscriptionRequest := types.MsgIncrementSubscriptionRequestAmount{Requester: Alice, SubscriptionRequestId: subReqId, Amount: 1000}
	_, err = ms.IncrementSubscriptionRequestAmount(ctx, &incrementSubscriptionRequest)

	require.NotNil(t, err)
}

func TestMsgServerLeaveJoinedSubscriptionRequestMsg(t *testing.T) {
	k, ms, ctx, _ := setupMsgServer(t)

	require.NotNil(t, ms)
	require.NotNil(t, ctx)
	require.NotEmpty(t, k)

	// Create a new subReq
	createSubscriptionRequest := types.MsgCreateSubscriptionRequest{Requester: Alice, DrpId: "alicedrp", Amount: 1000, StartBlock: 10, EndBlock: 20}
	createResponse, err := ms.CreateSubscriptionRequest(ctx, &createSubscriptionRequest)
	require.NoError(t, err)

	subReqId := createResponse.SubscriptionRequestId

	// Subscriber joins the subReq
	joinSubscriptionRequest := types.MsgJoinSubscriptionRequest{Subscriber: Bob, SubscriptionRequestId: subReqId}
	_, err = ms.JoinSubscriptionRequest(ctx, &joinSubscriptionRequest)

	require.NoError(t, err)

	leaveSubscriptionRequest := types.MsgLeaveSubscriptionRequest{Subscriber: Bob, SubscriptionRequestId: subReqId}
	// Subscriber tries to leave the subReq
	_, err = ms.LeaveSubscriptionRequest(ctx, &leaveSubscriptionRequest)

	require.NoError(t, err)
}

func TestMsgServerLeaveNotJoinedSubscriptionRequestMsg(t *testing.T) {
	k, ms, ctx, _ := setupMsgServer(t)

	require.NotNil(t, ms)
	require.NotNil(t, ctx)
	require.NotEmpty(t, k)

	// Create a new subReq
	createSubscriptionRequest := types.MsgCreateSubscriptionRequest{Requester: Alice, DrpId: "alicedrp", Amount: 1000, StartBlock: 10, EndBlock: 20}
	createResponse, err := ms.CreateSubscriptionRequest(ctx, &createSubscriptionRequest)
	require.NoError(t, err)

	subReqId := createResponse.SubscriptionRequestId

	leaveSubscriptionRequest := types.MsgLeaveSubscriptionRequest{Subscriber: Bob, SubscriptionRequestId: subReqId}
	// Subscriber tries to leave the subReq it has not joined
	_, err = ms.LeaveSubscriptionRequest(ctx, &leaveSubscriptionRequest)

	// It should error because you can't leave a subReq you did not join
	require.NotNil(t, err)
}

func TestMsgServerJoinLeaveJoinSubscriptionRequestlMsg(t *testing.T) {
	k, ms, ctx, am := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
	require.NotEmpty(t, k)

	// Create a new subReq
	createSubscriptionRequest := types.MsgCreateSubscriptionRequest{Requester: Alice, DrpId: "alicedrp", Amount: 1000, StartBlock: 10, EndBlock: 20}
	createResponse, err := ms.CreateSubscriptionRequest(ctx, &createSubscriptionRequest)
	require.NoError(t, err)

	subReqId := createResponse.SubscriptionRequestId

	// Subscriber joins the subReq
	joinSubscriptionRequest := types.MsgJoinSubscriptionRequest{Subscriber: Bob, SubscriptionRequestId: subReqId}
	_, err = ms.JoinSubscriptionRequest(ctx, &joinSubscriptionRequest)
	require.NoError(t, err)

	leaveSubscriptionRequest := types.MsgLeaveSubscriptionRequest{Subscriber: Bob, SubscriptionRequestId: subReqId}
	// Subscriber tries to leave the subReq it has not joined
	_, err = ms.LeaveSubscriptionRequest(ctx, &leaveSubscriptionRequest)
	require.NoError(t, err)

	// Jump one block forward
	ctx = MockBlockHeight(ctx, am, 1)
	// Subscriber joins the subReq again
	_, err = ms.JoinSubscriptionRequest(ctx, &joinSubscriptionRequest)

	require.NoError(t, err)
}
