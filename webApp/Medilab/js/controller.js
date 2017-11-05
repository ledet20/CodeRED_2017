var client = mqtt.connect("ws://iot.eclipse.org:80/ws")
client.on("connect", function() {
    console.log('connected');
    client.subscribe('distrk/patient');
});


client.on('message', function(topic, message) {

    console.log(JSON.parse(message.toString()));
     json = JSON.parse(message.toString());
	 var consults_id = json.Consults[0].id;

	/*
     window.id = json.id;
     name = json.name;
     completed = json.completed;
    var deparments_id = json.deparments[0].id;
    var deparments_name = json.deparments[0].name;
    var deparments_location = json.deparments[0].location;
	
	for each (variable in object) {
  statement
}
*/
   w3.displayObject("id01", json);
	w3.displayObject("id02", json);
	
	
	// w3.displayObject("id0300", json);
	//w3.displayObject("id04", json);
	//w3.displayObject("id05", json);
	//w3.displayObject("id06", json);
	
	console.log(json.Consults.length);
	
	for(var i = 0; i < 2; i++ ) {
		var value = (i + 1);
	//	console.log(value);
		console.log(value.toString());

		 // w3.displayObject(value.toString(), json );
	//	console.log(i);
	//	console.log(value);
	}
	
});

	for(var i = 0; i < 2; i++ ) {
		var value = (i + 1);
		//console.log(value);
		//console.log(value.toString());
		window.$vars = {
        value_out: value
    };

		 // w3.displayObject(value.toString(), json );
	//	console.log(i);
	//	console.log(value);
	}