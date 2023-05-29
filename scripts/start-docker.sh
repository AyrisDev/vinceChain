#!/bin/bash

KEY="dev0"
CHAINID="vince_1903-1"
MONIKER="mymoniker"
DATA_DIR=$(mktemp -d -t evmos-datadir.XXXXX)

echo "create and add new keys"
./vinced keys add $KEY --home $DATA_DIR --no-backup --chain-id $CHAINID --algo "eth_secp256k1" --keyring-backend test
echo "init Evmos with moniker=$MONIKER and chain-id=$CHAINID"
./vinced init $MONIKER --chain-id $CHAINID --home $DATA_DIR
echo "prepare genesis: Allocate genesis accounts"
./vinced add-genesis-account \
"$(./vinced keys show $KEY -a --home $DATA_DIR --keyring-backend test)" 1000000000000000000avce,1000000000000000000stake \
--home $DATA_DIR --keyring-backend test
echo "prepare genesis: Sign genesis transaction"
./vinced gentx $KEY 1000000000000000000stake --keyring-backend test --home $DATA_DIR --keyring-backend test --chain-id $CHAINID
echo "prepare genesis: Collect genesis tx"
./vinced collect-gentxs --home $DATA_DIR
echo "prepare genesis: Run validate-genesis to ensure everything worked and that the genesis file is setup correctly"
./vinced validate-genesis --home $DATA_DIR

echo "starting evmos node $i in background ..."
./vinced start --pruning=nothing --rpc.unsafe \
--keyring-backend test --home $DATA_DIR \
>$DATA_DIR/node.log 2>&1 & disown

echo "started evmos node"
tail -f /dev/null