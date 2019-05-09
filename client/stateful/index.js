const uuidv4 = require('uuid/v4')

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

function connect() {
  var call = stub.Begin({ clientId: clientId, maxNumbers: maxNumbers })

  call.on('data', number => {
    handleNumber(number)
  })
  call.on('end', () => {
    console.log('server has finished sending')
  })
  call.on('error', err => {
    console.error('error', err)
    call.destroy()
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
  }
}

var clientId
var maxNumbers
var totalReceived = 0

function run(args) {
  clientId = uuidv4()
  maxNumbers = args.maxReceive
  setupGrpc(args)
  connect()
}


module.exports = {
  run
}