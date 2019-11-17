export default abstract class Device {
    protected id: number = 0;
    protected URL: string = '';
    constructor(URL: string, id: number) {
      this.URL = URL;
      this.id = id;
    };
    abstract getData(): Promise<any>;
    abstract sendData(moistureLevel: number, serverTime: Date): Promise<any>;
  }