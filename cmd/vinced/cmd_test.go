package main_test

import (
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/client/flags"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	"github.com/stretchr/testify/require"

	"github.com/AyrisDev/vinceChain/v12/app"
	vinced "github.com/AyrisDev/vinceChain/v12/cmd/vinced"
	"github.com/AyrisDev/vinceChain/v12/utils"
)

func TestInitCmd(t *testing.T) {
	rootCmd, _ := vinced.NewRootCmd()
	rootCmd.SetArgs([]string{
		"init",       // Test the init cmd
		"vince-test", // Moniker
		fmt.Sprintf("--%s=%s", cli.FlagOverwrite, "true"), // Overwrite genesis.json, in case it already exists
		fmt.Sprintf("--%s=%s", flags.FlagChainID, utils.TestnetChainID+"-1"),
	})

	err := svrcmd.Execute(rootCmd, "vinced", app.DefaultNodeHome)
	require.NoError(t, err)
}

func TestAddKeyLedgerCmd(t *testing.T) {
	rootCmd, _ := vinced.NewRootCmd()
	rootCmd.SetArgs([]string{
		"keys",
		"add",
		"dev0",
		fmt.Sprintf("--%s", flags.FlagUseLedger),
	})

	err := svrcmd.Execute(rootCmd, "vinced", app.DefaultNodeHome)
	require.Error(t, err)
}
