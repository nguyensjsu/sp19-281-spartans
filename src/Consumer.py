import pika
import datetime
from riak.client import RiakClient
import os

riak_client = RiakClient(host=os.environ['RIAK_HOST'], pb_port=8087)

table_temperature = riak_client.table('Temperature')
table_cpu = riak_client.table('CPU')
table_mem = riak_client.table('Mem')

credentials = pika.PlainCredentials('guest', 'guest')
parameters = (
    pika.ConnectionParameters(host=os.environ['RABBIT_1'],credentials=credentials),
    pika.ConnectionParameters(host=os.environ['RABBIT_2'],
                              connection_attempts=5, retry_delay=1,credentials=credentials))


def callback(ch, method, properties, body):
    message = body.decode('utf-8')
    message = message.split(',')
    print(message)
    producer = int(message[0])
    date = datetime.datetime.strptime(message[1], '%Y-%m-%dT%H:%M:%S-07:00')
    temp = float(message[4])
    cpu_usage = float(message[2])
    mem_usage = float(message[3])
    ts_obj_temp = table_temperature.new([[producer,producer,date,temp]])
    print(ts_obj_temp.store())
    ts_obj_cpu = table_cpu.new([[producer,producer,date,cpu_usage]])
    print(ts_obj_cpu.store())
    ts_obj_mem = table_mem.new([[producer,producer,date,mem_usage]])
    print(ts_obj_mem.store())


connection = pika.BlockingConnection(parameters)
channel = connection.channel()

channel.basic_consume(queue='ha.spartans',
        auto_ack=True,
        on_message_callback=callback)

channel.start_consuming()
connection.close()

