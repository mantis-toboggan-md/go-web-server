# go-web-server
## Getting Started
clone repo into appropriate go src directory (https://golang.org/doc/code.html)
start web server from root directory:
```
go build
go run ../go_web_server
```
#### React web app
```
cd ui-desk
npm build
```
navigate to localhost:5000/

#### React Native app
```
cd ui-mobile
npm start
```


## RESTful go server
  * go net/http package
  * examples of static file serving and JSON serving
  ### SQL Database
  * connect to sql db using your chosen db drivers: (https://github.com/golang/go/wiki/SQLDrivers)
  * example routes for simple authn CRUD 
  ### Routing
  * routing via Gorilla mux https://www.gorillatoolkit.org/pkg/mux
  * example protected route, redirecting
  ### Authn
  * token-based authn
  * create token and validate token endpoints
  * validate token middleware example
  
## Example UIs
  ### React 16.9.0
  * served from go web server
  * react-router with protected routes using tokens
  * redirect to login
  
  ### React Native
  * technically there
