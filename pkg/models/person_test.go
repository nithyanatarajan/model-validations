package models_test

import (
	"context"
	"errors"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/nithyanatarajan/validations/pkg/models"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestPersonWithContext(t *testing.T) {
	person := models.Person{Age: 70}
	rule := validation.WithContext(func(ctx context.Context, value interface{}) error {
		ageLimit, _ := strconv.Atoi(ctx.Value("ageLimit").(string))
		if value.(models.Person).Age > ageLimit {
			return errors.New("person's age limit reached")
		}

		return nil
	})

	ctx := context.WithValue(context.Background(), "ageLimit", "60")

	err := validation.ValidateWithContext(ctx, person, rule)
	fmt.Println(err)

	assert.EqualError(t, err, "person's age limit reached")
}

func TestPerson_ValidateWithContext(t *testing.T) {
	ctx := context.WithValue(context.Background(), "ageLimit", "60")
	t.Run("age cannot be empty", func(t *testing.T) {
		person := models.Person{}

		err := person.ValidateWithContext(ctx)

		assert.EqualError(t, err, "Age: cannot be blank.")
	})

	t.Run("age cannot be greater than 60", func(t *testing.T) {
		person := models.Person{Age: 100}

		err := person.ValidateWithContext(ctx)

		assert.EqualError(t, err, "Age: person's age limit reached.")
	})
}

func TestAgeLimitCustomRuleDefinition(t *testing.T) {
	ctx := context.WithValue(context.Background(), "ageLimit", "60")

	t.Run("should return error when age limit reached", func(t *testing.T) {
		ruleDefinition := models.AgeLimitCustomRuleDefinition()

		err := ruleDefinition(ctx, 100)

		assert.EqualError(t, err, "person's age limit reached")
	})

	t.Run("should return nil when age limit not reached", func(t *testing.T) {
		ruleDefinition := models.AgeLimitCustomRuleDefinition()

		err := ruleDefinition(ctx, 60)

		assert.NoError(t, err)
	})

	t.Run("should return nil when no age limit is specified", func(t *testing.T) {
		ruleDefinition := models.AgeLimitCustomRuleDefinition()

		err := ruleDefinition(context.Background(), 60)

		assert.NoError(t, err)
	})

	t.Run("should return error when unable to parse age limit", func(t *testing.T) {
		ctx := context.WithValue(context.Background(), "ageLimit", "blah")
		ruleDefinition := models.AgeLimitCustomRuleDefinition()

		err := ruleDefinition(ctx, 60)

		assert.EqualError(t, err, "unable to validate age limit")
	})
}
