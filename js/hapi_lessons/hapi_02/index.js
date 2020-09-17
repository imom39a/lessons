'use strict'
const Hapi = require('hapi')
const server = new Hapi.Server()


var routes = [{
  method: 'GET',
  path: '/',
  handler: (request, reply) => {
    reply('hello hapi!');
  }
}, {
  method: 'GET',
  path: '/{name}',
  handler: (request, reply) => {
    server.log('error', 'error log');
    server.log('info', 'info log');
    server.log('om39a', 'custom log');
    reply(`hello ${request.params.name}!`)
  }
}];


let goodOptions = {
  reporters: [{
    reporter: require('good-console'),
    events: {
      log: ['om39a'],
      response: '*'
    }
  }, {
    reporter: require('good-file'),
    events: {
      ops: '*'
    },
    config: './test/awesome.log'
  }]
}

server.connection({
  host: 'localhost',
  port: 8000
});

server.register({
  register: require('good'),
  options: goodOptions
}, err => {
  server.route(routes);
  server.start(() => console.log('Started at:', server.info.uri))
});
