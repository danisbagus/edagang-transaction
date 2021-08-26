package port

type IAuthRepo interface {
	IsAuthorized(token string, routeName string, vars map[string]string) bool
}
