package handlers

import (
	"backend/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GEtHealt(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{"status":"ok"})
}

func GetWeather(c *gin.Context){
	city:=c.Query("city")
	if city ==""{
		c.JSON(http.StatusBadGateway, gin.H{"error":"quey param 'city' is required"})
		return	
	}

	units := c.DefaultQuery("units", "metric")

	data, err := service.FetchCurrentWeather(city, units)

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error":err.Error()})
		return
	}

	// pengkokndisian jika array kosong
	conditon, description, icon := "", "", ""

	if len(data.Weather)>0{
		conditon = data.Weather[0].Main
		description = data.Weather[0].Description
		icon = data.Weather[0].Icon
	}
	c.JSON(http.StatusOK, gin.H{
		"city" : data.NAME,
		"country": data.Sys.Country,
		"units": units,
		"temp": data.Main.Temp,
		"temp_min": data.Main.TempMin,
		"temp_max": data.Main.TempMax,
		"humidity" : data.Main.Humidity,
		"condition": conditon,
		"description":description,
		"icon":icon,
		"timestamp": data.Dt,
		"timezone": data.TimeZOne,

	})

}