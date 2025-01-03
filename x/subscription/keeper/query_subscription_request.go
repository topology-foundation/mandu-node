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

func (k Keeper) SubscriptionRequest(goCtx context.Context, req *types.QuerySubscriptionRequestRequest) (*types.QuerySubscriptionRequestResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	subReq, found := k.GetSubscriptionRequest(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QuerySubscriptionRequestResponse{SubscriptionRequest: subReq}, nil
}

func (k Keeper) SubscriptionRequestStatus(goCtx context.Context, req *types.QuerySubscriptionRequestStatusRequest) (*types.QuerySubscriptionRequestStatusResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	subReq, found := k.GetSubscriptionRequest(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QuerySubscriptionRequestStatusResponse{Status: subReq.Status}, nil
}

func (k Keeper) SubscriptionRequests(goCtx context.Context, req *types.QuerySubscriptionRequestsRequest) (*types.QuerySubscriptionRequestsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.GetRequesterStoreKey(req.Requester))

	var subReqs []types.SubscriptionRequest
	pageRes, err := query.Paginate(store, req.Pagination, func(key []byte, _ []byte) error {
		subReq, found := k.GetSubscriptionRequest(ctx, string(key))
		if !found {
			return sdkerrors.ErrKeyNotFound
		}

		subReqs = append(subReqs, subReq)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QuerySubscriptionRequestsResponse{SubscriptionRequests: subReqs, Pagination: pageRes}, nil
}
