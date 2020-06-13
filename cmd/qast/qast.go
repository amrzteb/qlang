package main

import (
	"fmt"
	"go/constant"
	"go/types"
)

var (
	pkg = types.NewPackage("github.com/qiniu/foo", "foo")
)

type myBool bool
type aliasBool = bool

func udbt() { // user define basic type
	tyBool := types.Typ[types.Bool]
	fmt.Println(tyBool)

	tyUntypedBool := types.Typ[types.UntypedBool]
	fmt.Println(tyUntypedBool)

	tyDefBool := types.Default(tyUntypedBool)
	fmt.Println("default of untyped bool:", tyDefBool)
	if tyBool != tyDefBool {
		fmt.Println("Failed: default of untyped bool != bool", tyDefBool)
	}

	myBoolName := types.NewTypeName(0, pkg, "myBool", nil)
	tyMyBool := types.NewNamed(myBoolName, tyBool, nil)
	fmt.Println(tyMyBool)
	fmt.Println("tyMyBool name:", myBoolName)

	if types.AssignableTo(tyMyBool, tyBool) {
		fmt.Println("Failed: assign tyMyBool to tyBool")
	}
	if types.AssignableTo(tyUntypedBool, tyMyBool) {
		fmt.Println("OK: assign tyUntypedBool to tyMyBool")
	}

	aliasBool := types.NewTypeName(0, pkg, "aliasBool", tyBool)
	fmt.Println("aliasBool name:", aliasBool, "isAlias:", aliasBool.IsAlias())

	if types.AssignableTo(tyUntypedBool, aliasBool.Type()) {
		fmt.Println("OK: assign tyUntypedBool to aliasBool")
	}
	if types.AssignableTo(aliasBool.Type(), tyBool) {
		fmt.Println("OK: assign aliasBool to tyBool")
	}

	aliasMyBool := types.NewTypeName(0, pkg, "aliasMyBool", tyMyBool)
	fmt.Println("aliasMyBool name:", aliasMyBool, "isAlias:", aliasMyBool.IsAlias())

	if types.AssignableTo(tyUntypedBool, aliasMyBool.Type()) {
		fmt.Println("OK: assign tyUntypedBool to aliasMyBool")
	}
	if types.AssignableTo(aliasMyBool.Type(), tyMyBool) {
		fmt.Println("OK: assign aliasMyBool to tyMyBool")
	}
	if types.AssignableTo(aliasMyBool.Type(), tyBool) {
		fmt.Println("Failed: assign aliasMyBool to tyBool")
	}
}

func testConst() {
	tyFloat64 := types.Typ[types.Float64]
	piVal := constant.MakeFloat64(3.14)
	piTypedConst := types.NewConst(0, pkg, "Pi", tyFloat64, piVal)
	fmt.Println(
		piTypedConst, "id:", piTypedConst.Id(),
		"val:", piTypedConst.Val(), "type:", piTypedConst.Type())

	tyUntypedFloat := types.Typ[types.UntypedFloat]
	piUntypedConst := types.NewConst(0, pkg, "Pi2", tyUntypedFloat, piVal)
	fmt.Println(
		piUntypedConst, "id:", piUntypedConst.Id(),
		"val:", piUntypedConst.Val(), "type:", piUntypedConst.Type())
}

func main() {
	testConst()
	udbt()
}
