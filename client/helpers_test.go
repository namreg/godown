package client

import context "context"

type ctxKey string

func contextWithValue(key string, value interface{}) context.Context {
	return context.WithValue(context.Background(), ctxKey(key), value)
}

func stringToPtr(str string) *string {
	return &str
}
