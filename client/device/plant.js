"use strict";
var __extends = (this && this.__extends) || (function () {
    var extendStatics = Object.setPrototypeOf ||
        ({ __proto__: [] } instanceof Array && function (d, b) { d.__proto__ = b; }) ||
        function (d, b) { for (var p in b) if (b.hasOwnProperty(p)) d[p] = b[p]; };
    return function (d, b) {
        extendStatics(d, b);
        function __() { this.constructor = d; }
        d.prototype = b === null ? Object.create(b) : (__.prototype = b.prototype, new __());
    };
})();
exports.__esModule = true;
var device_1 = require("./device");
var http_1 = require("../common/http");
var Plant = /** @class */ (function (_super) {
    __extends(Plant, _super);
    function Plant(URL, id) {
        return _super.call(this, URL, id) || this;
    }
    ;
    Plant.prototype.getData = function () {
        return http_1["default"].getData(this.URL, this.id);
    };
    Plant.prototype.sendData = function (moistureLevel, serverTime) {
        return http_1["default"].sendData(this.URL, this.id, moistureLevel, serverTime);
    };
    return Plant;
}(device_1["default"]));
exports["default"] = Plant;
