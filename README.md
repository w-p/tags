# tags
Gets a property from a filtered list of EC2 instances.

Filters are key value pairs using the keys listed here:

http://docs.aws.amazon.com/cli/latest/reference/ec2/describe-instances.html

### Usage
```
tags \
    select PrivateIpAddress \
    where instance-type t2.micro \
    --profile my-aws-profile \
    --json

[
    "10.0.3.23",
    "10.0.2.125",
    "10.0.2.110"
]
```

### Installation
```
git clone https://github.com/w-p/tags.git
cd tags
go clean && go install
```

### License
MIT
