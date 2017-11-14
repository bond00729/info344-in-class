//@ts-check
"use strict";

const express = require("express");

//value can be anything you want
module.exports = (mongoSession) => {
    if (!mongoSession) {
        throw new Error("where is my mongo session?")
    }

    let router = express.Router();
    router.get("/v1/channels", (req, res) => {
        //query mongo with mongoSession 
        res.json([{name: "general"}]);
    });

    return router;
}