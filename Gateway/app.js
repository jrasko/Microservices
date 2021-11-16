const express = require('express');
const {ApolloServer} = require('apollo-server');
const {ApolloGateway} = require('@apollo/gateway');
const client = require('prom-client');
client.collectDefaultMetrics()

const app = express();
app.get('/metrics', (req, res) => {
    res.set('Content-Type', client.register.contentType)
    client.register.metrics().then((str) => {
            res.end(str)
        }
    )
})

const gateway = new ApolloGateway({
    serviceList: [
        {name: 'orders', url: 'http://localhost:4001/query'},
        {name: 'shops', url: 'http://localhost:4002/query'},
        {name: 'customers', url: 'http://localhost:4003/query'}
    ],
});


const server = new ApolloServer({
        gateway: gateway,
        subscriptions: false,
    }
);

app.listen({port: 5000}, () => {
    console.log(`🚀 Metrics started`);
});

server.listen().then(({ url }) => {
    console.log(`🚀 Server ready at ${url}`);
});
