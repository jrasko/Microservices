const express = require('express');
const {ApolloServer} = require('apollo-server-express');
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
server.start().then(() => {
        server.applyMiddleware({app})
    }
)

app.listen({port: 4000}, () => {
    console.log(`🚀 Server ready at http://localhost:4000${server.graphqlPath}`);
});
