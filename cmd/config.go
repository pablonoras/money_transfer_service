package cmd

import (
	"database/sql"
)

type config struct{
	mySql map[string]mySql
	apiCall map[string]apiCall
}

type apiCall struct {
	baseURL string
}

type mySql struct{
	username string
	password string
	hostname string
	dbName   string
}

var environmentConfigs = map[string]config{
	"prod": {
		mySql: map[string]mySql{
			"transaction_mysql": {
				//TODO: prod mySQL-client configs
			},
			"user_mysql": {
				//TODO: prod mySQL-client configs
			},
		},
	},
	"local": {
		mySql: map[string]mySql{
			"test_mysql": {
				dbName:   "Test_founds_transfer",
				username: "root",
				password: "password",
				hostname: "127.0.0.1:3306",
			},
		},
	},
}

func initMySQL(config mySql) (*sql.DB, error){

	// TODO: this is just an implementation to run locally, for prod env we might use an existing db.

	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/")
	if err != nil {
		return nil, err
	}

	_,err = db.Exec("CREATE DATABASE IF NOT EXISTS " + config.dbName)
	if err != nil {
		return nil, err
	}

	_,err = db.Exec("USE " + config.dbName)
	if err != nil {
		return nil, err
	}

	dropTrxTable , _ := db.Prepare("DROP TABLE IF EXISTS transactions")
	_, err = dropTrxTable.Exec()
	if err != nil {
		return nil, err
	}

	dropUsersTable , _ := db.Prepare("DROP TABLE IF EXISTS users")
	_, err = dropUsersTable.Exec()
	if err != nil {
		return nil, err
	}

	transactionTable, err := db.Prepare("CREATE Table transactions (id int NOT NULL AUTO_INCREMENT,transaction_id varchar(50) NOT NULL, user_id varchar(50), receptor_id varchar(50), amount varchar(50), site_from varchar(50), site_to varchar(50), status varchar(50), creation_date varchar(50), last_modified_date  varchar(50), PRIMARY KEY (id));")
	if err != nil {
		return nil, err
	}

	_, err = transactionTable.Exec()
	if err != nil {
		return nil, err
	}


	userTable, err := db.Prepare("CREATE Table users(user_id varchar(50) , site varchar(50), balance int(50), PRIMARY KEY (user_id));")
	if err != nil {
		return nil, err
	}
	_, err = userTable.Exec()
	if err != nil {
		return nil, err
	}

	user1, err := db.Query("SELECT * FROM users WHERE user_id = '111';")
	if err != nil {
		return nil, err
	}

	if !user1.Next() {
		insertUser1, err := db.Prepare(" INSERT INTO users (user_id , site, balance) VALUES ('111','MLA','10000');")
		if err != nil {
			return nil, err
		}
		_, err = insertUser1.Exec()
		if err != nil {
			return nil, err
		}
	}

	_,err = db.Exec("USE " + config.dbName)
	if err != nil {
		return nil, err
	}

	user2, err := db.Query("SELECT * FROM users WHERE user_id = '222';")
	if err != nil {
		return nil, err
	}

	if !user2.Next() {
		insertUser2, err := db.Prepare(" INSERT INTO users (user_id , site, balance) VALUES ('222','MLB','20000');")
		if err != nil {
			return nil, err
		}
		_, err = insertUser2.Exec()
		if err != nil {
			return nil, err
		}
	}
	_,err = db.Exec("USE " + config.dbName)
	if err != nil {
		return nil, err
	}

	return db, nil
}
