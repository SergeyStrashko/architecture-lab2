'use strict';

import Plant from './device/plant';

const URL = 'http://localhost:8080';
const Plant1 = new Plant(URL, 1);
const Plant2 = new Plant(URL, 2);

// Plant1.sendData(0.1, new Date())
//   .then(console.log, console.error);

// Plant2.sendData(0.3, new Date())
//   .then(console.log, console.error);

Plant1.getData()
  .then(console.log, console.error);

Plant2.getData()
  .then(console.log, console.error);