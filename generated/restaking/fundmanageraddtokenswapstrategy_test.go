// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package restaking

import (
	"bytes"
	ag_gofuzz "github.com/gagliardetto/gofuzz"
	ag_require "github.com/stretchr/testify/require"
	"reflect"
	"strconv"
	"testing"
)

func TestEncodeDecode_FundManagerAddTokenSwapStrategy(t *testing.T) {
	fu := ag_gofuzz.New().NilChance(0)
	for i := 0; i < 1; i++ {
		t.Run("FundManagerAddTokenSwapStrategy"+strconv.Itoa(i), func(t *testing.T) {
			{
				{
					{
						params := new(FundManagerAddTokenSwapStrategyInstruction)
						fu.Fuzz(params)
						params.AccountMetaSlice = nil
						tmp := new(TokenSwapSourceOrcaDEXLiquidityPoolTuple)
						fu.Fuzz(tmp)
						params.SetSwapSource(TokenSwapSource{*tmp})
						buf := new(bytes.Buffer)
						err := encodeT(*params, buf)
						ag_require.NoError(t, err)
						got := new(FundManagerAddTokenSwapStrategyInstruction)
						err = decodeT(got, buf.Bytes())
						got.AccountMetaSlice = nil
						ag_require.NoError(t, err)
						// to prevent garbage buffer fill by fuzz
						if reflect.TypeOf(*tmp).Kind() != reflect.Struct {
							got.SwapSource = params.SwapSource
						}
						ag_require.Equal(t, params, got)
					}
				}
			}
		})
	}
}
