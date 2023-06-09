// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots.

package config

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
)

var dereferencableKindsConfig = map[reflect.Kind]struct{}{
	reflect.Array: {}, reflect.Chan: {}, reflect.Map: {}, reflect.Ptr: {}, reflect.Slice: {},
}

// Checks if t is a kind that can be dereferenced to get its underlying type.
func canGetElementConfig(t reflect.Kind) bool {
	_, exists := dereferencableKindsConfig[t]
	return exists
}

// This decoder hook tests types for json unmarshaling capability. If implemented, it uses json unmarshal to build the
// object. Otherwise, it'll just pass on the original data.
func jsonUnmarshalerHookConfig(_, to reflect.Type, data interface{}) (interface{}, error) {
	unmarshalerType := reflect.TypeOf((*json.Unmarshaler)(nil)).Elem()
	if to.Implements(unmarshalerType) || reflect.PtrTo(to).Implements(unmarshalerType) ||
		(canGetElementConfig(to.Kind()) && to.Elem().Implements(unmarshalerType)) {

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

func decode_Config(input, result interface{}) error {
	config := &mapstructure.DecoderConfig{
		TagName:          "json",
		WeaklyTypedInput: true,
		Result:           result,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToSliceHookFunc(","),
			jsonUnmarshalerHookConfig,
		),
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}

	return decoder.Decode(input)
}

func join_Config(arr interface{}, sep string) string {
	listValue := reflect.ValueOf(arr)
	strs := make([]string, 0, listValue.Len())
	for i := 0; i < listValue.Len(); i++ {
		strs = append(strs, fmt.Sprintf("%v", listValue.Index(i)))
	}

	return strings.Join(strs, sep)
}

func testDecodeJson_Config(t *testing.T, val, result interface{}) {
	assert.NoError(t, decode_Config(val, result))
}

func testDecodeRaw_Config(t *testing.T, vStringSlice, result interface{}) {
	assert.NoError(t, decode_Config(vStringSlice, result))
}

func TestConfig_GetPFlagSet(t *testing.T) {
	val := Config{}
	cmdFlags := val.GetPFlagSet("")
	assert.True(t, cmdFlags.HasFlags())
}

func TestConfig_SetFlags(t *testing.T) {
	actual := Config{}
	cmdFlags := actual.GetPFlagSet("")
	assert.True(t, cmdFlags.HasFlags())

	t.Run("Test_pod-application", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("pod-application", testValue)
			if vString, err := cmdFlags.GetString("pod-application"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.PodApplication)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_pod-template-container-name", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("pod-template-container-name", testValue)
			if vString, err := cmdFlags.GetString("pod-template-container-name"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.PodTemplateContainerName)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_pod-template-name", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("pod-template-name", testValue)
			if vString, err := cmdFlags.GetString("pod-template-name"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.PodTemplateName)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_pod-template-namespace", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("pod-template-namespace", testValue)
			if vString, err := cmdFlags.GetString("pod-template-namespace"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.PodTemplateNamespace)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_scan-interval", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := DefaultConfig.ScanInterval.String()

			cmdFlags.Set("scan-interval", testValue)
			if vString, err := cmdFlags.GetString("scan-interval"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.ScanInterval)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_shard.type", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("shard.type", testValue)
			if vString, err := cmdFlags.GetString("shard.type"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.ShardConfig.Type)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_shard.shard-count", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("shard.shard-count", testValue)
			if vInt, err := cmdFlags.GetInt("shard.shard-count"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vInt), &actual.ShardConfig.ShardCount)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
}
