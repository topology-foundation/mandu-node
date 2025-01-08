package keeper

import (
	"context"

	"mandu/x/subscription/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Subscription(goCtx context.Context, req *types.QuerySubscriptionRequest) (*types.QuerySubscriptionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	subscription, found := k.GetSubscription(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QuerySubscriptionResponse{Subscription: subscription}, nil
}

func (k Keeper) Subscriptions(goCtx context.Context, req *types.QuerySubscriptionsRequest) (*types.QuerySubscriptionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.GetSubscriberStoreKey(req.Subscriber))

	var subscriptions []types.Subscription
	pageRes, err := query.Paginate(store, req.Pagination, func(key []byte, _ []byte) error {
		subscription, found := k.GetSubscription(ctx, string(key))
		if !found {
			return sdkerrors.ErrKeyNotFound
		}

		subscriptions = append(subscriptions, subscription)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QuerySubscriptionsResponse{Subscriptions: subscriptions, Pagination: pageRes}, nil
}
