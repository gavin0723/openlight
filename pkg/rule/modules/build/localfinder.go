// Author: lipixun
// File Name: localfinder.go
// Description:

package build

import (
	"fmt"

	"github.com/yuin/gopher-lua"

	pbSpec "github.com/ops-openlight/openlight/protoc-gen-go/spec"

	LUA "github.com/ops-openlight/openlight/pkg/rule/modules/lua"
)

// Exposed lua infos
const (
	LUATypePythonLocalFinder = "Build-LocalFinder-Python"
)

// LocalFinderLUAFuncs defines all lua functions for local finder
var LocalFinderLUAFuncs = map[string]lua.LGFunction{
	"name":    LUALocalFinderName,
	"options": LUA.FuncObjectOptions,
}

// PythonLocalFinderLUAFuncs defines all lua functions for python local finder
var PythonLocalFinderLUAFuncs = map[string]lua.LGFunction{}

// RegisterPythonLocalFinderType registers local finder type in lua
func RegisterPythonLocalFinderType(L *lua.LState, mod *lua.LTable) {
	mt := L.NewTypeMetatable(LUATypePythonLocalFinder)
	var funcs = make(map[string]lua.LGFunction)
	for name, function := range LocalFinderLUAFuncs {
		funcs[name] = function
	}
	for name, function := range PythonLocalFinderLUAFuncs {
		funcs[name] = function
	}
	L.SetField(mt, "__index", L.SetFuncs(L.NewTable(), funcs))
}

// LocalFinder represents the local finder
type LocalFinder interface {
	LUA.Object
	// GetName returns name
	GetName() string
	// GetProto returns the protobuf object
	GetProto() (*pbSpec.LocalFinder, error)
}

// PythonLocalFinder represents python local finder
type PythonLocalFinder struct {
	LUA.Object
	Name   string
	Module string
}

// NewPythonLocalFinder creates a new PythonLocalFinder
func NewPythonLocalFinder(name, module string, options *lua.LTable) *PythonLocalFinder {
	var finder = PythonLocalFinder{
		Name:   name,
		Module: module,
	}
	finder.Object = LUA.NewObject(LUATypePythonLocalFinder, options, LocalFinder(&finder))
	// Done
	return &finder
}

// GetName returns name
func (finder *PythonLocalFinder) GetName() string {
	return finder.Name
}

// GetProto returns the protobuf object
func (finder *PythonLocalFinder) GetProto() (*pbSpec.LocalFinder, error) {
	// Get options
	var parent = 0
	if options := finder.GetOptions(); options != nil {
		var err error
		if parent, err = LUA.GetIntFromTable(options, "parent"); err != nil {
			return nil, fmt.Errorf("Failed to get options [parent], error: %s", err)
		}
	}
	// Create protobuf object
	var pbFinder pbSpec.LocalFinder
	pbFinder.Name = finder.Name
	pbFinder.Finder = &pbSpec.LocalFinder_Python{
		Python: &pbSpec.PythonLocalFinder{
			Module: finder.Module,
			Parent: int32(parent),
		},
	}
	// Done
	return &pbFinder, nil
}

//////////////////////////////////////// LUA functions ////////////////////////////////////////

// LUALocalFinderSelf get lua local finder self
func LUALocalFinderSelf(L *lua.LState) LocalFinder {
	ud := L.CheckUserData(1)
	if ref, ok := ud.Value.(LocalFinder); ok {
		return ref
	}
	L.ArgError(1, "LocalFinder expected")
	return nil
}

// LUALocalFinderName defines LocalFinder.name function in lua
func LUALocalFinderName(L *lua.LState) int {
	finder := LUALocalFinderSelf(L)
	if finder == nil {
		return 0
	}
	if L.GetTop() != 1 {
		L.ArgError(2, "Invalid arguments")
		return 0
	}
	L.Push(lua.LString(finder.GetName()))
	return 1
}
