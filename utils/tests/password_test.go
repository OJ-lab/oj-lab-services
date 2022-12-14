package tests

import (
	"github.com/OJ-lab/oj-lab-services/utils"
	"github.com/alexedwards/argon2id"
	"testing"
)

func TestPassword(t *testing.T) {
	hashedPassword, err := utils.GetHashedPassword("pa$$word", argon2id.DefaultParams)
	if err != nil {
		panic(err)
	}
	result, err := utils.CompareWithHashedPassword("pa$$word", hashedPassword)
	if err != nil {
		panic(err)
	}
	if !result {
		panic("Result is not true")
	}
}
