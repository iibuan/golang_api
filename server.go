package main

import (
	"fmt"
	"net/http"

	"github.com/iibuan/golang_api/controller"
	router "github.com/iibuan/golang_api/http"
	m "github.com/iibuan/golang_api/middleware"
	"github.com/iibuan/golang_api/repository"
	"github.com/iibuan/golang_api/service"
)

var (
	postRepository repository.PostRepository = repository.NewFireStoreRepository()
	postService    service.PostService       = service.NewPostService(postRepository)
	postController controller.PostController = controller.NewPostController(postService)
	httpRouter     router.Router             = router.NewMuxRouter()
)

func main() {
	const port string = ":8000"
	httpRouter.GET("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Up and running")
	})
	httpRouter.GET("/admin/posts", m.Chain(postController.GetPosts, m.Logging()))
	httpRouter.POST("/admin/posts", m.Chain(postController.AddPost, m.Logging()))
	httpRouter.SERVE(port)
}
