package pbauthentication

import (
	"context"

	"github.com/cerbos/cerbos-sdk-go/cerbos"
	"google.golang.org/grpc/metadata"
)

// UserInfo represents the user information extracted during authentication.
type UserInfo struct {
	// OwnerId represents the subject's organization identifier.
	OwnerId string
	// Subject represents the user ID.
	Subject string
	// Issuer of the token.
	Issuer string
	// Roles represents the roles the user has.
	Roles map[string]struct{}
	// Exp represents the expiration time of the token.
	Exp int64
	// Raw represents the raw token data.
	Raw map[string]interface{}
	// Principal represents the principal of the user to handle authorization.
	Principal *cerbos.Principal
}

// NewGrpcIncomingContext returns a new outgoing context with the subject and issuer in gRPC metadata.
func NewOutgoingContext(userInfo *UserInfo) context.Context {
	return NewOutgoingContextFromContext(context.Background(), userInfo)
}

// NewGrpcIncomingContextFromContext returns outgoing context derived from ctx with the subject and issuer in gRPC metadata.
func NewOutgoingContextFromContext(ctx context.Context, userInfo *UserInfo) context.Context {
	md := metadata.MD{}
	if userInfo != nil {
		if len(userInfo.Subject) > 0 {
			md.Append("io-clbs-openhes-auth-sub", userInfo.Subject)
		}
		if len(userInfo.Issuer) > 0 {
			md.Append("io-clbs-openhes-auth-iss", userInfo.Issuer)
		}
		if len(userInfo.Issuer) > 0 {
			md.Append("io-clbs-openhes-auth-ownerid", userInfo.OwnerId)
		}
	}
	return metadata.NewOutgoingContext(ctx, md)
}

// GetGrpcAuthInfo returns the subject and issuer from incoming gRPC context.
func FromIncomingContext(ctx context.Context) (userInfo *UserInfo) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil
	}
	userInfo = &UserInfo{}
	subjects := md.Get("io-clbs-openhes-auth-sub")
	if len(subjects) > 0 {
		userInfo.Subject = subjects[0]
	}
	issuers := md.Get("io-clbs-openhes-auth-iss")
	if len(issuers) > 0 {
		userInfo.Issuer = issuers[0]
	}
	ownerIds := md.Get("io-clbs-openhes-auth-ownerid")
	if len(ownerIds) > 0 {
		userInfo.OwnerId = ownerIds[0]
	}
	return
}
