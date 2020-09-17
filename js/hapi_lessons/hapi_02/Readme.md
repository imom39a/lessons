https://egghead.io/lessons/node-js-hapi-js-logging-with-good-and-good-console

hapi.js - Logging with good and good-console

hapi doesn't ship with logging support baked in. Luckily, hapi's rich plugin ecosystem includes everything needed to configure logging for your application. This video will introduce good, a process monitor for hapi, and good-console, a reporter for good that outputs to standard out.


npm init
npm install hapi --save
npm install nodemon
npm instal good good-console --save
npm install good-file good-http --save
nodemon -w ./ index.js


Good - https://www.npmjs.com/package/good

  good is a Hapi process monitor. It listens for events emitted by Hapi Server instances and allows custom reporters to be registered that output subscribed events.

Sets up the GoodConsole reporter listening for 'response' and 'log' events.

Sets up the GoodFile reporter to listen for 'ops' events and log them to ./test/fixtures/awesome_log according to the file rules listed in the good-file documentation.

Sets up the GoodHttp reporter to listen for error events and POSTs them to http://prod.logs:3000 with additional settings to pass into Wreck
