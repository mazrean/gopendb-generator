package config

import (
	"errors"
	"testing"

	"github.com/mazrean/gopendb-generator/cmd/domain"
	"github.com/mazrean/gopendb-generator/cmd/interfaces/config/models"
	"github.com/stretchr/testify/assert"
)

func configTest(t *testing.T) {
	t.Helper()

	assertion := assert.New(t)

	conf := NewConfig()

	type arg struct {
		config models.Root
	}
	type expect struct {
		config *domain.Config
		isErr  bool
		err    error
	}
	type test struct {
		description string
		arg
		expect
	}

	normalConfig := domain.Config{
		DBMS:     domain.MySQL,
		Version:  "8",
		Database: "test",
	}
	testCases := []test{
		{
			description: "normal",
			arg: arg{
				config: models.Root{
					Config: &normalConfig,
				},
			},
			expect: expect{
				config: &normalConfig,
			},
		},
		{
			description: "nil conf",
			arg: arg{
				config: models.Root{
					Config: nil,
				},
			},
			expect: expect{
				config: nil,
			},
		},
	}

	for _, testCase := range testCases {
		config = testCase.arg.config

		actualConfig, err := conf.Get()

		if err != nil {
			if testCase.isErr {
				if testCase.expect.err != nil {
					assertion.True(errors.Is(err, testCase.expect.err), testCase.description, "err")
				}
				continue
			} else {
				t.Fatalf("unexpected error(%s):%+v", testCase.description, err)
			}
		}
		if err == nil && testCase.expect.isErr {
			t.Fatalf("unexpected no error(%s)", testCase.description)
		}

		assertion.EqualValuesf(testCase.expect.config, actualConfig, testCase.description, "value")
	}
}
