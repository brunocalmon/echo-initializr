# echo-initializr

This tool was created to help developers create webservices in a faster way, withou needing to lose time with basic configurations.

## Usage

#### Available Commands:

  - help - Help about any command
  - init  - Create a basic project using echo framework from the scratch
    -   -d, --dependencies stringArray 
    
    Set the dependencies of your project. (default: empty)

    - -h, --help 
    
    help for init

    - -n, --namespace [string] 
    
    Set your project's name (default: github.com/example/sample)

    - -o, --outputDir [string]           
    
    Set the output directory to your project. (default "$HOME/echo_initializr")

    - -p, --port [int]                   
    
    Set the port of your project webserver. (default 8080)

    - -v, --version [string]             
    
    Set your project's version (default: installed)


### To try, just run this command on the root of the echo-initializr's project
```
go run main.go init --namespace "github.com/example/sample" -d "github.com/gorilla/mux,github.com/google/uuid" -p 8080
```