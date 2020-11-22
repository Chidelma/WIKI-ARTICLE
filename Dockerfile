#
# Build and bundle the Vue.js frontend SPA 
#
FROM node:14-alpine AS vue-build

WORKDIR /build

COPY package*.json ./

RUN npm install

COPY . .

RUN npm run build

#
# Build the Go server backend
#
FROM golang:1.15-alpine as go-build

WORKDIR /build/src/api

COPY api/*.go ./

RUN go build -o api

#
# Assemble the server binary and Vue bundle into a single app
#
# production stage
FROM nginx:stable-alpine as production-stage

COPY ./.nginx/nginx.conf /etc/nginx/nginx.conf

RUN rm -rf /usr/share/nginx/html/*

COPY --from=go-build /build/src/api/api /usr/share/nginx/html

COPY --from=vue-build /build/dist /usr/share/nginx/html

EXPOSE 8080

EXPOSE 9090

ENTRYPOINT ["nginx", "-g", "daemon off;"]