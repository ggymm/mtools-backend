package model
${$iLen := len .Imports}
${if gt $iLen 0}
import (
    ${range .Imports}"${.}"${end}
)
${end}

type ${TableMapper .Name} struct {
${range .fields}
    ${$col := $table.GetColumn .}	${ColumnMapper $col.Name}	${Type $col} `json:"${UnTitle (ColumnMapper $col.Name)}" ${Tag $table $col}`
${end}
}

func (m *${TableMapper .Name}) TableName() string {
    return "${$table.Name}"
}


