# Simple Blog using Golang

## API Endpoint
### Create an article
- Method: `POST`
- Path: `/articles`
- Response: 
  - `200` Server valid
  - `201` Article created
  - `503` Server not available

### Get article by ID
- Method: `GET`
- Path: `/articles/<article_id>`
- Response: 
  - `200` Server valid, Article Found
  - `404` Article not found in given id
  - `503` Server not available

### Get all articles
- Method: `Get`
- Path: `/articles`
- Response: 
  - `200` Server valid, Articles Found
  - `404` No artile found
  - `503` Server not available

## Packages Used
- Dependency management: `dep`
- Database: `MySql`
- Routing: `chi`

## Additional Setting in vim .bashrc
 export GOROOT=/usr/local/go</br>
 export PATH=$PATH:$GOROOT/bin</br>

 export GOPATH=/home/`username`/golib</br>
 export PATH=$PATH:$GOPATH/bin</br>
 export GOPATH=$GOPATH:/home/`username`/code</br>

## Project Setep Path
- code
  - bin
  - pkg
  - src
    - github.com
      - `username`
        - SimpleGoBlog

## How I lauch project
Terminal open at ~/code  
run `go build github/username/SimpleGoBlog`  
run `./SimpleGoBlog`

## Database setting
There are 4 constants in main.go which are dbNAME, dBPass, dbHost, and dbPort which can be modified before lauching application
