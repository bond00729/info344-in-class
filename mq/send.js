#!/usr/bin/env node
"use strict";

const amqp = require("amqplib");

const qName = "testQ";
const mqAddr = process.env.MQADDR || "localhost:5672";
const mqURL = `amqp://${mqAddr}`;

(async function() {
    try {
        console.log("connecting to %s", mqURL);
        let connection = await amqp.connect(mqURL);
        let channel = await connection.createChannel();
        let qConf = await channel.assertQueue(qName, {durable: false});
    
        console.log("starting to send message...");
        setInterval(() => {
            let message = "Current time is " + new Date().toLocaleTimeString();
            channel.sendToQueue(qName, Buffer.from(message));
        }, 1000);
    } catch (err) {
        console.log(err.stack);
    }
})();