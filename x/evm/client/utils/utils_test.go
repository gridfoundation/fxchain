package utils

import (
	"os"
	"testing"

	gridchain "github.com/gridfx/fxchain/app/types"
	sdk "github.com/gridfx/fxchain/libs/cosmos-sdk/types"
	"github.com/gridfx/fxchain/x/evm/types"
	"github.com/stretchr/testify/require"
)

const (
	expectedManageContractDeploymentWhitelistProposalJSON = `{
  "title": "default title",
  "description": "default description",
  "distributor_addresses": [
    "did:fury:ex1cftp8q8g4aa65nw9s5trwexe77d9t6crp8cm05",
    "did:fury:ex1k0wwsg7xf9tjt3rvxdewz42e74sp286awhuzqx"
  ],
  "is_added": true,
  "deposit": [
    {
      "denom": "fury",
      "amount": "100.000000000000000000"
    }
  ]
}`
	expectedManageContractBlockedListProposalJSON = `{
  "title": "default title",
  "description": "default description",
  "contract_addresses": [
    "did:fury:ex1cftp8q8g4aa65nw9s5trwexe77d9t6crp8cm05",
    "did:fury:ex1k0wwsg7xf9tjt3rvxdewz42e74sp286awhuzqx"
  ],
  "is_added": true,
  "deposit": [
    {
      "denom": "fury",
      "amount": "100.000000000000000000"
    }
  ]
}`
	expectedManageContractMethodBlockedListProposalJSON = `{
  "title": "default title",
  "description": "default description",
  "contract_addresses":[
        {
            "address": "did:fury:ex1cftp8q8g4aa65nw9s5trwexe77d9t6crp8cm05",
            "block_methods": [
                {
                    "sign": "0x371303c0",
                    "extra": "inc()"
                },
                {
                    "sign": "0x579be378",
                    "extra": "onc()"
                }
            ]
        },
		{
            "address": "did:fury:ex1k0wwsg7xf9tjt3rvxdewz42e74sp286awhuzqx",
            "block_methods": [
                {
                    "sign": "0x371303c0",
                    "extra": "inc()"
                },
                {
                    "sign": "0x579be378",
                    "extra": "onc()"
                }
            ]
        }
  ],
  "is_added": true,
  "deposit": [
    {
      "denom": "fury",
      "amount": "100.000000000000000000"
    }
  ]
}`
	fileName                 = "./proposal.json"
	expectedTitle            = "default title"
	expectedDescription      = "default description"
	expectedDistributorAddr1 = "did:fury:ex1cftp8q8g4aa65nw9s5trwexe77d9t6crp8cm05"
	expectedDistributorAddr2 = "did:fury:ex1k0wwsg7xf9tjt3rvxdewz42e74sp286awhuzqx"
	expectedMethodSign1      = "0x371303c0"
	expectedMethodExtra1     = "inc()"
	expectedMethodSign2      = "0x579be378"
	expectedMethodExtra2     = "onc()"
)

func init() {
	config := sdk.GetConfig()
	gridchain.SetBech32Prefixes(config)
}

func TestParseManageContractDeploymentWhitelistProposalJSON(t *testing.T) {
	// create JSON file
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0666)
	require.NoError(t, err)
	_, err = f.WriteString(expectedManageContractDeploymentWhitelistProposalJSON)
	require.NoError(t, err)
	require.NoError(t, f.Close())

	// remove the temporary JSON file
	defer os.Remove(fileName)

	proposal, err := ParseManageContractDeploymentWhitelistProposalJSON(types.ModuleCdc, fileName)
	require.NoError(t, err)
	require.Equal(t, expectedTitle, proposal.Title)
	require.Equal(t, expectedDescription, proposal.Description)
	require.True(t, proposal.IsAdded)
	require.Equal(t, 1, len(proposal.Deposit))
	require.Equal(t, sdk.DefaultBondDenom, proposal.Deposit[0].Denom)
	require.True(t, sdk.NewDec(100).Equal(proposal.Deposit[0].Amount))
	require.Equal(t, 2, len(proposal.DistributorAddrs))
	require.Equal(t, expectedDistributorAddr1, proposal.DistributorAddrs[0].String())
	require.Equal(t, expectedDistributorAddr2, proposal.DistributorAddrs[1].String())
}

func TestParseManageContractBlockedListProposalJSON(t *testing.T) {
	// create JSON file
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0666)
	require.NoError(t, err)
	_, err = f.WriteString(expectedManageContractBlockedListProposalJSON)
	require.NoError(t, err)
	require.NoError(t, f.Close())

	// remove the temporary JSON file
	defer os.Remove(fileName)

	proposal, err := ParseManageContractBlockedListProposalJSON(types.ModuleCdc, fileName)
	require.NoError(t, err)
	require.Equal(t, expectedTitle, proposal.Title)
	require.Equal(t, expectedDescription, proposal.Description)
	require.True(t, proposal.IsAdded)
	require.Equal(t, 1, len(proposal.Deposit))
	require.Equal(t, sdk.DefaultBondDenom, proposal.Deposit[0].Denom)
	require.True(t, sdk.NewDec(100).Equal(proposal.Deposit[0].Amount))
	require.Equal(t, 2, len(proposal.ContractAddrs))
	require.Equal(t, expectedDistributorAddr1, proposal.ContractAddrs[0].String())
	require.Equal(t, expectedDistributorAddr2, proposal.ContractAddrs[1].String())
}
func TestParseManageContractMethodBlockedListProposalJSON(t *testing.T) {
	// create JSON file
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0666)
	require.NoError(t, err)
	_, err = f.WriteString(expectedManageContractMethodBlockedListProposalJSON)
	require.NoError(t, err)
	require.NoError(t, f.Close())

	// remove the temporary JSON file
	defer os.Remove(fileName)

	proposal, err := ParseManageContractMethodBlockedListProposalJSON(types.ModuleCdc, fileName)
	require.NoError(t, err)
	require.Equal(t, expectedTitle, proposal.Title)
	require.Equal(t, expectedDescription, proposal.Description)
	require.True(t, proposal.IsAdded)
	require.Equal(t, 1, len(proposal.Deposit))
	require.Equal(t, sdk.DefaultBondDenom, proposal.Deposit[0].Denom)
	require.True(t, sdk.NewDec(100).Equal(proposal.Deposit[0].Amount))
	require.Equal(t, 2, len(proposal.ContractList))

	methods := types.ContractMethods{
		types.ContractMethod{
			Sign:  expectedMethodSign1,
			Extra: expectedMethodExtra1,
		},
		types.ContractMethod{
			Sign:  expectedMethodSign2,
			Extra: expectedMethodExtra2,
		},
	}
	addr1, err := sdk.AccAddressFromBech32(expectedDistributorAddr1)
	require.NoError(t, err)
	addr2, err := sdk.AccAddressFromBech32(expectedDistributorAddr2)
	require.NoError(t, err)
	expectBc1 := types.NewBlockContract(addr1, methods)
	expectBc2 := types.NewBlockContract(addr2, methods)
	ok := types.BlockedContractListIsEqual(t, proposal.ContractList, types.BlockedContractList{*expectBc1, *expectBc2})
	require.True(t, ok)
}

func TestParseManageSysContractAddressProposalJSON(t *testing.T) {
	defaultSysContractAddressProposalJSON := `{
  "title":"default title",
  "description":"default description",
  "contract_address": "0xA4FFCda536CC8fF1eeFe32D32EE943b9B4e70414",
  "is_added":true,
  "deposit": [
    {
      "denom": "fury",
      "amount": "100.000000000000000000"
    }
  ]
}`
	// create JSON file
	filePathName := "./defaultSysContractAddressProposalJSON.json"
	f, err := os.OpenFile(filePathName, os.O_RDWR|os.O_CREATE, 0666)
	require.NoError(t, err)
	_, err = f.WriteString(defaultSysContractAddressProposalJSON)
	require.NoError(t, err)
	require.NoError(t, f.Close())

	// remove the temporary JSON file
	defer os.Remove(filePathName)

	proposal, err := ParseManageSysContractAddressProposalJSON(types.ModuleCdc, filePathName)
	require.NoError(t, err)
	require.Equal(t, expectedTitle, proposal.Title)
	require.Equal(t, expectedDescription, proposal.Description)
	require.True(t, proposal.IsAdded)
	require.Equal(t, 1, len(proposal.Deposit))
	require.Equal(t, sdk.DefaultBondDenom, proposal.Deposit[0].Denom)
	require.True(t, sdk.NewDec(100).Equal(proposal.Deposit[0].Amount))
}
