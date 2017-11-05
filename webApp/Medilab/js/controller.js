var client = mqtt.connect("ws://iot.eclipse.org:80/ws")
client.on("connect", function() {
    console.log('connected');
    client.subscribe('distrk/patient/1');
});
client.on('message', function(topic, message) {
    console.log(topic);
    console.log(message);
});