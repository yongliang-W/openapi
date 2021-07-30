package server

import (
	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	proxy "github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	"github.com/lishimeng/go-log"
	pstore "github.com/lishimeng/openapi/internal/store"
	"sync"
	"time"
)

var singleton *proxy.Server
var once sync.Once

func Init() {
	once.Do(func() {
		log.Info("init oauth2 server")
		initExtraErrors()
		manager := _initConfig()
		singleton = proxy.NewDefaultServer(manager)
		singleton.SetAllowGetAccessRequest(true)
		singleton.SetAllowedResponseType()
		singleton.SetAllowedGrantType(oauth2.ClientCredentials)
		singleton.SetClientInfoHandler(proxy.ClientFormHandler)
	})
}

func _initConfig() (res oauth2.Manager) {
	manager := manage.NewDefaultManager()
	// 修改manage.DefaultXXX(最后几行)
	manager.SetClientTokenCfg(&manage.Config{
		AccessTokenExp:    time.Hour*24*7,
	})
	manager.MustTokenStorage(store.NewMemoryTokenStore())
	//clientStore := store.NewClientStore()
	//_ = clientStore.Set("00001", &models.Client{
	//	ID:     "00001",
	//	Secret: "asdasdasd",
	//	Domain: "http://localhost",
	//})
	clientStore := pstore.New()
	manager.MapClientStorage(clientStore)
	res = manager
	return
}

func GetServer() *proxy.Server {
	return singleton
}

func initExtraErrors() {
	var extraErrors []error
	extraErrors = append(extraErrors,
		errors.ErrInvalidRedirectURI,
		errors.ErrInvalidAuthorizeCode,
		errors.ErrInvalidAccessToken,
		errors.ErrInvalidRefreshToken,
		errors.ErrExpiredAccessToken,
		errors.ErrExpiredRefreshToken,
		errors.ErrMissingCodeVerifier,
		errors.ErrMissingCodeChallenge,
		errors.ErrInvalidCodeChallenge,
		)
	for _, e := range extraErrors{
		errors.Descriptions[e] = e.Error()
	}
}
