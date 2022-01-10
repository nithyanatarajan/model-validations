package models

import (
	"context"
	"errors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"strconv"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) ValidateWithContext(ctx context.Context) error {
	return validation.ValidateStructWithContext(ctx, &p,
		//validation.Field(&p.Age, validation.Required, validation.WithContext(AgeLimitCustomRuleDefinition())),
		validation.Field(&p.Age, AgeLimitCustomRuleDefinitionRules(ctx)...),
	)
}

func AgeLimitCustomRuleDefinitionRules(ctx context.Context) (rules []validation.Rule) {
	rules = append(rules, validation.Required)
	if data, ok := ctx.Value("ageLimit").(string); ok {
		ageLimit, _ := strconv.Atoi(data)
		rules = append(rules, validation.Max(ageLimit).Error("person's age limit reached"))
	}
	return
}

func AgeLimitCustomRuleDefinition() validation.RuleWithContextFunc {
	return func(ctx context.Context, ageValue interface{}) error {
		if data, ok := ctx.Value("ageLimit").(string); ok {
			ageLimit, err := strconv.Atoi(data)
			if err != nil {
				return errors.New("unable to validate age limit")
			}

			if ageValue.(int) > ageLimit {
				return errors.New("person's age limit reached")
			}
		}

		return nil
	}
}
