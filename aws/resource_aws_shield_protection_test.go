package aws

import (
	"fmt"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudfront"
	"github.com/aws/aws-sdk-go/service/shield"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccAWSShieldProtection_GlobalAccelerator(t *testing.T) {
	resourceName := "aws_shield_protection.test"
	rName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPartitionHasServicePreCheck(shield.EndpointsID, t)
			testAccPreCheckAWSShield(t)
		},
		ErrorCheck:   testAccErrorCheck(t, shield.EndpointsID),
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAWSShieldProtectionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccShieldProtectionGlobalAcceleratorConfig(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAWSShieldProtectionExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.%", "0"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccAWSShieldProtection_ElasticIPAddress(t *testing.T) {
	resourceName := "aws_shield_protection.test"
	rName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPartitionHasServicePreCheck(shield.EndpointsID, t)
			testAccPreCheckAWSShield(t)
		},
		ErrorCheck:   testAccErrorCheck(t, shield.EndpointsID),
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAWSShieldProtectionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccShieldProtectionElasticIPAddressConfig(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAWSShieldProtectionExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.%", "0"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccAWSShieldProtection_disappears(t *testing.T) {
	resourceName := "aws_shield_protection.test"
	rName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPartitionHasServicePreCheck(shield.EndpointsID, t)
			testAccPreCheckAWSShield(t)
		},
		ErrorCheck:   testAccErrorCheck(t, shield.EndpointsID),
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAWSShieldProtectionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccShieldProtectionElasticIPAddressConfig(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAWSShieldProtectionExists(resourceName),
					testAccCheckResourceDisappears(testAccProvider, resourceAwsShieldProtection(), resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccAWSShieldProtection_Alb(t *testing.T) {
	resourceName := "aws_shield_protection.test"
	rName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPartitionHasServicePreCheck(shield.EndpointsID, t)
			testAccPreCheckAWSShield(t)
		},
		ErrorCheck:   testAccErrorCheck(t, shield.EndpointsID),
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAWSShieldProtectionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccShieldProtectionAlbConfig(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAWSShieldProtectionExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.%", "0"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccAWSShieldProtection_Elb(t *testing.T) {
	resourceName := "aws_shield_protection.test"
	rName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPartitionHasServicePreCheck(shield.EndpointsID, t)
			testAccPreCheckAWSShield(t)
		},
		ErrorCheck:   testAccErrorCheck(t, shield.EndpointsID),
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAWSShieldProtectionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccShieldProtectionElbConfig(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAWSShieldProtectionExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.%", "0"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccAWSShieldProtection_Cloudfront(t *testing.T) {
	resourceName := "aws_shield_protection.test"
	rName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPartitionHasServicePreCheck(shield.EndpointsID, t)
			testAccPartitionHasServicePreCheck(cloudfront.EndpointsID, t)
			testAccPreCheckAWSShield(t)
		},
		ErrorCheck:   testAccErrorCheck(t, shield.EndpointsID),
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAWSShieldProtectionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccShieldProtectionCloudfrontConfig(rName, testAccShieldProtectionCloudfrontRetainConfig()),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAWSShieldProtectionExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.%", "0"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccAWSShieldProtection_Cloudfront_Tags(t *testing.T) {
	resourceName := "aws_shield_protection.test"
	rName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPartitionHasServicePreCheck(shield.EndpointsID, t)
			testAccPartitionHasServicePreCheck(cloudfront.EndpointsID, t)
			testAccPreCheckAWSShield(t)
		},
		ErrorCheck:   testAccErrorCheck(t, shield.EndpointsID),
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAWSShieldProtectionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccShieldProtectionCloudfrontConfigTags1(rName, testAccShieldProtectionCloudfrontRetainConfig(), "Key1", "value1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAWSShieldProtectionExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.Key1", "value1"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccShieldProtectionCloudfrontConfigTags2(rName, testAccShieldProtectionCloudfrontRetainConfig(), "Key1", "value1updated", "Key2", "value2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAWSShieldProtectionExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "tags.Key1", "value1updated"),
					resource.TestCheckResourceAttr(resourceName, "tags.Key2", "value2"),
				),
			},
			{
				Config: testAccShieldProtectionCloudfrontConfigTags1(rName, testAccShieldProtectionCloudfrontRetainConfig(), "Key2", "value2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAWSShieldProtectionExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.Key2", "value2"),
				),
			},
		},
	})
}

func TestAccAWSShieldProtection_Route53(t *testing.T) {
	resourceName := "aws_shield_protection.test"
	rName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPartitionHasServicePreCheck(shield.EndpointsID, t)
			testAccPreCheckAWSShield(t)
		},
		ErrorCheck:   testAccErrorCheck(t, shield.EndpointsID, "route53"),
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAWSShieldProtectionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccShieldProtectionRoute53HostedZoneConfig(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAWSShieldProtectionExists(resourceName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckAWSShieldProtectionDestroy(s *terraform.State) error {
	shieldconn := testAccProvider.Meta().(*AWSClient).shieldconn

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "aws_shield_protection" {
			continue
		}

		input := &shield.DescribeProtectionInput{
			ProtectionId: aws.String(rs.Primary.ID),
		}

		resp, err := shieldconn.DescribeProtection(input)

		if isAWSErr(err, shield.ErrCodeResourceNotFoundException, "") {
			continue
		}

		if err != nil {
			return err
		}

		if resp != nil && resp.Protection != nil && aws.StringValue(resp.Protection.Id) == rs.Primary.ID {
			return fmt.Errorf("The Shield protection with ID %v still exists", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckAWSShieldProtectionExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		conn := testAccProvider.Meta().(*AWSClient).shieldconn

		input := &shield.DescribeProtectionInput{
			ProtectionId: aws.String(rs.Primary.ID),
		}

		_, err := conn.DescribeProtection(input)

		if err != nil {
			return err
		}

		return nil
	}
}

func testAccPreCheckAWSShield(t *testing.T) {
	conn := testAccProvider.Meta().(*AWSClient).shieldconn

	input := &shield.ListProtectionsInput{}

	_, err := conn.ListProtections(input)

	if testAccPreCheckSkipError(err) || isAWSErr(err, shield.ErrCodeResourceNotFoundException, "subscription does not exist") {
		t.Skipf("skipping acceptance testing: %s", err)
	}

	if err != nil {
		t.Fatalf("unexpected PreCheck error: %s", err)
	}
}

// Set the environment variable TF_TEST_CLOUDFRONT_RETAIN
// when doing manual tests so that the test is not waiting for
// the distribution to be removed completely.
func testAccShieldProtectionCloudfrontRetainConfig() string {
	if _, ok := os.LookupEnv("TF_TEST_CLOUDFRONT_RETAIN"); ok {
		return "retain_on_delete = true"
	}
	return ""
}

func testAccShieldProtectionRoute53HostedZoneConfig(rName string) string {
	return fmt.Sprintf(`
resource "aws_route53_zone" "test" {
  name    = "%[1]s.com."
  comment = "Terraform Acceptance Tests"

  tags = {
    foo  = "bar"
    Name = %[1]q
  }
}

data "aws_partition" "current" {}

resource "aws_shield_protection" "test" {
  name         = %[1]q
  resource_arn = "arn:${data.aws_partition.current.partition}:route53:::hostedzone/${aws_route53_zone.test.zone_id}"
}
`, rName)
}

func testAccShieldProtectionElbConfig(rName string) string {
	return fmt.Sprintf(`
data "aws_availability_zones" "available" {
  state = "available"

  filter {
    name   = "opt-in-status"
    values = ["opt-in-not-required"]
  }
}

variable "subnets" {
  default = ["10.0.1.0/24", "10.0.2.0/24"]
  type    = list(string)
}

resource "aws_vpc" "test" {
  cidr_block = "10.0.0.0/16"

  tags = {
    foo  = "bar"
    Name = %[1]q
  }
}

resource "aws_subnet" "test" {
  count                   = 2
  vpc_id                  = aws_vpc.test.id
  cidr_block              = element(var.subnets, count.index)
  map_public_ip_on_launch = true
  availability_zone       = element(data.aws_availability_zones.available.names, count.index)

  tags = {
    foo  = "bar"
    Name = %[1]q
  }
}

resource "aws_elb" "test" {
  name = %[1]q

  subnets  = aws_subnet.test[*].id
  internal = true

  listener {
    instance_port     = 8000
    instance_protocol = "http"
    lb_port           = 80
    lb_protocol       = "http"
  }

  tags = {
    foo  = "bar"
    Name = %[1]q
  }

  cross_zone_load_balancing = true
}

resource "aws_shield_protection" "test" {
  name         = %[1]q
  resource_arn = aws_elb.test.arn
}
`, rName)
}

func testAccShieldProtectionAlbConfig(rName string) string {
	return fmt.Sprintf(`
data "aws_availability_zones" "available" {
  state = "available"

  filter {
    name   = "opt-in-status"
    values = ["opt-in-not-required"]
  }
}

variable "subnets" {
  default = ["10.0.1.0/24", "10.0.2.0/24"]
  type    = list(string)
}

resource "aws_lb" "test" {
  name            = %[1]q
  internal        = true
  security_groups = [aws_security_group.test.id]
  subnets         = aws_subnet.test[*].id

  idle_timeout               = 30
  enable_deletion_protection = false

  tags = {
    foo  = "bar"
    Name = %[1]q
  }
}

resource "aws_vpc" "test" {
  cidr_block = "10.0.0.0/16"

  tags = {
    foo  = "bar"
    Name = %[1]q
  }
}

resource "aws_subnet" "test" {
  count                   = 2
  vpc_id                  = aws_vpc.test.id
  cidr_block              = element(var.subnets, count.index)
  map_public_ip_on_launch = true
  availability_zone       = element(data.aws_availability_zones.available.names, count.index)

  tags = {
    foo  = "bar"
    Name = %[1]q
  }
}

resource "aws_security_group" "test" {
  name        = %[1]q
  description = "test"
  vpc_id      = aws_vpc.test.id

  ingress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    foo  = "bar"
    Name = %[1]q
  }
}

resource "aws_shield_protection" "test" {
  name         = %[1]q
  resource_arn = aws_lb.test.arn
}
`, rName)
}

func testAccShieldProtectionCloudfrontConfig(rName, retainOnDelete string) string {
	return fmt.Sprintf(`
resource "aws_cloudfront_distribution" "test" {
  origin {
    custom_origin_config {
      http_port              = 80
      https_port             = 443
      origin_protocol_policy = "https-only"

      origin_ssl_protocols = [
        "TLSv1",
        "TLSv1.1",
        "TLSv1.2",
      ]
    }

    # This is a fake origin and it's set to this name to indicate that.
    domain_name = "%[1]s.com"
    origin_id   = %[1]q
  }

  enabled             = false
  wait_for_deployment = false

  default_cache_behavior {
    allowed_methods  = ["HEAD", "DELETE", "POST", "GET", "OPTIONS", "PUT", "PATCH"]
    cached_methods   = ["GET", "HEAD"]
    target_origin_id = %[1]q

    forwarded_values {
      query_string = false
      headers      = ["*"]

      cookies {
        forward = "none"
      }
    }

    viewer_protocol_policy = "redirect-to-https"
    min_ttl                = 0
    default_ttl            = 0
    max_ttl                = 0
  }

  restrictions {
    geo_restriction {
      restriction_type = "none"
    }
  }

  tags = {
    foo  = "bar"
    Name = %[1]q
  }

  viewer_certificate {
    cloudfront_default_certificate = true
  }

  %[2]s
}

resource "aws_shield_protection" "test" {
  name         = %[1]q
  resource_arn = aws_cloudfront_distribution.test.arn

}
`, rName, retainOnDelete)
}

func testAccShieldProtectionCloudfrontConfigTags1(rName, retainOnDelete, tagKey string, tagValue string) string {
	return fmt.Sprintf(`
resource "aws_cloudfront_distribution" "test" {
  origin {
    custom_origin_config {
      http_port              = 80
      https_port             = 443
      origin_protocol_policy = "https-only"

      origin_ssl_protocols = [
        "TLSv1",
        "TLSv1.1",
        "TLSv1.2",
      ]
    }

    # This is a fake origin and it's set to this name to indicate that.
    domain_name = "%[1]s.com"
    origin_id   = %[1]q
  }

  enabled             = false
  wait_for_deployment = false

  default_cache_behavior {
    allowed_methods  = ["HEAD", "DELETE", "POST", "GET", "OPTIONS", "PUT", "PATCH"]
    cached_methods   = ["GET", "HEAD"]
    target_origin_id = %[1]q

    forwarded_values {
      query_string = false
      headers      = ["*"]

      cookies {
        forward = "none"
      }
    }

    viewer_protocol_policy = "redirect-to-https"
    min_ttl                = 0
    default_ttl            = 0
    max_ttl                = 0
  }

  restrictions {
    geo_restriction {
      restriction_type = "none"
    }
  }

  tags = {
    foo  = "bar"
    Name = %[1]q
  }

  viewer_certificate {
    cloudfront_default_certificate = true
  }

  %[2]s
}

resource "aws_shield_protection" "test" {
  name         = %[1]q
  resource_arn = aws_cloudfront_distribution.test.arn

  tags = {
    %[3]q = %[4]q
  }
}
`, rName, retainOnDelete, tagKey, tagValue)
}

func testAccShieldProtectionCloudfrontConfigTags2(rName, retainOnDelete, tagKey1 string, tagValue1 string, tagKey2 string, tagValue2 string) string {
	return fmt.Sprintf(`
resource "aws_cloudfront_distribution" "test" {
  origin {
    custom_origin_config {
      http_port              = 80
      https_port             = 443
      origin_protocol_policy = "https-only"

      origin_ssl_protocols = [
        "TLSv1",
        "TLSv1.1",
        "TLSv1.2",
      ]
    }

    # This is a fake origin and it's set to this name to indicate that.
    domain_name = "%[1]s.com"
    origin_id   = %[1]q
  }

  enabled             = false
  wait_for_deployment = false

  default_cache_behavior {
    allowed_methods  = ["HEAD", "DELETE", "POST", "GET", "OPTIONS", "PUT", "PATCH"]
    cached_methods   = ["GET", "HEAD"]
    target_origin_id = %[1]q

    forwarded_values {
      query_string = false
      headers      = ["*"]

      cookies {
        forward = "none"
      }
    }

    viewer_protocol_policy = "redirect-to-https"
    min_ttl                = 0
    default_ttl            = 0
    max_ttl                = 0
  }

  restrictions {
    geo_restriction {
      restriction_type = "none"
    }
  }

  tags = {
    foo  = "bar"
    Name = %[1]q
  }

  viewer_certificate {
    cloudfront_default_certificate = true
  }

  %[2]s
}

resource "aws_shield_protection" "test" {
  name         = %[1]q
  resource_arn = aws_cloudfront_distribution.test.arn

  tags = {
    %[3]q = %[4]q
    %[5]q = %[6]q
  }
}
`, rName, retainOnDelete, tagKey1, tagValue1, tagKey2, tagValue2)
}

func testAccShieldProtectionElasticIPAddressConfig(rName string) string {
	return fmt.Sprintf(`
data "aws_availability_zones" "available" {
  state = "available"

  filter {
    name   = "opt-in-status"
    values = ["opt-in-not-required"]
  }
}

data "aws_region" "current" {}

data "aws_caller_identity" "current" {}

data "aws_partition" "current" {}

resource "aws_eip" "test" {
  vpc = true

  tags = {
    foo  = "bar"
    Name = %[1]q
  }
}

resource "aws_shield_protection" "test" {
  name         = %[1]q
  resource_arn = "arn:${data.aws_partition.current.partition}:ec2:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:eip-allocation/${aws_eip.test.id}"
}
`, rName)
}

func testAccShieldProtectionGlobalAcceleratorConfig(rName string) string {
	return fmt.Sprintf(`
resource "aws_shield_protection" "test" {
  name         = %[1]q
  resource_arn = aws_globalaccelerator_accelerator.test.id
}

resource "aws_globalaccelerator_accelerator" "test" {
  name            = %[1]q
  ip_address_type = "IPV4"
  enabled         = true
}
`, rName)
}
