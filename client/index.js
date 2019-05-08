var parseArgs = require('minimist')
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

var call = stub.begin();

var totalReceived = 0
var sum = 0

call.on('data', number => {
  totalReceived++
  sum += number.value
  console.log('value:', number.value, 'received:', totalReceived, 'sum:', sum)
  if (totalReceived >= args.maxReceive) {
    console.log('sum:', sum)
    process.exit()
  }
})
call.on('end', () => {
  console.log('server has finished sending')
})
call.on('error', e => {
  console.error('error', e)
})
call.on('status', status => {
  console.log('status', status)
})