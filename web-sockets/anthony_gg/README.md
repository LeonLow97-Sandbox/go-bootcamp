## Commands ran to connect to WebSocket

```js
// Open 2 browsers on console devtools on chrome
// and run the following commands on both browser console
let socket = new WebSocket('ws://localhost:3000/ws');

// callback message from the client
socket.onmessage = (event) => {
  console.log('received from the server: ', event.data);
};

socket.send('Hello from client');
```

```js
let socket = new WebSocket('ws://localhost:3000/orderbookfeed');
```
