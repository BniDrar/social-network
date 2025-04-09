package entity

type contextKey string

const ContextID = contextKey("ID")
const IsAuthenticatedContextKey = contextKey("isAuthenticated") // same type
