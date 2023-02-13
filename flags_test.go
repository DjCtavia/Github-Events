package main

import (
	"fmt"
	"testing"
)

func TestHookURLFlag_Set(t *testing.T) {
	hookFlag := &hookURLFlag{}
	err := hookFlag.Set("invalid")

	if err != nil {
		t.Errorf("didn't expected error got: %s", err)
	}
}

func TestHookURLFlag_Set2(t *testing.T) {
	hookFlag := &hookURLFlag{}
	err := hookFlag.Set("ValidID@FakeURL")
	if err != nil {
		t.Errorf("didn't expected error got: %s", err)
	}

	err = hookFlag.Set("SecondValidID@FakeURL")
	if err != nil {
		t.Errorf("didn't expected error got: %s", err)
	}
}

func TestHookURLFlag_GetArrayOfDataHook(t *testing.T) {
	hookFlag := &hookURLFlag{}
	hookFlag.Set("invalid")

	_, err := hookFlag.GetArrayOfDataHook()
	if err == nil {
		t.Error("expected an error, got nil")
	}

	hookFlag = &hookURLFlag{}
	hookFlag.Set("ValidID@FakeURL")

	hookArray, err := hookFlag.GetArrayOfDataHook()
	if err != nil {
		t.Errorf("didn't expected error got: %s", err)
	}
	if len(hookArray) != 1 {
		t.Errorf("array should be equal to 1 but is: %d", len(hookArray))
	}
}

func TestHookURLFlag_GetArrayOfDataHook2(t *testing.T) {
	hookFlag := &hookURLFlag{}
	hookFlag.Set("ValidID@FakeURL")
	hookFlag.Set("SecondValidID@FakeURL")

	hookArray, err := hookFlag.GetArrayOfDataHook()
	if err != nil {
		t.Errorf("didn't expected error got: %s", err)
	}
	if len(hookArray) != 2 {
		t.Errorf("array should be equal to 2 but is: %d", len(hookArray))
	}
	if hookArray[0].ID != "ValidID" {
		t.Errorf("expected value: ValidID, got: %s", hookArray[0].ID)
	}
	if hookArray[0].hookURL != "FakeURL" {
		t.Errorf("expected value: FakeURL, got: %s", hookArray[0].hookURL)
	}
}

func TestHookURLFlag_String(t *testing.T) {
	hookFlag := &hookURLFlag{}
	rawFlag := "realID@https://localhost"
	hookFlag.Set(rawFlag)
	stringHook := hookFlag.String()

	if stringHook != rawFlag {
		t.Errorf("expected string equal to: %s, got: %s", rawFlag, stringHook)
	}
}

func TestHookURLFlag_String2(t *testing.T) {
	hookFlag := &hookURLFlag{}
	rawFlag := "realID@https://localhost"
	hookFlag.Set(rawFlag)
	rawFlag2 := "SomethingElese@https://otherhost"
	hookFlag.Set(rawFlag2)

	expectedString := fmt.Sprintf("%s,%s", rawFlag, rawFlag2)
	stringHook := hookFlag.String()

	if stringHook != expectedString {
		t.Errorf("expected string equal to: %s, got: %s", expectedString, stringHook)
	}
}
