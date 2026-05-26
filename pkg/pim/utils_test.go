/*
Copyright © 2024 netr0m <netr0m@pm.me>
*/
package pim

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParseDateTime(t *testing.T) {
	now := time.Now().Local()
	currentDate := now.Format("2006-01-02")
	currentTZ := now.Format("-07:00")
	errMsg := "resulting startDateTime does not match expected value"

	dateOnly, _ := parseDateTime("31/12/2024", "")
	timeOnly, _ := parseDateTime("", "13:37")
	dateTime, _ := parseDateTime("31/12/2024", "13:37")

	assert.Equal(t, fmt.Sprintf("2024-12-31T00:00:00%s", currentTZ), dateOnly, errMsg)
	assert.Equal(t, fmt.Sprintf("%sT13:37:00%s", currentDate, currentTZ), timeOnly, errMsg)
	assert.Equal(t, fmt.Sprintf("2024-12-31T13:37:00%s", currentTZ), dateTime, errMsg)
}

func TestNormalizeResourceScope(t *testing.T) {
	assert.Equal(t, "subscriptions/sub-1", NormalizeResourceScope("/subscriptions/sub-1/"))
	assert.Equal(t, "subscriptions/sub-1/resourceGroups/rg-1", NormalizeResourceScope("subscriptions/sub-1/resourceGroups/rg-1"))
}

func TestCreateResourceAssignmentRequestUsesEligibleAssignmentScope(t *testing.T) {
	resourceAssignment := &EligibleResourceAssignmentsDummyData.Value[0]
	scope, _ := CreateResourceAssignmentRequest(TEST_DUMMY_PRINCIPAL_ID, resourceAssignment, 30, "", "", "test", "Test", "1337")

	assert.Equal(t, TEST_DUMMY_SUBSCRIPTION_1_ID, scope)
}

func TestCreateResourceAssignmentRequestWithScopeOverridesEligibleAssignmentScope(t *testing.T) {
	resourceAssignment := &EligibleResourceAssignmentsDummyData.Value[0]
	activationScope := "/subscriptions/sub-1/resourceGroups/rg-1"
	scope, _ := CreateResourceAssignmentRequestWithScope(TEST_DUMMY_PRINCIPAL_ID, resourceAssignment, activationScope, 30, "", "", "test", "Test", "1337")

	assert.Equal(t, "subscriptions/sub-1/resourceGroups/rg-1", scope)
}
