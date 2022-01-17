package db

import (
	"database/sql"
	"fmt"
	"strings"
)

// FieldAttr
type FieldAttr struct {
	Field   string
	Type    sql.RawBytes
	Null    sql.RawBytes
	Key     sql.RawBytes
	Default sql.RawBytes
	Extra   sql.RawBytes
}

// CamelField
func (t FieldAttr) CamelField() (field string) {
	parts := strings.Split(string(t.Field), "_")
	for _, part := range parts {
		if part == "id" {
			field += "ID"
		} else {
			field += strings.Title(part)
		}
	}
	return
}

// FieldType
func (t FieldAttr) FieldType() string {
	var prefix string
	if string(t.Null) == "YES" {
		prefix = "*"
	}
	if strings.Contains(string(t.Type), "char") || strings.Contains(string(t.Type), "text") {
		return prefix + "string"
	}

	if strings.Contains(string(t.Type), "datetime") {
		return prefix + "time.Time"
	}

	return prefix + "int64"
}

// DBAttr
type DBAttr struct {
	ProviderName string
	TableName    string
	Package      string
	TableFields  []FieldAttr
}

// Table
func (d *DBAttr) Table() (name string) {
	idx := strings.LastIndex(d.TableName, "_tab")
	table := d.TableName
	if idx != -1 {
		table = table[:idx]
	}

	parts := strings.Split(table, "_")
	for _, part := range parts {
		name += strings.Title(part)
	}

	return name
}

// LowTable
func (d *DBAttr) LowTable() (name string) {
	idx := strings.LastIndex(d.TableName, "_tab")
	table := d.TableName
	if idx != -1 {
		table = table[:idx]
	}

	parts := strings.Split(table, "_")
	for i, part := range parts {
		if i != 0 {
			name += strings.Title(part)
		} else {
			name += part
		}
	}

	return name
}

// Inject
func (d *DBAttr) Inject() string {
	return "`"
}

// QueryDBAttr
func QueryDBAttr(conn *Connection, tableName string, packageName string) *DBAttr {
	attr := &DBAttr{ProviderName: conn.Name, TableName: tableName, Package: packageName}
	fmt.Printf("----- %v\n", *attr)
	row, err := conn.Query(fmt.Sprintf("desc %v;", tableName))
	if err != nil {
		panic(err)
	}

	// var tables []TableFileds
	for row.Next() {
		var table FieldAttr
		if err := row.Scan(&table.Field, &table.Type, &table.Null, &table.Key, &table.Default, &table.Extra); err != nil {
			panic(err)
		}

		attr.TableFields = append(attr.TableFields, table)
	}
	return attr
}
