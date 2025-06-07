package validation

import (
	"pirate-lang-go/core/utils"
	"pirate-lang-go/core/validation"
	"pirate-lang-go/modules/library/dto"
)

var ValidSkills = map[string]bool{
	"Listening": true,
	"Reading":   true,
	"Speaking":  true,
	"Writing":   true,
}

func ValidateCreatePart(dataRequest *dto.CreatePartRequest) *validation.ValidationResult {
	if dataRequest == nil {
		return nil
	}
	result := validation.NewValidationResult()
	// Validate name
	if utils.IsEmpty(dataRequest.Name) {
		result.AddError("full_name", "Name is required")
	}
	// Validate Skill
	if utils.IsEmpty(dataRequest.Skill) {
		result.AddError("full_name", "Name is required")
	} else {
		if !ValidSkills[dataRequest.Skill] {
			result.AddError("skill", "Skill must be one of 'Listening', 'Reading', 'Speaking', 'Writing'")
		}
	}
	// Validate Sequence
	if dataRequest.Sequence <= 0 {
		result.AddError("sequence", "Sequence must be a positive number")
	}
	return result
}
func ValidateUpdatePart(dataRequest *dto.UpdatePartRequest) *validation.ValidationResult {
	if dataRequest == nil {
		return nil
	}
	result := validation.NewValidationResult()
	// Validate name
	if utils.IsEmpty(dataRequest.Name) {
		result.AddError("full_name", "Name is required")
	}
	// Validate Skill
	if utils.IsEmpty(dataRequest.Skill) {
		result.AddError("full_name", "Name is required")
	} else {
		if !ValidSkills[dataRequest.Skill] {
			result.AddError("skill", "Skill must be one of 'Listening', 'Reading', 'Speaking', 'Writing'")
		}
	}
	// Validate Sequence
	if dataRequest.Sequence <= 0 {
		result.AddError("sequence", "Sequence must be a positive number")
	}
	return result
}
