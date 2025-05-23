// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package restaking

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// OperatorUpdateRewardPools is the `operator_update_reward_pools` instruction.
type OperatorUpdateRewardPoolsInstruction struct {

	// [0] = [WRITE, SIGNER] operator
	//
	// [1] = [] system_program
	//
	// [2] = [] receipt_token_mint
	//
	// [3] = [WRITE] reward_account
	//
	// [4] = [] event_authority
	//
	// [5] = [] program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewOperatorUpdateRewardPoolsInstructionBuilder creates a new `OperatorUpdateRewardPoolsInstruction` instruction builder.
func NewOperatorUpdateRewardPoolsInstructionBuilder() *OperatorUpdateRewardPoolsInstruction {
	nd := &OperatorUpdateRewardPoolsInstruction{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 6),
	}
	nd.AccountMetaSlice[1] = ag_solanago.Meta(Addresses["11111111111111111111111111111111"])
	return nd
}

// SetOperatorAccount sets the "operator" account.
func (inst *OperatorUpdateRewardPoolsInstruction) SetOperatorAccount(operator ag_solanago.PublicKey) *OperatorUpdateRewardPoolsInstruction {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(operator).WRITE().SIGNER()
	return inst
}

// GetOperatorAccount gets the "operator" account.
func (inst *OperatorUpdateRewardPoolsInstruction) GetOperatorAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetSystemProgramAccount sets the "system_program" account.
func (inst *OperatorUpdateRewardPoolsInstruction) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *OperatorUpdateRewardPoolsInstruction {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "system_program" account.
func (inst *OperatorUpdateRewardPoolsInstruction) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetReceiptTokenMintAccount sets the "receipt_token_mint" account.
func (inst *OperatorUpdateRewardPoolsInstruction) SetReceiptTokenMintAccount(receiptTokenMint ag_solanago.PublicKey) *OperatorUpdateRewardPoolsInstruction {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(receiptTokenMint)
	return inst
}

// GetReceiptTokenMintAccount gets the "receipt_token_mint" account.
func (inst *OperatorUpdateRewardPoolsInstruction) GetReceiptTokenMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetRewardAccountAccount sets the "reward_account" account.
func (inst *OperatorUpdateRewardPoolsInstruction) SetRewardAccountAccount(rewardAccount ag_solanago.PublicKey) *OperatorUpdateRewardPoolsInstruction {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(rewardAccount).WRITE()
	return inst
}

func (inst *OperatorUpdateRewardPoolsInstruction) findFindRewardAccountAddress(receiptTokenMint ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	var seeds [][]byte
	// const: reward
	seeds = append(seeds, []byte{byte(0x72), byte(0x65), byte(0x77), byte(0x61), byte(0x72), byte(0x64)})
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

// FindRewardAccountAddressWithBumpSeed calculates RewardAccount account address with given seeds and a known bump seed.
func (inst *OperatorUpdateRewardPoolsInstruction) FindRewardAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindRewardAccountAddress(receiptTokenMint, bumpSeed)
	return
}

func (inst *OperatorUpdateRewardPoolsInstruction) MustFindRewardAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindRewardAccountAddress(receiptTokenMint, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindRewardAccountAddress finds RewardAccount account address with given seeds.
func (inst *OperatorUpdateRewardPoolsInstruction) FindRewardAccountAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindRewardAccountAddress(receiptTokenMint, 0)
	return
}

func (inst *OperatorUpdateRewardPoolsInstruction) MustFindRewardAccountAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindRewardAccountAddress(receiptTokenMint, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetRewardAccountAccount gets the "reward_account" account.
func (inst *OperatorUpdateRewardPoolsInstruction) GetRewardAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetEventAuthorityAccount sets the "event_authority" account.
func (inst *OperatorUpdateRewardPoolsInstruction) SetEventAuthorityAccount(eventAuthority ag_solanago.PublicKey) *OperatorUpdateRewardPoolsInstruction {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(eventAuthority)
	return inst
}

func (inst *OperatorUpdateRewardPoolsInstruction) findFindEventAuthorityAddress(knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
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
func (inst *OperatorUpdateRewardPoolsInstruction) FindEventAuthorityAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindEventAuthorityAddress(bumpSeed)
	return
}

func (inst *OperatorUpdateRewardPoolsInstruction) MustFindEventAuthorityAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindEventAuthorityAddress(bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindEventAuthorityAddress finds EventAuthority account address with given seeds.
func (inst *OperatorUpdateRewardPoolsInstruction) FindEventAuthorityAddress() (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindEventAuthorityAddress(0)
	return
}

func (inst *OperatorUpdateRewardPoolsInstruction) MustFindEventAuthorityAddress() (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindEventAuthorityAddress(0)
	if err != nil {
		panic(err)
	}
	return
}

// GetEventAuthorityAccount gets the "event_authority" account.
func (inst *OperatorUpdateRewardPoolsInstruction) GetEventAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetProgramAccount sets the "program" account.
func (inst *OperatorUpdateRewardPoolsInstruction) SetProgramAccount(program ag_solanago.PublicKey) *OperatorUpdateRewardPoolsInstruction {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(program)
	return inst
}

// GetProgramAccount gets the "program" account.
func (inst *OperatorUpdateRewardPoolsInstruction) GetProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

func (inst OperatorUpdateRewardPoolsInstruction) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_OperatorUpdateRewardPools,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst OperatorUpdateRewardPoolsInstruction) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *OperatorUpdateRewardPoolsInstruction) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Operator is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.ReceiptTokenMint is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.RewardAccount is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.EventAuthority is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.Program is not set")
		}
	}
	return nil
}

func (inst *OperatorUpdateRewardPoolsInstruction) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("OperatorUpdateRewardPools")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=6]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("          operator", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("    system_program", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("receipt_token_mint", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("           reward_", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("   event_authority", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("           program", inst.AccountMetaSlice.Get(5)))
					})
				})
		})
}

func (obj OperatorUpdateRewardPoolsInstruction) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *OperatorUpdateRewardPoolsInstruction) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewOperatorUpdateRewardPoolsInstruction declares a new OperatorUpdateRewardPools instruction with the provided parameters and accounts.
func NewOperatorUpdateRewardPoolsInstruction(
	// Accounts:
	operator ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	receiptTokenMint ag_solanago.PublicKey,
	rewardAccount ag_solanago.PublicKey,
	eventAuthority ag_solanago.PublicKey,
	program ag_solanago.PublicKey) *OperatorUpdateRewardPoolsInstruction {
	return NewOperatorUpdateRewardPoolsInstructionBuilder().
		SetOperatorAccount(operator).
		SetSystemProgramAccount(systemProgram).
		SetReceiptTokenMintAccount(receiptTokenMint).
		SetRewardAccountAccount(rewardAccount).
		SetEventAuthorityAccount(eventAuthority).
		SetProgramAccount(program)
}
