import pika
import boto3

pika_conn_params = pika.ConnectionParameters(
    host='localhost', port=5672,
    credentials=pika.credentials.PlainCredentials('guest', 'guest'),
)
connection = pika.BlockingConnection(pika_conn_params)
channel = connection.channel()
queue = channel.queue_declare(
    queue="ha.spartans", durable=True,
    exclusive=False, auto_delete=False
)


value = queue.method.message_count
client = boto3.client('cloudwatch', region_name='us-west-2')
print(value)
client.put_metric_data(
    MetricData=[
        {
            'MetricName': 'OpsMetric',
            'Dimensions': [
                {
                    'Name': 'ConsumerQueue',
                    'Value': 'queueSize'
                }
             ],
             'Unit': 'Count',
             'Value': value
        }
    ],
    Namespace='SpartanUp'
)
