const WebSocket = require('ws');

const wss = new WebSocket.Server({ port: 8080 });

wss.on('connection', (ws, req) => {
    console.log(`connected: ${req.socket.remoteAddress} ${req.socket.remotePort}`);
    ws.on('message', (message) => {
        console.log(`received: ${message} from ${req.socket.remoteAddress} ${req.socket.remotePort}`);
        ws.send(`ack ${message}`);
    });
});
