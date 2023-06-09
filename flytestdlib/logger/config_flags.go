// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots.

package logger

import (
	"encoding/json"
	"reflect"

	"fmt"

	"github.com/spf13/pflag"
)

// If v is a pointer, it will get its element value or the zero value of the element type.
// If v is not a pointer, it will return it as is.
func (Config) elemValueOrNil(v interface{}) interface{} {
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

func (Config) mustJsonMarshal(v interface{}) string {
	raw, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	return string(raw)
}

func (Config) mustMarshalJSON(v json.Marshaler) string {
	raw, err := v.MarshalJSON()
	if err != nil {
		panic(err)
	}

	return string(raw)
}

// GetPFlagSet will return strongly types pflags for all fields in Config and its nested types. The format of the
// flags is json-name.json-sub-name... etc.
func (cfg Config) GetPFlagSet(prefix string) *pflag.FlagSet {
	cmdFlags := pflag.NewFlagSet("Config", pflag.ExitOnError)
	cmdFlags.Bool(fmt.Sprintf("%v%v", prefix, "show-source"), defaultConfig.IncludeSourceCode, "Includes source code location in logs.")
	cmdFlags.Bool(fmt.Sprintf("%v%v", prefix, "mute"), defaultConfig.Mute, "Mutes all logs regardless of severity. Intended for benchmarks/tests only.")
	cmdFlags.Int(fmt.Sprintf("%v%v", prefix, "level"), defaultConfig.Level, "Sets the minimum logging level.")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "formatter.type"), defaultConfig.Formatter.Type, "Sets logging format type.")
	return cmdFlags
}
