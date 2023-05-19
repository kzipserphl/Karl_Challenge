# Karl_Challenge

This repositortory contains Golang files that are designed to run using a command line interface (cli) command. You may already have Golang installed. To verify, launch a terminal and enter the following command:
```
go version
```

If a version is not displayed or "command not found" is displayed, go to https://go.dev/doc/install and follow the instructions. 

server.go
-----
To run, enter:
```
go run server.go
```
When executed, the following actions will occur:
- a route ("/hello") is created.
- a server (":8080") is created.
- a browser window will be opened.
- The text "Hello, World" will be rendered in the the browser window.

*Note: HTTPS (Hypertext Transfer Protocol Secure) is not included currently*

validate.go
-----
To run, enter:
```
go run validate.go
```
This script is based on the instructions provided here: https://www.hackerrank.com/challenges/validating-credit-card-number/problem
When executed, the script will validate 6 provided values according to established rules. 
- The 6 values are in an array at the beginning of the main() function.
- Each value is processed in the validateCC() function.
- If the value passes the validation, "Valid" will be displayed, otherwise "invalid" will be displayed.
- A variable called flagPrintDebug is initially set to false. When set to true, each value will be displayed and also the validation message will be displayed.
