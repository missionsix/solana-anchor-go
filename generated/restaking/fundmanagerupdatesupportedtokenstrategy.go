// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package restaking

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// FundManagerUpdateSupportedTokenStrategy is the `fund_manager_update_supported_token_strategy` instruction.
type FundManagerUpdateSupportedTokenStrategyInstruction struct {
	TokenMint                             *ag_solanago.PublicKey
	TokenDepositable                      *bool
	TokenAccumulatedDepositCapacityAmount *uint64
	TokenAccumulatedDepositAmount         *uint64 `bin:"optional"`
	TokenWithdrawable                     *bool
	TokenWithdrawalNormalReserveRateBps   *uint16
	TokenWithdrawalNormalReserveMaxAmount *uint64
	TokenRebalancingAmount                *uint64 `bin:"optional"`
	SolAllocationWeight                   *uint64
	SolAllocationCapacityAmount           *uint64

	// [0] = [SIGNER] fund_manager
	//
	// [1] = [] receipt_token_mint
	//
	// [2] = [WRITE] fund_account
	//
	// [3] = [] event_authority
	//
	// [4] = [] program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewFundManagerUpdateSupportedTokenStrategyInstructionBuilder creates a new `FundManagerUpdateSupportedTokenStrategyInstruction` instruction builder.
func NewFundManagerUpdateSupportedTokenStrategyInstructionBuilder() *FundManagerUpdateSupportedTokenStrategyInstruction {
	nd := &FundManagerUpdateSupportedTokenStrategyInstruction{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 5),
	}
	nd.AccountMetaSlice[0] = ag_solanago.Meta(Addresses["5UpLTLA7Wjqp7qdfjuTtPcUw3aVtbqFA5Mgm34mxPNg2"]).SIGNER()
	return nd
}

// SetTokenMint sets the "token_mint" parameter.
func (inst *FundManagerUpdateSupportedTokenStrategyInstruction) SetTokenMint(token_mint ag_solanago.PublicKey) *FundManagerUpdateSupportedTokenStrategyInstruction {
	inst.TokenMint = &token_mint
	return inst
}

// SetTokenDepositable sets the "token_depositable" parameter.
func (inst *FundManagerUpdateSupportedTokenStrategyInstruction) SetTokenDepositable(token_depositable bool) *FundManagerUpdateSupportedTokenStrategyInstruction {
	inst.TokenDepositable = &token_depositable
	return inst
}

// SetTokenAccumulatedDepositCapacityAmount sets the "token_accumulated_deposit_capacity_amount" parameter.
func (inst *FundManagerUpdateSupportedTokenStrategyInstruction) SetTokenAccumulatedDepositCapacityAmount(token_accumulated_deposit_capacity_amount uint64) *FundManagerUpdateSupportedTokenStrategyInstruction {
	inst.TokenAccumulatedDepositCapacityAmount = &token_accumulated_deposit_capacity_amount
	return inst
}

// SetTokenAccumulatedDepositAmount sets the "token_accumulated_deposit_amount" parameter.
func (inst *FundManagerUpdateSupportedTokenStrategyInstruction) SetTokenAccumulatedDepositAmount(token_accumulated_deposit_amount uint64) *FundManagerUpdateSupportedTokenStrategyInstruction {
	inst.TokenAccumulatedDepositAmount = &token_accumulated_deposit_amount
	return inst
}

// SetTokenWithdrawable sets the "token_withdrawable" parameter.
func (inst *FundManagerUpdateSupportedTokenStrategyInstruction) SetTokenWithdrawable(token_withdrawable bool) *FundManagerUpdateSupportedTokenStrategyInstruction {
	inst.TokenWithdrawable = &token_withdrawable
	return inst
}

// SetTokenWithdrawalNormalReserveRateBps sets the "token_withdrawal_normal_reserve_rate_bps" parameter.
func (inst *FundManagerUpdateSupportedTokenStrategyInstruction) SetTokenWithdrawalNormalReserveRateBps(token_withdrawal_normal_reserve_rate_bps uint16) *FundManagerUpdateSupportedTokenStrategyInstruction {
	inst.TokenWithdrawalNormalReserveRateBps = &token_withdrawal_normal_reserve_rate_bps
	return inst
}

// SetTokenWithdrawalNormalReserveMaxAmount sets the "token_withdrawal_normal_reserve_max_amount" parameter.
func (inst *FundManagerUpdateSupportedTokenStrategyInstruction) SetTokenWithdrawalNormalReserveMaxAmount(token_withdrawal_normal_reserve_max_amount uint64) *FundManagerUpdateSupportedTokenStrategyInstruction {
	inst.TokenWithdrawalNormalReserveMaxAmount = &token_withdrawal_normal_reserve_max_amount
	return inst
}

// SetTokenRebalancingAmount sets the "token_rebalancing_amount" parameter.
func (inst *FundManagerUpdateSupportedTokenStrategyInstruction) SetTokenRebalancingAmount(token_rebalancing_amount uint64) *FundManagerUpdateSupportedTokenStrategyInstruction {
	inst.TokenRebalancingAmount = &token_rebalancing_amount
	return inst
}

// SetSolAllocationWeight sets the "sol_allocation_weight" parameter.
func (inst *FundManagerUpdateSupportedTokenStrategyInstruction) SetSolAllocationWeight(sol_allocation_weight uint64) *FundManagerUpdateSupportedTokenStrategyInstruction {
	inst.SolAllocationWeight = &sol_allocation_weight
	return inst
}

// SetSolAllocationCapacityAmount sets the "sol_allocation_capacity_amount" parameter.
func (inst *FundManagerUpdateSupportedTokenStrategyInstruction) SetSolAllocationCapacityAmount(sol_allocation_capacity_amount uint64) *FundManagerUpdateSupportedTokenStrategyInstruction {
	inst.SolAllocationCapacityAmount = &sol_allocation_capacity_amount
	return inst
}

// SetFundManagerAccount sets the "fund_manager" account.
func (inst *FundManagerUpdateSupportedTokenStrategyInstruction) SetFundManagerAccount(fundManager ag_solanago.PublicKey) *FundManagerUpdateSupportedTokenStrategyInstruction {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(fundManager).SIGNER()
	return inst
}

// GetFundManagerAccount gets the "fund_manager" account.
func (inst *FundManagerUpdateSupportedTokenStrategyInstruction) GetFundManagerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetReceiptTokenMintAccount sets the "receipt_token_mint" account.
func (inst *FundManagerUpdateSupportedTokenStrategyInstruction) SetReceiptTokenMintAccount(receiptTokenMint ag_solanago.PublicKey) *FundManagerUpdateSupportedTokenStrategyInstruction {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(receiptTokenMint)
	return inst
}

// GetReceiptTokenMintAccount gets the "receipt_token_mint" account.
func (inst *FundManagerUpdateSupportedTokenStrategyInstruction) GetReceiptTokenMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetFundAccountAccount sets the "fund_account" account.
func (inst *FundManagerUpdateSupportedTokenStrategyInstruction) SetFundAccountAccount(fundAccount ag_solanago.PublicKey) *FundManagerUpdateSupportedTokenStrategyInstruction {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(fundAccount).WRITE()
	return inst
}

func (inst *FundManagerUpdateSupportedTokenStrategyInstruction) findFindFundAccountAddress(receiptTokenMint ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	var seeds [][]byte
	// const: fund
	seeds = append(seeds, []byte{byte(0x66), byte(0x75), byte(0x6e), byte(0x64)})
	// path: receiptTokenMint
	seeds = append(seeds, receiptTokenMint.Bytes())

	if knownBumpSeed != 0 {
		seeds = append(seeds, []byte{byte(bumpSeed)})
		pda, err = ag_solanago.CreateProgramAddress(seeds, ProgramID)
	} else {
		pda, bumpSeed, err = ag_solanago.FindProgramAddress(seeds, ProgramID)
	}
	return
}

// FindFundAccountAddressWithBumpSeed calculates FundAccount account address with given seeds and a known bump seed.
func (inst *FundManagerUpdateSupportedTokenStrategyInstruction) FindFundAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindFundAccountAddress(receiptTokenMint, bumpSeed)
	return
}

func (inst *FundManagerUpdateSupportedTokenStrategyInstruction) MustFindFundAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindFundAccountAddress(receiptTokenMint, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindFundAccountAddress finds FundAccount account address with given seeds.
func (inst *FundManagerUpdateSupportedTokenStrategyInstruction) FindFundAccountAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindFundAccountAddress(receiptTokenMint, 0)
	return
}

func (inst *FundManagerUpdateSupportedTokenStrategyInstruction) MustFindFundAccountAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindFundAccountAddress(receiptTokenMint, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetFundAccountAccount gets the "fund_account" account.
func (inst *FundManagerUpdateSupportedTokenStrategyInstruction) GetFundAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetEventAuthorityAccount sets the "event_authority" account.
func (inst *FundManagerUpdateSupportedTokenStrategyInstruction) SetEventAuthorityAccount(eventAuthority ag_solanago.PublicKey) *FundManagerUpdateSupportedTokenStrategyInstruction {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(eventAuthority)
	return inst
}

func (inst *FundManagerUpdateSupportedTokenStrategyInstruction) findFindEventAuthorityAddress(knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	var seeds [][]byte
	// const: __event_authority
	seeds = append(seeds, []byte{byte(0x5f), byte(0x5f), byte(0x65), byte(0x76), byte(0x65), byte(0x6e), byte(0x74), byte(0x5f), byte(0x61), byte(0x75), byte(0x74), byte(0x68), byte(0x6f), byte(0x72), byte(0x69), byte(0x74), byte(0x79)})

	if knownBumpSeed != 0 {
		seeds = append(seeds, []byte{byte(bumpSeed)})
		pda, err = ag_solanago.CreateProgramAddress(seeds, ProgramID)
	} else {
		pda, bumpSeed, err = ag_solanago.FindProgramAddress(seeds, ProgramID)
	}
	return
}

// FindEventAuthorityAddressWithBumpSeed calculates EventAuthority account address with given seeds and a known bump seed.
func (inst *FundManagerUpdateSupportedTokenStrategyInstruction) FindEventAuthorityAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindEventAuthorityAddress(bumpSeed)
	return
}

func (inst *FundManagerUpdateSupportedTokenStrategyInstruction) MustFindEventAuthorityAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindEventAuthorityAddress(bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindEventAuthorityAddress finds EventAuthority account address with given seeds.
func (inst *FundManagerUpdateSupportedTokenStrategyInstruction) FindEventAuthorityAddress() (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindEventAuthorityAddress(0)
	return
}

func (inst *FundManagerUpdateSupportedTokenStrategyInstruction) MustFindEventAuthorityAddress() (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindEventAuthorityAddress(0)
	if err != nil {
		panic(err)
	}
	return
}

// GetEventAuthorityAccount gets the "event_authority" account.
func (inst *FundManagerUpdateSupportedTokenStrategyInstruction) GetEventAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetProgramAccount sets the "program" account.
func (inst *FundManagerUpdateSupportedTokenStrategyInstruction) SetProgramAccount(program ag_solanago.PublicKey) *FundManagerUpdateSupportedTokenStrategyInstruction {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(program)
	return inst
}

// GetProgramAccount gets the "program" account.
func (inst *FundManagerUpdateSupportedTokenStrategyInstruction) GetProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

func (inst FundManagerUpdateSupportedTokenStrategyInstruction) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_FundManagerUpdateSupportedTokenStrategy,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst FundManagerUpdateSupportedTokenStrategyInstruction) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *FundManagerUpdateSupportedTokenStrategyInstruction) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.TokenMint == nil {
			return errors.New("TokenMint parameter is not set")
		}
		if inst.TokenDepositable == nil {
			return errors.New("TokenDepositable parameter is not set")
		}
		if inst.TokenAccumulatedDepositCapacityAmount == nil {
			return errors.New("TokenAccumulatedDepositCapacityAmount parameter is not set")
		}
		if inst.TokenWithdrawable == nil {
			return errors.New("TokenWithdrawable parameter is not set")
		}
		if inst.TokenWithdrawalNormalReserveRateBps == nil {
			return errors.New("TokenWithdrawalNormalReserveRateBps parameter is not set")
		}
		if inst.TokenWithdrawalNormalReserveMaxAmount == nil {
			return errors.New("TokenWithdrawalNormalReserveMaxAmount parameter is not set")
		}
		if inst.SolAllocationWeight == nil {
			return errors.New("SolAllocationWeight parameter is not set")
		}
		if inst.SolAllocationCapacityAmount == nil {
			return errors.New("SolAllocationCapacityAmount parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.FundManager is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.ReceiptTokenMint is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.FundAccount is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.EventAuthority is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.Program is not set")
		}
	}
	return nil
}

func (inst *FundManagerUpdateSupportedTokenStrategyInstruction) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("FundManagerUpdateSupportedTokenStrategy")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=10]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("                                 TokenMint", *inst.TokenMint))
						paramsBranch.Child(ag_format.Param("                          TokenDepositable", *inst.TokenDepositable))
						paramsBranch.Child(ag_format.Param("     TokenAccumulatedDepositCapacityAmount", *inst.TokenAccumulatedDepositCapacityAmount))
						paramsBranch.Child(ag_format.Param("             TokenAccumulatedDepositAmount (OPT)", inst.TokenAccumulatedDepositAmount))
						paramsBranch.Child(ag_format.Param("                         TokenWithdrawable", *inst.TokenWithdrawable))
						paramsBranch.Child(ag_format.Param("       TokenWithdrawalNormalReserveRateBps", *inst.TokenWithdrawalNormalReserveRateBps))
						paramsBranch.Child(ag_format.Param("     TokenWithdrawalNormalReserveMaxAmount", *inst.TokenWithdrawalNormalReserveMaxAmount))
						paramsBranch.Child(ag_format.Param("                    TokenRebalancingAmount (OPT)", inst.TokenRebalancingAmount))
						paramsBranch.Child(ag_format.Param("                       SolAllocationWeight", *inst.SolAllocationWeight))
						paramsBranch.Child(ag_format.Param("               SolAllocationCapacityAmount", *inst.SolAllocationCapacityAmount))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=5]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("      fund_manager", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("receipt_token_mint", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("             fund_", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("   event_authority", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("           program", inst.AccountMetaSlice.Get(4)))
					})
				})
		})
}

func (obj FundManagerUpdateSupportedTokenStrategyInstruction) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `TokenMint` param:
	err = encoder.Encode(obj.TokenMint)
	if err != nil {
		return err
	}
	// Serialize `TokenDepositable` param:
	err = encoder.Encode(obj.TokenDepositable)
	if err != nil {
		return err
	}
	// Serialize `TokenAccumulatedDepositCapacityAmount` param:
	err = encoder.Encode(obj.TokenAccumulatedDepositCapacityAmount)
	if err != nil {
		return err
	}
	// Serialize `TokenAccumulatedDepositAmount` param (optional):
	{
		if obj.TokenAccumulatedDepositAmount == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.TokenAccumulatedDepositAmount)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `TokenWithdrawable` param:
	err = encoder.Encode(obj.TokenWithdrawable)
	if err != nil {
		return err
	}
	// Serialize `TokenWithdrawalNormalReserveRateBps` param:
	err = encoder.Encode(obj.TokenWithdrawalNormalReserveRateBps)
	if err != nil {
		return err
	}
	// Serialize `TokenWithdrawalNormalReserveMaxAmount` param:
	err = encoder.Encode(obj.TokenWithdrawalNormalReserveMaxAmount)
	if err != nil {
		return err
	}
	// Serialize `TokenRebalancingAmount` param (optional):
	{
		if obj.TokenRebalancingAmount == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.TokenRebalancingAmount)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `SolAllocationWeight` param:
	err = encoder.Encode(obj.SolAllocationWeight)
	if err != nil {
		return err
	}
	// Serialize `SolAllocationCapacityAmount` param:
	err = encoder.Encode(obj.SolAllocationCapacityAmount)
	if err != nil {
		return err
	}
	return nil
}
func (obj *FundManagerUpdateSupportedTokenStrategyInstruction) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `TokenMint`:
	err = decoder.Decode(&obj.TokenMint)
	if err != nil {
		return err
	}
	// Deserialize `TokenDepositable`:
	err = decoder.Decode(&obj.TokenDepositable)
	if err != nil {
		return err
	}
	// Deserialize `TokenAccumulatedDepositCapacityAmount`:
	err = decoder.Decode(&obj.TokenAccumulatedDepositCapacityAmount)
	if err != nil {
		return err
	}
	// Deserialize `TokenAccumulatedDepositAmount` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.TokenAccumulatedDepositAmount)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `TokenWithdrawable`:
	err = decoder.Decode(&obj.TokenWithdrawable)
	if err != nil {
		return err
	}
	// Deserialize `TokenWithdrawalNormalReserveRateBps`:
	err = decoder.Decode(&obj.TokenWithdrawalNormalReserveRateBps)
	if err != nil {
		return err
	}
	// Deserialize `TokenWithdrawalNormalReserveMaxAmount`:
	err = decoder.Decode(&obj.TokenWithdrawalNormalReserveMaxAmount)
	if err != nil {
		return err
	}
	// Deserialize `TokenRebalancingAmount` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.TokenRebalancingAmount)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `SolAllocationWeight`:
	err = decoder.Decode(&obj.SolAllocationWeight)
	if err != nil {
		return err
	}
	// Deserialize `SolAllocationCapacityAmount`:
	err = decoder.Decode(&obj.SolAllocationCapacityAmount)
	if err != nil {
		return err
	}
	return nil
}

// NewFundManagerUpdateSupportedTokenStrategyInstruction declares a new FundManagerUpdateSupportedTokenStrategy instruction with the provided parameters and accounts.
func NewFundManagerUpdateSupportedTokenStrategyInstruction(
	// Parameters:
	token_mint ag_solanago.PublicKey,
	token_depositable bool,
	token_accumulated_deposit_capacity_amount uint64,
	token_accumulated_deposit_amount uint64,
	token_withdrawable bool,
	token_withdrawal_normal_reserve_rate_bps uint16,
	token_withdrawal_normal_reserve_max_amount uint64,
	token_rebalancing_amount uint64,
	sol_allocation_weight uint64,
	sol_allocation_capacity_amount uint64,
	// Accounts:
	fundManager ag_solanago.PublicKey,
	receiptTokenMint ag_solanago.PublicKey,
	fundAccount ag_solanago.PublicKey,
	eventAuthority ag_solanago.PublicKey,
	program ag_solanago.PublicKey) *FundManagerUpdateSupportedTokenStrategyInstruction {
	return NewFundManagerUpdateSupportedTokenStrategyInstructionBuilder().
		SetTokenMint(token_mint).
		SetTokenDepositable(token_depositable).
		SetTokenAccumulatedDepositCapacityAmount(token_accumulated_deposit_capacity_amount).
		SetTokenAccumulatedDepositAmount(token_accumulated_deposit_amount).
		SetTokenWithdrawable(token_withdrawable).
		SetTokenWithdrawalNormalReserveRateBps(token_withdrawal_normal_reserve_rate_bps).
		SetTokenWithdrawalNormalReserveMaxAmount(token_withdrawal_normal_reserve_max_amount).
		SetTokenRebalancingAmount(token_rebalancing_amount).
		SetSolAllocationWeight(sol_allocation_weight).
		SetSolAllocationCapacityAmount(sol_allocation_capacity_amount).
		SetFundManagerAccount(fundManager).
		SetReceiptTokenMintAccount(receiptTokenMint).
		SetFundAccountAccount(fundAccount).
		SetEventAuthorityAccount(eventAuthority).
		SetProgramAccount(program)
}
