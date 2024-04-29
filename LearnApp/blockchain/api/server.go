package api

import (
	"encoding/gob"
	"encoding/hex"
	"log"
	"net/http"

	"github.com/anthdm/projectx/core"
	"github.com/anthdm/projectx/types"
	"github.com/go-kit/log"
	"github.com/labstack/echo/v4"
)

type TxResponse struct {
	TxCount uint
	Hashes  []string
}

type APIError struct {
	Error string
}

type Blogc struct {
	Hash          string
	Version       uint32
	DataHash      string
	PrevBlockHash string
	Height        uint32
	Timestamp     int64
	Validator     string
	Signature     string
	TxResponse    TxResponse
}

type ServerConfig struct {
	Logger     log.Logger
	ListenAddr string
}

type Server struct {
	txChan chan *core.Transaction
	ServerConfig
	bc *core.Blockchain
}

func NewServer(cfg ServerConfig, bc *core.Blockchain, txChan chan *core.Transaction) *Server {
	return &Server{
		ServerConfig: cfg,
		bc:           bc,
		txChan:       txChan,
	}
}

func (s *Server) Start() error {
	e := echo.New()
	e.Get("/block/:hashorid", s.handleGetBlock)
	e.Get("/tx/:hash", s.handleGetTx)
	e.Post("/tx", s.handlePostTx)
	return e.Start(s.ListenAddr)
}

func (s *Server) handlePostTx(c echo.Context) error {
	tx := &core.Transaction{}
	if err := gob.NewDecoder(c.Request().Body).Decode(tx); err != nil {
		return c.JSON(http.StatusBadRequest, APIError{Error: err.Error()})
	}
	s.txChan <- tx
	return nil
}

func (s *Server) handleGetTx(c echo.Context) error {
	hash := c.Param("hash")
	b, err := hex.DecodeString(hash)
	if err != nil {
		return c.JSON(http.StatusBadRequest, APIError{Error: err.Error()})
	}
	tx, err := s.bc.GetTxByHash(types.HashFromBytes(b))
	if err != nil {
		return c.JSON(http.StatusBadRequest, APIError{Error: err.Error()})
	}

	return c.JSON(http.StatusOK, tx)
}
