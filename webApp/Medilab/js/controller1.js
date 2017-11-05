//Create a new Client object with your broker's hostname, port and your own clientId
var client = new Paho.MQTT.Client("ws://iot.eclipse.org:80/ws",  "myclientid_" + parseInt(Math.random() * 100, 10));
// ("broker.mqttdashboard.com", 8000, "myclientid_" + parseInt(Math.random() * 100, 10))
// ("iot.eclipse.org", 1883,  "myclientid_" + parseInt(Math.random() * 100, 10)); iot.eclipse.org

// set callback handlers
client.onConnectionLost = onConnectionLost;
client.onMessageArrived = onMessageArrived;

var options = {

     //connection attempt timeout in seconds
     timeout: 3,

     //Gets Called if the connection has successfully been established
     onSuccess: function () {
         alert("Connected");
     },

     //Gets Called if the connection could not be established
     onFailure: function (message) {
         alert("Connection failed: " + message.errorMessage);
     }

 };

//Attempt to connect
client.connect(options);

// Called when the connection is made
// function onConnect(){
//     console.log("Connected!");
// }

// Connect the client, providing an onConnect callback
// client.connect({
//     onSuccess: onConnect
// });


// called when the client connects
function onConnect() {
  // Once a connection has been made, make a subscription and send a message.
  console.log("onConnect");
  client.subscribe("distrk/patient/1");
  message = new Paho.MQTT.Message("Hello");
  message.destinationName = "World";
  client.send(message);
}

// called when the client loses its connection
function onConnectionLost(responseObject) {
  if (responseObject.errorCode !== 0) {
    console.log("onConnectionLost:"+responseObject.errorMessage);
  }
}

// called when a message arrives
function onMessageArrived(message) {
  console.log("onMessageArrived:"+message.payloadString);
}