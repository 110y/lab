const {EchoRequest, EchoResponse} = require('../../proto/echo/server_pb.js');
const {EchoServiceClient} = require('../../proto/echo/server_grpc_web_pb.js');

var client = new EchoServiceClient('http://localhost:40000');

var request = new EchoRequest();

client.echo(request, {}, (err, response) => {
  console.log(response.getHost());
});
