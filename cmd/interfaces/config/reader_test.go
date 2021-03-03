package config

import (
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mazrean/gopendb-generator/cmd/domain"
	"github.com/mazrean/gopendb-generator/cmd/interfaces/config/models"
	"github.com/mazrean/gopendb-generator/cmd/interfaces/reader/mock_reader"
	conf "github.com/mazrean/gopendb-generator/cmd/usecases/config"
	"github.com/stretchr/testify/assert"
)

func readYAMLTest(t *testing.T) {
	t.Helper()

	assertion := assert.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fileReader := mock_reader.NewMockReader(ctrl)

	reader := NewReader(fileReader)

	type arg struct {
		path string
	}
	type mock struct {
		isExist  bool
		yamlBody string
	}
	type expect struct {
		config models.Root
		isErr  bool
		err    error
	}
	type test struct {
		description string
		arg
		mock
		expect
	}

	testCases := []test{
		{
			description: "normal",
			arg: arg{
				path: "test.yaml",
			},
			mock: mock{
				isExist: true,
				yamlBody: `config:
  dbms: mysql
  version: 8.0
  database: test
tables:
  - id: user
    description: ユーザーテーブル
    name: users
    engin: InnoDB
    char_set: utf8mb4
    max_rows: 100000
    min_rows: 10
    avg_row_length: 100
    primary_key:
      columns:
        - name: id
          length: 10
          direction: asc
      options:
        type: btree
        key_block_size: 0
        parser: my_parser
    columns:
      - id: id
        description: ユーザーID
        name: id
        type:
          name: int
          length: 11
        extra:
          auto_increment: true
          format: dynamic
          storage: disk
      - id: name
        description: ユーザー名
        name: name
        type:
          name: text
      - id: created_at
        description: ユーザー作成日時
        name: created_at
        type:
          name: datetime
      - id: deleted_at
        description: ユーザー削除日時
        name: deleted_at
        type:
          name: datetime
  - id: message
    description: メッセージテーブル
    name: messages
    primary_key:
      columns:
        - name: id
    columns:
      - id: id
        description: メッセージID
        name: id
        type:
          name: int
          length: 11
        extra:
          auto_increment: true
      - id: user_id
        description: ユーザーID
        name: user_id
        type:
          name: int
          length: 11
        reference:
          - table: user
            match: full
            on_delete: restrict
            on_update: restrict
            columns:
              - name: id
      - id: body
        description: メッセージ本体
        name: body
        type:
          name: text
      - id: created_at
        description: メッセージ作成日時
        name: created_at
        type:
          name: datetime
      - id: deleted_at
        description: メッセージ削除日時
        name: deleted_at
        type:
          name: datetime`,
			},
			expect: expect{
				config: models.Root{
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
				},
			},
		},
		{
			description: "file not found",
			arg: arg{
				path: "test.yaml",
			},
			mock: mock{
				isExist: false,
			},
			expect: expect{
				isErr: true,
				err:   conf.ErrFileNotFound,
			},
		},
	}

	for _, testCase := range testCases {
		fileReader.EXPECT().IsExist(testCase.arg.path).Return(testCase.mock.isExist, nil).Times(1)
		if testCase.mock.isExist {
			fileReader.EXPECT().Read(testCase.arg.path).Return(io.NopCloser(strings.NewReader(testCase.mock.yamlBody)), nil)
		}

		err := reader.ReadYAML(testCase.arg.path)

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

		assertion.Equalf(testCase.expect.config, config, testCase.description, "config")
	}
}
