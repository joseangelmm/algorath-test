package repository

var scripts = []DbScript{
	{
		Script:      scriptCredential,
		Description: "add cretendials",
	},
}

var scriptCredential = `CREATE TABLE IF NOT EXISTS credentials (
	ID int PRIMARY KEY,	
	APIKEY text ,
	APISECRET text
);

INSERT INTO users (ID,APIKEY,APISECRET) VALUES (0,"","");

`