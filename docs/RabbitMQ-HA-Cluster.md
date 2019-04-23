# RabbitMQ HA Cluster

In our architecture, each queue is set up in a HA cluster. Reason being that the queues should be highly available and even if one queue goes down, the other queue will serve the requests.

### **Setup**

```
AMI: rabbitmq-docker-host-ami
Instance: 2
Private Subnet
```

### Starting RabbitMQ docker containers

**DOCKERFILE**

```
FROM rabbitmq:3-management

RUN apt-get update
RUN apt-get install -y vim
```

**DOCKER BUILD**

```
sudo docker build -t rabbitmq:spartanup .
```

**DOCKER RUN**

In the first instance run:

```
sudo docker run -d -h node-1.rabbit                                 \
           --name rabbit                                            \
           -p "4369:4369"                                           \
           -p "5672:5672"                                           \
           -p "15672:15672"                                         \
           -p "25672:25672"                                         \
           -p "35197:35197"                                         \
           -e "RABBITMQ_USE_LONGNAME=true"                          \
           -e "RABBITMQ_LOGS=/var/log/rabbitmq/rabbit.log"          \
           -e RABBITMQ_ERLANG_COOKIE='abcd'                         \
           -v /data:/var/lib/rabbitmq \
           -v /data/logs:/var/log/rabbitmq \
           rabbitmq:spartanup
```

This sets up the first node.

Similarly spin up the second AWS instance and run:

```
docker run -d -h node-2.rabbit                                      \
           --name rabbit                                            \
           -p "4369:4369"                                           \
           -p "5672:5672"                                           \
           -p "15672:15672"                                         \
           -p "25672:25672"                                         \
           -p "35197:35197"                                         \
           -e "RABBITMQ_USE_LONGNAME=true"                          \
           -e "RABBITMQ_LOGS=/var/log/rabbitmq/rabbit.log"          \
           -e RABBITMQ_ERLANG_COOKIE='abcd'                         \
           -v /data:/var/lib/rabbitmq \
           -v /data/logs:/var/log/rabbitmq \
           rabbitmq:spartanup
```

### Configuring our RabbitMQ cluster

##### **ERLANG COOKIE**

To form a cluster, both the containers on the 2 instances should have the same Erlang Cookie.

Erlang Cookie is located on the /var/lib/rabbitmq/.erlang.cookie

We are setting this up by passing it as an Environment variable in the docker run command. 

##### **/etc/hosts**

The /etc/hosts file on both the instances should have the hosts information.

As well as the /etc/hosts inside the rabbitmq containers on both the nodes.

```
10.1.1.154 node-1.rabbit node-1
10.1.1.159 node-2.rabbit node-2
```

##### **Joining Node1**

From rabbit-node2, join the first node to form the cluster:

```
docker exec rabbit rabbitmqctl stop_app
docker exec rabbit rabbitmqctl join_cluster rabbit@node-1.rabbit
docker exec rabbit rabbitmqctl start_app
```

Now we can see that both the nodes joined to form the cluster.

##### Mirrored Queues

To have a replica-set of the queues in both of the instances, we need to set up the Mirrored queues. This will ensure the high availability of the queues.

```
docker exec rabbit rabbitmqctl set_policy ha "." '{"ha-mode":"all"}'
```

The above ha policy ensures that all queues starting with the name "ha." will be mirrored across the 2 nodes in the cluster.

### Create the HA queue

From any node, create a queue:

```
docker exec rabbit rabbitmqadmin declare queue name=ha.spartans durable=true
```

The above queue gets mirrored in the other node as well.

**Another HA cluster**

Similarly, create another HA cluster of RabbitMQ.

### **Challenges Faced**

As the 2 docker containers of the 2 rabbitmqs are not on the same instance, we were not being able to form the cluster. The join command threw the following error:

```
root@node-2:/# rabbitmqctl join_cluster rabbit@10.250.126.204
Clustering node rabbit@node-2.rabbit with rabbit@10.250.126.204
Error: unable to perform an operation on node 'rabbit@10.250.126.204'. Please see diagnostics information and suggestions below.

Most common reasons for this are:

 * Target node is unreachable (e.g. due to hostname resolution, TCP connection or firewall issues)
 * CLI tool fails to authenticate with the server (e.g. due to CLI tool's Erlang cookie not matching that of the server)
 * Target node is not running

In addition to the diagnostics info below:

 * See the CLI, clustering and networking guides on https://rabbitmq.com/documentation.html to learn more
 * Consult server logs on node rabbit@10.250.126.204
 * If target node is configured to use long node names, don't forget to use --longnames with CLI tools

DIAGNOSTICS
===========

attempted to contact: ['rabbit@10.250.126.204']

rabbit@10.250.126.204:
  * connected to epmd (port 4369) on 10.250.126.204
  * epmd reports node 'rabbit' uses port 25672 for inter-node and CLI tool traffic 
  * TCP connection succeeded but Erlang distribution failed 

  * Node name (or hostname) mismatch: node "rabbit@node-1.rabbit" believes its node name is not "rabbit@node-1.rabbit" but something else.
    All nodes and CLI tools must refer to node "rabbit@node-1.rabbit" using the same name the node itself uses (see its logs to find out what it is)


Current node details:
 * node name: 'rabbitmqcli-1257-rabbit@node-2.rabbit'
 * effective user's home directory: /var/lib/rabbitmq
 * Erlang cookie hash: PXMlDtXQP3bsTm4AKCAwkA==
```

The problem was that the nodes did not know about each other. Modified the /etc/hosts file on the instances, plus inside the containers (as written above in the /etc/hosts section) to solve the problem.
