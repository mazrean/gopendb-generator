package config

import (
	"testing"

	"github.com/mazrean/gopendb-generator/cmd/domain"
	"github.com/mazrean/gopendb-generator/cmd/interfaces/config/models"
	"github.com/stretchr/testify/assert"
)

func TestReader(t *testing.T) {
	t.Parallel()

	t.Run("ReadYAML", readYAMLTest)
}

func readYAMLTest(t *testing.T) {
	t.Parallel()
	t.Helper()

	assertion := assert.New(t)

	reader := NewReader()

	path := "../../test_data/test.yaml"

	err := reader.ReadYAML(path)
	if err != nil {
		t.Errorf("failed to read yaml: %w", err)
	}

	expectConfig := models.Root{
		Config: &domain.Config{
			DBMS:     domain.MySQL,
			Version:  "8.0",
			Database: "test",
		},
		Tables: []*models.Table{
			{
				Table: &domain.Table{
					ID:           "user",
					Description:  "ユーザーテーブル",
					Name:         "users",
					Engin:        "InnoDB",
					CharSet:      "utf8mb4",
					MaxRows:      100000,
					MinRows:      10,
					AvgRowLength: 100,
				},
				PrimaryKey: &models.PrimaryKey{
					Columns: []*domain.IndexColumn{
						{
							Name:      "id",
							Length:    10,
							Direction: domain.Asc,
						},
					},
					Options: &models.IndexOption{
						IndexOption: &domain.IndexOption{
							KeyBlockSize: 0,
							Parser:       "my_parser",
						},
						Type: domain.Btree,
					},
				},
				Columns: []*models.Column{
					{
						Column: &domain.Column{
							ID:          "id",
							Description: "ユーザーID",
							Name:        "id",
							Type: domain.Type{
								Name:   domain.Int,
								Length: 11,
							},
						},
						Reference: nil,
						Extra: &domain.Extra{
							AutoIncrement: true,
							Format:        domain.Dynamic,
							Storage:       domain.Disk,
						},
					},
					{
						Column: &domain.Column{
							ID:          "name",
							Description: "ユーザー名",
							Name:        "name",
							Type: domain.Type{
								Name: domain.Text,
							},
						},
						Reference: nil,
						Extra:     nil,
					},
					{
						Column: &domain.Column{
							ID:          "created_at",
							Description: "ユーザー作成日時",
							Name:        "created_at",
							Type: domain.Type{
								Name: domain.Datetime,
							},
						},
						Reference: nil,
						Extra:     nil,
					},
					{
						Column: &domain.Column{
							ID:          "deleted_at",
							Description: "ユーザー削除日時",
							Name:        "deleted_at",
							Type: domain.Type{
								Name: domain.Datetime,
							},
						},
						Reference: nil,
						Extra:     nil,
					},
				},
			},
			{
				Table: &domain.Table{
					ID:           "message",
					Description:  "メッセージテーブル",
					Name:         "messages",
					Engin:        "",
					CharSet:      "",
					MaxRows:      0,
					MinRows:      0,
					AvgRowLength: 0,
				},
				PrimaryKey: &models.PrimaryKey{
					Columns: []*domain.IndexColumn{
						{
							Name:      "id",
							Length:    0,
							Direction: domain.IndexDirectionDefault,
						},
					},
					Options: nil,
				},
				Columns: []*models.Column{
					{
						Column: &domain.Column{
							ID:          "id",
							Description: "メッセージID",
							Name:        "id",
							Type: domain.Type{
								Name:   domain.Int,
								Length: 11,
							},
						},
						Reference: nil,
						Extra: &domain.Extra{
							AutoIncrement: true,
							Format:        domain.ColumnFormatDefault,
							Storage:       domain.StorageDefault,
						},
					},
					{
						Column: &domain.Column{
							ID:          "user_id",
							Description: "ユーザーID",
							Name:        "user_id",
							Type: domain.Type{
								Name:   domain.Int,
								Length: 11,
							},
						},
						Reference: []*models.Reference{
							{
								Reference: &domain.Reference{
									Table:    "user",
									Match:    domain.Full,
									OnDelete: domain.Restrict,
									OnUpdate: domain.Restrict,
								},
								Columns: []*domain.IndexColumn{
									{
										Name:      "id",
										Length:    0,
										Direction: domain.IndexDirectionDefault,
									},
								},
							},
						},
						Extra: nil,
					},
					{
						Column: &domain.Column{
							ID:          "body",
							Description: "メッセージ本体",
							Name:        "body",
							Type: domain.Type{
								Name:   domain.Text,
								Length: 0,
							},
						},
						Reference: nil,
						Extra:     nil,
					},
					{
						Column: &domain.Column{
							ID:          "created_at",
							Description: "メッセージ作成日時",
							Name:        "created_at",
							Type: domain.Type{
								Name: domain.Datetime,
							},
						},
						Reference: nil,
						Extra:     nil,
					},
					{
						Column: &domain.Column{
							ID:          "deleted_at",
							Description: "メッセージ削除日時",
							Name:        "deleted_at",
							Type: domain.Type{
								Name: domain.Datetime,
							},
						},
						Reference: nil,
						Extra:     nil,
					},
				},
			},
		},
	}

	assertion.Equal(expectConfig, config)
}
