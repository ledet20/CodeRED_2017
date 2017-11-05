var client = mqtt.connect("ws://iot.eclipse.org:80/ws")

client.on("connect", function() {
   console.log('connected');
   client.subscribe('distrk/patient');
   client.subscribe('distrk/patient/not-found');
});


client.on('message', function(topic, message) {
   if (topic == "distrk/patient") {
      json = JSON.parse(message.toString());
      console.log(json);

      var consultSource = document.getElementById("consult-template").innerHTML
      var consultTemplate = Handlebars.compile(consultSource)

      $('#patient-name').html(json.Name);
      $('.containerDischarge').html('');
      for (var i = 0; i < json.Consults.length; i++) {
	 var html = consultTemplate(json.Consults[i])
	 $('.containerDischarge').append(html);
      }
   } else if (topic == "distrk/patient/not-found") {
      $('#patient-name').html("No patient");
      $('.containerDischarge').html('');
   }
});
