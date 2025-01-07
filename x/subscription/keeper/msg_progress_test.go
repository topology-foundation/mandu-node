package keeper_test

import (
	"testing"

	"mandu/x/subscription/types"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/sha3"
)

func ObfuscatedDataHashHelper(verticesHashes []string, subscriber string) string {
	hasher := sha3.New256()
	for _, hash := range verticesHashes {
		hasher.Write([]byte(hash))
	}
	hasher.Write([]byte(subscriber))
	hashBytes := hasher.Sum(nil)
	return string(hashBytes)
}

func TestSubmitProgress(t *testing.T) {
	_, ms, ctx, am := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)

	response, err := ms.CreateSubscriptionRequest(ctx, &types.MsgCreateSubscriptionRequest{Requester: Alice, DrpIds: []string{"alicedrp"}, Amount: 10000, StartBlock: 10})
	require.NoError(t, err)

	subReqId := response.SubscriptionRequestId

	// Jump to block 12 to initiate the subReq
	ctx = MockBlockHeight(ctx, am, 12)
	// Subscriber joins the subReq after it is initiated
	joinSubscriptionRequest := types.MsgJoinSubscriptionRequest{Subscriber: Bob, SubscriptionRequestId: subReqId}
	joinResponse, err := ms.JoinSubscriptionRequest(ctx, &joinSubscriptionRequest)
	require.NoError(t, err)

	subscriptionId := joinResponse.SubscriptionId
	subscriberId := Bob

	// create mock vertices hashes and the corresponding obfuscated hash
	verticesHashes1 := []string{"000", "111", "222", "333", "444", "555"}
	obfuscatedHash1 := ObfuscatedDataHashHelper(verticesHashes1, subscriberId)

	// submit progress
	_, err = ms.SubmitProgress(ctx, &types.MsgSubmitProgress{Subscriber: subscriberId, SubscriptionId: subscriptionId, ObfuscatedVerticesHash: obfuscatedHash1})
	// There should not be any error
	require.NoError(t, err)
	// Jump to block 13
	ctx = MockBlockHeight(ctx, am, 13)

	// create mock vertices hashes and the corresponding obfuscated hash
	verticesHashes2 := []string{"666", "777", "888", "999", "1010"}
	obfuscatedHash2 := ObfuscatedDataHashHelper(verticesHashes2, subscriberId)

	_, err = ms.SubmitProgress(ctx, &types.MsgSubmitProgress{Subscriber: subscriberId, SubscriptionId: subscriptionId, PreviousVerticesHashes: verticesHashes1, ObfuscatedVerticesHash: obfuscatedHash2})

	// There should not be any error
	require.NoError(t, err)
}

func TestSubmitProgressWithIncorrectObfuscatedHash(t *testing.T) {
	_, ms, ctx, am := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)

	response, err := ms.CreateSubscriptionRequest(ctx, &types.MsgCreateSubscriptionRequest{Requester: Alice, DrpId: "alicedrp", Amount: 10000, StartBlock: 10, EndBlock: 20})
	require.NoError(t, err)

	subReqId := response.SubscriptionRequestId

	// Jump to block 12 to initiate the subReq
	ctx = MockBlockHeight(ctx, am, 12)
	// Subscriber joins the subReq after it is initiated
	joinSubscriptionRequest := types.MsgJoinSubscriptionRequest{Subscriber: Bob, SubscriptionRequestId: subReqId}
	joinResponse, err := ms.JoinSubscriptionRequest(ctx, &joinSubscriptionRequest)
	require.NoError(t, err)

	subscriptionId := joinResponse.SubscriptionId
	subscriberId := Bob

	// create mock vertices hashes and the corresponding obfuscated hash
	verticesHashes1 := []string{"000", "111", "222", "333", "444", "555"}
	// obfuscatedHash1 := MockObfuscatedDataHash(verticesHashes1, subscriberId)
	obfuscatedHash1 := "oogabooga"

	// submit progress
	_, err = ms.SubmitProgress(ctx, &types.MsgSubmitProgress{Subscriber: subscriberId, SubscriptionId: subscriptionId, ObfuscatedVerticesHash: obfuscatedHash1})
	// There should not be any error
	require.NoError(t, err)
	// Jump to block 13
	ctx = MockBlockHeight(ctx, am, 13)

	// create mock vertices hashes and the corresponding obfuscated hash
	verticesHashes2 := []string{"666", "777", "888", "999", "1010"}
	obfuscatedHash2 := ObfuscatedDataHashHelper(verticesHashes2, subscriberId)

	_, err = ms.SubmitProgress(ctx, &types.MsgSubmitProgress{Subscriber: subscriberId, SubscriptionId: subscriptionId, PreviousVerticesHashes: verticesHashes1, ObfuscatedVerticesHash: obfuscatedHash2})

	// There should be an error because you submitted the wrong vertices hashes
	require.Error(t, err)
}

func TestSubmitProgressAfterEpochDeadline(t *testing.T) {
	_, ms, ctx, am := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)

	response, err := ms.CreateSubscriptionRequest(ctx, &types.MsgCreateSubscriptionRequest{Requester: Alice, DrpId: "alicedrp", Amount: 10000, StartBlock: 10, EndBlock: 25})
	require.NoError(t, err)

	subReqId := response.SubscriptionRequestId

	// Jump to block 12 to initiate the subReq
	ctx = MockBlockHeight(ctx, am, 12)
	// Subscriber joins the subReq after it is initiated
	joinSubscriptionRequest := types.MsgJoinSubscriptionRequest{Subscriber: Bob, SubscriptionRequestId: subReqId}
	joinResponse, err := ms.JoinSubscriptionRequest(ctx, &joinSubscriptionRequest)
	require.NoError(t, err)

	subscriptionId := joinResponse.SubscriptionId
	subscriberId := Bob

	// create mock vertices hashes and the corresponding obfuscated hash
	verticesHashes1 := []string{"000", "111", "222", "333", "444", "555"}
	// obfuscatedHash1 := MockObfuscatedDataHash(verticesHashes1, subscriberId)
	obfuscatedHash1 := ObfuscatedDataHashHelper(verticesHashes1, subscriberId)

	// submit progress
	_, err = ms.SubmitProgress(ctx, &types.MsgSubmitProgress{Subscriber: subscriberId, SubscriptionId: subscriptionId, ObfuscatedVerticesHash: obfuscatedHash1})
	// There should not be any error
	require.NoError(t, err)
	// Jump to block 23
	ctx = MockBlockHeight(ctx, am, 23)

	// create mock vertices hashes and the corresponding obfuscated hash
	verticesHashes2 := []string{"666", "777", "888", "999", "1010"}
	obfuscatedHash2 := ObfuscatedDataHashHelper(verticesHashes2, subscriberId)

	_, err = ms.SubmitProgress(ctx, &types.MsgSubmitProgress{Subscriber: subscriberId, SubscriptionId: subscriptionId, PreviousVerticesHashes: verticesHashes1, ObfuscatedVerticesHash: obfuscatedHash2})

	// There should be an error beacuse you submit after the epoch deadline
	require.Error(t, err)
}
