package voting

import (
	"github.com/arterynetwork/artr/x/voting/types"
	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"

	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/arterynetwork/artr/x/voting/client/cli"
	"github.com/arterynetwork/artr/x/voting/client/rest"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
)

// TypeCode check to ensure the interface is properly implemented
var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
)

// AppModuleBasic defines the basic application module used by the voting module.
type AppModuleBasic struct{}

// Name returns the voting module's name.
func (AppModuleBasic) Name() string {
	return ModuleName
}

// RegisterCodec registers the voting module's types for the given codec.
func (AppModuleBasic) RegisterCodec(cdc *codec.Codec) {
	RegisterCodec(cdc)
}

// DefaultGenesis returns default genesis state as raw bytes for the voting
// module.
func (AppModuleBasic) DefaultGenesis() json.RawMessage {
	return ModuleCdc.MustMarshalJSON(DefaultGenesisState())
}

// ValidateGenesis performs genesis state validation for the voting module.
func (AppModuleBasic) ValidateGenesis(bz json.RawMessage) error {
	var data GenesisState
	err := ModuleCdc.UnmarshalJSON(bz, &data)
	if err != nil {
		return err
	}
	return ValidateGenesis(data)
}

// RegisterRESTRoutes registers the REST routes for the voting module.
func (AppModuleBasic) RegisterRESTRoutes(ctx context.CLIContext, rtr *mux.Router) {
	rest.RegisterRoutes(ctx, rtr)
}

// GetTxCmd returns the root tx command for the voting module.
func (AppModuleBasic) GetTxCmd(cdc *codec.Codec) *cobra.Command {
	return cli.GetTxCmd(cdc)
}

// GetQueryCmd returns no root query command for the voting module.
func (AppModuleBasic) GetQueryCmd(cdc *codec.Codec) *cobra.Command {
	return cli.GetQueryCmd(StoreKey, cdc)
}

//____________________________________________________________________________

// AppModule implements an application module for the voting module.
type AppModule struct {
	AppModuleBasic

	keeper             Keeper
	scheduleKeeper     types.ScheduleKeeper
	upgradeKeeper      types.UprgadeKeeper
	nodingKeeper       types.NodingKeeper
	delegatingKeeper   types.DelegatingKeeper
	referralKeeper     types.ReferralKeeper
	subscriptionKeeper types.SubscriptionKeeper
	profileKeeper      types.ProfileKeeper
	earningKeeper      types.EarningKeeper
	vpnKeeper          types.VpnKeeper
}

// NewAppModule creates a new AppModule object
func NewAppModule(k Keeper,
	scheduleKeeper types.ScheduleKeeper,
	uprgadeKeeper types.UprgadeKeeper,
	nodingKeeper types.NodingKeeper,
	delegatingKeeper types.DelegatingKeeper,
	referralKeeper types.ReferralKeeper,
	subscriptionKeeper types.SubscriptionKeeper,
	profileKeeper types.ProfileKeeper,
	earningKeeper types.EarningKeeper,
	vpnKeeper types.VpnKeeper,
) AppModule {
	return AppModule{
		AppModuleBasic:     AppModuleBasic{},
		keeper:             k,
		scheduleKeeper:     scheduleKeeper,
		upgradeKeeper:      uprgadeKeeper,
		nodingKeeper:       nodingKeeper,
		delegatingKeeper:   delegatingKeeper,
		referralKeeper:     referralKeeper,
		subscriptionKeeper: subscriptionKeeper,
		profileKeeper:      profileKeeper,
		earningKeeper:      earningKeeper,
		vpnKeeper:          vpnKeeper,
	}
}

// Name returns the voting module's name.
func (AppModule) Name() string {
	return ModuleName
}

// RegisterInvariants registers the voting module invariants.
func (am AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {}

// Route returns the message routing key for the voting module.
func (AppModule) Route() string {
	return RouterKey
}

// NewHandler returns an sdk.Handler for the voting module.
func (am AppModule) NewHandler() sdk.Handler {
	return NewHandler(am.keeper)
}

// QuerierRoute returns the voting module's querier route name.
func (AppModule) QuerierRoute() string {
	return QuerierRoute
}

// NewQuerierHandler returns the voting module sdk.Querier.
func (am AppModule) NewQuerierHandler() sdk.Querier {
	return NewQuerier(am.keeper)
}

// InitGenesis performs genesis initialization for the voting module. It returns
// no validator updates.
func (am AppModule) InitGenesis(ctx sdk.Context, data json.RawMessage) []abci.ValidatorUpdate {
	var genesisState GenesisState
	ModuleCdc.MustUnmarshalJSON(data, &genesisState)
	InitGenesis(ctx, am.keeper, genesisState)
	return []abci.ValidatorUpdate{}
}

// ExportGenesis returns the exported genesis state as raw bytes for the voting
// module.
func (am AppModule) ExportGenesis(ctx sdk.Context) json.RawMessage {
	gs := ExportGenesis(ctx, am.keeper)
	return ModuleCdc.MustMarshalJSON(gs)
}

// BeginBlock returns the begin blocker for the voting module.
func (am AppModule) BeginBlock(ctx sdk.Context, req abci.RequestBeginBlock) {
}

// EndBlock returns the end blocker for the voting module. It returns no validator
// updates.
func (AppModule) EndBlock(_ sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}
