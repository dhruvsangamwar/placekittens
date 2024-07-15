package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/nfnt/resize"
)

// GetRandomFile returns the filename of a random file in the specified directory
func GetRandomImage(dir string) (string, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return "", err
	}

	var fileList []string
	for _, file := range files {
		if !file.IsDir() {
			fileList = append(fileList, filepath.Join(dir, file.Name()))
		}
	}

	if len(fileList) == 0 {
		return "", fmt.Errorf("no files found in directory")
	}

	randomIndex := rand.Intn(len(fileList))

	return fileList[randomIndex], nil
}

// This function should return a resized color image file
func returnColorImage(c *gin.Context) {
	w := c.Param("width")
	h := c.Param("height")

	height, err := strconv.Atoi(h)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid height"})
		return
	}

	width, err := strconv.Atoi(w)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid width"})
		return
	}

	filename, err := GetRandomImage("../public")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to find image"})
		return
	}

	filepath := "../public/" + filename

	file, err := os.Open(filepath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to open image"})
		return
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to decode image"})
		return
	}

	resizedImg := resize.Resize(uint(width), uint(height), img, resize.Lanczos3)

	c.Writer.Header().Set("Content-Type", "image/jpeg")
	jpeg.Encode(c.Writer, resizedImg, nil)
}

func returnGreyImage(c *gin.Context) {
	w := c.Param("width")
	h := c.Param("height")

	height, err := strconv.Atoi(h)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid height"})
		return
	}

	width, err := strconv.Atoi(w)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid width"})
		return
	}

	filename, err := GetRandomImage("../public")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to find image"})
		return
	}

	filepath := "../public/" + filename

	file, err := os.Open(filepath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to open image"})
		return
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to decode image"})
		return
	}

	resizedImg := resize.Resize(uint(width), uint(height), img, resize.Lanczos3)

	grayImg := image.NewGray(resizedImg.Bounds())
	for y := 0; y < resizedImg.Bounds().Dy(); y++ {
		for x := 0; x < resizedImg.Bounds().Dx(); x++ {
			originalColor := resizedImg.At(x, y)
			grayColor := color.GrayModel.Convert(originalColor)
			grayImg.Set(x, y, grayColor)
		}
	}

	c.Writer.Header().Set("Content-Type", "image/jpeg")
	jpeg.Encode(c.Writer, grayImg, nil)

}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("./static/*.html")
    router.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", nil)
    })

	router.GET("/:width/:height", returnColorImage)
	router.GET("/g/:width/:height", returnGreyImage)
	router.Run("localhost:8080")
}
