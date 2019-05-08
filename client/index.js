var parseArgs = require('minimist')
var backoff = require('backoff');

var args = parseArgs(process.argv.slice(2), { default: { port: 8888, maxReceive: 0xffff } });

var PROTO_PATH = __dirname + '/../protos/numbers.proto';
var grpc = require('grpc');
var protoLoader = require('@grpc/proto-loader');
// Suggested options for similarity to existing grpc.load behavior
var packageDefinition = protoLoader.loadSync(
  PROTO_PATH,
  {
    keepCase: true,
    longs: String,
    enums: String,
    defaults: true,
    oneofs: true
  });
var numbers = grpc.loadPackageDefinition(packageDefinition).numbers;

var stub = new numbers.Generator(`localhost:${args['port']}`, grpc.credentials.createInsecure())

var connectBackoff = backoff.exponential({ randomisationFactor: 0.1 })
connectBackoff.failAfter(10)

connectBackoff.on('backoff', (number, delay) => {
  console.log('backoff:start', 'number:', number, 'delay:', delay)
})
connectBackoff.on('ready', (number, delay) => {
  console.log('backoff:ready', 'number:', number, 'delay:', delay)
  connect()
})
connectBackoff.on('fail', () => {
  console.log('backoff:fail')
})

var totalReceived = 0
var sum = 0

connect()

function connect() {
  var call = totalReceived == 0 ? stub.Begin() : stub.Resume({ seed: sum })

  call.on('data', number => {
    connectBackoff.reset()
    handleNumber(number)
  })
  call.on('end', () => {
    console.log('server has finished sending')
  })
  call.on('error', err => {
    console.error('error', err)
    call.destroy()
    connectBackoff.backoff(err)
  })
  call.on('status', status => {
    console.log('status', status)
  })
}

function handleNumber(number) {
  totalReceived++
  sum += number.value
  console.log('value:', number.value, 'received:', totalReceived, 'sum:', sum)
  if (totalReceived >= args.maxReceive) {
    console.log('sum:', sum)
    process.exit()
  }
}
