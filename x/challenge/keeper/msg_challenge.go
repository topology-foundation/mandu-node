package keeper

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/gob"
	"encoding/json"
	"fmt"

	manduTypes "mandu/types"
	"mandu/x/challenge/types"
	sTypes "mandu/x/subscription/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/google/uuid"
)

func (k msgServer) Challenge(goCtx context.Context, msg *types.MsgChallenge) (*types.MsgChallengeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	requester, err := sdk.AccAddressFromBech32(msg.Challenger)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "invalid challenger address")
	}

	currentBlock := ctx.BlockHeight()
	var hashes sTypes.Set[string]

	for _, hash := range msg.VerticesHashes {
		block, found := k.GetHashSubmissionBlock(ctx, msg.SubscriberId, hash)
		if !found {
			return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("hash %s not found", hash))
		} else if currentBlock-block > ChallengePeriod {
			return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("hash %s was submitted more than %d blocks ago", hash, ChallengePeriod))
		} else {
			hashes.Add(hash)
		}
	}

	totalChallengePrice := k.PricePerVertexChallenge(ctx, msg.Challenger, msg.SubscriberId) * int64(len(msg.VerticesHashes))
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, requester, types.ModuleName, sdk.NewCoins(sdk.NewInt64Coin(manduTypes.TokenDenom, totalChallengePrice)))
	if err != nil {
		return nil, errorsmod.Wrap(err, "failed to send coins to module account")
	}

	id := uuid.NewString()
	buf := &bytes.Buffer{}
	err = gob.NewEncoder(buf).Encode(hashes)
	if err != nil {
		return nil, errorsmod.Wrap(err, "failed to encode challenged hashes")
	}

	k.SetChallenge(ctx, types.Challenge{
		Id:               id,
		Challenger:       msg.Challenger,
		Subscriber:         msg.SubscriberId,
		Amount:           uint64(totalChallengePrice),
		LastActive:       uint64(currentBlock),
		ChallengedHashes: buf.Bytes(),
	})

	return &types.MsgChallengeResponse{ChallengeId: id}, nil
}

func (k msgServer) SubmitProof(goCtx context.Context, msg *types.MsgSubmitProof) (*types.MsgSubmitProofResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	challenge, found := k.GetChallenge(ctx, msg.ChallengeId)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("challenge %s not found", msg.ChallengeId))
	}
	if challenge.Subscriber != msg.Subscriber {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "unauthorized subscriber")
	}
	if k.isChallengeExpired(ctx, challenge) {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "challenge is expired")
	}

	buf := bytes.NewBuffer(challenge.ChallengedHashes)
	var challengedHashes sTypes.Set[string]
	err := gob.NewDecoder(buf).Decode(&challengedHashes)
	if err != nil {
		return nil, errorsmod.Wrap(err, "failed to decode challenged hashes")
	}

	for _, vertex := range msg.Vertices {
		if challengedHashes.Has(vertex.Hash) {
			vertexData := map[string]interface{}{
				"operation": vertex.Operation,
				"deps":      vertex.Dependencies,
				"nodeId":    vertex.NodeId,
			}
			stringified, err := json.Marshal(vertexData)
			if err != nil {
				return nil, errorsmod.Wrap(err, fmt.Sprintf("failed to marshal vertex with hash %s", vertex.Hash))
			}
			computedHash := sha256.Sum256(stringified)

			if !bytes.Equal(computedHash[:], []byte(vertex.Hash)) {
				return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("hash %s does not match the computed hash", vertex.Hash))
			}

			k.SetProof(ctx, msg.ChallengeId, *vertex)
			challengedHashes.Remove(vertex.Hash)
		}
	}

	challenge.LastActive = uint64(ctx.BlockHeight())
	buf = &bytes.Buffer{}
	err = gob.NewEncoder(buf).Encode(challengedHashes)
	if err != nil {
		return nil, errorsmod.Wrap(err, "failed to encode challenged hashes")
	}

	challenge.ChallengedHashes = buf.Bytes()

	k.SetChallenge(ctx, challenge)

	return &types.MsgSubmitProofResponse{}, nil
}

func (k msgServer) RequestDependencies(goCtx context.Context, msg *types.MsgRequestDependencies) (*types.MsgRequestDependenciesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	challenge, found := k.GetChallenge(ctx, msg.ChallengeId)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("challenge %s not found", msg.ChallengeId))
	}
	if challenge.Challenger != msg.Challenger {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "unauthorized challenger")
	}
	if k.isChallengeExpired(ctx, challenge) {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "challenge is expired")
	}

	requester, err := sdk.AccAddressFromBech32(msg.Challenger)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "invalid challenger address")
	}

	fee := k.PricePerVertexChallenge(ctx, msg.Challenger, challenge.Subscriber) * int64(len(msg.VerticesHashes))
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, requester, types.ModuleName, sdk.NewCoins(sdk.NewInt64Coin(manduTypes.TokenDenom, fee)))
	if err != nil {
		return nil, errorsmod.Wrap(err, "failed to send coins to module account")
	}

	buf := bytes.NewBuffer(challenge.ChallengedHashes)
	var challengedHashes sTypes.Set[string]
	err = gob.NewDecoder(buf).Decode(&challengedHashes)
	if err != nil {
		return nil, errorsmod.Wrap(err, "failed to decode challenged hashes")
	}

	currentBlock := ctx.BlockHeight()
	for _, hash := range msg.VerticesHashes {
		block, found := k.GetHashSubmissionBlock(ctx, challenge.Subscriber, hash)
		if !found {
			return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("hash %s not found", hash))
		}
		if currentBlock-block > ChallengePeriod {
			return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("hash %s was submitted more than %d blocks ago", hash, ChallengePeriod))
		} else {
			challengedHashes.Add(hash)
		}
	}

	challenge.LastActive = uint64(currentBlock)
	challenge.Amount += uint64(fee)

	buf = &bytes.Buffer{}
	err = gob.NewEncoder(buf).Encode(challengedHashes)
	if err != nil {
		return nil, errorsmod.Wrap(err, "failed to encode challenged hashes")
	}

	challenge.ChallengedHashes = buf.Bytes()

	k.SetChallenge(ctx, challenge)

	return &types.MsgRequestDependenciesResponse{}, nil
}

func (k msgServer) SettleChallenge(goCtx context.Context, msg *types.MsgSettleChallenge) (*types.MsgSettleChallengeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	challenge, found := k.GetChallenge(ctx, msg.ChallengeId)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("challenge %s not found", msg.ChallengeId))
	}
	if msg.Requester != challenge.Challenger && msg.Requester != challenge.Subscriber {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "not the challenger or the subscriber")
	}
	if k.isChallengeExpired(ctx, challenge) {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "challenge is not yet expired")
	}

	buf := bytes.NewBuffer(challenge.ChallengedHashes)
	var challengedHashes sTypes.Set[string]
	err := gob.NewDecoder(buf).Decode(&challengedHashes)
	if err != nil {
		return nil, errorsmod.Wrap(err, "failed to decode challenged hashes")
	}

	coins := sdk.NewCoins(sdk.NewInt64Coin(manduTypes.TokenDenom, int64(challenge.Amount)))
	if len(challengedHashes) == 0 {
		// all hashes were verified - send coins to subscriber, remove challenge
		err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sdk.AccAddress(challenge.Subscriber), coins)
		if err != nil {
			return nil, errorsmod.Wrap(err, "failed to send coins to subscriber")
		}
	} else {
		// some hashes were not verified - send coins to challenger
		err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sdk.AccAddress(challenge.Challenger), coins)
		if err != nil {
			return nil, errorsmod.Wrap(err, "failed to send coins to challenger")
		}
	}
	k.RemoveChallenge(ctx, challenge.Id)

	return &types.MsgSettleChallengeResponse{}, nil
}
