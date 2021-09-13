package tpl

import (
	"strings"
	"text/template"
)

func Funcs() template.FuncMap {
	result := template.FuncMap{}
	result["SmallCamel"] = SmallCamel
	result["BigCamel"] = BigCamel
	result["FormatJavaDataType"] = FormatJavaDataType
	return result
}

func SmallCamel(str string) string {
	data := make([]byte, 0, len(str))
	j := false
	k := true
	num := len(str) - 1
	for i := 0; i <= num; i++ {
		d := str[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && str[i+1] >= 'a' && str[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}

func BigCamel(str string) string {
	str = strings.Replace(str, "_", " ", -1)
	str = strings.Title(str)
	return strings.Replace(str, " ", "", -1)
}

func FormatJavaDataType(str string) string {
	str = strings.ToLower(str)
	types := map[string]string{
		"int":        "Integer",
		"tinyint":    "Byte",
		"smallint":   "Short",
		"mediumint":  "Integer",
		"bigint":     "Long",
		"bit":        "Boolean",
		"double":     "Double",
		"float":      "Float",
		"number":     "java.math.BigDecimal",
		"decimal":    "java.math.BigDecimal",
		"char":       "String",
		"varchar":    "String",
		"date":       "java.util.Date",
		"time":       "java.util.Date",
		"year":       "java.util.Date",
		"timestamp":  "java.util.Date",
		"datetime":   "java.util.Date",
		"tinyblob":   "Byte[]",
		"blob":       "Byte[]",
		"mediumblob": "Byte[]",
		"longblob":   "Byte[]",
		"tinytext":   "String",
		"text":       "String",
		"mediumtext": "String",
		"longtext":   "String",
		"binary":     "Byte[]",
		"varbinary":  "Byte[]",
	}
	return types[str]
}
