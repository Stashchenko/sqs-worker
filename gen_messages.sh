aws --endpoint http://localhost:4576 sqs create-queue --queue-name task-queue.fifo --attributes '{"FifoQueue": "true", "ContentBasedDeduplication":"true"}'

COUNTER=1
for ((n=0;n<100;n++))
do
 aws --endpoint http://localhost:4576 sqs send-message --queue-url http://localhost:4576/queue/task-queue.fifo --message-body "{'msg': 'hello${COUNTER}'}" --message-group-id "a-string"
 COUNTER=$[$COUNTER +1]
done

