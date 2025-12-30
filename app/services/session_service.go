package services

//
//func MustGetUserSession(ctx context.Context) *pkgjwt.JwtResponse {
//	val, ok := ctx.Value(middleware.BearerToken).(*pkgjwt.JwtResponse)
//	if !ok {
//		return nil
//	}
//
//	return val
//}
//
//func GetClientAuth(s *Service) (*models.Platforms, error) {
//	session := MustGetClientAuth(s.ctx)
//	if session == nil {
//		return nil, helpers.NewErrorTrace(errors.New("client session forbidden"), "").WithStatusCode(http.StatusForbidden)
//	}
//
//	return session, nil
//}
//
//func MustGetClientAuth(ctx context.Context) *models.Platforms {
//	val, ok := ctx.Value(middleware.Client).(*models.Platforms)
//	if !ok {
//		return nil
//	}
//
//	return val
//}
