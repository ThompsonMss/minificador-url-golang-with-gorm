package services

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/thompsonmss/minificador-url-golang-with-gorm/entities"
	"github.com/thompsonmss/minificador-url-golang-with-gorm/repositories"
)

type RedirectService struct {
	Response     http.ResponseWriter
	Request      *http.Request
	Repo         repositories.UrlRepositoryInterface
	RepoClickUrl repositories.ClickUrlRepositoryInterface
}

func (service *RedirectService) Index() {

	hashURL := strings.Replace(fmt.Sprint(service.Request.URL), "/", "", 1)

	urlModel, _ := service.Repo.FindHash(hashURL)

	host := urlModel.Original

	if host == "" {
		host = "/index" // Redirect for my site
	} else {

		// Checking if url is from my site to avoid infinite loop
		myUrlSite := os.Getenv("URL_CHECK")

		res, _ := regexp.MatchString(myUrlSite, urlModel.Original)

		if res == true {
			host = "/index" // Redirect for my site
		}
	}

	channel := make(chan string, 1)

	// Recording click statistics

	ip, _, err := net.SplitHostPort(service.Request.RemoteAddr)
	if err != nil {
		ip = ""
	}

	clickUrlModel := entities.ClickUrl{
		IP:    ip,
		UrlID: urlModel.ID,
	}

	go registerClick(service, clickUrlModel, channel)

	http.Redirect(service.Response, service.Request, host, http.StatusSeeOther)

	<-channel
}

func registerClick(service *RedirectService, clickurl entities.ClickUrl, channel chan string) {

	_, err := service.RepoClickUrl.Store(clickurl)
	if err != nil {
		println("Could not register click")
	}

	channel <- "Finish"

}
