// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"github.com/GoBotApiOfficial/gobotapi/types"
	rawTypes "github.com/GoBotApiOfficial/gobotapi/types/raw"
	"reflect"
)

// DeleteMyCommands Use this method to delete the list of the bot's commands for the given scope and user language
// After deletion, higher level commands will be shown to affected users
// Returns True on success.
type DeleteMyCommands struct {
	LanguageCode string                 `json:"language_code,omitempty"`
	Scope        *types.BotCommandScope `json:"scope,omitempty"`
}

func (entity *DeleteMyCommands) ProgressCallable() rawTypes.ProgressCallable {
	return nil
}

func (entity *DeleteMyCommands) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (entity DeleteMyCommands) MarshalJSON() ([]byte, error) {
	nilCheck := func(val any) bool {
		if val == nil {
			return true
		}
		v := reflect.ValueOf(val)
		k := v.Kind()
		switch k {
		case reflect.Chan, reflect.Func, reflect.Map, reflect.Pointer, reflect.UnsafePointer, reflect.Interface, reflect.Slice:
			return v.IsNil()
		default:
			return false
		}
	}
	_ = nilCheck
	if nilCheck(entity.Scope) {
		entity.Scope = nil
	}
	type x0 DeleteMyCommands
	return json.Marshal((x0)(entity))
}

func (DeleteMyCommands) MethodName() string {
	return "deleteMyCommands"
}

func (DeleteMyCommands) ParseResult(response []byte) (*rawTypes.Result, error) {
	var x1 struct {
		Result bool `json:"result"`
	}
	err := json.Unmarshal(response, &x1)
	if err != nil {
		return nil, err
	}
	result := rawTypes.Result{
		Kind:   types.TypeBoolean,
		Result: x1.Result,
	}
	return &result, nil
}
