package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/kardianos/service"
	_ "github.com/lib/pq"
	"net/http"
)

const serviceName = "Medium service"
const serviceDescription = "Simple service, just for fun"

type program struct{}

func (p program) Start(s service.Service) error {
	fmt.Println(s.String() + " started")

	router := httprouter.New()
	router.ServeFiles("/js/*filepath", http.Dir("js"))
	router.ServeFiles("/css/*filepath", http.Dir("css"))
	router.ServeFiles("/img/*filepath", http.Dir("img"))
	router.GET("/", authHandler)
	router.POST("/", authHandler)
	router.GET("/homepage", serveHomepage)
	router.GET("/seasons", serveSeasons)
	router.GET("/profile", servePro)
	router.GET("/predict", serveProTwo)
	router.POST("/get_dinamic", getTime)
	router.POST("/get_season_class", getSeasonClass)
	router.POST("/get_season", getSeason)
	router.POST("/get_profile", getProfile)
	router.POST("/get_predict", getPredict)
	router.POST("/get_class", getClass)
	router.GET("/download_dinamic", downloadDinamic)
	router.GET("/download_season", downloadSeason)
	router.GET("/download_profile", downloadProfile)
	router.GET("/download_predict", downloadPredict)
	err := http.ListenAndServe(":3000", router)
	if err != nil {
	}
	return nil
}

func (p program) Stop(s service.Service) error {
	return nil
}

func main() {
	serviceConfig := &service.Config{
		Name:        serviceName,
		DisplayName: serviceName,
		Description: serviceDescription,
	}
	prg := &program{}
	s, err := service.New(prg, serviceConfig)
	if err != nil {
		fmt.Println("Cannot create the service: " + err.Error())
	}
	err = s.Run()
	if err != nil {
		fmt.Println("Cannot start the service: " + err.Error())
	}
}
