# sp19-281-spartans

# Team Project - Device Monitoring System

## 

## Team Name

Spartans

## 

## Team Members

- [Busi Pallavi Reddy](https://github.com/busipallavi-reddy)
- [Maunil Swadas](https://github.com/maunilswadas)
- [Rakesh Amireddy](https://github.com/rakeshamireddy)

## 

## Project Name

### Device Monitoring System

## Project Description

Device Monitoring System is a SAAS app where one can monitor the usage of the 1000's of desktops in an organization. The UI shows the the following 3 statistics of a desktop(producer):

* Device Temperature
* Device CPU Usage
* Device RAM Memory

To keep track of so much amount of data which is being pushed regularly by a producer, we need a reliable and stable architecture.

**Architecture**

![](https://github.com/nguyensjsu/sp19-281-spartans/blob/develop/docs/Architecture.png)

**Design**

* Bottom in the architecture is our Producers which are the desktops. These keep on producing the CPU Stats very frequently. [Producer Code](https://github.com/nguyensjsu/sp19-281-spartans/blob/develop/src/producer.go)

* Then we have load balanced RabbitMQ HA clusters of mirrored queues. These queues serve as Message queues in which the producers produce the data and consumers consume the data. These are mirrored queues(where one is master and the other is a mirrored queue). Both the queues will have the same data, hence forming a Highly Available cluster. [RabbitMQ Cluster setup](https://github.com/nguyensjsu/sp19-281-spartans/blob/develop/docs/RabbitMQ-HA-Cluster.md)

* Consuming the messages is an auto scaled group of consumers, which consume the data from the RabbitMQ cluster and push it to the NoSQL Riak TS database cluster. [Consumer code](https://github.com/nguyensjsu/sp19-281-spartans/blob/develop/src/Consumer.py)

* The auto scaling event of the consumers should be triggered when there is an overwhelming incoming of messages from the queue or any consumer goes down. The AS event is triggered when the RabbitMQ queue length exceeds a certain threshold. To do this, we created an alarm based on a custom metric of queue length on the Cloud Watch console. The RabbitMQ hosted instances keeps on pushing the queue length data to the CloudWatch. Whenever the queue length exceeds a threshold, alarm is triggered and the consumers are scaled out. [CloudWatch Custom Alerts and Consumer ASG](https://github.com/nguyensjsu/sp19-281-spartans/blob/develop/docs/CloudWatch_custom_alerts_and_Consumer_ASG.md)

* The Riak TS is a timeseries NoSQL database which is a perfect match for storing Timeseries data like sensor data. Plus its query language is like SQL and we created 3 tables each for the device stats. [Riak TS setup](https://github.com/nguyensjsu/sp19-281-spartans/blob/develop/docs/RiakTS.md)

* In front of the Riak TS, we hosted our microservices which read the individual stats from the RiakTS. There is another microservice for maintaining all the producers on MongoDB.

  [Microservice for CPU Usage stats](https://github.com/nguyensjsu/sp19-281-spartans/tree/develop/src/microservice_cpu)

  [Microservice for Memory stats](https://github.com/nguyensjsu/sp19-281-spartans/tree/develop/src/microservice_memory)

  [Microservice for temperature stats](https://github.com/nguyensjsu/sp19-281-spartans/tree/develop/src/microservice_temperature)

  [Microservice for maintaining the producers](https://github.com/nguyensjsu/sp19-281-spartans/tree/develop/src/microservice_producers)

* The SAAS app front end is deployed on Heroku which queries the above microservices to fetch the different statistics. [Front End Code](https://github.com/nguyensjsu/sp19-281-spartans/tree/develop/src/frontend)

* SAAS App on Heroku

  ![](https://github.com/nguyensjsu/sp19-281-spartans/blob/develop/docs/FrontEnd.jpeg)

## GitHub Repo:

<https://github.com/nguyensjsu/sp19-281-spartans>

## 

## Project Task Board:

<https://github.com/nguyensjsu/sp19-281-spartans/projects/2>

## 

## Weekly Journals:

**Week 1**

https://github.com/nguyensjsu/sp19-281-spartans/blob/develop/journal/Week1-RakeshAmireddy.md

https://github.com/nguyensjsu/sp19-281-spartans/blob/develop/journal/Week1-busipallavi-reddy.md

https://github.com/nguyensjsu/sp19-281-spartans/blob/develop/journal/Week1-maunil-swadas.md



**Week 2**

https://github.com/nguyensjsu/sp19-281-spartans/blob/develop/journal/Week2-RakeshAmireddy.md

https://github.com/nguyensjsu/sp19-281-spartans/blob/develop/journal/Week2-busipallavi-reddy.md

https://github.com/nguyensjsu/sp19-281-spartans/blob/develop/journal/Week2-maunil-swadas.md



**Week 3**

https://github.com/nguyensjsu/sp19-281-spartans/blob/develop/journal/Week3-RakeshAmireddy.md

https://github.com/nguyensjsu/sp19-281-spartans/blob/develop/journal/Week3-busipallavi-reddy.md

https://github.com/nguyensjsu/sp19-281-spartans/blob/develop/journal/Week3-maunil-swadas.md



**Week 4**

https://github.com/nguyensjsu/sp19-281-spartans/blob/develop/journal/Week4-RakeshAmireddy.md

https://github.com/nguyensjsu/sp19-281-spartans/blob/develop/journal/Week4-busipallavi-reddy.md

https://github.com/nguyensjsu/sp19-281-spartans/blob/develop/journal/Week4-maunil-swadas.md





