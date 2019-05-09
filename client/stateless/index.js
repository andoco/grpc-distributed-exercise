var backoff = require('backoff');

var stub

function setupGrpc(args) {
  var PROTO_PATH = __dirname + '/../../protos/numbers.proto';
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

  stub = new numbers.Generator(`localhost:${args['port']}`, grpc.credentials.createInsecure())
}

var connectBackoff

function setupConnectBackoff() {
  connectBackoff = backoff.exponential({ randomisationFactor: 0.1 })
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
}

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
  if (totalReceived >= maxReceive) {
    console.log('sum:', sum)
    process.exit()
  }
}

var maxReceive
var totalReceived = 0
var sum = 0

function run(args) {
  maxReceive = args.maxReceive
  setupGrpc(args)
  setupConnectBackoff()
  connect()
}

module.exports = {
  run
}