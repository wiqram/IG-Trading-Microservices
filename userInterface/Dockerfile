FROM node:15

# Create app directory
WORKDIR /app

# Install app dependencies
#RUN npm install webpack -g
# Update container and install deps
RUN apt-get update && apt-get install -y build-essential
# A wildcard is used to ensure both package.json AND package-lock.json are copied
# where available (npm@5+)
#COPY package*.json ./
COPY . ./
RUN npm install --quiet
# If you are building your code for production
# RUN npm ci --only=production

# Bundle app source
#COPY . ./

#CMD ./node_modules/.bin/webpack-dev-server --hot --inline --config ./webpack.config.js --host 0.0.0.0

EXPOSE 8092
CMD [ "npm", "start" ]