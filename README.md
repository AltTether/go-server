# GO TEST SERVER

## usage

### build image
    docker build -t go_server .

### run container
    docker run -it --rm -v $(pwd):/work -p 8080:8080 go_server /bin/bash

### run server
    go run ./src/app.go

### access
http://localhost:8080  
http://localhost:8080\hoge
