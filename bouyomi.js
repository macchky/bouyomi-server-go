var net = require("net");
var socket = new net.Socket();
var msg = "テストだよん";
console.log(msg);
var bmsg = new Buffer(msg);
console.log(bmsg);
var bmsg_length = new Buffer(new Int32Array([bmsg.length]));
console.log(bmsg_length);
var bCode = new Buffer([0]);
console.log(bCode);
var iVoice = new Buffer(new Int16Array([-30000]));
console.log(iVoice);
var iVolume = new Buffer(new Int16Array([-1]));
console.log(iVolume);
var iSpeed = new Buffer(new Int16Array([-1]));
console.log(iSpeed);
var iTone = new Buffer(new Int16Array([-1]));
console.log(iTone);
var iCommand = new Buffer(new Int16Array([1]));
console.log(iCommand);


//var type = new Buffer([0, -1, -1, -1, 0, 0]);

socket.on("connect", function() {
	console.log("connected.");
	socket.write(iCommand);
	console.log(1);
	socket.write(iSpeed);
	console.log(2);
	socket.write(iTone);
	console.log(3);
	socket.write(iVolume);
	console.log(4);
	socket.write(iVoice);
	console.log(5);
	socket.write(bCode);
	console.log(6);
	socket.write(bmsg_length);
	console.log(7);
	socket.write(bmsg);
	console.log(8);
	socket.end();
});

socket.connect(50001, "192.168.0.2");