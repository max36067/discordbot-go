package apps

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"
	"time"
)


func PictureGenerate(text string) io.Reader {

  url := "https://ai-api.magicstudio.com/api/ai-art-generator"
  method := "POST"

  payload := &bytes.Buffer{}
  writer := multipart.NewWriter(payload)
  _ = writer.WriteField("prompt", text)
  _ = writer.WriteField("output_format", "bytes")
  _ = writer.WriteField("user_profile_id", "null")
  _ = writer.WriteField("anonymous_user_id", "767ebe14-db87-496e-8915-939eac32e426")
  _ = writer.WriteField("request_timestamp", strconv.Itoa(int(time.Now().Unix())))
  _ = writer.WriteField("user_is_subscribed", "false")
  _ = writer.WriteField("client_id", "pSgX7WgjukXCBoYwDM8G8GLnRRkvAoJlqa5eAVvj95o")
  err := writer.Close()
  if err != nil {
    log.Fatal(err)
  }


  client := &http.Client {}
  req, err := http.NewRequest(method, url, payload)

  if err != nil {
    log.Fatal(err)
  }
  req.Header.Set("Content-Type", writer.FormDataContentType())
  res, err := client.Do(req)
  if err != nil {
    log.Fatal(err)
  }
  defer res.Body.Close()
  imageData, err:= io.ReadAll(res.Body)
  if err != nil {
	log.Fatal(err)
  }
  file := bytes.NewReader(imageData)
  return file
}