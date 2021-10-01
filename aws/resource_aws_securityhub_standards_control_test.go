package aws

import (
	"context"
	"fmt"
	"regexp"
	"testing"

	"github.com/aws/aws-sdk-go/service/securityhub"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	tfsecurityhub "github.com/hashicorp/terraform-provider-aws/aws/internal/service/securityhub"
	"github.com/hashicorp/terraform-provider-aws/aws/internal/service/securityhub/finder"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/provider"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
)

func testAccAWSSecurityHubStandardsControl_basic(t *testing.T) {
	var standardsControl securityhub.StandardsControl
	resourceName := "aws_securityhub_standards_control.test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acctest.PreCheck(t) },
		ErrorCheck:   acctest.ErrorCheck(t, securityhub.EndpointsID),
		Providers:    acctest.Providers,
		CheckDestroy: nil, //lintignore:AT001
		Steps: []resource.TestStep{
			{
				Config: testAccAWSSecurityHubStandardsControlConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckAWSSecurityHubStandardsControlExists(resourceName, &standardsControl),
					resource.TestCheckResourceAttr(resourceName, "control_id", "CIS.1.10"),
					resource.TestCheckResourceAttr(resourceName, "control_status", "ENABLED"),
					resource.TestCheckResourceAttrSet(resourceName, "control_status_updated_at"),
					resource.TestCheckResourceAttr(resourceName, "description", "IAM password policies can prevent the reuse of a given password by the same user. It is recommended that the password policy prevent the reuse of passwords."),
					resource.TestCheckResourceAttr(resourceName, "disabled_reason", ""),
					resource.TestCheckResourceAttr(resourceName, "related_requirements.0", "CIS AWS Foundations 1.10"),
					resource.TestCheckResourceAttrSet(resourceName, "remediation_url"),
					resource.TestCheckResourceAttr(resourceName, "severity_rating", "LOW"),
					resource.TestCheckResourceAttr(resourceName, "title", "Ensure IAM password policy prevents password reuse"),
				),
			},
		},
	})
}

func testAccAWSSecurityHubStandardsControl_disabledControlStatus(t *testing.T) {
	var standardsControl securityhub.StandardsControl
	resourceName := "aws_securityhub_standards_control.test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acctest.PreCheck(t) },
		ErrorCheck:   acctest.ErrorCheck(t, securityhub.EndpointsID),
		Providers:    acctest.Providers,
		CheckDestroy: nil, //lintignore:AT001
		Steps: []resource.TestStep{
			{
				Config: testAccAWSSecurityHubStandardsControlConfig_disabledControlStatus(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckAWSSecurityHubStandardsControlExists(resourceName, &standardsControl),
					resource.TestCheckResourceAttr(resourceName, "control_status", "DISABLED"),
					resource.TestCheckResourceAttr(resourceName, "disabled_reason", "We handle password policies within Okta"),
				),
			},
		},
	})
}

func testAccAWSSecurityHubStandardsControl_enabledControlStatusAndDisabledReason(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acctest.PreCheck(t) },
		ErrorCheck:   acctest.ErrorCheck(t, securityhub.EndpointsID),
		Providers:    acctest.Providers,
		CheckDestroy: nil, //lintignore:AT001
		Steps: []resource.TestStep{
			{
				Config:      testAccAWSSecurityHubStandardsControlConfig_enabledControlStatus(),
				ExpectError: regexp.MustCompile("InvalidInputException: DisabledReason should not be given for action other than disabling control"),
			},
		},
	})
}

func testAccCheckAWSSecurityHubStandardsControlExists(n string, control *securityhub.StandardsControl) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Security Hub Standards Control ID is set")
		}

		conn := acctest.Provider.Meta().(*conns.AWSClient).SecurityHubConn

		standardsSubscriptionARN, err := tfsecurityhub.StandardsControlARNToStandardsSubscriptionARN(rs.Primary.ID)

		if err != nil {
			return err
		}

		output, err := finder.StandardsControlByStandardsSubscriptionARNAndStandardsControlARN(context.TODO(), conn, standardsSubscriptionARN, rs.Primary.ID)

		if err != nil {
			return err
		}

		*control = *output

		return nil
	}
}

func testAccAWSSecurityHubStandardsControlConfig_basic() string {
	return acctest.ConfigCompose(
		testAccAWSSecurityHubStandardsSubscriptionConfig_basic,
		`
resource aws_securityhub_standards_control test {
  standards_control_arn = format("%s/1.10", replace(aws_securityhub_standards_subscription.test.id, "subscription", "control"))
  control_status        = "ENABLED"
}
`)
}

func testAccAWSSecurityHubStandardsControlConfig_disabledControlStatus() string {
	return acctest.ConfigCompose(
		testAccAWSSecurityHubStandardsSubscriptionConfig_basic,
		`
resource aws_securityhub_standards_control test {
  standards_control_arn = format("%s/1.11", replace(aws_securityhub_standards_subscription.test.id, "subscription", "control"))
  control_status        = "DISABLED"
  disabled_reason       = "We handle password policies within Okta"
}
`)
}

func testAccAWSSecurityHubStandardsControlConfig_enabledControlStatus() string {
	return acctest.ConfigCompose(
		testAccAWSSecurityHubStandardsSubscriptionConfig_basic,
		`
resource aws_securityhub_standards_control test {
  standards_control_arn = format("%s/1.12", replace(aws_securityhub_standards_subscription.test.id, "subscription", "control"))
  control_status        = "ENABLED"
  disabled_reason       = "We handle password policies within Okta"
}
`)
}
