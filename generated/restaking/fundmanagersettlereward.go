// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package restaking

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// FundManagerSettleReward is the `fund_manager_settle_reward` instruction.
type FundManagerSettleRewardInstruction struct {
	RewardPoolId *uint8
	RewardId     *uint16
	Amount       *uint64

	// [0] = [SIGNER] fund_manager
	//
	// [1] = [] receipt_token_mint
	//
	// [2] = [WRITE] reward_account
	//
	// [3] = [] reward_token_mint
	//
	// [4] = [] reward_token_program
	//
	// [5] = [] event_authority
	//
	// [6] = [] program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewFundManagerSettleRewardInstructionBuilder creates a new `FundManagerSettleRewardInstruction` instruction builder.
func NewFundManagerSettleRewardInstructionBuilder() *FundManagerSettleRewardInstruction {
	nd := &FundManagerSettleRewardInstruction{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 7),
	}
	nd.AccountMetaSlice[0] = ag_solanago.Meta(Addresses["5UpLTLA7Wjqp7qdfjuTtPcUw3aVtbqFA5Mgm34mxPNg2"]).SIGNER()
	return nd
}

// SetRewardPoolId sets the "reward_pool_id" parameter.
func (inst *FundManagerSettleRewardInstruction) SetRewardPoolId(reward_pool_id uint8) *FundManagerSettleRewardInstruction {
	inst.RewardPoolId = &reward_pool_id
	return inst
}

// SetRewardId sets the "reward_id" parameter.
func (inst *FundManagerSettleRewardInstruction) SetRewardId(reward_id uint16) *FundManagerSettleRewardInstruction {
	inst.RewardId = &reward_id
	return inst
}

// SetAmount sets the "amount" parameter.
func (inst *FundManagerSettleRewardInstruction) SetAmount(amount uint64) *FundManagerSettleRewardInstruction {
	inst.Amount = &amount
	return inst
}

// SetFundManagerAccount sets the "fund_manager" account.
func (inst *FundManagerSettleRewardInstruction) SetFundManagerAccount(fundManager ag_solanago.PublicKey) *FundManagerSettleRewardInstruction {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(fundManager).SIGNER()
	return inst
}

// GetFundManagerAccount gets the "fund_manager" account.
func (inst *FundManagerSettleRewardInstruction) GetFundManagerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetReceiptTokenMintAccount sets the "receipt_token_mint" account.
func (inst *FundManagerSettleRewardInstruction) SetReceiptTokenMintAccount(receiptTokenMint ag_solanago.PublicKey) *FundManagerSettleRewardInstruction {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(receiptTokenMint)
	return inst
}

// GetReceiptTokenMintAccount gets the "receipt_token_mint" account.
func (inst *FundManagerSettleRewardInstruction) GetReceiptTokenMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetRewardAccountAccount sets the "reward_account" account.
func (inst *FundManagerSettleRewardInstruction) SetRewardAccountAccount(rewardAccount ag_solanago.PublicKey) *FundManagerSettleRewardInstruction {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(rewardAccount).WRITE()
	return inst
}

func (inst *FundManagerSettleRewardInstruction) findFindRewardAccountAddress(receiptTokenMint ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
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
func (inst *FundManagerSettleRewardInstruction) FindRewardAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindRewardAccountAddress(receiptTokenMint, bumpSeed)
	return
}

func (inst *FundManagerSettleRewardInstruction) MustFindRewardAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindRewardAccountAddress(receiptTokenMint, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindRewardAccountAddress finds RewardAccount account address with given seeds.
func (inst *FundManagerSettleRewardInstruction) FindRewardAccountAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindRewardAccountAddress(receiptTokenMint, 0)
	return
}

func (inst *FundManagerSettleRewardInstruction) MustFindRewardAccountAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindRewardAccountAddress(receiptTokenMint, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetRewardAccountAccount gets the "reward_account" account.
func (inst *FundManagerSettleRewardInstruction) GetRewardAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetRewardTokenMintAccount sets the "reward_token_mint" account.
func (inst *FundManagerSettleRewardInstruction) SetRewardTokenMintAccount(rewardTokenMint ag_solanago.PublicKey) *FundManagerSettleRewardInstruction {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(rewardTokenMint)
	return inst
}

// GetRewardTokenMintAccount gets the "reward_token_mint" account (optional).
func (inst *FundManagerSettleRewardInstruction) GetRewardTokenMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetRewardTokenProgramAccount sets the "reward_token_program" account.
func (inst *FundManagerSettleRewardInstruction) SetRewardTokenProgramAccount(rewardTokenProgram ag_solanago.PublicKey) *FundManagerSettleRewardInstruction {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(rewardTokenProgram)
	return inst
}

// GetRewardTokenProgramAccount gets the "reward_token_program" account (optional).
func (inst *FundManagerSettleRewardInstruction) GetRewardTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetEventAuthorityAccount sets the "event_authority" account.
func (inst *FundManagerSettleRewardInstruction) SetEventAuthorityAccount(eventAuthority ag_solanago.PublicKey) *FundManagerSettleRewardInstruction {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(eventAuthority)
	return inst
}

func (inst *FundManagerSettleRewardInstruction) findFindEventAuthorityAddress(knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
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
func (inst *FundManagerSettleRewardInstruction) FindEventAuthorityAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindEventAuthorityAddress(bumpSeed)
	return
}

func (inst *FundManagerSettleRewardInstruction) MustFindEventAuthorityAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindEventAuthorityAddress(bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindEventAuthorityAddress finds EventAuthority account address with given seeds.
func (inst *FundManagerSettleRewardInstruction) FindEventAuthorityAddress() (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindEventAuthorityAddress(0)
	return
}

func (inst *FundManagerSettleRewardInstruction) MustFindEventAuthorityAddress() (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindEventAuthorityAddress(0)
	if err != nil {
		panic(err)
	}
	return
}

// GetEventAuthorityAccount gets the "event_authority" account.
func (inst *FundManagerSettleRewardInstruction) GetEventAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetProgramAccount sets the "program" account.
func (inst *FundManagerSettleRewardInstruction) SetProgramAccount(program ag_solanago.PublicKey) *FundManagerSettleRewardInstruction {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(program)
	return inst
}

// GetProgramAccount gets the "program" account.
func (inst *FundManagerSettleRewardInstruction) GetProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

func (inst FundManagerSettleRewardInstruction) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_FundManagerSettleReward,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst FundManagerSettleRewardInstruction) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *FundManagerSettleRewardInstruction) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.RewardPoolId == nil {
			return errors.New("RewardPoolId parameter is not set")
		}
		if inst.RewardId == nil {
			return errors.New("RewardId parameter is not set")
		}
		if inst.Amount == nil {
			return errors.New("Amount parameter is not set")
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
			return errors.New("accounts.RewardAccount is not set")
		}

		// [3] = RewardTokenMint is optional

		// [4] = RewardTokenProgram is optional

		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.EventAuthority is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.Program is not set")
		}
	}
	return nil
}

func (inst *FundManagerSettleRewardInstruction) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("FundManagerSettleReward")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=3]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("  RewardPoolId", *inst.RewardPoolId))
						paramsBranch.Child(ag_format.Param("      RewardId", *inst.RewardId))
						paramsBranch.Child(ag_format.Param("        Amount", *inst.Amount))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=7]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("        fund_manager", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("  receipt_token_mint", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("             reward_", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("   reward_token_mint", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("reward_token_program", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("     event_authority", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("             program", inst.AccountMetaSlice.Get(6)))
					})
				})
		})
}

func (obj FundManagerSettleRewardInstruction) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `RewardPoolId` param:
	err = encoder.Encode(obj.RewardPoolId)
	if err != nil {
		return err
	}
	// Serialize `RewardId` param:
	err = encoder.Encode(obj.RewardId)
	if err != nil {
		return err
	}
	// Serialize `Amount` param:
	err = encoder.Encode(obj.Amount)
	if err != nil {
		return err
	}
	return nil
}
func (obj *FundManagerSettleRewardInstruction) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `RewardPoolId`:
	err = decoder.Decode(&obj.RewardPoolId)
	if err != nil {
		return err
	}
	// Deserialize `RewardId`:
	err = decoder.Decode(&obj.RewardId)
	if err != nil {
		return err
	}
	// Deserialize `Amount`:
	err = decoder.Decode(&obj.Amount)
	if err != nil {
		return err
	}
	return nil
}

// NewFundManagerSettleRewardInstruction declares a new FundManagerSettleReward instruction with the provided parameters and accounts.
func NewFundManagerSettleRewardInstruction(
	// Parameters:
	reward_pool_id uint8,
	reward_id uint16,
	amount uint64,
	// Accounts:
	fundManager ag_solanago.PublicKey,
	receiptTokenMint ag_solanago.PublicKey,
	rewardAccount ag_solanago.PublicKey,
	rewardTokenMint ag_solanago.PublicKey,
	rewardTokenProgram ag_solanago.PublicKey,
	eventAuthority ag_solanago.PublicKey,
	program ag_solanago.PublicKey) *FundManagerSettleRewardInstruction {
	return NewFundManagerSettleRewardInstructionBuilder().
		SetRewardPoolId(reward_pool_id).
		SetRewardId(reward_id).
		SetAmount(amount).
		SetFundManagerAccount(fundManager).
		SetReceiptTokenMintAccount(receiptTokenMint).
		SetRewardAccountAccount(rewardAccount).
		SetRewardTokenMintAccount(rewardTokenMint).
		SetRewardTokenProgramAccount(rewardTokenProgram).
		SetEventAuthorityAccount(eventAuthority).
		SetProgramAccount(program)
}
