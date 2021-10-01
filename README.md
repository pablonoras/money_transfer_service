# Money-Transfer Api

## Descriptor
- This api simulates a money transfer service with the following supported operations: 
    - Send money to other users 
    - Check a user balance
    - Check a user transactions 
    
## How to Run the App? 

### Considerations: 
    
- A local mySQL-DB will be created when you run the app, also the 'user' and 'transaction' tables will be created inside the DB, so make sure you have installed mySQL in your computer. 
- The DB is created with the following configuration: 
                    - dbName:   "Test_founds_transfer",
                    - username: "root",
                    - password: "password",
                    - hostname: "127.0.0.1:3306",
        
- Also, it will INSERT 2 rows in the 'user' table (simulates 2 users of the wallet) with the following features: 
    - User 1: {user_id:'111' , balance:'10000', site: 'ARG'}
    - User 2: {user_id: '222', balance:'20000', site: 'BRA'}

- You are able to create transactions, retrieve the balance and get the operations of this 2 sample users (user_id: '111' and user_id: '222').


### Steps: 

1) Install mysql in your computer
    - In this link https://dev.mysql.com/downloads/mysql/
    - For MacOS you can do it with home brew just by running the command: brew install mysql
    
2) Check MySQL root password: Run the command "mysql -u root -p" and enter the password "password" 
- If it's okey great!, continue with step 3). 
- If It's not, check "https://stackoverflow.com/questions/7534056/mysql-root-password-change" and change it into password: "password". 

3) In the root of the project run the command: go run main.go

4) Check a user balance 
   In any browser, postman or just by running the following curl in the terminal:  
   
   curl --request GET 'http://localhost:8080/user/{user_id}'

   Example: 
   - curl --request GET 'http://localhost:8080/user/222'
   - curl --request GET 'http://localhost:8080/user/111'

5) Send Money to other user: 
   
   curl --request POST 'http://localhost:8080/transactions/{user_id}/{receptor_id}/{amount}'

   Example:
   - curl --request POST 'http://localhost:8080/transactions/111/222/250'
    
6) Check user's transactions:
    
   curl --request GET 'http://localhost:8080/transactions/{user_id}'
   
   Example:  
   - curl --request GET http://localhost:8080/transactions/111


    
## Project Structure

- The project uses a Hexagonal architecture, that is surrounding by the core of the application [core](internal/core) . It is a technology agnostic component that contains all the business logic.
- The Actors are real world things that want to interact with the core (e.g: MySQL DB, KVS or a Queue)
- Ports [ports](internal/core/ports) are interfaces that define how the communication between an actor and the core has to be done. 
- Adapter for a driven port [repositores](internal/repositories) transforms a technology agnostic request from the core into an a specific technology request on the actor.
- Adapter for a driver port  [handlers](internal/handlers) transforms a specific technology request into a call on a core service.
- Dependency injection [dependency injection](cmd/dependencies.go): connection of the adapters to the corresponding ports when the application starts. 
    
## Monitoring strategy

- I would add metrics (well-specified tags and metric names) and logs (structured logging, with x-request-id and context tracking), using Datadog and Kibanna as the monitoring platforms in order to trace every error and insight.

## Test strategy

- Before the deploy I would implement integration and functional tests, by mocking repositories and services, with more than 80% of coverage. 

- [Integration Test Example](internal/core/service/userService_test.go)

## Changelog 

### `0.0.1` - 13/09/2021 Pablo Noras (pablonoras@gmail.com) 
  - First functional version