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

var ValidPlan = map[string]bool{
	"SUBSCRIPTION": true,
	"FREE":         true,
}
var ValidGroup = map[string]bool{
	"MULTIPLE_CHOICE":        true,
	"MULTIPLE_CHOICE_HIDDEN": true,
	"ESSAY":                  true,
}
var ValidLang = map[string]bool{
	"vn":  true,
	"eng": true,
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
		result.AddError("skill", "Skill is required")
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
		result.AddError("skill", "Name is required")
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
func ValidateCreateQuestionGroup(dataRequest *dto.CreateQuestionGroupRequest) *validation.ValidationResult {
	if dataRequest == nil {
		return nil
	}
	result := validation.NewValidationResult()
	// Validate name
	if utils.IsEmpty(dataRequest.Name) {
		result.AddError("full_name", "Name is required")
	}
	// Validate Plan
	if utils.IsEmpty(dataRequest.PlanType) {
		result.AddError("plan_type", "Plan is required")
	} else {
		if !ValidPlan[dataRequest.PlanType] {
			result.AddError("plan_type", "Plan must be one of 'SUBSCRIPTION', 'FREE'")
		}
	}
	// Validate Group
	if utils.IsEmpty(dataRequest.GroupType) {
		result.AddError("group_type", "Group is required")
	} else {
		if !ValidGroup[dataRequest.GroupType] {
			result.AddError("group_type", "Group must be one of 'MULTIPLE_CHOICE', 'MULTIPLE_CHOICE_HIDDEN','ESSAY'")
		}
	}
	return result
}
func ValidateUpdateQuestionGroup(dataRequest *dto.UpdateQuestionGroupRequest) *validation.ValidationResult {
	if dataRequest == nil {
		return nil
	}
	result := validation.NewValidationResult()
	// Validate name
	if utils.IsEmpty(dataRequest.Name) {
		result.AddError("full_name", "Name is required")
	}
	// Validate Plan
	if utils.IsEmpty(dataRequest.PlanType) {
		result.AddError("plan_type", "Plan is required")
	} else {
		if !ValidPlan[dataRequest.PlanType] {
			result.AddError("plan_type", "Plan must be one of 'SUBSCRIPTION', 'FREE'")
		}
	}
	// Validate Group
	if utils.IsEmpty(dataRequest.GroupType) {
		result.AddError("group_type", "Group is required")
	} else {
		if !ValidGroup[dataRequest.GroupType] {
			result.AddError("group_type", "Group must be one of 'MULTIPLE_CHOICE', 'MULTIPLE_CHOICE_HIDDEN','ESSAY'")
		}
	}
	return result
}
func ValidateLang(lang string) bool {
	return !ValidLang[lang]
}
