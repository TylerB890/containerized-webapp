FROM node:16-alpine as build

WORKDIR /vue-frontend/

COPY package*.json ./
COPY yarn.lock ./

RUN yarn

COPY . .

RUN yarn build

FROM nginx:1.21.6-alpine

WORKDIR /usr/share/nginx/html

RUN rm -rf ./*
COPY --from=build /vue-frontend/dist /app
COPY nginx/nginx.conf /etc/nginx/nginx.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]