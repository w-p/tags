

package awsapi


import (
    "fmt"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/ec2"
)


func FilterInstances(s *session.Session, kvp []string) ([]ec2.Instance, error) {
    var err error
    client := ec2.New(s)
    filters := []*ec2.Filter{}

    for i := 0; i < len(kvp); i += 2 {
        key := aws.String(kvp[i])
        value := []*string{&kvp[i + 1]}
        filter := ec2.Filter{Name: key, Values: value}
        filters = append(filters, &filter)
    }

    req := ec2.DescribeInstancesInput{Filters: filters}
    res, err := client.DescribeInstances(&req)

    if err != nil {
        return nil, err
    }

    instances := []ec2.Instance{}
    for _, reservation := range res.Reservations {
        for _, instance := range reservation.Instances {
            instances = append(instances, *instance)
        }
    }
    return instances, err
}


func SelectProperty(instances *[]ec2.Instance, property string) []string {
    var results []string
    var value string
    for _, instance := range *instances {
        value = getPropertyByName(&instance, &property)
        results = append(results, fmt.Sprintf(value))
    }
    return results
}


func getPropertyByName(instance *ec2.Instance, property *string) string {
    switch *property {
    case "State":
        return *instance.State.Name
    case "VpcId":
        return *instance.VpcId
    case "ImageId":
        return *instance.ImageId
    case "SubnetId":
        return *instance.SubnetId
    case "InstanceId":
        return *instance.InstanceId
    case "InstanceType":
        return *instance.InstanceType
    case "PublicDnsName":
        return *instance.PublicDnsName
    case "PublicIpAddress":
        return *instance.PublicIpAddress
    case "PrivateDnsName":
        return *instance.PrivateDnsName
    case "PrivateIpAddress":
        return *instance.PrivateIpAddress
    default:
        return ""
    }
}
