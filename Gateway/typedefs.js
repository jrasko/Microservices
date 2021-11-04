const {gql} = require('apollo-server');

const typedefs = gql`
	type Query{
		allCustomers: [Customer!]!
		customer(id: ID!): Customer!
	}
	type Customer{
		id: ID!
		username: String!
		firstname: String!
		lastname: String!
	}
	scalar SWTime
	type Order{
		id: ID!
		time: SWTime!
		state: OrderState!
		shopID: ID!
		customerID: ID!
		products: [OrderProducts!]!
	}
	type OrderProducts{
		order: Order!
		productId: ID!
		quantity: Int!
		price: Float!
	}
	enum OrderState{
		ACCEPTED
		IN_PROGRESS
		DONE
		READY_FOR_PICK_UP
		RETRIEVED
	}
	type Product{
		id: ID!
		name: String!
		description: String
		shops: [ProductInShop!]!
		image: String!
	}
	scalar Coordinates
	type Shop{
		id: ID!
		name: String!
		location: Coordinates
		products: [ProductInShop!]!
	}
	type ProductInShop{
		Product: Product!
		Shop: Shop!
		price: Float!
		currency: String!
		sales_unit: Int!
	}
`
module.exports = typedefs
