# placekittens

This is a simple kitten image resizer API built with Go and the Gin framework. It provides endpoints to serve images in color or grayscale, resized to the specified dimensions. The images are served from a specified directory, and the API supports both color and grayscale transformations.

## Getting Started

### Prerequisites

- Go 1.18 or later
- [Gin Framework](https://github.com/gin-gonic/gin)
- [nfnt/resize](https://github.com/nfnt/resize)

### Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/dhruvsangamwar/placekittens.git
   cd palcekittens
   ```

2. **Install dependencies:**

   ```bash
   go mod tidy
   ```

### Configuration

1. **Set up the image directory:**
   
   Ensure you have a directory named `public` (or change the path in the `GetRandomImage` function) with images that will be used by the API.

2. **Modify paths if necessary:**

   The API expects images to be located in `../public`. Adjust the paths in the `GetRandomImage` function if your directory structure is different.

### Running the API

1. **Launch the server:**

   ```bash
   go run main.go
   ```

2. **Access the API:**

   - The server will start on `localhost:8080`.

### Endpoints

- **Home Page**

  `GET /`

  Renders an HTML page located at `./static/index.html`.

- **Resize Image (Color)**

  `GET /:width/:height`

  Returns a color image resized to the specified width and height.

  **Parameters:**
  - `width`: Desired width of the image.
  - `height`: Desired height of the image.

  **Example:**

  ```bash
  curl http://localhost:8080/300/200
  ```

- **Resize Image (Grayscale)**

  `GET /g/:width/:height`

  Returns a grayscale image resized to the specified width and height.

  **Parameters:**
  - `width`: Desired width of the image.
  - `height`: Desired height of the image.

  **Example:**

  ```bash
  curl http://localhost:8080/g/300/200
  ```

### Error Handling

- **Invalid Dimensions:**

  If the width or height provided is not a valid integer, a `400 Bad Request` error will be returned with a message indicating the error.

- **File Access Errors:**

  If there are issues accessing or decoding images, a `500 Internal Server Error` will be returned.

### License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
