package main

import "testing"

func TestGithubEventString_ToString(t *testing.T) {
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

	hookFlag = &hookURLFlag{}
	hookFlag.Set("ValidID@FakeURL")
	hookFlag.Set("SecondValidID@FakeURL")

	hookArray, err = hookFlag.GetArrayOfDataHook()
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
