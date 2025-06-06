// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package restaking

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// FundManagerInitializeFundJitoRestakingVaultDelegation is the `fund_manager_initialize_fund_jito_restaking_vault_delegation` instruction.
type FundManagerInitializeFundJitoRestakingVaultDelegationInstruction struct {

	// [0] = [SIGNER] fund_manager
	//
	// [1] = [WRITE] fund_account
	//
	// [2] = [] receipt_token_mint
	//
	// [3] = [] vault_account
	//
	// [4] = [] operator_account
	//
	// [5] = [] vault_operator_delegation
	//
	// [6] = [] event_authority
	//
	// [7] = [] program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewFundManagerInitializeFundJitoRestakingVaultDelegationInstructionBuilder creates a new `FundManagerInitializeFundJitoRestakingVaultDelegationInstruction` instruction builder.
func NewFundManagerInitializeFundJitoRestakingVaultDelegationInstructionBuilder() *FundManagerInitializeFundJitoRestakingVaultDelegationInstruction {
	nd := &FundManagerInitializeFundJitoRestakingVaultDelegationInstruction{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 8),
	}
	nd.AccountMetaSlice[0] = ag_solanago.Meta(Addresses["5UpLTLA7Wjqp7qdfjuTtPcUw3aVtbqFA5Mgm34mxPNg2"]).SIGNER()
	return nd
}

// SetFundManagerAccount sets the "fund_manager" account.
func (inst *FundManagerInitializeFundJitoRestakingVaultDelegationInstruction) SetFundManagerAccount(fundManager ag_solanago.PublicKey) *FundManagerInitializeFundJitoRestakingVaultDelegationInstruction {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(fundManager).SIGNER()
	return inst
}

// GetFundManagerAccount gets the "fund_manager" account.
func (inst *FundManagerInitializeFundJitoRestakingVaultDelegationInstruction) GetFundManagerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetFundAccountAccount sets the "fund_account" account.
func (inst *FundManagerInitializeFundJitoRestakingVaultDelegationInstruction) SetFundAccountAccount(fundAccount ag_solanago.PublicKey) *FundManagerInitializeFundJitoRestakingVaultDelegationInstruction {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(fundAccount).WRITE()
	return inst
}

func (inst *FundManagerInitializeFundJitoRestakingVaultDelegationInstruction) findFindFundAccountAddress(receiptTokenMint ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
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
func (inst *FundManagerInitializeFundJitoRestakingVaultDelegationInstruction) FindFundAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindFundAccountAddress(receiptTokenMint, bumpSeed)
	return
}

func (inst *FundManagerInitializeFundJitoRestakingVaultDelegationInstruction) MustFindFundAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindFundAccountAddress(receiptTokenMint, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindFundAccountAddress finds FundAccount account address with given seeds.
func (inst *FundManagerInitializeFundJitoRestakingVaultDelegationInstruction) FindFundAccountAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindFundAccountAddress(receiptTokenMint, 0)
	return
}

func (inst *FundManagerInitializeFundJitoRestakingVaultDelegationInstruction) MustFindFundAccountAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindFundAccountAddress(receiptTokenMint, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetFundAccountAccount gets the "fund_account" account.
func (inst *FundManagerInitializeFundJitoRestakingVaultDelegationInstruction) GetFundAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetReceiptTokenMintAccount sets the "receipt_token_mint" account.
func (inst *FundManagerInitializeFundJitoRestakingVaultDelegationInstruction) SetReceiptTokenMintAccount(receiptTokenMint ag_solanago.PublicKey) *FundManagerInitializeFundJitoRestakingVaultDelegationInstruction {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(receiptTokenMint)
	return inst
}

// GetReceiptTokenMintAccount gets the "receipt_token_mint" account.
func (inst *FundManagerInitializeFundJitoRestakingVaultDelegationInstruction) GetReceiptTokenMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetVaultAccountAccount sets the "vault_account" account.
func (inst *FundManagerInitializeFundJitoRestakingVaultDelegationInstruction) SetVaultAccountAccount(vaultAccount ag_solanago.PublicKey) *FundManagerInitializeFundJitoRestakingVaultDelegationInstruction {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(vaultAccount)
	return inst
}

// GetVaultAccountAccount gets the "vault_account" account.
func (inst *FundManagerInitializeFundJitoRestakingVaultDelegationInstruction) GetVaultAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetOperatorAccountAccount sets the "operator_account" account.
func (inst *FundManagerInitializeFundJitoRestakingVaultDelegationInstruction) SetOperatorAccountAccount(operatorAccount ag_solanago.PublicKey) *FundManagerInitializeFundJitoRestakingVaultDelegationInstruction {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(operatorAccount)
	return inst
}

// GetOperatorAccountAccount gets the "operator_account" account.
func (inst *FundManagerInitializeFundJitoRestakingVaultDelegationInstruction) GetOperatorAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetVaultOperatorDelegationAccount sets the "vault_operator_delegation" account.
func (inst *FundManagerInitializeFundJitoRestakingVaultDelegationInstruction) SetVaultOperatorDelegationAccount(vaultOperatorDelegation ag_solanago.PublicKey) *FundManagerInitializeFundJitoRestakingVaultDelegationInstruction {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(vaultOperatorDelegation)
	return inst
}

func (inst *FundManagerInitializeFundJitoRestakingVaultDelegationInstruction) findFindVaultOperatorDelegationAddress(vaultAccount ag_solanago.PublicKey, operatorAccount ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	var seeds [][]byte
	// const: vault_operator_delegation
	seeds = append(seeds, []byte{byte(0x76), byte(0x61), byte(0x75), byte(0x6c), byte(0x74), byte(0x5f), byte(0x6f), byte(0x70), byte(0x65), byte(0x72), byte(0x61), byte(0x74), byte(0x6f), byte(0x72), byte(0x5f), byte(0x64), byte(0x65), byte(0x6c), byte(0x65), byte(0x67), byte(0x61), byte(0x74), byte(0x69), byte(0x6f), byte(0x6e)})
	// path: vaultAccount
	seeds = append(seeds, vaultAccount.Bytes())
	// path: operatorAccount
	seeds = append(seeds, operatorAccount.Bytes())

	programID := Addresses["Vau1t6sLNxnzB7ZDsef8TLbPLfyZMYXH8WTNqUdm9g8"]

	if knownBumpSeed != 0 {
		seeds = append(seeds, []byte{byte(bumpSeed)})
		pda, err = ag_solanago.CreateProgramAddress(seeds, programID)
	} else {
		pda, bumpSeed, err = ag_solanago.FindProgramAddress(seeds, programID)
	}
	return
}

// FindVaultOperatorDelegationAddressWithBumpSeed calculates VaultOperatorDelegation account address with given seeds and a known bump seed.
func (inst *FundManagerInitializeFundJitoRestakingVaultDelegationInstruction) FindVaultOperatorDelegationAddressWithBumpSeed(vaultAccount ag_solanago.PublicKey, operatorAccount ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindVaultOperatorDelegationAddress(vaultAccount, operatorAccount, bumpSeed)
	return
}

func (inst *FundManagerInitializeFundJitoRestakingVaultDelegationInstruction) MustFindVaultOperatorDelegationAddressWithBumpSeed(vaultAccount ag_solanago.PublicKey, operatorAccount ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindVaultOperatorDelegationAddress(vaultAccount, operatorAccount, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindVaultOperatorDelegationAddress finds VaultOperatorDelegation account address with given seeds.
func (inst *FundManagerInitializeFundJitoRestakingVaultDelegationInstruction) FindVaultOperatorDelegationAddress(vaultAccount ag_solanago.PublicKey, operatorAccount ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindVaultOperatorDelegationAddress(vaultAccount, operatorAccount, 0)
	return
}

func (inst *FundManagerInitializeFundJitoRestakingVaultDelegationInstruction) MustFindVaultOperatorDelegationAddress(vaultAccount ag_solanago.PublicKey, operatorAccount ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindVaultOperatorDelegationAddress(vaultAccount, operatorAccount, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetVaultOperatorDelegationAccount gets the "vault_operator_delegation" account.
func (inst *FundManagerInitializeFundJitoRestakingVaultDelegationInstruction) GetVaultOperatorDelegationAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetEventAuthorityAccount sets the "event_authority" account.
func (inst *FundManagerInitializeFundJitoRestakingVaultDelegationInstruction) SetEventAuthorityAccount(eventAuthority ag_solanago.PublicKey) *FundManagerInitializeFundJitoRestakingVaultDelegationInstruction {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(eventAuthority)
	return inst
}

func (inst *FundManagerInitializeFundJitoRestakingVaultDelegationInstruction) findFindEventAuthorityAddress(knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
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
func (inst *FundManagerInitializeFundJitoRestakingVaultDelegationInstruction) FindEventAuthorityAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindEventAuthorityAddress(bumpSeed)
	return
}

func (inst *FundManagerInitializeFundJitoRestakingVaultDelegationInstruction) MustFindEventAuthorityAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindEventAuthorityAddress(bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindEventAuthorityAddress finds EventAuthority account address with given seeds.
func (inst *FundManagerInitializeFundJitoRestakingVaultDelegationInstruction) FindEventAuthorityAddress() (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindEventAuthorityAddress(0)
	return
}

func (inst *FundManagerInitializeFundJitoRestakingVaultDelegationInstruction) MustFindEventAuthorityAddress() (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindEventAuthorityAddress(0)
	if err != nil {
		panic(err)
	}
	return
}

// GetEventAuthorityAccount gets the "event_authority" account.
func (inst *FundManagerInitializeFundJitoRestakingVaultDelegationInstruction) GetEventAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetProgramAccount sets the "program" account.
func (inst *FundManagerInitializeFundJitoRestakingVaultDelegationInstruction) SetProgramAccount(program ag_solanago.PublicKey) *FundManagerInitializeFundJitoRestakingVaultDelegationInstruction {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(program)
	return inst
}

// GetProgramAccount gets the "program" account.
func (inst *FundManagerInitializeFundJitoRestakingVaultDelegationInstruction) GetProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

func (inst FundManagerInitializeFundJitoRestakingVaultDelegationInstruction) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_FundManagerInitializeFundJitoRestakingVaultDelegation,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst FundManagerInitializeFundJitoRestakingVaultDelegationInstruction) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *FundManagerInitializeFundJitoRestakingVaultDelegationInstruction) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.FundManager is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.FundAccount is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.ReceiptTokenMint is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.VaultAccount is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.OperatorAccount is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.VaultOperatorDelegation is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.EventAuthority is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.Program is not set")
		}
	}
	return nil
}

func (inst *FundManagerInitializeFundJitoRestakingVaultDelegationInstruction) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("FundManagerInitializeFundJitoRestakingVaultDelegation")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=8]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("             fund_manager", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("                    fund_", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("       receipt_token_mint", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("                   vault_", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("                operator_", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("vault_operator_delegation", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("          event_authority", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("                  program", inst.AccountMetaSlice.Get(7)))
					})
				})
		})
}

func (obj FundManagerInitializeFundJitoRestakingVaultDelegationInstruction) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *FundManagerInitializeFundJitoRestakingVaultDelegationInstruction) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewFundManagerInitializeFundJitoRestakingVaultDelegationInstruction declares a new FundManagerInitializeFundJitoRestakingVaultDelegation instruction with the provided parameters and accounts.
func NewFundManagerInitializeFundJitoRestakingVaultDelegationInstruction(
	// Accounts:
	fundManager ag_solanago.PublicKey,
	fundAccount ag_solanago.PublicKey,
	receiptTokenMint ag_solanago.PublicKey,
	vaultAccount ag_solanago.PublicKey,
	operatorAccount ag_solanago.PublicKey,
	vaultOperatorDelegation ag_solanago.PublicKey,
	eventAuthority ag_solanago.PublicKey,
	program ag_solanago.PublicKey) *FundManagerInitializeFundJitoRestakingVaultDelegationInstruction {
	return NewFundManagerInitializeFundJitoRestakingVaultDelegationInstructionBuilder().
		SetFundManagerAccount(fundManager).
		SetFundAccountAccount(fundAccount).
		SetReceiptTokenMintAccount(receiptTokenMint).
		SetVaultAccountAccount(vaultAccount).
		SetOperatorAccountAccount(operatorAccount).
		SetVaultOperatorDelegationAccount(vaultOperatorDelegation).
		SetEventAuthorityAccount(eventAuthority).
		SetProgramAccount(program)
}
