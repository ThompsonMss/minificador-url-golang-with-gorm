package services

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"time"

	"github.com/thompsonmss/minificador-url-golang-with-gorm/entities"
	"github.com/thompsonmss/minificador-url-golang-with-gorm/repositories"
)

type HomeService struct {
	Response http.ResponseWriter
	Request  *http.Request
	RepoUrl  repositories.UrlRepositoryInterface
}

type ResponseData struct {
	Obs   string
	Error bool
	Success bool
}

func (service *HomeService) Index() {

	erroResponse := ResponseData{
		Obs:   "",
		Error: false,
		Success: true,
	}

	viewHome(service, erroResponse)
}

func (service *HomeService) Register() {

	service.Request.ParseForm()
	urlData := service.Request.Form.Get("url")

	// Checking if url is valid
	res, _ := regexp.MatchString(`^(?:http(s)?:\/\/)?[\w.-]+(?:\.[\w\.-]+)+[\w\-\._~:/?#[\]@!\$&'\(\)\*\+,;=.]+$`, urlData)

	if res == true {

		// Checking if url is from my site to avoid infinite loop
		myUrlSite := os.Getenv("URL_CHECK")
		res, _ := regexp.MatchString(myUrlSite, urlData)

		if res {

			erroResponse := ResponseData{
				Obs:   "Você não pode encurtar URLs deste site.",
				Error: true,
				Success: false,
			}

			viewHome(service, erroResponse)
		} else {

			// Gerando um novo HASH
			hash := GenerateHash(2)
			hash = verifyHash(hash, service.RepoUrl, 2)

			urlOficial := os.Getenv("URL_OFICIAL")

			urlModel := entities.NewUrl(urlData, hash)
			_, err := service.RepoUrl.Store(urlModel)

			if err != nil {

				fmt.Printf("Error: %v", err)

				data := ResponseData{
					Error: true,
					Obs:   "Não foi possível gerar a URL",
					Success: false,
				}

				viewHome(service, data)
			} else {
				data := ResponseData{
					Obs: os.Getenv("PROTOCOL") + "://" + urlOficial + "/" + hash,
				}

				viewHome(service, data)
			}

		}

	} else {

		erroResponse := ResponseData{
			Obs:   "Por favor, insira uma URL válida.",
			Error: true,
			Success: false,
		}

		viewHome(service, erroResponse)
	}
}

func viewHome(service *HomeService, data interface{}) {
	t := template.New("home-template")

	filePrefix, _ := filepath.Abs("./views/html/home")

	t, err := template.ParseFiles(filePrefix + "/index.html")

	if err != nil {
		log.Fatalf("Erro: %v", err)
	}

	t.Execute(service.Response, data)
}

func GenerateHash(size int) string {

	charValidos := "abcdefghijklmnopqrstuvwxyz123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var arrChar []string

	for _, v := range charValidos {
		arrChar = append(arrChar, string(v))
	}

	lenCharValidos := len(arrChar)

	hash := ""

	for i := 0; i < size; i++ {

		rand.Seed(time.Now().UnixNano())
		char := rand.Intn(lenCharValidos - 1)

		hash = hash + "" + arrChar[char]
	}

	return hash

}

func verifyHash(hash string, repo repositories.UrlRepositoryInterface, size int) string {

	urlModel, _ := repo.FindHash(hash)
	size++

	if urlModel.Hash != "" {
		hash = verifyHash(GenerateHash(size), repo, size)
	}

	return hash
}
