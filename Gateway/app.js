import { ApolloServer } from 'apollo-server'
import gateway from './createSubs.js'

const runServer = async () => {
    // Get newly merged schema
    // start server with the new schema
    const server = new ApolloServer({schema: gateway });
    server.listen().then(({url}) => {
        console.log(`Running at ${url}`);
    });
}

try {
    runServer().then(r => console.log(r));
} catch (err) {
    console.error(err);
}
