//@ts-check
"use strict";

//load the express and morgan modules
const express = require("express");
const morgan = require("morgan");

const addr = process.env.ADDR || ":80";
const [host, port] = addr.split(":");
const portNum = parseInt(port);

const app = express();
//app.use() is like a handler in go, it is called on each request
//in this case, morgan is a middleware handler that we are adding 
//can use multiple .use() and they will be executed in order
app.use(morgan(process.env.LOG_FORMAT || "dev"));

app.get("/", (req, res) => {
    res.set("Content-Type", "text/plain");
    res.send("Hello world, from Node.js!")
});

app.get("/v1/users/me/hello", (req, res) => {
    let userJSON = req.get("X-User");
    if (!userJSON) {
        throw new Error("no X-User header provided");
    }
    let user = JSON.parse(userJSON);
    res.json({
        message: `Hello, ${user.firstName} ${user.lastName}`
    });
});

const handlers = require("./handlers");
app.use(handlers({}));

app.use((err, req, res, next) => {
    console.error(err.stack);
    res.set("Content-Type", "text/plain");
    res.send(err.message);
});

app.listen(portNum, host, () => {
    console.log(`server is listening at http://${addr}...`)
});