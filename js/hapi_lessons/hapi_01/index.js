'use strict'
const Hapi = require('hapi')

const server = new Hapi.Server()
server.connection({
  host: 'localhost',
  port: 8000
})

server.route({
  method: 'GET',
  path: '/',
  handler: (request, reply) => {
    reply('hello hapi!')
  }
})

server.route({
  method: 'GET',
  path: '/{name}',
  handler: (request, reply) => {
    reply(`hello ${request.params.name}!`)
  }
})

server.start(() => console.log('Started at:', server.info.uri))
