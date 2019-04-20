package schema

// GetSchema : get graphql schema
func GetSchema() string {
	schema := `
		schema{
			query: Query
			mutation: Mutation
		}
		type Query {
			getMyProfile: GetMyProfileResponse!
		}
		type Mutation {
			signUp(email:String!, password:String!, firstName:String!, lastName:String!): SignUpResponse!
		}
		type SignUpResponse {
			ok: Boolean!
			error: String
			user: User
		}
		type GetMyProfileResponse {
			ok: Boolean!
			error: String
			user: User
		}
	  type User {
	  	id: ID!
	  	email: String!
	  	password: String!
	  	firstName: String!
	  	lastName: String!
	  	bio: String
	  	avatar: String
	  	createdAt: String!
	  	updatedAt: String!
	  }
	`

	return schema
}