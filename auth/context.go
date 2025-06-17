package auth

import "context"

// Criamos uma chave "tipo seguro" para evitar colisão com outras chaves do context
type contextKey string

const contextKeyUserID = contextKey("userID")

// Adiciona o userID ao context
func WithUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, contextKeyUserID, userID)
}

// Recupera o userID do context (retorna vazio se não existir)
func GetUserID(ctx context.Context) (string, bool) {
	userID, ok := ctx.Value(contextKeyUserID).(string)
	return userID, ok
}
