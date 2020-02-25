package types

import (
	"fmt"
	"regexp"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBuilderRegexp(t *testing.T) {
	cases := []struct {
		example string
		valid   bool
	}{
		{"fedora/httpd:version1.0", true},
		{"cosmwasm-opt:0.6.3", true},
		{"cosmwasm-opt-:0.6.3", false},
		{"confio/js-builder-1:test", true},
		{"confio/.builder-1:manual", false},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			ok, err := regexp.MatchString(BuildTagRegexp, tc.example)
			assert.NoError(t, err)
			assert.Equal(t, tc.valid, ok)
		})

	}
}

func TestStoreCodeValidation(t *testing.T) {
	badAddress, err := sdk.AccAddressFromHex("012345")
	require.NoError(t, err)
	// proper address size
	goodAddress := sdk.AccAddress(make([]byte, 20))

	cases := map[string]struct {
		msg   MsgStoreCode
		valid bool
	}{
		"empty": {
			msg:   MsgStoreCode{},
			valid: false,
		},
		"correct minimal": {
			msg: MsgStoreCode{
				Sender:       goodAddress,
				WASMByteCode: []byte("foo"),
			},
			valid: true,
		},
		"missing code": {
			msg: MsgStoreCode{
				Sender: goodAddress,
			},
			valid: false,
		},
		"bad sender minimal": {
			msg: MsgStoreCode{
				Sender:       badAddress,
				WASMByteCode: []byte("foo"),
			},
			valid: false,
		},
		"correct maximal": {
			msg: MsgStoreCode{
				Sender:       goodAddress,
				WASMByteCode: []byte("foo"),
				Builder:      "confio/cosmwasm-opt:0.6.2",
				Source:       "https://crates.io/api/v1/crates/cw-erc20/0.1.0/download",
			},
			valid: true,
		},
		"invalid builder": {
			msg: MsgStoreCode{
				Sender:       goodAddress,
				WASMByteCode: []byte("foo"),
				Builder:      "-bad-opt:0.6.2",
				Source:       "https://crates.io/api/v1/crates/cw-erc20/0.1.0/download",
			},
			valid: false,
		},
		"invalid source scheme": {
			msg: MsgStoreCode{
				Sender:       goodAddress,
				WASMByteCode: []byte("foo"),
				Builder:      "cosmwasm-opt:0.6.2",
				Source:       "ftp://crates.io/api/download.tar.gz",
			},
			valid: false,
		},
		"invalid source format": {
			msg: MsgStoreCode{
				Sender:       goodAddress,
				WASMByteCode: []byte("foo"),
				Builder:      "cosmwasm-opt:0.6.2",
				Source:       "/api/download-ss",
			},
			valid: false,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			err := tc.msg.ValidateBasic()
			if tc.valid {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}

}
