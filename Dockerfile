# Client App
FROM node:12-alpine as client-app
LABEL authors="Vasu Vanka"
WORKDIR /usr/src/app
COPY ["package.json", "npm-shrinkwrap.json*", "./"]
RUN npm install --silent
RUN npm install -g @angular/cli
COPY . .
RUN npm run build
EXPOSE 3000
CMD ["npm", "start"]




