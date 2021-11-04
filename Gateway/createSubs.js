import {print} from 'graphql';
import {introspectSchema} from '@graphql-tools/wrap';
import fetch from 'node-fetch';
import {stitchSchemas} from "@graphql-tools/stitch";

const graphqlApis = [
    {
        uri: 'http://localhost:4001/query',
    },
    {
        uri: 'http://localhost:4002/query',
    },
    {
        uri: 'http://localhost:4003/query',
    }
];

let executorFunctions = []
for (let i = 0; i < graphqlApis.length; i++) {
    executorFunctions.push(
        async ({document, variables}) => {
            const query = print(document);
            const fetchResult = await fetch(graphqlApis[i].uri, {
                method: 'POST',
                headers: {'Content-Type': 'application/json'},
                body: JSON.stringify({query, variables}),
            });
            return fetchResult.json();
        }
    )
}

const subschemas = []
for (let i = 0; i < executorFunctions.length; i++) {
    subschemas.push(
        {
            executor: executorFunctions[i],
            schema: await introspectSchema(executorFunctions[i])
        }
    )
}
export default await stitchSchemas({
    subschemas: subschemas
})

