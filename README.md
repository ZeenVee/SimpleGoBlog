# Simple Blog using Golang

## API Endpoint
### Create an article
- Method: `POST`
- Path: `/articles`

### Get article by ID
- Method: `GET`
- Path: `/articles/<article_id>`

### Get all articles
- Method: `Get`
- Path: `/articles`

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

