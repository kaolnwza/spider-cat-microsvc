# Spider Cat Microservice

<br/>

## Description
Cat attack spider with API, on service would be publisher from Cat-Service to attacking Spider on Spider-Consumer. Whole events control by Kafka.

<br/>


## Objective
- I just want to learning Microservice with Kafka.
- Practice more Go.
- I've learning kafka from here [CodeBankok](https://www.youtube.com/watch?v=RjtIdUOpH04) (GREATE TEACHER 👍).

<br/>



## Annotation
- I'm new for Kafka and Microservice.
- Hex Arch some point won't be loosely couple between adapter and domain.
- Some point not completed at all, i will be continue them in the future.
- ⭐️⭐️⭐️ Just Have Fun xD ⭐️⭐️⭐️.

## Try IT!!

Start Kafka & Zookeep.

    docker compose up

<br/>

Start Cat Service.

    go run cat-service/cmd/main.go

<br/>

Start Cat Service.

    go run spider-service/cmd/main.go

<br/>

Request to Attack Spider.

    curl -d 'content-type:application/json' localhost:1111/spiders/c9e5c3d8-c1a4-47ca-b086-fa9767280153