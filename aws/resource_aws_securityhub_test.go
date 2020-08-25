package aws

import (
	"testing"
)

func TestAccAWSSecurityHub_serial(t *testing.T) {
	testCases := map[string]map[string]func(t *testing.T){
		"Account": {
			"basic": testAccAWSSecurityHubAccount_basic,
		},
		"Member": {
			"basic":  testAccAWSSecurityHubMember_basic,
			"invite": testAccAWSSecurityHubMember_invite,
		},
		"CustomAction": {
			"basic": testAccAwsSecurityHubActionTarget_basic,
		},
		"ProductSubscription": {
			"basic": testAccAWSSecurityHubProductSubscription_basic,
		},
		"StandardsSubscription": {
			"basic": testAccAWSSecurityHubStandardsSubscription_basic,
		},
	}

	for group, m := range testCases {
		m := m
		t.Run(group, func(t *testing.T) {
			for name, tc := range m {
				tc := tc
				t.Run(name, func(t *testing.T) {
					tc(t)
				})
			}
		})
	}
}
