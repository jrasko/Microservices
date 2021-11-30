const express = require('express');
const {ApolloServer} = require('apollo-server');
const client = require('prom-client');


module.exports = async ({getBuiltMesh}) => {
    client.collectDefaultMetrics()
    const app = express();
    app.get('/metrics', (req, res) => {
        res.set('Content-Type', client.register.contentType)
        client.register.metrics().then((str) => {
                res.end(str)
            }
        )
    })
    app.listen({port: 5000}, () => {
        console.log(`🚀 Metrics started`);
    });

    const {schema} = await getBuiltMesh();
    const server = new ApolloServer({
            schema,
            subscriptions: false,
        }
    );
    const {url} = await server.listen(4000)
    console.info(`🚀 Server ready at ${url}`);
}

