package entity

type contextKey string

const ContextID = contextKey("ID")

// const IsAuthenticatedContextKey = contextKey("isAuthenticated") // same type
// const IsAuthenticatedContextKey = contextKey("authenticatedUserID")
const IsAuthenticatedContextKey = contextKey("isAuthenticated")
const AuthenticatedUserID = contextKey("authenticatedUserID")
