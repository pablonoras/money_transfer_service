# Founds Transfer Api

## Descriptor
- This api simulates a money transfer service with the following supported operations: 
    - Send money to other users 
    - Check a user balance
    - Check a user transactions 
    
## Running the API
    
- When running the api locally, a local mySQL DB will be created with the 'user' and 'transaction' tables, so you might have installed mySQL in your computer. 
- The DB is created with the following configuration: 
                    dbName:   "Test_founds_transfer",
                    username: "root",
                    password: "password",
                    hostname: "127.0.0.1:3306",
        
- Also, it will INSERT 2 rows in the 'user' table (simulates 2 users of the wallet) with the following features: 
    - User 1: {user_id:'111' , balance:'10000', site: 'MLA'}
    - User 2: {user_id: '222', balance:'20000', site: 'MLB'}
    
    
 ## Endpoints
 
- Check a user balance: GET http://localhost:8080/user/{user_id}
    example:  GET http://localhost:8080/user/222
    
- Check a user transactions: GET http://localhost:8080/transactions/{user_id}
    example: GET http://localhost:8080/transactions/222
    
- Send Money to other users: POST http://localhost:8080/transactions/{user_id}/{receptor_id}/{amount}
    example: POST http://localhost:8080/transactions/111/222/50000
    
## Monitoring strategy

- I would add some metrics and logs, and use Datadog and Kibanna as the monitoring platforms in order to trace every error and insight.

## Changelog 

### `0.0.1 - 13/09/2021 Pablo Noras (pablonoras@gmail.com) 
  - First functional version