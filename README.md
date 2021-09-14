# Founds Transfer Api

## Descriptor
- This api simulates a money transfer service with the following supported operations: 
    - Send money to other users 
    - Check a user balance
    - Check a user transactions 
    
## Running the API
    
- A local mySQL-DB will be created when you run the app, also the 'user' and 'transaction' tables will be created inside the DB, so make sure you have installed mySQL in your computer. 
- The DB is created with the following configuration: 
                    - dbName:   "Test_founds_transfer",
                    - username: "root",
                    - password: "password",
                    - hostname: "127.0.0.1:3306",
        
- Also, it will INSERT 2 rows in the 'user' table (simulates 2 users of the wallet) with the following features: 
    - User 1: {user_id:'111' , balance:'10000', site: 'ARG'}
    - User 2: {user_id: '222', balance:'20000', site: 'BRA'}

- You are able to create transactions, retrieve the balance and get the operations of just this 2 sample users (are the only one registered).
    
 ## Endpoints
 
- Check a user balance: GET http://localhost:8080/user/{user_id}
    example:  GET http://localhost:8080/user/222
    
- Check a user transactions: GET http://localhost:8080/transactions/{user_id}
    example: GET http://localhost:8080/transactions/222
    
- Send Money to other users: POST http://localhost:8080/transactions/{user_id}/{receptor_id}/{amount}
    example: POST http://localhost:8080/transactions/111/222/50000
    
## Monitoring strategy

- I would add metrics (well-specified tags and metric names) and logs (structured logging, with x-request-id and context tracking), using Datadog and Kibanna as the monitoring platforms in order to trace every error and insight.

## Changelog 

### `0.0.1 - 13/09/2021 Pablo Noras (pablonoras@gmail.com) 
  - First functional version