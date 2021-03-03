package config

import (
	"testing"

	"github.com/mazrean/gopendb-generator/cmd/domain"
	"github.com/mazrean/gopendb-generator/cmd/interfaces/config/models"
	conf "github.com/mazrean/gopendb-generator/cmd/usecases/config"
	"github.com/stretchr/testify/assert"
)

func tableTest(t *testing.T) {
	t.Helper()

	t.Run("GetAll", getAllTest)
}

func getAllTest(t *testing.T) {
	t.Helper()

	assertion := assert.New(t)

	table := NewTable()

	type arg struct {
		config models.Root
	}
	type expect struct {
		tableDetails []*conf.TableDetail
	}
	type test struct {
		description string
		arg
		expect
	}

	testCases := []test{
		{
			description: "normal",
			arg: arg{
				config: models.Root{
					Tables: []*models.Table{
						{
							Table: &domain.Table{
								ID:          "user",
								Description: "ユーザーテーブル",
								Name:        "users",
							},
							PrimaryKey: &models.PrimaryKey{
								Columns: []*domain.IndexColumn{
									{
										Name: "id",
									},
								},
							},
						},
					},
				},
			},
			expect: expect{
				tableDetails: []*conf.TableDetail{
					{
						Table: &domain.Table{
							ID:          "user",
							Description: "ユーザーテーブル",
							Name:        "users",
						},
						PrimaryKeyColumnNames: []string{"id"},
					},
				},
			},
		},
		{
			description: "nil table",
			arg: arg{
				config: models.Root{
					Tables: nil,
				},
			},
			expect: expect{
				tableDetails: []*conf.TableDetail{},
			},
		},
		{
			description: "nil primary key",
			arg: arg{
				config: models.Root{
					Tables: []*models.Table{
						{
							Table: &domain.Table{
								ID:          "user",
								Description: "ユーザーテーブル",
								Name:        "users",
							},
							PrimaryKey: nil,
						},
					},
				},
			},
			expect: expect{
				tableDetails: []*conf.TableDetail{
					{
						Table: &domain.Table{
							ID:          "user",
							Description: "ユーザーテーブル",
							Name:        "users",
						},
						PrimaryKeyColumnNames: []string{},
					},
				},
			},
		},
		{
			description: "multi column primary key",
			arg: arg{
				config: models.Root{
					Tables: []*models.Table{
						{
							Table: &domain.Table{
								ID:          "user",
								Description: "ユーザーテーブル",
								Name:        "users",
							},
							PrimaryKey: &models.PrimaryKey{
								Columns: []*domain.IndexColumn{
									{
										Name: "id",
									},
									{
										Name: "name",
									},
								},
							},
						},
					},
				},
			},
			expect: expect{
				tableDetails: []*conf.TableDetail{
					{
						Table: &domain.Table{
							ID:          "user",
							Description: "ユーザーテーブル",
							Name:        "users",
						},
						PrimaryKeyColumnNames: []string{"id", "name"},
					},
				},
			},
		},
	}

	for _, testCase := range testCases {
		config = testCase.arg.config

		actualTableDetails := table.GetAll()

		assertion.Equal(testCase.expect.tableDetails, actualTableDetails, testCase.description)
	}
}
