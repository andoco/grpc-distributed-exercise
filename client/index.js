var parseArgs = require('minimist')

var args = parseArgs(process.argv.slice(2), { default: { port: 8888, maxReceive: 0xffff } });

function usage() {
  console.log('USAGE: [OPTS] <stateless|stateful>')
}

var cmd = args._.length > 0 ? args._[0] : ""

switch (cmd) {
  case 'stateless':
    require('./stateless').run(args)
    break
  case 'stateful':
    require('./stateful').run(args)
    break
  default:
    usage()
    break
}