package dbr

var mockTmpl = `// Code generated by gen-db. DO NOT EDIT.

package {{.Package}}

import (
	"reflect"
	"testing"
	"time"
)
// holder start
type fields struct {
	{{range .TableFields}} {{.CamelField}} {{.FieldType}} 
	{{end}}
	updateInfo map[string]interface{}
}
// holder end

// fields start
{{range .TableFields}}{{.CamelField}}: tt.fields.{{.CamelField}}, 
{{end}}
// fields end

{{range .TableFields}}
func Test{{$.Table}}Entity_Set{{.CamelField}}(t *testing.T) {
	placeholder
	type args struct {
		v interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *{{$.Table}}Entity
	}{
		{
			name: "normal",
			args: args{v: 0},
			want: &{{$.Table}}Entity{updateInfo: map[string]interface{}{
				"{{.Field}}": 0,
			}},
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &{{$.Table}}Entity{
				fields-holder
				updateInfo: tt.fields.updateInfo,
			}
			if got := e.Set{{.CamelField}}(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Set{{.CamelField}}() = %v, want %v", got, tt.want)
			}
		})
	}
}
{{end}}

`
