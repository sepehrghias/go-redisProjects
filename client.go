package main

import (
	pb "awesomeProject1/golang_pbs"
	"context"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
	"strconv"
	"time"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	ctx  = context.Background()
)

type ads struct {
	Title string `json:"title"`
	Image string `json:"image"`
}

func getAdd(c *gin.Context) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	id := c.Param("id")
	cpc, err := rdb.Get(ctx, id).Result()
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "this id not found"})
		return
	}
	flag.Parse()
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		c.IndentedJSON(http.StatusServiceUnavailable, gin.H{"message": "service not available"})
		return
	}
	defer conn.Close()
	connection := pb.NewAdsWithCpcClient(conn)
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	newId, err := strconv.ParseInt(id, 10, 0)
	min_cpc, err2 := strconv.ParseInt(cpc, 10, 0)
	if err != nil || err2 != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "you enter wrong id"})
		return
	}
	r, err := connection.GetAdByCpc(ctx, &pb.Ad_Request{Id: newId, MinCpc: min_cpc})
	if err != nil {
		c.IndentedJSON(http.StatusServiceUnavailable, gin.H{"message": "service not available"})
		return
	}
	if r == nil {
		fmt.Println("P1:", r)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "ads not found"})
		return
	}
	if r.GetTitle() == "" || r.GetImage() == "" {
		fmt.Printf("P2: %+v \n", r)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "ads not found"})
		return
	}
	fmt.Println(r)
	adResponse := ads{
		Title: r.GetTitle(),
		Image: r.GetImage(),
	}
	c.IndentedJSON(http.StatusOK, adResponse)
}

func main() {
	fmt.Println("start project")
	router := gin.Default()
	router.GET("/min_cpc/:id", getAdd)
	router.Run("localhost:8080")
}
