package keeper

import (
	"mandu/utils"
	"mandu/x/subscription/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SetSubscriptionRequest(ctx sdk.Context, subReq types.SubscriptionRequest) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SubscriptionRequestKeyPrefix))

	appendedValue := k.cdc.MustMarshal(&subReq)
	store.Set([]byte(subReq.Id), appendedValue)

	providerStore := prefix.NewStore(storeAdapter, types.GetRequesterStoreKey(subReq.Requester))
	providerStore.Set([]byte(subReq.Id), []byte{})
}

func (k Keeper) GetSubscriptionRequest(ctx sdk.Context, subReqId string) (subReq types.SubscriptionRequest, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SubscriptionRequestKeyPrefix))

	subReqBytes := store.Get([]byte(subReqId))
	if subReqBytes == nil {
		return subReq, false
	}

	k.cdc.MustUnmarshal(subReqBytes, &subReq)
	return subReq, true
}

// check if at least one subscription is active
func (k Keeper) IsSubscriptionRequestActive(ctx sdk.Context, subReq types.SubscriptionRequest) bool {
	for _, subscriptionId := range subReq.SubscriptionIds {
		subscription, found := k.GetSubscription(ctx, subscriptionId)
		if !found {
			continue
		}
		if subscription.StartEpoch <= utils.BlockToEpoch(ctx.BlockHeight(), subReq.EpochSize) &&
			subscription.EndEpoch >= utils.BlockToEpoch(ctx.BlockHeight(), subReq.EpochSize) {
			return true
		}
	}
	return false
}

// returns a map of subscription to provider
func (k Keeper) GetAllActiveSubscriptions(ctx sdk.Context, subReq types.SubscriptionRequest) map[string]string {
	subscriptions := make(map[string]string)
	for _, subscriptionId := range subReq.SubscriptionIds {
		subscription, found := k.GetSubscription(ctx, subscriptionId)
		if !found {
			continue
		}
		if subscription.StartEpoch <= utils.BlockToEpoch(ctx.BlockHeight(), subReq.EpochSize) &&
			subscription.EndEpoch >= utils.BlockToEpoch(ctx.BlockHeight(), subReq.EpochSize) {
			subscriptions[subscriptionId] = subscription.Subscriber
		}
	}
	return subscriptions
}

func (k Keeper) CalculateMinimumStake(ctx sdk.Context, subReq types.SubscriptionRequest) int64 {
	return 0
}

func (k Keeper) IsSubscriptionRequestUnavailable(status types.SubscriptionRequest_Status) bool {
	switch status {
	case types.SubscriptionRequest_CANCELLED, types.SubscriptionRequest_EXPIRED:
		return true
	default:
		return false
	}
}

func (k Keeper) SubscriptionRequestHasSubscriber(ctx sdk.Context, subReq types.SubscriptionRequest, subscriber string) bool {
	for _, subscriptionId := range subReq.SubscriptionIds {
		sub, _ := k.GetSubscription(ctx, subscriptionId)
		if sub.Subscriber == subscriber &&
			utils.BlockToEpoch(ctx.BlockHeight(), subReq.EpochSize) <= sub.EndEpoch {
			return true
		}
	}
	return false
}

func (k Keeper) CalculateBlockReward(ctx sdk.Context, subReq types.SubscriptionRequest) int64 {
	remainingEpochs := subReq.EndEpoch - utils.BlockToEpoch(ctx.BlockHeight(), subReq.EpochSize)
	return int64(subReq.AvailableAmount) / remainingEpochs
}

// Iterate over all subReqs and apply the given callback function
func (k Keeper) IterateSubscriptionRequests(ctx sdk.Context, shouldBreak func(subReq types.SubscriptionRequest) (bool, error)) error {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	iterator := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SubscriptionRequestKeyPrefix)).Iterator(nil, nil)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var subReq types.SubscriptionRequest
		k.cdc.MustUnmarshal(iterator.Value(), &subReq)
		stop, err := shouldBreak(subReq)
		if err != nil || stop {
			return err
		}
	}
	return nil
}

func (k Keeper) SetSubscription(ctx sdk.Context, subscription types.Subscription) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SubscriptionKeyPrefix))

	appendedValue := k.cdc.MustMarshal(&subscription)
	store.Set([]byte(subscription.Id), appendedValue)

	providerStore := prefix.NewStore(storeAdapter, types.GetSubscriberStoreKey(subscription.Subscriber))
	providerStore.Set([]byte(subscription.Id), []byte{})
}

func (k Keeper) GetSubscription(ctx sdk.Context, subscriptionId string) (subscription types.Subscription, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SubscriptionKeyPrefix))

	subscriptionBytes := store.Get([]byte(subscriptionId))
	if subscriptionBytes == nil {
		return subscription, false
	}

	k.cdc.MustUnmarshal(subscriptionBytes, &subscription)
	return subscription, true
}
