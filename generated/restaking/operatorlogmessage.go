// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package restaking

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// OperatorLogMessage is the `operator_log_message` instruction.
type OperatorLogMessageInstruction struct {
	Message *string

	// [0] = [] event_authority
	//
	// [1] = [] program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewOperatorLogMessageInstructionBuilder creates a new `OperatorLogMessageInstruction` instruction builder.
func NewOperatorLogMessageInstructionBuilder() *OperatorLogMessageInstruction {
	nd := &OperatorLogMessageInstruction{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 2),
	}
	return nd
}

// SetMessage sets the "message" parameter.
func (inst *OperatorLogMessageInstruction) SetMessage(message string) *OperatorLogMessageInstruction {
	inst.Message = &message
	return inst
}

// SetEventAuthorityAccount sets the "event_authority" account.
func (inst *OperatorLogMessageInstruction) SetEventAuthorityAccount(eventAuthority ag_solanago.PublicKey) *OperatorLogMessageInstruction {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(eventAuthority)
	return inst
}

func (inst *OperatorLogMessageInstruction) findFindEventAuthorityAddress(knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
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
func (inst *OperatorLogMessageInstruction) FindEventAuthorityAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindEventAuthorityAddress(bumpSeed)
	return
}

func (inst *OperatorLogMessageInstruction) MustFindEventAuthorityAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindEventAuthorityAddress(bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindEventAuthorityAddress finds EventAuthority account address with given seeds.
func (inst *OperatorLogMessageInstruction) FindEventAuthorityAddress() (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindEventAuthorityAddress(0)
	return
}

func (inst *OperatorLogMessageInstruction) MustFindEventAuthorityAddress() (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindEventAuthorityAddress(0)
	if err != nil {
		panic(err)
	}
	return
}

// GetEventAuthorityAccount gets the "event_authority" account.
func (inst *OperatorLogMessageInstruction) GetEventAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetProgramAccount sets the "program" account.
func (inst *OperatorLogMessageInstruction) SetProgramAccount(program ag_solanago.PublicKey) *OperatorLogMessageInstruction {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(program)
	return inst
}

// GetProgramAccount gets the "program" account.
func (inst *OperatorLogMessageInstruction) GetProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

func (inst OperatorLogMessageInstruction) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_OperatorLogMessage,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst OperatorLogMessageInstruction) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *OperatorLogMessageInstruction) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Message == nil {
			return errors.New("Message parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.EventAuthority is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Program is not set")
		}
	}
	return nil
}

func (inst *OperatorLogMessageInstruction) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("OperatorLogMessage")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Message", *inst.Message))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=2]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("event_authority", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("        program", inst.AccountMetaSlice.Get(1)))
					})
				})
		})
}

func (obj OperatorLogMessageInstruction) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Message` param:
	err = encoder.Encode(obj.Message)
	if err != nil {
		return err
	}
	return nil
}
func (obj *OperatorLogMessageInstruction) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Message`:
	err = decoder.Decode(&obj.Message)
	if err != nil {
		return err
	}
	return nil
}

// NewOperatorLogMessageInstruction declares a new OperatorLogMessage instruction with the provided parameters and accounts.
func NewOperatorLogMessageInstruction(
	// Parameters:
	message string,
	// Accounts:
	eventAuthority ag_solanago.PublicKey,
	program ag_solanago.PublicKey) *OperatorLogMessageInstruction {
	return NewOperatorLogMessageInstructionBuilder().
		SetMessage(message).
		SetEventAuthorityAccount(eventAuthority).
		SetProgramAccount(program)
}
