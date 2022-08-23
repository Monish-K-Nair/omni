// Main server file
package omni

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error: failed to load the env file")
	}

	if os.Getenv("ENV") == "PRODUCTION" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	r.LoadHTMLGlob("./public/html/*")
	port := os.Getenv("PORT")
	log.Printf("\n\n PORT: %s \n ENV: %s \n SSL: %s \n Version: %s \n\n", port, os.Getenv("ENV"), os.Getenv("SSL"), os.Getenv("API_VERSION"))
	if os.Getenv("SSL") == "TRUE" {
		SSLKeys := &struct {
			CERT string
			KEY  string
		}{
			CERT: "./cert/omni.cer",
			KEY:  "./cert/omni.key",
		}
		r.RunTLS(":"+port, SSLKeys.CERT, SSLKeys.KEY)
	} else {
		r.Run(":" + port)
	}
}
