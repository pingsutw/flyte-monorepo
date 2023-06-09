package api

import (
	"bytes"
	"text/template"
)

func GenerateCodeFile(buffer *bytes.Buffer, info TypeInfo) error {
	return mainTmpl.Execute(buffer, info)
}

func GenerateTestFile(buffer *bytes.Buffer, info TypeInfo) error {
	return testTmpl.Execute(buffer, info)
}

var mainTmpl = template.Must(template.New("MainFile").Parse(
	`// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots.

package {{ .Package }}

import (
	"encoding/json"

	"github.com/spf13/pflag"
	"fmt"
{{range $path, $name := .Imports}}
	{{$name}} "{{$path}}"{{end}}
)

// If v is a pointer, it will get its element value or the zero value of the element type.
// If v is not a pointer, it will return it as is.
func ({{ .Name }}) elemValueOrNil(v interface{}) interface{} {
	if t := reflect.TypeOf(v); t.Kind() == reflect.Ptr {
		if reflect.ValueOf(v).IsNil() {
			return reflect.Zero(t.Elem()).Interface()
		} else {
			return reflect.ValueOf(v).Interface()
		}
	} else if v == nil {
		return reflect.Zero(t).Interface()
	}

	return v
}

func ({{ .Name }}) mustJsonMarshal(v interface{}) string {
	raw, err := json.Marshal(v)
    if err != nil {
        panic(err)
    }

    return string(raw)
}

func ({{ .Name }}) mustMarshalJSON(v json.Marshaler) string {
    raw, err := v.MarshalJSON()
    if err != nil {
        panic(err)
    }

    return string(raw)
}

// GetPFlagSet will return strongly types pflags for all fields in {{ .Name }} and its nested types. The format of the
// flags is json-name.json-sub-name... etc.
func (cfg {{ .Name }}) GetPFlagSet(prefix string) *pflag.FlagSet {
	cmdFlags := pflag.NewFlagSet("{{ .Name }}", pflag.ExitOnError)
	{{- range .Fields }}
	{{- if eq .FlagMethodName "" }}
	{{- if .ShouldBindDefault }}
	cmdFlags.{{ .FlagMethodName }}Var(&{{ .DefaultValue }}, fmt.Sprintf("%v%v", prefix, "{{ .Name }}"), {{ .UsageString }})
	{{- else }}
	cmdFlags.{{ .FlagMethodName }}(fmt.Sprintf("%v%v", prefix, "{{ .Name }}"), {{ .UsageString }})
	{{- end }}
	{{- else }}
	{{- if .ShouldBindDefault }}
	cmdFlags.{{ .FlagMethodName }}Var(&{{ .DefaultValue }}, fmt.Sprintf("%v%v", prefix, "{{ .Name }}"), {{ .DefaultValue }}, {{ .UsageString }})
	{{- else }}
	cmdFlags.{{ .FlagMethodName }}(fmt.Sprintf("%v%v", prefix, "{{ .Name }}"), {{ .DefaultValue }}, {{ .UsageString }})
	{{- end }}
	{{- end }}
	{{- end }}
	return cmdFlags
}
`))

var testTmpl = template.Must(template.New("TestFile").Parse(
	`// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots.

package {{ .Package }}

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
{{- range $path, $name := .Imports}}
	{{$name}} "{{$path}}"
{{- end}}
)

var dereferencableKinds{{ .Name }} = map[reflect.Kind]struct{}{
	reflect.Array: {}, reflect.Chan: {}, reflect.Map: {}, reflect.Ptr: {}, reflect.Slice: {},
}

// Checks if t is a kind that can be dereferenced to get its underlying type.
func canGetElement{{ .Name }}(t reflect.Kind) bool {
	_, exists := dereferencableKinds{{ .Name }}[t]
	return exists
}

// This decoder hook tests types for json unmarshaling capability. If implemented, it uses json unmarshal to build the
// object. Otherwise, it'll just pass on the original data.
func jsonUnmarshalerHook{{ .Name }}(_, to reflect.Type, data interface{}) (interface{}, error) {
	unmarshalerType := reflect.TypeOf((*json.Unmarshaler)(nil)).Elem()
	if to.Implements(unmarshalerType) || reflect.PtrTo(to).Implements(unmarshalerType) ||
		(canGetElement{{ .Name }}(to.Kind()) && to.Elem().Implements(unmarshalerType)) {

		raw, err := json.Marshal(data)
		if err != nil {
			fmt.Printf("Failed to marshal Data: %v. Error: %v. Skipping jsonUnmarshalHook", data, err)
			return data, nil
		}

		res := reflect.New(to).Interface()
		err = json.Unmarshal(raw, &res)
		if err != nil {
			fmt.Printf("Failed to umarshal Data: %v. Error: %v. Skipping jsonUnmarshalHook", data, err)
			return data, nil
		}

		return res, nil
	}

	return data, nil
}

func decode_{{ .Name }}(input, result interface{}) error {
	config := &mapstructure.DecoderConfig{
		TagName:          "json",
		WeaklyTypedInput: true,
		Result:           result,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToSliceHookFunc(","),
			jsonUnmarshalerHook{{ .Name }},
		),
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}

	return decoder.Decode(input)
}

func join_{{ .Name }}(arr interface{}, sep string) string {
	listValue := reflect.ValueOf(arr)
	strs := make([]string, 0, listValue.Len())
	for i := 0; i < listValue.Len(); i++ {
		strs = append(strs, fmt.Sprintf("%v", listValue.Index(i)))
	}

	return strings.Join(strs, sep)
}

func testDecodeJson_{{ .Name }}(t *testing.T, val, result interface{}) {
	assert.NoError(t, decode_{{ .Name }}(val, result))
}

func testDecodeRaw_{{ .Name }}(t *testing.T, vStringSlice, result interface{}) {
	assert.NoError(t, decode_{{ .Name }}(vStringSlice, result))
}

func Test{{ .Name }}_GetPFlagSet(t *testing.T) {
	val := {{ .Name }}{}
	cmdFlags := val.GetPFlagSet("")
	assert.True(t, cmdFlags.HasFlags())
}

func Test{{ .Name }}_SetFlags(t *testing.T) {
	actual := {{ .Name }}{}
	cmdFlags := actual.GetPFlagSet("")
	assert.True(t, cmdFlags.HasFlags())

	{{ $ParentName := .Name }}
	{{- range .Fields }}
	t.Run("Test_{{ .Name }}", func(t *testing.T) { {{ $varName := print "v" .FlagMethodName }}
		{{- if .ShouldTestDefault }}
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if {{ $varName }}, err := cmdFlags.Get{{ .FlagMethodName }}("{{ .Name }}"); err == nil {
				assert.Equal(t, {{ .Typ }}({{ .DefaultValue }}), {{ $varName }})
			} else {
				assert.FailNow(t, err.Error())
			}
		})
		{{- end }}

		t.Run("Override", func(t *testing.T) {
			{{ if eq .TestStrategy "Json" }}testValue := {{ .TestValue }}
			{{ else if eq .TestStrategy "Raw" }}testValue := {{ .TestValue }}
			{{ else }}testValue := join_{{ $ParentName }}({{ .TestValue }}, ",")
			{{ end }}
			cmdFlags.Set("{{ .Name }}", testValue)
			{{- if eq .FlagMethodName "" }}
			if {{ $varName }} := cmdFlags.Lookup("{{ .Name }}"); {{ $varName }} != nil {
				{{ if eq .TestStrategy "Json" }}testDecodeJson_{{ $ParentName }}(t, fmt.Sprintf("%v", v.Value.String()), &actual.{{ .GoName }})
				{{ else if eq .TestStrategy "Raw" }}testDecodeRaw_{{ $ParentName }}(t, v.Value.String(), &actual.{{ .GoName }})
				{{ else }}testDecodeRaw_{{ $ParentName }}(t, join_{{ $ParentName }}({{ print "v" .FlagMethodName }}, ",").Value.String(), &actual.{{ .GoName }})
				{{ end }}
			}
			{{- else }}
			if {{ $varName }}, err := cmdFlags.Get{{ .TestFlagMethodName }}("{{ .Name }}"); err == nil {
				{{ if eq .TestStrategy "Json" }}testDecodeJson_{{ $ParentName }}(t, fmt.Sprintf("%v", {{ print "v" .FlagMethodName }}), &actual.{{ .GoName }})
				{{ else if eq .TestStrategy "Raw" }}testDecodeRaw_{{ $ParentName }}(t, {{ print "v" .FlagMethodName }}, &actual.{{ .GoName }})
				{{ else }}testDecodeRaw_{{ $ParentName }}(t, join_{{ $ParentName }}({{ print "v" .FlagMethodName }}, ","), &actual.{{ .GoName }})
				{{ end }}
			} else {
				assert.FailNow(t, err.Error())
			}
			{{- end }}
		})
	})
	{{- end }}
}
`))
