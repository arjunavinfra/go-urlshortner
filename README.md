# go-urlshortner
Url shortner


Pre request : 
 1. Mysql database on local 
 2. create a database with name "indexer" 
 3. Use the url.sql for doing the same
 4. update the username and passworon main.go [db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/indexer")]
 
 How to run the project 
 1. Install the dependencies
 
 ```ruby
 go get -u github.com/gorilla/mux```
 
 
 ```ruby
 go get -u github.com/go-sql-driver/mysql```
 

 ```ruby
 go run main.go```
 
 pass the json data to localhost:9908/shrink for shrinking the url 
 

![image](https://user-images.githubusercontent.com/118735091/233076295-0ac70578-aae7-4dca-b088-0db4c4000e66.png)


 pass the json data to localhost:9908/resolve for resolving the url 



![image](https://user-images.githubusercontent.com/118735091/233076420-af459255-0552-423e-8f40-8354eb1fdc21.png)
