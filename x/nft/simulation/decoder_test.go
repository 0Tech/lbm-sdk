package simulation_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/line/lbm-sdk/crypto/keys/ed25519"
	"github.com/line/lbm-sdk/simapp"
	sdk "github.com/line/lbm-sdk/types"
	"github.com/line/lbm-sdk/types/kv"
	"github.com/line/lbm-sdk/x/nft"
	"github.com/line/lbm-sdk/x/nft/keeper"
	"github.com/line/lbm-sdk/x/nft/simulation"
)

var (
	ownerPk1   = ed25519.GenPrivKey().PubKey()
	ownerAddr1 = sdk.AccAddress(ownerPk1.Address())
)

func TestDecodeStore(t *testing.T) {
	cdc := simapp.MakeTestEncodingConfig().Marshaler
	dec := simulation.NewDecodeStore(cdc)

	class := nft.Class{
		Id:          "ClassID",
		Name:        "ClassName",
		Symbol:      "ClassSymbol",
		Description: "ClassDescription",
		Uri:         "ClassURI",
	}
	classBz, err := cdc.Marshal(&class)
	require.NoError(t, err)

	nft := nft.NFT{
		ClassId: "ClassID",
		Id:      "NFTID",
		Uri:     "NFTURI",
	}
	nftBz, err := cdc.Marshal(&nft)
	require.NoError(t, err)

	nftOfClassByOwnerValue := []byte{0x01}

	totalSupply := 1
	totalSupplyBz := sdk.Uint64ToBigEndian(1)

	kvPairs := kv.Pairs{
		Pairs: []kv.Pair{
			{Key: keeper.ClassKey, Value: classBz},
			{Key: keeper.NFTKey, Value: nftBz},
			{Key: keeper.NFTOfClassByOwnerKey, Value: nftOfClassByOwnerValue},
			{Key: keeper.OwnerKey, Value: ownerAddr1},
			{Key: keeper.ClassTotalSupply, Value: totalSupplyBz},
			{Key: []byte{0x99}, Value: []byte{0x99}},
		},
	}

	tests := []struct {
		name        string
		expectErr   bool
		expectedLog string
	}{
		{"Class", false, fmt.Sprintf("%v\n%v", class, class)},
		{"NFT", false, fmt.Sprintf("%v\n%v", nft, nft)},
		{"NFTOfClassByOwnerKey", false, fmt.Sprintf("%v\n%v", nftOfClassByOwnerValue, nftOfClassByOwnerValue)},
		{"OwnerKey", false, fmt.Sprintf("%v\n%v", ownerAddr1, ownerAddr1)},
		{"ClassTotalSupply", false, fmt.Sprintf("%v\n%v", totalSupply, totalSupply)},
		{"other", true, ""},
	}

	for i, tt := range tests {
		i, tt := i, tt
		t.Run(tt.name, func(t *testing.T) {
			if tt.expectErr {
				require.Panics(t, func() { dec(kvPairs.Pairs[i], kvPairs.Pairs[i]) }, tt.name)
			} else {
				require.Equal(t, tt.expectedLog, dec(kvPairs.Pairs[i], kvPairs.Pairs[i]), tt.name)
			}
		})
	}
}
