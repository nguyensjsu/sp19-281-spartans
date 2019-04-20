# Week 2



### What tasks did I work on / complete?

- Tried out the Amazon SQS with alerting, in which when the number of messages in the SQS reaches a certain threshold, the CloudWatch which monitors this metric, raises an alert which autoscales the consumer instances.
- This turned out to be pretty simple and on further discussions, came to a conclusion that we will implement an architecture like the SQS using RabbitMQ cluster.
- Started exploring the RabbitMQ HA cluster.
- Setup a RabbitMQ HA cluster, where the queues are mirrored.
- Worked with the team to setup the architecture.

### What am I planning to work on next?

* Integrating the RabbitMQ HA cluster with Cloudwatch, where a custom metric of reading the queue size will raise an alarm when the queue size exceeds a threshold and trigger an Autoscaling event of the consumers.  

### What tasks are blocked waiting on another team member?

* Waited on Rakesh to write the consumer and producer GO APIs, which I got during mid week.

