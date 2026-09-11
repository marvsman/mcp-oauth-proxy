package providers

import "context"

type codeVerifierKey struct{}

// WithCodeVerifier returns a context carrying the PKCE code verifier that
// ExchangeCodeForToken sends to the upstream provider as code_verifier.
func WithCodeVerifier(ctx context.Context, verifier string) context.Context {
	if verifier == "" {
		return ctx
	}
	return context.WithValue(ctx, codeVerifierKey{}, verifier)
}

// CodeVerifierFromContext returns the PKCE code verifier stored by WithCodeVerifier, if any.
func CodeVerifierFromContext(ctx context.Context) string {
	if v, ok := ctx.Value(codeVerifierKey{}).(string); ok {
		return v
	}
	return ""
}
