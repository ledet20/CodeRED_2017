# Distrk

We were inspired to create this simple tool that would make communication between patients more effective in terms of automating the discharge process. By allowing the physician to provide the patient with a detailed list of procedures we would be able to make the discharge process faster. Our main concern when building this product was to bridge the gap between any miscommunication that may occur between the patient and the physician.

## What your submission does?

he basic architecture of the app is as follows:

NFC Bootstrapper (phone-app) <--MQTT--> server --MQTT--> web-app

The basic idea is that the user (patient) uses NFC from his phone to quickly pull the required data from the server. NFC is only using for bootstrapping and connecting the web app to the server. Once the app is connected, the NFC has not used anymore.

The Web App, in its current state, shows the required steps that the patient needs to follow in order to complete the discharge process. However, updating the state of each step is not possible yet.

As outlined in the diagram above, MQTT was used for all communication between devices.

## How you built it?
All of our modules make use of a MQTT library in order to connect to the public sandbox MQTT broker iot.eclipse.org on port 1883.

### Phone App
The Phone App, which **simulates** the NFC bootstrapping between the web app and the server is written in Python and uses PAHO MQTT for connecting to the broker.

### Web App
Listens to the MQTT broker for changes in the current patient. Thanks to subscribing feature of the MQTT protocol real-time changes were easy to render. The web technologies used were HTML, CSS, and Javascript. Javascript libraries include jQuery, Handlebars, and MQTT.js.

### Server
The server is the component of our app that takes care of information routing between the components. The server connects to a database where all the patient data is stored. Once a pairing request is received, it validates that the patient exists and sends the patient information back to the web app.

The server is written in Golang, and uses gobot for MQTT communication and gorm for database manipulations.

## Challenges you ran into, and how you overcame them?
One of the biggest challenges was to decide which programming languages and frameworks to use that would be able to effectively integrate with our database and each other. We also were able to integrate MQTT into our program even though no one in our group had prior knowledge communication protocol before let alone network programming. Also, we had to effectively send data to and from different programming languages and databases.

## Accomplishments that you're proud of

Effectively creating a system design and its implementation that we were that was able to serve as a guide throughout the whole project. As a team that had formed the day of the hackathon, we were able to effectively collaborate with each other and learn from each other's skillsets.

## What's next for your product?

What we have today is a very simple prototype. For the next version, we hope to have a GUI to load and modify data. We would also want to use the actual device as a link between the app and the server.

## Diagrams

## ER Model
![er model](https://raw.githubusercontent.com/Christopher-Hill/CodeRED_2017/master/diagrams/er_model.jpg)

## Architecture
![architecture](https://raw.githubusercontent.com/Christopher-Hill/CodeRED_2017/master/diagrams/architeture.jpg)

## Sequence
![sequence](https://raw.githubusercontent.com/Christopher-Hill/CodeRED_2017/master/diagrams/sequence.jpg)
