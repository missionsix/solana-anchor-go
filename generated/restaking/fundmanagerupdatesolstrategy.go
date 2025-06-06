// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package restaking

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// FundManagerUpdateSolStrategy is the `fund_manager_update_sol_strategy` instruction.
type FundManagerUpdateSolStrategyInstruction struct {
	SolDepositable                      *bool
	SolAccumulatedDepositCapacityAmount *uint64
	SolAccumulatedDepositAmount         *uint64 `bin:"optional"`
	SolWithdrawable                     *bool
	SolWithdrawalNormalReserveRateBps   *uint16
	SolWithdrawalNormalReserveMaxAmount *uint64

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

// NewFundManagerUpdateSolStrategyInstructionBuilder creates a new `FundManagerUpdateSolStrategyInstruction` instruction builder.
func NewFundManagerUpdateSolStrategyInstructionBuilder() *FundManagerUpdateSolStrategyInstruction {
	nd := &FundManagerUpdateSolStrategyInstruction{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 5),
	}
	nd.AccountMetaSlice[0] = ag_solanago.Meta(Addresses["5UpLTLA7Wjqp7qdfjuTtPcUw3aVtbqFA5Mgm34mxPNg2"]).SIGNER()
	return nd
}

// SetSolDepositable sets the "sol_depositable" parameter.
func (inst *FundManagerUpdateSolStrategyInstruction) SetSolDepositable(sol_depositable bool) *FundManagerUpdateSolStrategyInstruction {
	inst.SolDepositable = &sol_depositable
	return inst
}

// SetSolAccumulatedDepositCapacityAmount sets the "sol_accumulated_deposit_capacity_amount" parameter.
func (inst *FundManagerUpdateSolStrategyInstruction) SetSolAccumulatedDepositCapacityAmount(sol_accumulated_deposit_capacity_amount uint64) *FundManagerUpdateSolStrategyInstruction {
	inst.SolAccumulatedDepositCapacityAmount = &sol_accumulated_deposit_capacity_amount
	return inst
}

// SetSolAccumulatedDepositAmount sets the "sol_accumulated_deposit_amount" parameter.
func (inst *FundManagerUpdateSolStrategyInstruction) SetSolAccumulatedDepositAmount(sol_accumulated_deposit_amount uint64) *FundManagerUpdateSolStrategyInstruction {
	inst.SolAccumulatedDepositAmount = &sol_accumulated_deposit_amount
	return inst
}

// SetSolWithdrawable sets the "sol_withdrawable" parameter.
func (inst *FundManagerUpdateSolStrategyInstruction) SetSolWithdrawable(sol_withdrawable bool) *FundManagerUpdateSolStrategyInstruction {
	inst.SolWithdrawable = &sol_withdrawable
	return inst
}

// SetSolWithdrawalNormalReserveRateBps sets the "sol_withdrawal_normal_reserve_rate_bps" parameter.
func (inst *FundManagerUpdateSolStrategyInstruction) SetSolWithdrawalNormalReserveRateBps(sol_withdrawal_normal_reserve_rate_bps uint16) *FundManagerUpdateSolStrategyInstruction {
	inst.SolWithdrawalNormalReserveRateBps = &sol_withdrawal_normal_reserve_rate_bps
	return inst
}

// SetSolWithdrawalNormalReserveMaxAmount sets the "sol_withdrawal_normal_reserve_max_amount" parameter.
func (inst *FundManagerUpdateSolStrategyInstruction) SetSolWithdrawalNormalReserveMaxAmount(sol_withdrawal_normal_reserve_max_amount uint64) *FundManagerUpdateSolStrategyInstruction {
	inst.SolWithdrawalNormalReserveMaxAmount = &sol_withdrawal_normal_reserve_max_amount
	return inst
}

// SetFundManagerAccount sets the "fund_manager" account.
func (inst *FundManagerUpdateSolStrategyInstruction) SetFundManagerAccount(fundManager ag_solanago.PublicKey) *FundManagerUpdateSolStrategyInstruction {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(fundManager).SIGNER()
	return inst
}

// GetFundManagerAccount gets the "fund_manager" account.
func (inst *FundManagerUpdateSolStrategyInstruction) GetFundManagerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetReceiptTokenMintAccount sets the "receipt_token_mint" account.
func (inst *FundManagerUpdateSolStrategyInstruction) SetReceiptTokenMintAccount(receiptTokenMint ag_solanago.PublicKey) *FundManagerUpdateSolStrategyInstruction {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(receiptTokenMint)
	return inst
}

// GetReceiptTokenMintAccount gets the "receipt_token_mint" account.
func (inst *FundManagerUpdateSolStrategyInstruction) GetReceiptTokenMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetFundAccountAccount sets the "fund_account" account.
func (inst *FundManagerUpdateSolStrategyInstruction) SetFundAccountAccount(fundAccount ag_solanago.PublicKey) *FundManagerUpdateSolStrategyInstruction {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(fundAccount).WRITE()
	return inst
}

func (inst *FundManagerUpdateSolStrategyInstruction) findFindFundAccountAddress(receiptTokenMint ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
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
func (inst *FundManagerUpdateSolStrategyInstruction) FindFundAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindFundAccountAddress(receiptTokenMint, bumpSeed)
	return
}

func (inst *FundManagerUpdateSolStrategyInstruction) MustFindFundAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindFundAccountAddress(receiptTokenMint, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindFundAccountAddress finds FundAccount account address with given seeds.
func (inst *FundManagerUpdateSolStrategyInstruction) FindFundAccountAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindFundAccountAddress(receiptTokenMint, 0)
	return
}

func (inst *FundManagerUpdateSolStrategyInstruction) MustFindFundAccountAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindFundAccountAddress(receiptTokenMint, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetFundAccountAccount gets the "fund_account" account.
func (inst *FundManagerUpdateSolStrategyInstruction) GetFundAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetEventAuthorityAccount sets the "event_authority" account.
func (inst *FundManagerUpdateSolStrategyInstruction) SetEventAuthorityAccount(eventAuthority ag_solanago.PublicKey) *FundManagerUpdateSolStrategyInstruction {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(eventAuthority)
	return inst
}

func (inst *FundManagerUpdateSolStrategyInstruction) findFindEventAuthorityAddress(knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
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
func (inst *FundManagerUpdateSolStrategyInstruction) FindEventAuthorityAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindEventAuthorityAddress(bumpSeed)
	return
}

func (inst *FundManagerUpdateSolStrategyInstruction) MustFindEventAuthorityAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindEventAuthorityAddress(bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindEventAuthorityAddress finds EventAuthority account address with given seeds.
func (inst *FundManagerUpdateSolStrategyInstruction) FindEventAuthorityAddress() (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindEventAuthorityAddress(0)
	return
}

func (inst *FundManagerUpdateSolStrategyInstruction) MustFindEventAuthorityAddress() (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindEventAuthorityAddress(0)
	if err != nil {
		panic(err)
	}
	return
}

// GetEventAuthorityAccount gets the "event_authority" account.
func (inst *FundManagerUpdateSolStrategyInstruction) GetEventAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetProgramAccount sets the "program" account.
func (inst *FundManagerUpdateSolStrategyInstruction) SetProgramAccount(program ag_solanago.PublicKey) *FundManagerUpdateSolStrategyInstruction {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(program)
	return inst
}

// GetProgramAccount gets the "program" account.
func (inst *FundManagerUpdateSolStrategyInstruction) GetProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

func (inst FundManagerUpdateSolStrategyInstruction) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_FundManagerUpdateSolStrategy,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst FundManagerUpdateSolStrategyInstruction) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *FundManagerUpdateSolStrategyInstruction) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.SolDepositable == nil {
			return errors.New("SolDepositable parameter is not set")
		}
		if inst.SolAccumulatedDepositCapacityAmount == nil {
			return errors.New("SolAccumulatedDepositCapacityAmount parameter is not set")
		}
		if inst.SolWithdrawable == nil {
			return errors.New("SolWithdrawable parameter is not set")
		}
		if inst.SolWithdrawalNormalReserveRateBps == nil {
			return errors.New("SolWithdrawalNormalReserveRateBps parameter is not set")
		}
		if inst.SolWithdrawalNormalReserveMaxAmount == nil {
			return errors.New("SolWithdrawalNormalReserveMaxAmount parameter is not set")
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

func (inst *FundManagerUpdateSolStrategyInstruction) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("FundManagerUpdateSolStrategy")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=6]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("                          SolDepositable", *inst.SolDepositable))
						paramsBranch.Child(ag_format.Param("     SolAccumulatedDepositCapacityAmount", *inst.SolAccumulatedDepositCapacityAmount))
						paramsBranch.Child(ag_format.Param("             SolAccumulatedDepositAmount (OPT)", inst.SolAccumulatedDepositAmount))
						paramsBranch.Child(ag_format.Param("                         SolWithdrawable", *inst.SolWithdrawable))
						paramsBranch.Child(ag_format.Param("       SolWithdrawalNormalReserveRateBps", *inst.SolWithdrawalNormalReserveRateBps))
						paramsBranch.Child(ag_format.Param("     SolWithdrawalNormalReserveMaxAmount", *inst.SolWithdrawalNormalReserveMaxAmount))
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

func (obj FundManagerUpdateSolStrategyInstruction) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `SolDepositable` param:
	err = encoder.Encode(obj.SolDepositable)
	if err != nil {
		return err
	}
	// Serialize `SolAccumulatedDepositCapacityAmount` param:
	err = encoder.Encode(obj.SolAccumulatedDepositCapacityAmount)
	if err != nil {
		return err
	}
	// Serialize `SolAccumulatedDepositAmount` param (optional):
	{
		if obj.SolAccumulatedDepositAmount == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.SolAccumulatedDepositAmount)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `SolWithdrawable` param:
	err = encoder.Encode(obj.SolWithdrawable)
	if err != nil {
		return err
	}
	// Serialize `SolWithdrawalNormalReserveRateBps` param:
	err = encoder.Encode(obj.SolWithdrawalNormalReserveRateBps)
	if err != nil {
		return err
	}
	// Serialize `SolWithdrawalNormalReserveMaxAmount` param:
	err = encoder.Encode(obj.SolWithdrawalNormalReserveMaxAmount)
	if err != nil {
		return err
	}
	return nil
}
func (obj *FundManagerUpdateSolStrategyInstruction) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `SolDepositable`:
	err = decoder.Decode(&obj.SolDepositable)
	if err != nil {
		return err
	}
	// Deserialize `SolAccumulatedDepositCapacityAmount`:
	err = decoder.Decode(&obj.SolAccumulatedDepositCapacityAmount)
	if err != nil {
		return err
	}
	// Deserialize `SolAccumulatedDepositAmount` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.SolAccumulatedDepositAmount)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `SolWithdrawable`:
	err = decoder.Decode(&obj.SolWithdrawable)
	if err != nil {
		return err
	}
	// Deserialize `SolWithdrawalNormalReserveRateBps`:
	err = decoder.Decode(&obj.SolWithdrawalNormalReserveRateBps)
	if err != nil {
		return err
	}
	// Deserialize `SolWithdrawalNormalReserveMaxAmount`:
	err = decoder.Decode(&obj.SolWithdrawalNormalReserveMaxAmount)
	if err != nil {
		return err
	}
	return nil
}

// NewFundManagerUpdateSolStrategyInstruction declares a new FundManagerUpdateSolStrategy instruction with the provided parameters and accounts.
func NewFundManagerUpdateSolStrategyInstruction(
	// Parameters:
	sol_depositable bool,
	sol_accumulated_deposit_capacity_amount uint64,
	sol_accumulated_deposit_amount uint64,
	sol_withdrawable bool,
	sol_withdrawal_normal_reserve_rate_bps uint16,
	sol_withdrawal_normal_reserve_max_amount uint64,
	// Accounts:
	fundManager ag_solanago.PublicKey,
	receiptTokenMint ag_solanago.PublicKey,
	fundAccount ag_solanago.PublicKey,
	eventAuthority ag_solanago.PublicKey,
	program ag_solanago.PublicKey) *FundManagerUpdateSolStrategyInstruction {
	return NewFundManagerUpdateSolStrategyInstructionBuilder().
		SetSolDepositable(sol_depositable).
		SetSolAccumulatedDepositCapacityAmount(sol_accumulated_deposit_capacity_amount).
		SetSolAccumulatedDepositAmount(sol_accumulated_deposit_amount).
		SetSolWithdrawable(sol_withdrawable).
		SetSolWithdrawalNormalReserveRateBps(sol_withdrawal_normal_reserve_rate_bps).
		SetSolWithdrawalNormalReserveMaxAmount(sol_withdrawal_normal_reserve_max_amount).
		SetFundManagerAccount(fundManager).
		SetReceiptTokenMintAccount(receiptTokenMint).
		SetFundAccountAccount(fundAccount).
		SetEventAuthorityAccount(eventAuthority).
		SetProgramAccount(program)
}
