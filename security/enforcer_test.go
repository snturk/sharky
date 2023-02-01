package security

import (
	"fmt"
	"github.com/voicera/gooseberry/urn"
	model2 "sharky/model"
	"testing"
)

// Example Role & action definitions
var (
	ACTION_VIEW   = model2.SharkyAction{Name: "VIEW"}
	ACTION_DELETE = model2.SharkyAction{Name: "DELETE"}
)

// TestEnforcer_Enforce tests the Enforcer.Enforce() method
// to ensure that it returns true for a valid permission
// and false for an invalid permission
func TestEnforcer_Enforce(t *testing.T) {
	// Test a permission that should be allowed
	objUrn := *urn.NewURN("sharkyobj", "testobj")
	fmt.Println(objUrn)
	obj := model2.SharkyObject{}.Builder().SetName("TestObject").SetUrn(objUrn).Build()

	perm := model2.Permission{}.Builder().SetUrn(objUrn).SetAction(ACTION_VIEW).Build()

	role := model2.SharkyRole{
		Name:        "TestRole",
		Permissions: []model2.Permission{perm},
	}

	context := model2.NewPermissionDecisionContext(role, ACTION_VIEW, obj)
	enforcer, err := GetEnforcer()
	if err != nil {
		t.Errorf("Error getting enforcer: %s", err)
	}

	if !enforcer.Enforce(context) {
		t.Errorf("Enforcer did not return true for valid permission")
	} else {
		t.Logf("Enforcer returned true for valid permission")
	}

	// Test a permission that should be denied
	context = model2.NewPermissionDecisionContext(role, ACTION_DELETE, obj)
	if enforcer.Enforce(context) {
		t.Errorf("Enforcer did not return false for invalid permission")
	} else {
		t.Logf("Enforcer returned false for invalid permission")
	}
}
