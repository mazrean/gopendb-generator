package config

import (
	conf "github.com/mazrean/gopendb-generator/cmd/usecases/config"
)

// Table テーブル関連の情報取り出しの構造体
type Table struct{}

// NewTable Tableのコンストラクタ
func NewTable() *Table {
	return &Table{}
}

// GetAll 全テーブル情報の取得
func (*Table) GetAll() []*conf.TableDetail {
	tableDetails := make([]*conf.TableDetail, 0, len(config.Tables))

	for _, table := range config.Tables {
		var primaryKeycolumnNames []string
		if table != nil && table.PrimaryKey != nil && table.PrimaryKey.Columns != nil {
			primaryKeycolumnNames = make([]string, 0, len(table.PrimaryKey.Columns))
			for _, column := range table.PrimaryKey.Columns {
				primaryKeycolumnNames = append(primaryKeycolumnNames, column.Name)
			}
		} else {
			primaryKeycolumnNames = []string{}
		}

		tableDetails = append(tableDetails, &conf.TableDetail{
			Table:                 table.Table,
			PrimaryKeyColumnNames: primaryKeycolumnNames,
		})
	}

	return tableDetails
}
