// GENERATED. DO NOT MODIFY!

package postgres

import (
	"github.com/jmoiron/sqlx"

	"github.com/Peripli/service-manager/pkg/types"
	
)

func (Platform) Empty() Entity {
	return Platform{}
}

func (Platform) PrimaryColumn() string {
	return "id"
}

func (Platform) TableName() string {
	return "platforms"
}

func (e Platform) GetID() string {
	return e.ID
}

func (e Platform) Labels() EntityLabels {
    return nil
}

func (e Platform) RowsToList(rows *sqlx.Rows) (types.ObjectList, error) {
    result := &types.Platforms{}
	for rows.Next() {
		var item Platform
		if err := rows.StructScan(&item); err != nil {
			return nil, err
		}
		result.Add(item.ToObject())
	}
	return result, nil
}



