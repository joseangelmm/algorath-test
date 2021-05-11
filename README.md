# algorath-test
1) Run Main function from main.go gile

2) Make a PUT request to /credentials in order to update credential from database with json 
{
   "APIKey": "value",
   "APISecret": "value"
} 
//If you don't have it let me know to give you

3) Make a GET request to endpoint /start to conect with websocket

4) Make a GET request to endpoint /shut-down to finish the procedure


To run tests of package repository it must be with working directory /algorath/ not /algorath/repository/ in order to use the same database (bad practise because it will write bad data over well api credentials)
