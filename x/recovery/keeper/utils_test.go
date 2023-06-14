package keeper_test

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	transfertypes "github.com/cosmos/ibc-go/v6/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v6/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v6/modules/core/04-channel/types"
	ibcgotesting "github.com/cosmos/ibc-go/v6/testing"
	"github.com/AyrisDev/vinceChain/v12/app"
	ibctesting "github.com/AyrisDev/vinceChain/v12/ibc/testing"
	"github.com/AyrisDev/vinceChain/v12/utils"
	claimstypes "github.com/AyrisDev/vinceChain/v12/x/claims/types"
	inflationtypes "github.com/AyrisDev/vinceChain/v12/x/inflation/types"
	"github.com/AyrisDev/vinceChain/v12/x/recovery/types"
)

func CreatePacket(amount, denom, sender, receiver, srcPort, srcChannel, dstPort, dstChannel string, seq, timeout uint64) channeltypes.Packet {
	transfer := transfertypes.FungibleTokenPacketData{
		Amount:   amount,
		Denom:    denom,
		Receiver: sender,
		Sender:   receiver,
	}
	return channeltypes.NewPacket(
		transfer.GetBytes(),
		seq,
		srcPort,
		srcChannel,
		dstPort,
		dstChannel,
		clienttypes.ZeroHeight(), // timeout height disabled
		timeout,
	)
}

func (suite *IBCTestingSuite) SetupTest() {
	// initializes 3 test chains
	suite.coordinator = ibctesting.NewCoordinator(suite.T(), 1, 2)
	suite.vinceChain = suite.coordinator.GetChain(ibcgotesting.GetChainID(1))
	suite.IBCOsmosisChain = suite.coordinator.GetChain(ibcgotesting.GetChainID(2))
	suite.IBCCosmosChain = suite.coordinator.GetChain(ibcgotesting.GetChainID(3))
	suite.coordinator.CommitNBlocks(suite.vinceChain, 2)
	suite.coordinator.CommitNBlocks(suite.IBCOsmosisChain, 2)
	suite.coordinator.CommitNBlocks(suite.IBCCosmosChain, 2)

	// Mint coins locked on the vince account generated with secp.
	amt, ok := sdk.NewIntFromString("1000000000000000000000")
	suite.Require().True(ok)
	coinvince := sdk.NewCoin(utils.BaseDenom, amt)
	coins := sdk.NewCoins(coinvince)
	err := suite.vinceChain.App.(*app.Vince).BankKeeper.MintCoins(suite.vinceChain.GetContext(), inflationtypes.ModuleName, coins)
	suite.Require().NoError(err)

	// Fund sender address to pay fees
	err = suite.vinceChain.App.(*app.Vince).BankKeeper.SendCoinsFromModuleToAccount(suite.vinceChain.GetContext(), inflationtypes.ModuleName, suite.vinceChain.SenderAccount.GetAddress(), coins)
	suite.Require().NoError(err)

	coinvince = sdk.NewCoin(utils.BaseDenom, sdk.NewInt(10000))
	coins = sdk.NewCoins(coinvince)
	err = suite.vinceChain.App.(*app.Vince).BankKeeper.MintCoins(suite.vinceChain.GetContext(), inflationtypes.ModuleName, coins)
	suite.Require().NoError(err)
	err = suite.vinceChain.App.(*app.Vince).BankKeeper.SendCoinsFromModuleToAccount(suite.vinceChain.GetContext(), inflationtypes.ModuleName, suite.IBCOsmosisChain.SenderAccount.GetAddress(), coins)
	suite.Require().NoError(err)

	// Mint coins on the osmosis side which we'll use to unlock our avce
	coinOsmo := sdk.NewCoin("uosmo", sdk.NewInt(10))
	coins = sdk.NewCoins(coinOsmo)
	err = suite.IBCOsmosisChain.GetSimApp().BankKeeper.MintCoins(suite.IBCOsmosisChain.GetContext(), minttypes.ModuleName, coins)
	suite.Require().NoError(err)
	err = suite.IBCOsmosisChain.GetSimApp().BankKeeper.SendCoinsFromModuleToAccount(suite.IBCOsmosisChain.GetContext(), minttypes.ModuleName, suite.IBCOsmosisChain.SenderAccount.GetAddress(), coins)
	suite.Require().NoError(err)

	// Mint coins on the cosmos side which we'll use to unlock our avce
	coinAtom := sdk.NewCoin("uatom", sdk.NewInt(10))
	coins = sdk.NewCoins(coinAtom)
	err = suite.IBCCosmosChain.GetSimApp().BankKeeper.MintCoins(suite.IBCCosmosChain.GetContext(), minttypes.ModuleName, coins)
	suite.Require().NoError(err)
	err = suite.IBCCosmosChain.GetSimApp().BankKeeper.SendCoinsFromModuleToAccount(suite.IBCCosmosChain.GetContext(), minttypes.ModuleName, suite.IBCCosmosChain.SenderAccount.GetAddress(), coins)
	suite.Require().NoError(err)

	// Mint coins for IBC tx fee on Osmosis and Cosmos chains
	stkCoin := sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, amt))

	err = suite.IBCOsmosisChain.GetSimApp().BankKeeper.MintCoins(suite.IBCOsmosisChain.GetContext(), minttypes.ModuleName, stkCoin)
	suite.Require().NoError(err)
	err = suite.IBCOsmosisChain.GetSimApp().BankKeeper.SendCoinsFromModuleToAccount(suite.IBCOsmosisChain.GetContext(), minttypes.ModuleName, suite.IBCOsmosisChain.SenderAccount.GetAddress(), stkCoin)
	suite.Require().NoError(err)

	err = suite.IBCCosmosChain.GetSimApp().BankKeeper.MintCoins(suite.IBCCosmosChain.GetContext(), minttypes.ModuleName, stkCoin)
	suite.Require().NoError(err)
	err = suite.IBCCosmosChain.GetSimApp().BankKeeper.SendCoinsFromModuleToAccount(suite.IBCCosmosChain.GetContext(), minttypes.ModuleName, suite.IBCCosmosChain.SenderAccount.GetAddress(), stkCoin)
	suite.Require().NoError(err)

	claimparams := claimstypes.DefaultParams()
	claimparams.AirdropStartTime = suite.vinceChain.GetContext().BlockTime()
	claimparams.EnableClaims = true
	err = suite.vinceChain.App.(*app.Vince).ClaimsKeeper.SetParams(suite.vinceChain.GetContext(), claimparams)
	suite.Require().NoError(err)

	params := types.DefaultParams()
	params.EnableRecovery = true
	err = suite.vinceChain.App.(*app.Vince).RecoveryKeeper.SetParams(suite.vinceChain.GetContext(), params)
	suite.Require().NoError(err)

	evmParams := suite.vinceChain.App.(*app.Vince).EvmKeeper.GetParams(s.vinceChain.GetContext())
	evmParams.EvmDenom = utils.BaseDenom
	err = suite.vinceChain.App.(*app.Vince).EvmKeeper.SetParams(s.vinceChain.GetContext(), evmParams)
	suite.Require().NoError(err)

	suite.pathOsmosisvince = ibctesting.NewTransferPath(suite.IBCOsmosisChain, suite.vinceChain) // clientID, connectionID, channelID empty
	suite.pathCosmosvince = ibctesting.NewTransferPath(suite.IBCCosmosChain, suite.vinceChain)
	suite.pathOsmosisCosmos = ibctesting.NewTransferPath(suite.IBCCosmosChain, suite.IBCOsmosisChain)
	ibctesting.SetupPath(suite.coordinator, suite.pathOsmosisvince) // clientID, connectionID, channelID filled
	ibctesting.SetupPath(suite.coordinator, suite.pathCosmosvince)
	ibctesting.SetupPath(suite.coordinator, suite.pathOsmosisCosmos)
	suite.Require().Equal("07-tendermint-0", suite.pathOsmosisvince.EndpointA.ClientID)
	suite.Require().Equal("connection-0", suite.pathOsmosisvince.EndpointA.ConnectionID)
	suite.Require().Equal("channel-0", suite.pathOsmosisvince.EndpointA.ChannelID)
}

var timeoutHeight = clienttypes.NewHeight(1000, 1000)

func (suite *IBCTestingSuite) SendAndReceiveMessage(path *ibctesting.Path, origin *ibcgotesting.TestChain, coin string, amount int64, sender string, receiver string, seq uint64) {
	// Send coin from A to B
	transferMsg := transfertypes.NewMsgTransfer(path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, sdk.NewCoin(coin, sdk.NewInt(amount)), sender, receiver, timeoutHeight, 0, "")
	_, err := ibctesting.SendMsgs(origin, ibctesting.DefaultFeeAmt, transferMsg)
	suite.Require().NoError(err) // message committed
	// Recreate the packet that was sent
	transfer := transfertypes.NewFungibleTokenPacketData(coin, strconv.Itoa(int(amount)), sender, receiver, "")
	packet := channeltypes.NewPacket(transfer.GetBytes(), seq, path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, path.EndpointB.ChannelConfig.PortID, path.EndpointB.ChannelID, timeoutHeight, 0)
	// Receive message on the counterparty side, and send ack
	err = path.RelayPacket(packet)
	suite.Require().NoError(err)
}
