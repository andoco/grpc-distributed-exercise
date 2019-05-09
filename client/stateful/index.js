const uuidv4 = require('uuid/v4')
var backoff = require('backoff');

var stub

function setupGrpc(args) {
  var PROTO_PATH = __dirname + '/../../protos/randstream.proto';
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
  var randstream = grpc.loadPackageDefinition(packageDefinition).randstream;

  stub = new randstream.Generator(`localhost:${args['port']}`, grpc.credentials.createInsecure())
}

var connectBackoff

function setupConnectBackoff() {
  connectBackoff = backoff.exponential({ initialDelay: 1000, maxDelay: 10000, randomisationFactor: 0.1 })
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
  var call = totalReceived == 0
    ? stub.Begin({ clientId: clientId, maxNumbers: maxNumbers })
    : stub.Resume({ clientId: clientId })

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
  console.log('value:', number.value, 'received:', totalReceived)

  if (number.checksum) {
    console.log('received checksum:', number.checksum)
    process.exit()
  }
}

var clientId
var maxNumbers
var totalReceived = 0

function run(args) {
  clientId = uuidv4()
  maxNumbers = args.maxReceive
  setupGrpc(args)
  setupConnectBackoff()
  connect()
}


module.exports = {
  run
}