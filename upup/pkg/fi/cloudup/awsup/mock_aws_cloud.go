package awsup

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/autoscaling"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/elb"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/route53"
	"github.com/golang/glog"
	"k8s.io/kops/upup/pkg/fi"
	"k8s.io/kubernetes/federation/pkg/dnsprovider"
)

type MockAWSCloud struct {
	MockCloud
	region string
	tags   map[string]string

	zones []*ec2.AvailabilityZone
}

var _ fi.Cloud = (*MockAWSCloud)(nil)

func InstallMockAWSCloud(region string, zoneLetters string) {
	i := &MockAWSCloud{region: region}
	awsCloudInstances[region] = i
	for _, c := range zoneLetters {
		azName := fmt.Sprintf("%s%c", region, c)
		az := &ec2.AvailabilityZone{
			RegionName: aws.String(region),
			ZoneName:   aws.String(azName),
			State:      aws.String("available"),
		}
		i.zones = append(i.zones, az)
	}

	allRegions = []*ec2.Region{
		{RegionName: aws.String(region)},
	}
}

type MockCloud struct {
}

func (c *MockCloud) ProviderID() fi.CloudProviderID {
	return "mock"
}

func (c *MockCloud) FindDNSHostedZone(dnsName string) (string, error) {
	return "", fmt.Errorf("MockCloud FindDNSHostedZone not implemented")
}

func (c *MockCloud) DNS() (dnsprovider.Interface, error) {
	return nil, fmt.Errorf("MockCloud DNS not implemented")
}

func (c *MockAWSCloud) Region() string {
	return c.region
}

func (c *MockAWSCloud) DescribeAvailabilityZones() ([]*ec2.AvailabilityZone, error) {
	return c.zones, nil
}

func (c *MockAWSCloud) AddTags(name *string, tags map[string]string) {
	glog.Fatalf("MockAWSCloud AddTags not implemented")
}

func (c *MockAWSCloud) BuildFilters(name *string) []*ec2.Filter {
	glog.Fatalf("MockAWSCloud BuildFilters not implemented")
	return nil
}

func (c *MockAWSCloud) AddAWSTags(id string, expected map[string]string) error {
	return fmt.Errorf("MockAWSCloud AddAWSTags not implemented")
}

func (c *MockAWSCloud) BuildTags(name *string) map[string]string {
	glog.Fatalf("MockAWSCloud BuildTags not implemented")
	return nil
}

func (c *MockAWSCloud) Tags() map[string]string {
	glog.Fatalf("MockAWSCloud Tags not implemented")
	return nil
}

func (c *MockAWSCloud) CreateTags(resourceId string, tags map[string]string) error {
	return fmt.Errorf("MockAWSCloud CreateTags not implemented")
}

func (c *MockAWSCloud) GetTags(resourceID string) (map[string]string, error) {
	return nil, fmt.Errorf("MockAWSCloud GetTags not implemented")
}

func (c *MockAWSCloud) GetELBTags(loadBalancerName string) (map[string]string, error) {
	return nil, fmt.Errorf("MockAWSCloud GetELBTags not implemented")
}

func (c *MockAWSCloud) CreateELBTags(loadBalancerName string, tags map[string]string) error {
	return fmt.Errorf("MockAWSCloud CreateELBTags not implemented")
}

func (c *MockAWSCloud) DescribeInstance(instanceID string) (*ec2.Instance, error) {
	return nil, fmt.Errorf("MockAWSCloud DescribeInstance not implemented")
}

func (c *MockAWSCloud) DescribeVPC(vpcID string) (*ec2.Vpc, error) {
	return nil, fmt.Errorf("MockAWSCloud DescribeVPC not implemented")
}

func (c *MockAWSCloud) ResolveImage(name string) (*ec2.Image, error) {
	return nil, fmt.Errorf("MockAWSCloud ResolveImage not implemented")
}

func (c *MockAWSCloud) WithTags(tags map[string]string) AWSCloud {
	m := &MockAWSCloud{}
	*m = *c
	m.tags = tags
	return m
}

func (c *MockAWSCloud) EC2() *ec2.EC2 {
	glog.Fatalf("MockAWSCloud EC2 not implemented")
	return nil
}

func (c *MockAWSCloud) IAM() *iam.IAM {
	glog.Fatalf("MockAWSCloud IAM not implemented")
	return nil
}

func (c *MockAWSCloud) ELB() *elb.ELB {
	glog.Fatalf("MockAWSCloud ELB not implemented")
	return nil
}

func (c *MockAWSCloud) Autoscaling() *autoscaling.AutoScaling {
	glog.Fatalf("MockAWSCloud Autoscaling not implemented")
	return nil
}

func (c *MockAWSCloud) Route53() *route53.Route53 {
	glog.Fatalf("MockAWSCloud Route53 not implemented")
	return nil
}
