var parseArgs = require('minimist')
var args = parseArgs(process.argv.slice(2), { default: { port: 8888 } });

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

call.on('data', number => {
  console.log('received data', number.value)
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