package keeper

import (
	"context"
	"math"

	manduTypes "mandu/types"
	"mandu/utils"
	"mandu/utils/validation"
	"mandu/x/subscription/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/google/uuid"
)

func validateMsgCreateSubscriptionRequest(msg *types.MsgCreateSubscriptionRequest) error {
	startEpoch := utils.BlockToEpoch(msg.StartBlock, msg.EpochSize)
	endEpoch := startEpoch + msg.Duration
	if err := validation.ValidateEpochRange(startEpoch, endEpoch); err != nil {
		return err
	}
	if err := validation.ValidatePositiveAmount(msg.Amount); err != nil {
		return err
	}
	for _, v := range msg.DrpIds {
		if err := validation.ValidateNonEmptyString(v); err != nil {
			return err
		}
	}
	if err := validation.ValidateAddress(msg.Requester); err != nil {
		return err
	}
	return nil
}

func (k msgServer) CreateSubscriptionRequest(goCtx context.Context, msg *types.MsgCreateSubscriptionRequest) (*types.MsgCreateSubscriptionRequestResponse, error) {
	err := validateMsgCreateSubscriptionRequest(msg)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	id := uuid.NewString()
	startEpoch := utils.BlockToEpoch(msg.StartBlock, msg.EpochSize)
	subReq := types.SubscriptionRequest{
		Id:              id,
		Requester:       msg.Requester,
		Status:          types.SubscriptionRequest_SCHEDULED,
		TotalAmount:     msg.Amount,
		AvailableAmount: msg.Amount,
		StartBlock:      msg.StartBlock,
		EpochSize:       msg.EpochSize,
		Duration:        msg.Duration,
		StartEpoch:      startEpoch,
		EndEpoch:        startEpoch + msg.Duration,
		DrpIds:          msg.DrpIds,
		Writers:         msg.Writers,
		InitialFrontier: msg.InitialFrontier,
		SubscriptionIds: []string{},
	}

	requester, err := sdk.AccAddressFromBech32(msg.Requester)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "invalid requester address")
	}

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, requester, types.ModuleName, sdk.NewCoins(sdk.NewInt64Coin(manduTypes.TokenDenom, int64(msg.Amount))))
	if err != nil {
		return nil, errorsmod.Wrap(err, "failed to send coins to module account")
	}

	k.SetSubscriptionRequest(ctx, subReq)

	return &types.MsgCreateSubscriptionRequestResponse{SubscriptionRequestId: id}, nil
}

func validateMsgCancelSubscriptionRequest(msg *types.MsgCancelSubscriptionRequest) error {
	if err := validation.ValidateNonEmptyString(msg.SubscriptionRequestId); err != nil {
		return err
	}
	if err := validation.ValidateAddress(msg.Requester); err != nil {
		return err
	}
	return nil
}

func (k msgServer) CancelSubscriptionRequest(goCtx context.Context, msg *types.MsgCancelSubscriptionRequest) (*types.MsgCancelSubscriptionRequestResponse, error) {
	err := validateMsgCancelSubscriptionRequest(msg)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	subReq, found := k.GetSubscriptionRequest(ctx, msg.SubscriptionRequestId)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "subReq with id "+msg.SubscriptionRequestId+" not found")
	}
	if msg.Requester != subReq.Requester {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "only the requester can cancel the subReq")
	}
	if subReq.Status == types.SubscriptionRequest_SCHEDULED || subReq.Status == types.SubscriptionRequest_INITIALIZED {
		subReq.Status = types.SubscriptionRequest_CANCELLED
		k.SetSubscriptionRequest(ctx, subReq)
		// return the remaining amount to the requester
		err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sdk.AccAddress(subReq.Requester), sdk.NewCoins(sdk.NewInt64Coin(manduTypes.TokenDenom, int64(subReq.AvailableAmount))))
		if err != nil {
			return nil, errorsmod.Wrap(err, "failed to send coins to module account")
		}
		return &types.MsgCancelSubscriptionRequestResponse{}, nil
	}
	if subReq.Status == types.SubscriptionRequest_INACTIVE || subReq.Status == types.SubscriptionRequest_ACTIVE {
		subReq.Status = types.SubscriptionRequest_CANCELLED
		k.SetSubscriptionRequest(ctx, subReq)
		for _, subscriptionId := range subReq.SubscriptionIds {
			subscription, found := k.GetSubscription(ctx, subscriptionId)
			if !found {
				return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "SHOULD NOT HAPPEN: subscription with id "+subscriptionId+" not found")
			}
			subscription.EndEpoch = utils.BlockToEpoch(ctx.BlockHeight(), subReq.EpochSize)
			k.SetSubscription(ctx, subscription)
		}
		// return the remaining amount to the requester
		err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sdk.AccAddress(subReq.Requester), sdk.NewCoins(sdk.NewInt64Coin(manduTypes.TokenDenom, int64(subReq.AvailableAmount))))
		if err != nil {
			return nil, errorsmod.Wrap(err, "failed to send coins to module account")
		}
	}

	return &types.MsgCancelSubscriptionRequestResponse{}, nil
}

func validateMsgUpdateSubscriptionRequest(msg *types.MsgUpdateSubscriptionRequest) error {
	if err := validation.ValidateNonEmptyString(msg.SubscriptionRequestId); err != nil {
		return err
	}
	if err := validation.ValidateAddress(msg.Requester); err != nil {
		return err
	}
	if err := validation.ValidatePositiveAmount(msg.Amount); err != nil {
		return err
	}
	return nil
}

func (k msgServer) UpdateSubscriptionRequest(goCtx context.Context, msg *types.MsgUpdateSubscriptionRequest) (*types.MsgUpdateSubscriptionRequestResponse, error) {
	err := validateMsgUpdateSubscriptionRequest(msg)
	if err != nil {
		return nil, err
	}
	requester, err := sdk.AccAddressFromBech32(msg.Requester)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "invalid requester address")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	subReq, found := k.GetSubscriptionRequest(ctx, msg.SubscriptionRequestId)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "subReq with id "+msg.SubscriptionRequestId+" not found")
	}
	if msg.Requester != subReq.Requester {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "only the requester can update the subReq")
	}

	// Amount, StartBlock, and EndBlock are optional arguments an default to 0 if not provided
	if ctx.BlockHeight() < int64(subReq.StartBlock) {
		if msg.Amount != 0 {
			if msg.Amount < subReq.TotalAmount {
				amountToReturn := subReq.TotalAmount - msg.Amount
				err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, requester, sdk.NewCoins(sdk.NewInt64Coin(manduTypes.TokenDenom, int64(amountToReturn))))
				if err != nil {
					return nil, errorsmod.Wrap(err, "failed to send coins to module account")
				}
			} else if msg.Amount > subReq.TotalAmount {
				amountToDeposit := msg.Amount - subReq.TotalAmount
				sdkError := k.bankKeeper.SendCoinsFromAccountToModule(ctx, requester, types.ModuleName, sdk.NewCoins(sdk.NewInt64Coin(manduTypes.TokenDenom, int64(amountToDeposit))))
				if sdkError != nil {
					return nil, errorsmod.Wrap(sdkError, "failed to send coins to module account")
				}
			}
			subReq.TotalAmount = msg.Amount
			subReq.AvailableAmount = msg.Amount
		}
		if msg.StartBlock != 0 {
			if int64(msg.StartBlock) < ctx.BlockHeight() {
				return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "start block must be greater than current block height")
			}
			subReq.StartBlock = msg.StartBlock
		}
	} else {
		if msg.Amount != 0 {
			if msg.Amount < subReq.TotalAmount {
				return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "amount must be greater than initial amount")
			}
			amountToDeposit := msg.Amount - subReq.TotalAmount
			requester, err := sdk.AccAddressFromBech32(msg.Requester)
			if err != nil {
				return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "invalid requester address")
			}
			sdkError := k.bankKeeper.SendCoinsFromAccountToModule(ctx, requester, types.ModuleName, sdk.NewCoins(sdk.NewInt64Coin(manduTypes.TokenDenom, int64(amountToDeposit))))
			if sdkError != nil {
				return nil, errorsmod.Wrap(sdkError, "failed to send coins to module account")
			}
			subReq.TotalAmount = msg.Amount
			subReq.AvailableAmount += amountToDeposit
		}
		if msg.StartBlock != 0 {
			return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "cannot update start block after subReq has started")
		}
	}

	if msg.Duration != 0 {
		endEpoch := subReq.StartEpoch + msg.Duration
		currentEpoch := utils.BlockToEpoch(ctx.BlockHeight(), subReq.EpochSize)
		if endEpoch <= currentEpoch {
			return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "end block must be greater than current block height")
		}
		subReq.Duration = msg.Duration
		subReq.EndEpoch = endEpoch
	}

	if len(msg.Writers) != 0 {
		subReq.Writers = msg.Writers
	}

	k.SetSubscriptionRequest(ctx, subReq)

	return &types.MsgUpdateSubscriptionRequestResponse{}, nil
}

func validateMsgIncrementSubscriptionRequestAmount(msg *types.MsgIncrementSubscriptionRequestAmount) error {
	if err := validation.ValidatePositiveAmount(msg.Amount); err != nil {
		return err
	}
	if err := validation.ValidateNonEmptyString(msg.SubscriptionRequestId); err != nil {
		return err
	}
	if err := validation.ValidateAddress(msg.Requester); err != nil {
		return err
	}
	return nil
}

func (k msgServer) IncrementSubscriptionRequestAmount(goCtx context.Context, msg *types.MsgIncrementSubscriptionRequestAmount) (*types.MsgIncrementSubscriptionRequestAmountResponse, error) {
	if err := validateMsgIncrementSubscriptionRequestAmount(msg); err != nil {
		return nil, err
	}

	requester, err := sdk.AccAddressFromBech32(msg.Requester)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "invalid requester address")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	subReq, found := k.GetSubscriptionRequest(ctx, msg.SubscriptionRequestId)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "subReq with id "+msg.SubscriptionRequestId+" not found")
	}
	if msg.Requester != subReq.Requester {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "only the requester can increment the subReq amount")
	}

	if k.IsSubscriptionRequestUnavailable(subReq.Status) {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "cannot topup the expired subReq with id "+msg.SubscriptionRequestId)
	}

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, requester, types.ModuleName, sdk.NewCoins(sdk.NewInt64Coin(manduTypes.TokenDenom, int64(msg.Amount))))
	if err != nil {
		return nil, errorsmod.Wrap(err, "failed to send coins to module account")
	}
	subReq.TotalAmount += msg.Amount
	subReq.AvailableAmount += msg.Amount

	k.SetSubscriptionRequest(ctx, subReq)
	return &types.MsgIncrementSubscriptionRequestAmountResponse{}, nil
}

func validateMsgJoinSubscriptionRequest(msg *types.MsgJoinSubscriptionRequest) error {
	if err := validation.ValidateNonEmptyString(msg.SubscriptionRequestId); err != nil {
		return err
	}
	if err := validation.ValidateAddress(msg.Subscriber); err != nil {
		return err
	}
	return nil
}

func (k msgServer) JoinSubscriptionRequest(goCtx context.Context, msg *types.MsgJoinSubscriptionRequest) (*types.MsgJoinSubscriptionRequestResponse, error) {
	err := validateMsgJoinSubscriptionRequest(msg)
	if err != nil {
		return nil, err
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	subReq, found := k.GetSubscriptionRequest(ctx, msg.SubscriptionRequestId)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "subReq with id "+msg.SubscriptionRequestId+" not found")
	}

	if k.IsSubscriptionRequestUnavailable(subReq.Status) {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidVersion, "subReq with id "+msg.SubscriptionRequestId+"is not available to join")
	}

	if k.SubscriptionRequestHasSubscriber(ctx, subReq, msg.Subscriber) {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidVersion, "subscriber is already subscribed to the subReq with id "+msg.SubscriptionRequestId)
	}

	delegations, err := k.stakingKeeper.GetDelegatorDelegations(ctx, sdk.AccAddress(msg.Subscriber), math.MaxUint16)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "subscriber "+msg.Subscriber+" not found")
	}

	var totalStake int64 = 0
	for _, delegation := range delegations {
		totalStake += delegation.GetShares().TruncateInt64()
	}

	// need a formula to determine the necessary amount to join the subReq, currently always accepted
	if totalStake < k.CalculateMinimumStake(ctx, subReq) {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "insufficient stake to join subReq")
	}

	id := uuid.NewString()
	var subscriptionStartEpoch int64
	if subReq.StartBlock < ctx.BlockHeight() {
		subscriptionStartEpoch = utils.BlockToEpoch(ctx.BlockHeight(), subReq.EpochSize)
		subReq.Status = types.SubscriptionRequest_ACTIVE
	} else {
		subscriptionStartEpoch = subReq.StartEpoch
	}

	subscription := types.Subscription{
		Id:                    id,
		SubscriptionRequestId: msg.SubscriptionRequestId,
		Subscriber:            msg.Subscriber,
		StartEpoch:            subscriptionStartEpoch,
		EndEpoch:              subReq.EndEpoch,
	}
	k.SetSubscription(ctx, subscription)
	subReq.SubscriptionIds = append(subReq.SubscriptionIds, subscription.Id)

	k.SetSubscriptionRequest(ctx, subReq)

	return &types.MsgJoinSubscriptionRequestResponse{SubscriptionId: id}, nil
}

func validateMsgLeaveSubscriptionRequest(msg *types.MsgLeaveSubscriptionRequest) error {
	if err := validation.ValidateNonEmptyString(msg.SubscriptionRequestId); err != nil {
		return err
	}
	if err := validation.ValidateAddress(msg.Subscriber); err != nil {
		return err
	}
	return nil
}

func (k msgServer) LeaveSubscriptionRequest(goCtx context.Context, msg *types.MsgLeaveSubscriptionRequest) (*types.MsgLeaveSubscriptionRequestResponse, error) {
	err := validateMsgLeaveSubscriptionRequest(msg)
	if err != nil {
		return nil, err
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	subReq, found := k.GetSubscriptionRequest(ctx, msg.SubscriptionRequestId)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "subReq with id "+msg.SubscriptionRequestId+" not found")
	}

	isSubscribed := false
	for _, subscriptionId := range subReq.SubscriptionIds {
		subscription, found := k.GetSubscription(ctx, subscriptionId)
		if !found {
			return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "SHOULD NOT HAPPEN: subscription with id "+subscriptionId+" not found")
		}
		if subscription.Subscriber == msg.Subscriber {
			isSubscribed = true
			subscription.EndEpoch = utils.BlockToEpoch(ctx.BlockHeight(), subReq.EpochSize)
			k.SetSubscription(ctx, subscription)
		}
	}

	if !isSubscribed {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "subscriber must be subscribed to the subReq with id "+msg.SubscriptionRequestId+" to leave it")
	}

	if !k.IsSubscriptionRequestActive(ctx, subReq) {
		subReq.Status = types.SubscriptionRequest_INACTIVE
		k.SetSubscriptionRequest(ctx, subReq)
	}

	return &types.MsgLeaveSubscriptionRequestResponse{}, nil
}
