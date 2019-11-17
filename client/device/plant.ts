import Device from './device';
import requests from '../common/http';

export default class Plant extends Device {
  constructor(URL: string, id: number) {
    super(URL, id);
  };
  public getData() {
    return requests.getData(this.URL, this.id);
  }
  public sendData = (moistureLevel: number, serverTime: Date) => {
    return requests.sendData(this.URL, this.id, moistureLevel, serverTime);
  }
}