package subscription

import (
	"context"
	"encoding/json"
	"fmt"

	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/core/store"
	"cosmossdk.io/depinject"
	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	modulev1 "mandu/api/mandu/subscription/module"
	manduTypes "mandu/types"
	"mandu/utils"
	"mandu/x/subscription/keeper"
	"mandu/x/subscription/types"
)

var (
	_ module.AppModuleBasic      = (*AppModule)(nil)
	_ module.AppModuleSimulation = (*AppModule)(nil)
	_ module.HasGenesis          = (*AppModule)(nil)
	_ module.HasInvariants       = (*AppModule)(nil)
	_ module.HasConsensusVersion = (*AppModule)(nil)

	_ appmodule.AppModule       = (*AppModule)(nil)
	_ appmodule.HasBeginBlocker = (*AppModule)(nil)
	_ appmodule.HasEndBlocker   = (*AppModule)(nil)
)

// ----------------------------------------------------------------------------
// AppModuleBasic
// ----------------------------------------------------------------------------

// AppModuleBasic implements the AppModuleBasic interface that defines the
// independent methods a Cosmos SDK module needs to implement.
type AppModuleBasic struct {
	cdc codec.BinaryCodec
}

func NewAppModuleBasic(cdc codec.BinaryCodec) AppModuleBasic {
	return AppModuleBasic{cdc: cdc}
}

// Name returns the name of the module as a string.
func (AppModuleBasic) Name() string {
	return types.ModuleName
}

// RegisterLegacyAminoCodec registers the amino codec for the module, which is used
// to marshal and unmarshal structs to/from []byte in order to persist them in the module's KVStore.
func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {}

// RegisterInterfaces registers a module's interface types and their concrete implementations as proto.Message.
func (a AppModuleBasic) RegisterInterfaces(reg cdctypes.InterfaceRegistry) {
	types.RegisterInterfaces(reg)
}

// DefaultGenesis returns a default GenesisState for the module, marshalled to json.RawMessage.
// The default GenesisState need to be defined by the module developer and is primarily used for testing.
func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(types.DefaultGenesis())
}

// ValidateGenesis used to validate the GenesisState, given in its json.RawMessage form.
func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, config client.TxEncodingConfig, bz json.RawMessage) error {
	var genState types.GenesisState
	if err := cdc.UnmarshalJSON(bz, &genState); err != nil {
		return fmt.Errorf("failed to unmarshal %s genesis state: %w", types.ModuleName, err)
	}
	return genState.Validate()
}

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the module.
func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	if err := types.RegisterQueryHandlerClient(context.Background(), mux, types.NewQueryClient(clientCtx)); err != nil {
		panic(err)
	}
}

// ----------------------------------------------------------------------------
// AppModule
// ----------------------------------------------------------------------------

// AppModule implements the AppModule interface that defines the inter-dependent methods that modules need to implement
type AppModule struct {
	AppModuleBasic

	keeper        keeper.Keeper
	accountKeeper types.AccountKeeper
	bankKeeper    types.BankKeeper
	stakingKeeper types.StakingKeeper
}

func NewAppModule(
	cdc codec.Codec,
	keeper keeper.Keeper,
	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
	stakingKeeper types.StakingKeeper,
) AppModule {
	return AppModule{
		AppModuleBasic: NewAppModuleBasic(cdc),
		keeper:         keeper,
		accountKeeper:  accountKeeper,
		bankKeeper:     bankKeeper,
		stakingKeeper:  stakingKeeper,
	}
}

// RegisterServices registers a gRPC query service to respond to the module-specific gRPC queries
func (am AppModule) RegisterServices(cfg module.Configurator) {
	types.RegisterMsgServer(cfg.MsgServer(), keeper.NewMsgServerImpl(am.keeper))
	types.RegisterQueryServer(cfg.QueryServer(), am.keeper)
}

// RegisterInvariants registers the invariants of the module. If an invariant deviates from its predicted value, the InvariantRegistry triggers appropriate logic (most often the chain will be halted)
func (am AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {}

// InitGenesis performs the module's genesis initialization. It returns no validator updates.
func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, gs json.RawMessage) {
	var genState types.GenesisState
	// Initialize global index to index in genesis state
	cdc.MustUnmarshalJSON(gs, &genState)

	InitGenesis(ctx, am.keeper, genState)
}

// ExportGenesis returns the module's exported genesis state as raw JSON bytes.
func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	genState := ExportGenesis(ctx, am.keeper)
	return cdc.MustMarshalJSON(genState)
}

// ConsensusVersion is a sequence number for state-breaking change of the module.
// It should be incremented on each consensus-breaking change introduced by the module.
// To avoid wrong/empty versions, the initial version should be set to 1.
func (AppModule) ConsensusVersion() uint64 { return 1 }

// BeginBlock contains the logic that is automatically triggered at the beginning of each block.
// The begin block implementation is optional.
func (am AppModule) BeginBlock(_ context.Context) error {
	return nil
}

// EndBlock contains the logic that is automatically triggered at the end of each block.
// The end block implementation is optional.
func (am AppModule) EndBlock(goCtx context.Context) error {
	ctx := sdk.UnwrapSDKContext(goCtx)
	err := am.keeper.IterateSubscriptionRequests(ctx, func(subReq types.SubscriptionRequest) (bool, error) {
		// subReq status updates
		// return false to callback to continue iteration
		switch subReq.Status {
		case types.SubscriptionRequest_EXPIRED:
			return false, nil
		case types.SubscriptionRequest_CANCELLED:
			return false, nil
		case types.SubscriptionRequest_SCHEDULED:
			if utils.BlockToEpoch(ctx.BlockHeight(), subReq.EpochSize) < subReq.StartEpoch {
				return false, nil
			}

			if am.keeper.IsSubscriptionRequestActive(ctx, subReq) {
				subReq.Status = types.SubscriptionRequest_ACTIVE
				updatedSubscriptionRequest, err := am.PayActiveSubscribersPerBlock(ctx, subReq)
				if err != nil {
					return true, err
				}
				subReq = *updatedSubscriptionRequest
			} else {
				subReq.Status = types.SubscriptionRequest_INITIALIZED
			}
		case types.SubscriptionRequest_INITIALIZED:
			if utils.BlockToEpoch(ctx.BlockHeight(), subReq.EpochSize) > subReq.EndEpoch {
				subReq.Status = types.SubscriptionRequest_EXPIRED
				// return the remaining amount to the requester
				err := am.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sdk.AccAddress(subReq.Requester), sdk.NewCoins(sdk.NewInt64Coin(manduTypes.TokenDenom, int64(subReq.AvailableAmount))))
				if err != nil {
					return true, err
				}
			}
		case types.SubscriptionRequest_ACTIVE:
			if utils.BlockToEpoch(ctx.BlockHeight(), subReq.EpochSize) > subReq.EndEpoch {
				subReq.Status = types.SubscriptionRequest_EXPIRED
				// return the remaining amount to the requester
				err := am.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sdk.AccAddress(subReq.Requester), sdk.NewCoins(sdk.NewInt64Coin(manduTypes.TokenDenom, int64(subReq.AvailableAmount))))
				if err != nil {
					return true, err
				}
			} else {
				updatedSubscriptionRequest, err := am.PayActiveSubscribersPerBlock(ctx, subReq)
				if err != nil {
					return true, err
				}
				subReq = *updatedSubscriptionRequest
			}
		case types.SubscriptionRequest_INACTIVE:
			if utils.BlockToEpoch(ctx.BlockHeight(), subReq.EpochSize) > subReq.EndEpoch {
				subReq.Status = types.SubscriptionRequest_EXPIRED
				// return the remaining amount to the requester
				err := am.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sdk.AccAddress(subReq.Requester), sdk.NewCoins(sdk.NewInt64Coin(manduTypes.TokenDenom, int64(subReq.AvailableAmount))))
				if err != nil {
					return true, err
				}
			}
		default:
			return false, nil
		}

		am.keeper.SetSubscriptionRequest(ctx, subReq)
		return false, nil
	})
	return err
}

func (am AppModule) PayActiveSubscribersPerBlock(ctx sdk.Context, subReq types.SubscriptionRequest) (*types.SubscriptionRequest, error) {
	activeSubscriptions := am.keeper.GetAllActiveSubscriptions(ctx, subReq)
	blockReward := am.keeper.CalculateBlockReward(ctx, subReq)
	currentBlock := ctx.BlockHeight()
	// iterate through the progress to get the total while recording the progress of each subscriber
	subscriberProgress := make(map[string]int)
	totalProgress := 0
	for subscription, subscriber := range activeSubscriptions {
		progress, found := am.keeper.GetProgressSize(ctx, subscription, currentBlock)
		if !found {
			subscriberProgress[subscriber] = 0
		}
		subscriberProgress[subscriber] = progress
		totalProgress += progress
	}

	totalRewardSent := int64(0)
	for subscription, subscriber := range activeSubscriptions {
		// reward based on the progress size
		reward := int64(float64(blockReward) * float64(subscriberProgress[activeSubscriptions[subscription]]) / float64(totalProgress))
		err := am.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sdk.AccAddress(subscriber), sdk.NewCoins(sdk.NewInt64Coin(manduTypes.TokenDenom, reward)))
		if err != nil {
			return nil, err
		}
		totalRewardSent += reward
	}

	subReq.AvailableAmount -= totalRewardSent
	return &subReq, nil
}

// IsOnePerModuleType implements the depinject.OnePerModuleType interface.
func (am AppModule) IsOnePerModuleType() {}

// IsAppModule implements the appmodule.AppModule interface.
func (am AppModule) IsAppModule() {}

// ----------------------------------------------------------------------------
// App Wiring Setup
// ----------------------------------------------------------------------------

func init() {
	appmodule.Register(
		&modulev1.Module{},
		appmodule.Provide(ProvideModule),
	)
}

type ModuleInputs struct {
	depinject.In

	StoreService store.KVStoreService
	Cdc          codec.Codec
	Config       *modulev1.Module
	Logger       log.Logger

	AccountKeeper types.AccountKeeper
	BankKeeper    types.BankKeeper
	StakingKeeper types.StakingKeeper
}

type ModuleOutputs struct {
	depinject.Out

	SubscriptionKeeper keeper.Keeper
	Module             appmodule.AppModule
}

func ProvideModule(in ModuleInputs) ModuleOutputs {
	// default to governance authority if not provided
	authority := authtypes.NewModuleAddress(govtypes.ModuleName)
	if in.Config.Authority != "" {
		authority = authtypes.NewModuleAddressOrBech32Address(in.Config.Authority)
	}

	moduleAddress := authtypes.NewModuleAddress(types.ModuleName)

	k := keeper.NewKeeper(
		in.Cdc,
		in.StoreService,
		in.Logger,
		authority.String(),
		moduleAddress.String(),
		in.AccountKeeper,
		in.BankKeeper,
		in.StakingKeeper,
	)
	m := NewAppModule(
		in.Cdc,
		k,
		in.AccountKeeper,
		in.BankKeeper,
		in.StakingKeeper,
	)

	return ModuleOutputs{SubscriptionKeeper: k, Module: m}
}
