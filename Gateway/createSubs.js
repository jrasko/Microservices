import {print} from 'graphql';
import {introspectSchema} from '@graphql-tools/wrap';
import fetch from 'node-fetch';
import {stitchSchemas} from "@graphql-tools/stitch";
import {delegateToSchema} from '@graphql-tools/delegate';


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
    subschemas: subschemas,
    typeDefs: `
    extend type Order{
        customer: Customer!
        shop: Shop!
    }
    extend type OrderProducts{
        product: Product!
    }
    `,
    resolvers: {
        Order: {
            customer: {
                selectionSet: `{ customerID }`,
                resolve(customer, args, context, info) {
                    return delegateToSchema({
                        schema: subschemas[2],
                        operation: 'query',
                        fieldName: 'customer',
                        args: {id: "00001"},
                        context,
                        info,
                    });
                }
            },
            shop: {
                selectionSet: `{ shopID }`,
                resolve(shop, args, context, info) {
                    return delegateToSchema({
                        schema: subschemas[1],
                        operation: 'query',
                        fieldName: 'shop',
                        args: {id: shop.shopID},
                        context,
                        info,
                    });
                }
            }
        },
        OrderProducts: {
            product: {
                selectionSet: `{ productId }`,
                resolve(product, args, context, info) {
                    return delegateToSchema({
                        schema: subschemas[1],
                        operation: 'query',
                        fieldName: 'product',
                        args: {id: product.productId},
                        context,
                        info,
                    });
                }
            }
        }
    }

})

