var client = mqtt.connect("ws://iot.eclipse.org:80/ws")
client.on("connect", function() {
    console.log('connected');
    client.subscribe('distrk/patient/1');
});
client.on('message', function(topic, message) {
    console.log(message.toString());
    console.log(typeof message.toString())
    var msg = JSON.stringify(message.toString())
    var msg1 = JSON.parse(msg);

    console.log(typeof msg)
    console.log(typeof msg1)
    console.log(msg1);
    console.log(msg1.name);


});