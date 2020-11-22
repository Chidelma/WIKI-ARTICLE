# WIKI-VUE-GO WEB APPLICATION

## Go Server (Part I)

Open a terminal

```
cd api
```

### Run Tests
```
go test
```

### Run Build
```
go build main -o api 
```
then
```
./api
```

Use curl or broswer (localhost:9090/articles/) to see output

## Vue SPA (Part II)

Open another terminal

### Install Dependencies
```
npm install
```

### Run Tests
```
npm run test:unit
```

### Run Build
```
npm run build
```

## Docker (Part III)

Open Terminal

### Run Build
```
sudo docker build --tag md-wiki:2019 .
```

### Run Docker Container 
```
sudo docker run --name vuego -ti -p 8080:8080 md-wiki:2019
```

Open Another Terminal

### Start Up Go Server
```
sudo docker container exec -it vuego /usr/share/nginx/html/api /usr/share/nginx/html/
```
