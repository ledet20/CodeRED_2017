# -*- coding: utf-8 -*-
import json
import paho.mqtt.client as mqtt
import webbrowser
import time
from tkinter import *
import tkinter as ttk
from random import randint

root = Tk()
root.title("NFC Simulator")
root.minsize(400, 200)
res_width = root.winfo_screenwidth()
res_height = root.winfo_screenheight()
root.geometry("%dx%d+0+0" % (res_width, res_height))

# Known and given values
phone_device = "A1EKAC1"
patientID = 11923801
bounds = (1, 3)

# Functions
def scanDevices(event):
    simulated_number = 5
    while (True):
        simulated_number = randint(1, 3)
        if(simulated_number == 1):
            text_broker_activity.insert(END, "\nSwiped device %s found." % phone_device)
            client.publish("distrk/new", patientID)
            # MQTT verify(input_ID)
            break


def swipe(event):
    scanDevices(event)

def manual_connect():
    inputID = entry_manualID.get()
    if (len(str(inputID)) == 0):
        text_broker_activity.insert(END, "\nDischarge request rejected with missing ID.")
    else:
        text_broker_activity.insert(END, "\nDischarge of patient requested")

    client.publish("distrk/new", inputID)

def back():
    label_error.place_forget()
    button_ok.place_forget()

    label_phone_device.place(x=5, y=5)
    button_go.place(x=620, y=800, width=100)
    label_patientID.place(x=15, y=800)
    entry_manualID.place(x=130, y=800)
    label_swipe.place(x=325, y=765)

def show_error():
    label_phone_device.place_forget()
    button_go.place_forget()
    label_patientID.place_forget()
    entry_manualID.place_forget()
    label_swipe.place_forget()

    label_error.place(x=320, y=100)
    button_ok.place(x=350, y=800, width=50, height=50)

    # Config
    label_error.config(font=("Helvetica", 100))
    button_ok.config(command =lambda: back())

    # invalid ID



# Main Frame
frame_main = Frame(root)
frame_main.place(x=10, y=20, height=res_height-150, width=res_width)

# Phone Frame and its Children
frame_phone = Frame(frame_main, background="#25383C")
frame_phone.place(x=0, y=10, height=res_height-20, width=750, bordermode=OUTSIDE)

label_phone_device = Label(frame_phone, text=phone_device, font="120", background="#25383C", foreground="white")
button_go = Button(frame_phone, text="GO", background = "yellowgreen")
label_patientID = Label(frame_phone, font="50", text="Patient ID ")
entry_manualID = Entry(frame_phone, font="50", width=45)
label_swipe = Label(frame_phone, text="Swipe or", background="#25383C", foreground="white")

label_error = Label(frame_phone, text="X")
button_ok = Button(frame_phone, text="Ok", background='yellowgreen')

label_phone_device.place(x=5, y=5)
button_go.place(x=620, y=800, width=100)
label_patientID.place(x=15, y=800)
entry_manualID.place(x=130, y=800)
label_swipe.place(x=325, y=765)

# Config Changes
button_go.config(command=lambda: manual_connect())
label_phone_device.config(font=("Helvetica",44))
label_swipe.config(font=("Helvetica",15))

# Broker Frame and its Children
frame_broker = Frame(frame_main, background="grey")
frame_broker.place(x=780, y=10, height=res_height-20, width=500, bordermode=OUTSIDE)

text_broker_activity = Text(frame_broker)
label_brokerID = Label(frame_broker, text="Broker ID #1923810928197917319")

label_brokerID.place(x=10, y=100, height=20, width=475)
text_broker_activity.place(x=10, y=150, height=700, width=475)

## MQTT
def on_connect(client, userdata, flags, rc):
    text_broker_activity.insert(END, "Broker connected with result code %s." % rc)
    client.subscribe("distrk/nfc-swipe")

def on_publish(client, userdata, result):
    print("Data published with result code %s" % result)

def on_message(client, userdata, msg):
    if msg.topic == "distrk/nfc-swipe":
        print("Received NFC Swipe, forwarding to server...")
        client.publish("distrk/new", msg.payload)

client = mqtt.Client()
client.connect("iot.eclipse.org", 1883, 60)
client.on_connect = on_connect
client.on_publish = on_publish
client.on_message = on_message
client.loop_start()

root.bind("<Up>", swipe)
root.mainloop()
client.loop_stop()
