const WebSocket = require('ws');
const cron = require('node-cron');

const ws = new WebSocket('ws://localhost:8080');

ws.on('open', () => {
    ws.on('message', (message) => {
        console.log(`received: ${message}`);
    });
    cron.schedule('*/5 * * * * *', (now) => {
        ws.send(`event at \`${now}\``);
    });
});
