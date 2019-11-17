import fetch from 'node-fetch';

export default {
  sendData: (URL: string, id: number, moistureLevel: number, serverTime: Date) => (
    fetch(`${URL}/sendData`, {
      method: 'POST',
      body: JSON.stringify({ id, moistureLevel, serverTime }),
      headers: { 'Content-Type': 'application/json' }
    }).then((res: any) => res.json())
  ),
  getData: (URL: string, id: number) => (
    fetch(`${URL}/getData`, {
      method: 'POST',
      body: JSON.stringify({name}),
      headers: { 'Content-Type': 'application/json' }
    }).then((res: any) => res.json())
  )
};