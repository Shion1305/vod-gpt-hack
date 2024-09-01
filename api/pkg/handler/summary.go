package handler

import (
	"api/pkg/domain"
	infraDynamo "api/pkg/infra/dynamo"
	"api/pkg/schema"
	"api/pkg/service"
	"api/pkg/uc"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	"log"
	"net/http"
	"sync"
)

type SummaryHandler struct {
	d     *uc.GetTranscript
	cache reqSummaryCache
	br    *service.BedrockService
}

func NewSummaryHandler(d *infraDynamo.Dynamo) *SummaryHandler {
	br, err := service.NewBedrockService()
	if err != nil {
		log.Fatalf("failed to init bedrock service: %v", err)
	}
	return &SummaryHandler{
		d: uc.NewGetTranscript(d),
		cache: reqSummaryCache{
			store: make(map[uuid.UUID]domain.SummaryRequest),
		},
		br: br,
	}
}

func (h *SummaryHandler) Start() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req schema.SummaryRequest
		if err := c.BindJSON(&req); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		reqVID, err := uuid.Parse(req.VID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("invalid uuid, %v: %v", req.VID, err).Error()})
			return
		}
		summaryRequest := domain.SummaryRequest{
			From: req.From,
			To:   req.To,
			VID:  reqVID,
		}
		// issue UUID
		reqID := summaryRequest.VID
		h.cache.add(reqID, summaryRequest)
		// send redirect to GET /api/v1/summary/:id
		c.AbortWithStatusJSON(http.StatusCreated, gin.H{"id": reqID})
		return
	}
}

func (h *SummaryHandler) Send() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqID := c.Param("id")
		reqUUID, err := uuid.Parse(reqID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println("reqID: ", reqID)
		req := h.cache.pop(reqUUID)
		if req == nil {
			req = &domain.SummaryRequest{
				From: 0,
				To:   100,
				VID:  reqUUID,
			}
		}
		//// クエリの実行
		resp, err := h.d.ExecuteGrouped(c, req.VID.String(), req.From, req.To)
		if err != nil {
			fmt.Println(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		systemInput := "あなたは優秀なAIアシスタントです。ユーザーからは、動画の文字起こしデータが提供されます。簡潔に要約してください。なお、回答の冒頭には、「はい、わかりました」などの前置きは不要です。\n"
		fmt.Println("start streaming")
		completion, err := h.br.ClaudeMessageStreamCompletion(systemInput, resp.Content)
		if err != nil {
			fmt.Println(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		fullResp := ""
		c.Stream(func(w io.Writer) bool {
			if msg, ok := <-completion; ok {
				c.SSEvent("delta", msg)
				fullResp += msg
				log.Printf("received delta resp: %s\n", msg)
				return true
			}
			log.Printf("received all messages\n")
			return false
		})
		fmt.Println(fullResp)
	}
}

type reqSummaryCache struct {
	m     sync.Mutex
	store map[uuid.UUID]domain.SummaryRequest
}

func (c *reqSummaryCache) add(id uuid.UUID, req domain.SummaryRequest) {
	c.m.Lock()
	defer c.m.Unlock()
	c.store[id] = req
}

func (c *reqSummaryCache) pop(id uuid.UUID) *domain.SummaryRequest {
	c.m.Lock()
	defer c.m.Unlock()
	req, ok := c.store[id]
	if !ok {
		return nil
	}
	delete(c.store, id)
	return &req
}
