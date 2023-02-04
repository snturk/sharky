package security

import (
	"github.com/voicera/gooseberry/urn"
	"log"
	"sharky/model"
	"testing"
)

// Example Role & action definitions
var (
	ACTION_VIEW   = model.SharkyAction{Name: "VIEW"}
	ACTION_DELETE = model.SharkyAction{Name: "DELETE"}
)

const NamespaceId = "sharky:obj"

var logger = log.New(log.Writer(), "TestEnforcer_Enforce ", log.LstdFlags)

// TestEnforcer_Enforce tests the Enforcer.Enforce() method
// to ensure that it returns true for a valid permission
// and false for an invalid permission
func TestEnforcer_Enforce(t *testing.T) {
	// Test a permission that should be allowed

	objUrn := *urn.NewURN(NamespaceId, "testobj:1234")
	obj := model.SharkyDomain{}.Builder().SetName("TestObject").SetUrn(objUrn).Build()
	perm := model.Permission{}.Builder().SetUrn(objUrn).SetAction(ACTION_VIEW).Build()
	role := model.SharkyRole{
		Name:        "TestRole",
		Permissions: []model.Permission{perm},
	}

	context := model.NewPermissionDecisionContext(role, ACTION_VIEW, obj)

	enforcer, err := GetEnforcer(*logger)
	if err != nil {
		t.Errorf("Error getting enforcer: %s", err)
	}

	if !enforcer.Enforce(context) {
		t.Errorf("Enforcer did not return true for valid permission")
	} else {
		logger.Printf("Enforcer returned true for valid permission")
	}

	// Test a permission that should be denied
	context = model.NewPermissionDecisionContext(role, ACTION_DELETE, obj)
	if enforcer.Enforce(context) {
		t.Errorf("Enforcer did not return false for invalid permission")
	} else {
		logger.Printf("Enforcer returned false for invalid permission")
	}
}
