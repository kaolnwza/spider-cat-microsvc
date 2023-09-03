
<!-- get topic list -->
kafka-topics --bootstrap-server=localhost:9092 --list

<!-- create new topic -->
kafka-topics --bootstrap-server=localhost:9092 --topic=hihi --create

<!-- conn subscribe topic -->
kafka-console-consumer --bootstrap-server=localhost:9092 --topic=hihi 

<!-- subscribe log msg -->
kafka-console-consumer --bootstrap-server=localhost:9092 --topic=hihi --group=log

<!-- conn publish topic -->
kafka-console-producer --bootstrap-server=localhost:9092 --topic=hihi 

<!-- get consumer group -->
kafka-consumer-groups --bootstrap-server=localhost:9092 --list



{"UUID": "f226518e-9096-4576-b229-03af17b78594"}