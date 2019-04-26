#!/usr/bin/ python
import pika

pika_conn_params = pika.ConnectionParameters(
    host='localhost', port=5672,
    credentials=pika.credentials.PlainCredentials('guest', 'guest'),
)
connection = pika.BlockingConnection(pika_conn_params)
channel = connection.channel()
queue = channel.queue_declare(
    queue="sensordata2", durable=True,
    exclusive=False, auto_delete=False
)

print(queue.method.message_count)