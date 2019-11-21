'use strict';
exports.__esModule = true;
var plant_1 = require("./device/plant");
var URL = 'http://localhost:8080';
var Plant1 = new plant_1["default"](URL, 1);
var Plant2 = new plant_1["default"](URL, 2);
Plant1.sendData(Math.random(), new Date().toString())
  .then(console.log, console.error);
Plant2.sendData(Math.random(), new Date())
  .then(console.log, console.error);
Plant1.getData()
    .then(console.log, console.error);

