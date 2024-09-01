package handler

import (
	"api/pkg/domain"
	infraDynamo "api/pkg/infra/dynamo"
	"api/pkg/schema"
	"api/pkg/service"
	"api/pkg/uc"
	"fmt"
	"github.com/google/uuid"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type ChatHandler struct {
	d     *uc.GetTranscript
	cache reqCache
	br    *service.BedrockService
}

func NewChatHandler(d *infraDynamo.Dynamo) *ChatHandler {
	br, err := service.NewBedrockService()
	if err != nil {
		log.Fatalf("failed to init bedrock service: %v", err)
	}
	return &ChatHandler{
		d: uc.NewGetTranscript(d),
		cache: reqCache{
			store: make(map[uuid.UUID]domain.ChatRequest),
		},
		br: br,
	}
}

func (h *ChatHandler) Start() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req schema.ChatRequest
		if err := c.BindJSON(&req); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		reqVID, err := uuid.Parse(req.VID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("invalid uuid, %v: %v", req.VID, err).Error()})
			return
		}
		chatRequest := domain.ChatRequest{
			Question: req.Question,
			From:     req.From,
			To:       req.To,
			VID:      reqVID,
		}
		// issue UUID
		reqID := chatRequest.VID
		h.cache.add(reqID, chatRequest)
		// send redirect to GET /api/v1/chat/:id
		c.AbortWithStatusJSON(http.StatusCreated, gin.H{"id": reqID})
		return
	}
}

func (h *ChatHandler) Send() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqID := c.Param("id")
		reqUUID, err := uuid.Parse(reqID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		req := h.cache.pop(reqUUID)
		//// クエリの実行
		//resp, err := h.d.Execute(c, req.VID.String(), req.From, req.To)
		//if err != nil {
		//	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		//	return
		//}
		sampleTranscript := "先ほどの議題に戻りますが、この部分についてはもう少し詳細に検討する必要があるかと思います。特に、リスク管理の観点から見た場合、現状のプロセスでは対応が難しい点がいくつか見受けられます。そこで、次のステップとして、各チームから具体的な課題と改善提案を集めて、来週のミーティングで共有することを提案します。また、コストの見積もりについても再度精査が必要ですので、財務チームと連携して進めていきます。皆さんの意見をお聞かせください。"
		sampleTranscript += "\n---\n"
		sampleTranscript += "これはある動画の文字起こしです。あなたは優秀なAIアシスタントです。文字起こしの内容を元に、ユーザーの質問に答えてください。\n\n"
		//userMessage := "文字起こしで触れられていることを要約して簡潔に答えてください。\n"

		completion, err := h.br.ClaudeMessageStreamCompletion(sampleTranscript, req.Question)
		if err != nil {
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

func (ch *ChatHandler) SendDummy() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqID := c.Param("id")
		reqUUID, err := uuid.Parse(reqID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		req := ch.cache.pop(reqUUID)
		if req == nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "request not found"})
			return
		}
		sampleText := "あなたのリクエストを受け付けました。uuidは" + reqID + "です。質問は" + req.Question + "です。開始時間は" + fmt.Sprintf("%f", req.From) + "です。終了時間は" + fmt.Sprintf("%f", req.To) + "です。"
		sampleText += sampleText
		sampleText += sampleText
		sampleText += sampleText
		sampleText += sampleText
		sampleRune := []rune(sampleText)

		respCh := make(chan string)
		go func() {
			time.Sleep(50 * time.Millisecond)
			// 2文字ずつ返す
			for i := 0; i < len(sampleRune); i += 2 {
				if i+2 > len(sampleRune) {
					respCh <- string(sampleRune[i:])
					break
				}
				respCh <- string(sampleRune[i : i+2])
				fmt.Println(string(sampleRune[i : i+2]))
				time.Sleep(100 * time.Millisecond)
			}
			close(respCh)
		}()

		fullResp := ""
		c.Stream(func(w io.Writer) bool {
			if msg, ok := <-respCh; ok {
				c.SSEvent("delta", msg)
				fullResp += msg
				log.Printf("received delta resp: %s\n", msg)
				return true
			}
			log.Printf("received all messages\n")
			return false
		})
	}
}

type reqCache struct {
	m     sync.Mutex
	store map[uuid.UUID]domain.ChatRequest
}

func (c *reqCache) add(id uuid.UUID, req domain.ChatRequest) {
	c.m.Lock()
	defer c.m.Unlock()
	c.store[id] = req
}

func (c *reqCache) pop(id uuid.UUID) *domain.ChatRequest {
	c.m.Lock()
	defer c.m.Unlock()
	req, ok := c.store[id]
	if !ok {
		return nil
	}
	delete(c.store, id)
	return &req
}
