# AWS CloudWatch Custom Alerts
    + For auto scaling rabbitMQ consumers when queue size exceeds the threshold we have to create custom alerts
    + To do that we need to create custom matrics in CloudWatch

## Steps to create custom metric in CloudWatch

### Create IAM policy
```
Service:        CloudWatch
Access level:   Write
                PutMetricData
Name:           QueueMetricData281
```
### Create IAM Role
```
Service type  :   EC2
policy        :   QueueMetricData281
Name          :   QueueMetricRole281
```
### Launch EC2 Instance
```
Name            :   RabbitMQ-node-1
Instance Type   :   t2.micro
AMI             :   rabbitmq-docker-host-ami
IAM             :   QueueMetricRole281
AMI             :   Riak KV 2.2 Series
VPC             :   cmpe281
Network         :   Private Subnet
Auto Public IP  :   no
Security Group  :   rabbit-sg 
SG Open Ports   :   4369, 5672, 15672, 25672, 35197
```
### SSH into instance
```
ssh -i <key> ec2-user@<jumpbox public ip>
ssh -i <key> centos@<RabbitMQ-node-1 private ip>
```

