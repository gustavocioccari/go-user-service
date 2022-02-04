## :file_folder: The Project
This project is a simple user API developed to study microservice architecture using Kafka for asynchronous communication

## :rocket: Technologies
|   Back-End   |
| :---:        |
| Go           |
| MongoDB      |
| Kafka        |
| Docker        |

## :computer: Installation
Set the environment variables in a .env file and run
```bash
docker-compose up -d --build
```
It will build the user API, a mongodb to store the users and a kafka broker with a new-users topic already created
___
To test the async communication with other service you can clone the repository https://github.com/gustavocioccari/go-mail and run
```bash
go run sendEmail.go
```
It will start listening to kafka broker and every new user creation will be read and an email will be sent to the given email
___