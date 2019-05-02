# Riak TS

After playing around with Riak KV buckets for storing the sensor data, we figured out that data accessibility and manageability was becoming default, and Basho's documentation recommends to use Riak TS in sensor data use cases.

Riak TS is simpler to use as its query language is almost like SQL.

### Cluster Setup

```
1. AMI:             Riak TS 1.x Series
2. Instance Type:   t2.medium
3. VPC:             spartan
4. Network:         private subnet
5. Auto Public IP:  no
6. Security Group:  riak cluster 
7. SG Open Ports:   (see below)
8. Key Pair:        cmpe281-us-west-1
9. Subnet  
	Private subnet 10.0.1.0/24 - 2 nodes - us-west-2c
	Private subnet 10.0.3.0/24 - 3 nodes - us-west-2a

Riak Cluster Security Group (Open Ports):

    22 (SSH)
    8087 (Riak Protocol Buffers Interface)
    8098 (Riak HTTP Interface)
    Port range: 0-65535 (Source: Security Group ID) // riak cluster SG

```

On each node run

```
sudo riak start
sudo riak ping
sudo riak-admin status
```

From 2nd and 3rd node

```
sudo riak-admin cluster join riak@<ip.of.first.node>
```

On the 1st node

```
sudo riak-admin cluster plan
sudo riak-admin cluster commit
sudo riak-admin member_status
```

On each node's /etc/riak/riak_shell.config, add the other 2 IPs

```
%%% -*- erlang -*-
[
 {riak_shell, [
              {logging, off},
              {cookie, riak},
              {show_connection_status, false},
              {nodes, [
                   riak@<ip.of.first.node>,
                   riak@<ip.of.second.node>,
                   riak@<ip.of.third.node>
                      ]}
             ]}
].
```

On each node check riak-shell

```
sudo riak-shell
riak-shell(1)>ping;
'riak@<ip.of.first.node>':  (connected)
'riak@<ip.of.second.node>':  (connected)
'riak@<ip.of.third.node>':  (connected)
riak-shell(2)>
```



### Creating Tables

After the cluster has been setup, we created 3 tables for each of the metric we are measuring:

**Temperature**

```
CREATE TABLE Temperature 
( 
   id           SINT64    NOT NULL, 
   producerid   SINT64    NOT NULL,
   time         TIMESTAMP NOT NULL,
   temperature  DOUBLE,
   PRIMARY KEY (
     (id, QUANTUM(time, 15, 'm')),
      id, time
   )
);
```

**CPU Usage**

```
CREATE TABLE CPU 
( 
   id           SINT64    NOT NULL, 
   producerid   SINT64    NOT NULL,
   time         TIMESTAMP NOT NULL,
   usage  DOUBLE,
   PRIMARY KEY (
     (id, QUANTUM(time, 15, 'm')),
      id, time
   )
);
```

**RAM Usage**

```
CREATE TABLE Mem
( 
   id           SINT64    NOT NULL, 
   producerid   SINT64    NOT NULL,
   time         TIMESTAMP NOT NULL,
   usage  DOUBLE,
   PRIMARY KEY (
     (id, QUANTUM(time, 15, 'm')),
      id, time
   )
);
```



### Writing Data

We are writing data into the respective tables using the consumer Python script, which fetches the data from RabbitMQ clusters.



### Querying Data

We are querying the data from the Riak cluster using the GO Microservices.



### Reference

https://docs.riak.com/riak/ts/1.5.2/index.html