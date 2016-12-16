package saltstack

/**
 * My saltstack installation do not support websocket and i could not figure out
 * why, so use events instead. May change into ws request later.
 */

import (
  "io"
	"fmt"
  "time"
  "bufio"
  "bytes"
  "errors"
  "io/ioutil"
  "net/http"
	log "github.com/Sirupsen/logrus"
  "github.com/parnurzeal/gorequest"
	"github.com/jellybean4/gosalt/util"
  "strings"
)

func startEventReceiver() {
	outer := make(chan []byte, 1024)
	go consumeEvent(outer)
	defer close(outer)

	defer func() {
		if err := recover(); err != nil {
			log.WithFields(log.Fields{
				"func":   "saltstack.startEventReceiver",
				"reason": err,
			}).Error(util.UNKNOWN_ERROR_LOG)
			go startEventReceiver()
		}
	}()

	for {
		url := fmt.Sprintf("http://%s:%s/events", util.GetConfig(util.MASTER),
			util.GetConfig(util.NETAPI_PORT))

		token, err := saltToken(false)
		if err != nil {
			time.Sleep(30 * time.Second)
			continue
		}
		agent := gorequest.New().Get(url).
			SetDebug(true).
			Set("Aceept", "application/json").
			Set("X-Auth-Token", token)
		if req, err := agent.MakeRequest(); err != nil {
			log.WithFields(log.Fields{
				"url":    url,
				"reason": err.Error(),
			}).Error(util.SALT_EVENT_LISTEN_LOG)
		} else if resp, err := agent.Client.Do(req); err != nil {
			log.WithFields(log.Fields{
				"req":    req,
				"reason": err.Error(),
			}).Error(util.SALT_EVENT_LISTEN_LOG)
		} else if resp.StatusCode != http.StatusOK {
			if body, err := ioutil.ReadAll(resp.Body); err != nil {
				log.WithFields(log.Fields{
					"req":    req,
					"resp":   err.Error(),
					"status": resp.Status,
				}).Error(util.SALT_EXEC_LOG)
			} else {
				log.WithFields(log.Fields{
					"req":    req,
					"resp":   body,
					"status": resp.Status,
				}).Error(util.SALT_EXEC_LOG)
			}
			resp.Body.Close()
		} else {
      receiveEvent(resp.Body, outer)
      resp.Body.Close()
		}
		time.Sleep(30 * time.Second)
	}
}

func receiveEvent(source io.ReadCloser, outer chan []byte) error {
	defer util.Recover()
	reader := bufio.NewReaderSize(source, 4096)

  if line, err := readLine(source); err != nil {
    return err
  } else if strings.Compare(string(line), "retry: 400\n") != 0 {
    reason := "first line not retry:400"
    log.WithFields(log.Fields{
      "line": line,
      "reason": reason,
    }).Error(util.SALT_EVENT_RECV_LOG)
    return errors.New(reason)
  }

	for {
		if tag, err := readLine(reader); err != nil {
			return err
		} else if event, err := readLine(reader); err != nil {
			return err
		} else if  !strings.HasPrefix("tag:") {
      log.WithFields(log.Fields{
        ""
      })
    } else if !strings.HasPrefix("data:") {

    } else {

    }
	}
}

func readLine(reader *bufio.Reader) ([]byte, error) {
	var buffer *bytes.Buffer

	for {
		if line, err := reader.ReadSlice('\n'); err == bufio.ErrBufferFull {
			if buffer == nil {
				buffer = new(bytes.Buffer)
			}
			buffer.Write(line)
		} else if err != nil {
			log.WithFields(log.Fields{
				"reason": err.Error(),
			}).Error(util.SALT_EVENT_RECV_LOG)
			return nil, err
		} else if buffer != nil {
			buffer.Write(line)
			return buffer.Bytes(), nil
		} else {
			return line, nil
		}
	}
}

func consumeEvent(input chan []byte) {
	for {
		if message, ok := <-input; !ok {
			return
		} else {
			log.WithFields(log.Fields{
				"message": string(message),
			}).Debug("salt consumer receive message")
		}
	}
}
