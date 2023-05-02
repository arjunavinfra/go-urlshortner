# go-urlshortner
Url shortner


Pre request : 
 1. Mysql database on local 
 2. create a database with name "indexer" 
 3. Use the url.sql for doing the same
 4. update the username and passworon main.go [db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/indexer")]
 
 How to run the project 
 1. Install the dependencies
 
 ```go get -u github.com/gorilla/mux```

 ```go run main.go```

![image](https://user-images.githubusercontent.com/118735091/233076295-0ac70578-aae7-4dca-b088-0db4c4000e66.png)



![image](https://user-images.githubusercontent.com/118735091/233076420-af459255-0552-423e-8f40-8354eb1fdc21.png)
