1) Run Main function from main.go gile

2) Make a PUT request to /credentials in order to update credential from database with json 
{
   "APIKey": "value",
   "APISecret": "value"
}

3) Make a GET request to endpoint /start to conect with websocket

4) Make a GET request to endpoint /shut-down to finish the procedure