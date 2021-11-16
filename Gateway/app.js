const express = require('express');
import { ApolloServer } from 'apollo-server'
import gateway from './createSubs.js'
const client = require('prom-client')
client.collectDefaultMetrics()

const app = express();
app.get('/metrics', (req, res) => {
    res.set('Content-Type', client.register.contentType)
    client.register.metrics().then((str) => {
            res.end(str)
        }
    )
})

const runServer = async () => {
    // Get newly merged schema
    // start server with the new schema
    const server = new ApolloServer({schema: gateway });
    server.listen().then(({url}) => {
        console.log(`Running at ${url}`);
    });
}
app.listen({port: 5000}, () => {
    console.log(`🚀 Metrics started`);
});


try {
    runServer().then(r => console.log(r));
} catch (err) {
    console.error(err);
}
