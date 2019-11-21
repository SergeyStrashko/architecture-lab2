"use strict";
exports.__esModule = true;
var node_fetch_1 = require("node-fetch");
exports["default"] = {
    sendData: function (URL, id, moistureLevel, serverTime) { return (node_fetch_1["default"](URL + "/sendData", {
        method: 'POST',
        body: JSON.stringify({ id: id, moistureLevel: moistureLevel, serverTime: serverTime }),
        headers: { 'Content-Type': 'application/json' }
    }).then(function (res) { return res.json(); })); },
    getData: function (URL, id) { return (node_fetch_1["default"](URL + "/getData", {
        method: 'POST',
        body: JSON.stringify({ id: id }),
        headers: { 'Content-Type': 'application/json' }
    }).then(function (res) { return res.json(); })); }
};
