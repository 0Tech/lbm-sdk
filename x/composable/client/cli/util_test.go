package cli_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	sdk "github.com/line/lbm-sdk/types"
	sdkerrors "github.com/line/lbm-sdk/types/errors"
	"github.com/line/lbm-sdk/x/composable"
	"github.com/line/lbm-sdk/x/composable/client/cli"
)

func createAddresses(size int, prefix string) []sdk.AccAddress {
	addrs := make([]sdk.AccAddress, size)
	for i := range addrs {
		addrs[i] = sdk.AccAddress(fmt.Sprintf("%s%d", prefix, i))
	}

	return addrs
}

func createClassIDs(size int, prefix string) []string {
	owners := createAddresses(size, prefix)
	ids := make([]string, len(owners))
	for i, owner := range owners {
		ids[i] = composable.ClassIDFromOwner(owner)
	}

	return ids
}

func TestParseFullID(t *testing.T) {
	classID := createClassIDs(1, "class")[0]
	id := make([]rune, 78)
	for i := range id {
		id[i] = '0'
	}
	id[0] = '1'

	testCases := map[string]struct {
		classID   string
		delimiter string
		id        string
		err       error
	}{
		"valid id": {
			classID:   classID,
			delimiter: ":",
			id:        string(id),
		},
		"invalid format": {
			classID: classID,
			id:      string(id),
			err:     sdkerrors.ErrInvalidType,
		},
		"invalid uint": {
			classID:   classID,
			delimiter: ":",
			id:        string(id) + "0",
			err:       composable.ErrInvalidNFTID,
		},
		"invalid class id": {
			delimiter: ":",
			id:        string(id),
			err:       composable.ErrInvalidClassID,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			fullIDStr := fmt.Sprintf("%s%s%s", tc.classID, tc.delimiter, tc.id)

			fullID, err := cli.ParseFullID(fullIDStr)
			require.ErrorIs(t, err, tc.err)
			if tc.err != nil {
				return
			}

			require.Equal(t, tc.classID, fullID.ClassId)
			require.Equal(t, sdk.NewUintFromString(tc.id), fullID.Id)
		})
	}
}
