FROM node:15
 
ADD package.json /tmp/package.json
RUN cd /tmp && npm install --quiet

RUN mkdir -p /opt/app && cp -a /tmp/node_modules /opt/app/

WORKDIR /opt/app
ADD . /opt/app
RUN npm install -g nodemon
#RUN npm run knex migrate:latest
#RUN npm run knex seed:run
EXPOSE 3073

CMD ["npm", "start"]
