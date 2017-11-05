var client = mqtt.connect("ws://iot.eclipse.org:80/ws")
client.on("connect", function() {
    console.log('connected');
    client.subscribe('distrk/patient/1');
});


client.on('message', function(topic, message) {

    console.log(JSON.parse(message.toString()));
     json = JSON.parse(message.toString());

     window.id = json.id;
     name = json.name;
     completed = json.completed;
    var deparments_id = json.deparments[0].id;
    var deparments_name = json.deparments[0].name;
    var deparments_location = json.deparments[0].location;
});
