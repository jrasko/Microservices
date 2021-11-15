const { ApolloServer } = require('apollo-server');
const { ApolloGateway } = require("@apollo/gateway");

const gateway = new ApolloGateway({
    serviceList: [
        { name: 'orders', url: 'http://localhost:4001/query' },
        { name: 'shops', url: 'http://localhost:4002/query' },
        { name: 'customers', url: 'http://localhost:4003/query' }
    ],
});

const server = new ApolloServer({
    gateway,
    subscriptions: false,
});

server.listen().then(({ url }) => {
    console.log(`🚀 Server ready at ${url}`);
});
