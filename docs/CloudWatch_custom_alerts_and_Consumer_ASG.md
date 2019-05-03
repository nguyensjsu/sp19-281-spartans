# AWS CloudWatch Custom Alerts and Consumer Auto Scaling Group

For auto scaling RabbitMQ consumers when queue size exceeds the threshold we have to create custom alerts. To do that we need to create custom metrics in CloudWatch which measures the queue length from the rabbitMQ HA cluster and then raise an alarm which would trigger the auto scaling event that scales RabbitMQ consumers.

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

On the instances whose metric has to be measured in our case the RabbitMQ HA cluster, IAM Role should be associated with it. While launching the RabbitMQ instance, associate the QueueMetricRole281 with the instance. 

### SSH into instance

```
ssh -i <key> ec2-user@<jumpbox public ip>
ssh -i <key> centos@<RabbitMQ-node-1 private ip>
```

In the instance, we wrote a small python script which uses the boto library to keep sending the queue length data to Cloud Watch.

[Link to Code](https://github.com/nguyensjsu/sp19-281-spartans/blob/develop/src/CWQueueMetric.py)

This script was added to cron tab to run every 1 minute. It sends the queue size every 1 minute to Cloud Watch.

### CloudWatch Alarm

On the CW console in Metrics tab we can see the QueueSize Metric under the spartanup namespace.

To create the alarm, 

```
select metric: queuesize(in spartan up namespace)
Name: QueueThresholdAlarm
Whenever: queuesize >= 20
Action: None
```

## Consumer AMI

```
Ubuntu 18.04 LTS Free Tier
Instance: 1
Private Subnet
SG Open Ports:   cmpe281-dmz
```

Created a script which consumes messages from RabbitMQ and pushes the data to the Riak TS NoSQL database.

[Link to Consumer code](https://github.com/nguyensjsu/sp19-281-spartans/blob/develop/src/Consumer.py) 

Added this script to init.d to run on startup and enabled the service.

Created the AMI for this image named, consumer-ami

## AutoScaling Group

```
Create a Launch Configuration

Select My AMI:              consumer-ami
Instance Type:              T2-Micro (Free Tier)
Launch Configuration Name:  comsumer-lc
Enable Monitoring:          Enable CloudWatch detailed monitoring
Select Public IP:           Private Subnet
Security Group:             cmpe281-dmz (SG)
Select Key Pair:            cmpe281-us-west-2
Select VPC:                 Spartan-up VPC

Create an Auto Scaling Group

Create Auto Scale Group:        consumer-asg
Group Size (Starts with):       1
Network:                        Spartan-up VPC | Private Subnets

Use scaling policies to adjust the capacity of this group

Scale between:                  1 - 3 instances
Increase 1 instance when:       QueueSize >= 20 
Decrease 1 instance when:       QueueSize <= 20
```
