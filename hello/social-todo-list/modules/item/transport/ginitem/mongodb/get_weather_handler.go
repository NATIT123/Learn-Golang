package ginitem

import (
	"encoding/json"
	"fmt"
	"main/common"
	mmongodb "main/modules/item/models/mongodb"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetWeather(API_KEY string) func(*gin.Context) {
	return func(c *gin.Context) {

		var data mmongodb.WeatherData

		city := c.Query("city")

		value, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=" + city + "&units=metric&appid=" + API_KEY)
		defer value.Body.Close() // Ensure body is closed

		if err != nil {
			fmt.Println("Error making HTTP request:", err)
			c.JSON(http.StatusBadRequest, err)
			return
		}

		if value.StatusCode != 200 {
			fmt.Println("Error status code from API:", value.StatusCode)
			c.JSON(http.StatusInternalServerError, "Error fetching weather data")
			return
		}

		// Print response for debugging (replace later):
		// body, _ := ioutil.ReadAll(value.Body)
		// fmt.Println("API response:", string(body))

		err = json.NewDecoder(value.Body).Decode(&data)
		if err != nil {
			fmt.Println("Error decoding JSON:", err)
			c.JSON(http.StatusBadRequest, "Error parsing weather data")
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
