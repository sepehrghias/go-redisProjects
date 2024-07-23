package main

import (
	pb "awesomeProject1/protos"
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
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0, // use default DB
	})
	id := c.Param("id")
	cpc, err := redisClient.HGet(ctx, "positions", id).Result()
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "position not found"})
		return
	}
	flag.Parse()
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "page not found"})
		return
	}
	defer conn.Close()
	connection := pb.NewAdRetrieverClient(conn)
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	positionId, err := strconv.ParseInt(id, 10, 0)
	minCpc, err2 := strconv.ParseInt(cpc, 10, 0)
	if err != nil || err2 != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "you enter wrong id"})
		return
	}
	r, err := connection.GetAds(ctx, &pb.TargetingRequest{Id: positionId, MinCpc: minCpc})
	if err != nil {
		c.IndentedJSON(http.StatusServiceUnavailable, gin.H{"message": "service not available"})
		return
	}
	if r == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "ads not found"})
		return
	}
	if r.GetTitle() == "" || r.GetImage() == "" {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "ads not found"})
		return
	}
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
