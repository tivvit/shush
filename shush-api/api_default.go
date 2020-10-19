/*
 * Shush API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package shush_api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/tivvit/shush/shush/backend"
	"github.com/tivvit/shush/shush/model"
	"io/ioutil"
	"net/http"
	"strconv"
)

func UrlsGet(w http.ResponseWriter, r *http.Request) {
	all, err := bck.GetAll()
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	js, err := json.Marshal(all)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	returnJson(w, http.StatusOK, js)
	return
}

func returnJson(w http.ResponseWriter, statusCode int, message []byte) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
	_, err := w.Write(message)
	if err != nil {
		log.Warnf("response write failed %s", err.Error())
	}
}

func UrlsPost(w http.ResponseWriter, r *http.Request) {
	url := model.Url{}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Warn("malformed request body", err)
		returnJson(w, http.StatusBadRequest, []byte(`{"error": "malformed request body"}`))
		return
	}
	err = json.Unmarshal(b, &url)
	if err != nil {
		log.Warn("malformed json", err)
		returnJson(w, http.StatusBadRequest, []byte(`{"error": "malformed json body"}`))
		return
	}
	// shortUrl us provided
	if url.ShortUrl != "" {
		// todo inform user (about using bad method?)
		if !short.IsValidShort(url.ShortUrl) {
			returnJson(w, http.StatusBadRequest, []byte(fmt.Sprintf(`{"error": "invalid short_url %s"}`, url.ShortUrl)))
			return
		}
		err := storeUrl(bck, url.ShortUrl, url)
		if err != nil {
			log.Error(err)
			// todo internal?
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else {
		shortener := short.Conf.DefaultShortener
		shortenerParam := r.URL.Query().Get("shortener")
		if shortenerParam != "" {
			if allowed, ok := short.Conf.AllowedShorteners[shortenerParam]; ok && allowed {
				shortener = shortenerParam
			} else {
				returnJson(w, http.StatusBadRequest, []byte(fmt.Sprintf(`{"error": "unknown shortener type %s"}`, shortenerParam)))
				return
			}
		}
		hashAlgo := short.Conf.DefaultHashAlgo
		hashAlgoParam := r.URL.Query().Get("algo")

		if hashAlgoParam != "" {
			if shortener != "shortner" {
				returnJson(w, http.StatusBadRequest, []byte(`{"error": "defined hash type for non-hash shortener"}`))
				return
			}
			if allowed, ok := short.Conf.AllowedHashAlgo[hashAlgoParam]; ok && allowed {
				hashAlgo = hashAlgoParam
			} else {
				returnJson(w, http.StatusBadRequest, []byte(fmt.Sprintf(`{"error": "unknown hash type %s"}`, hashAlgoParam)))
				return
			}
		}
		ln := short.Conf.DefaultLen
		lenParamStr := r.URL.Query().Get("len")
		if lenParamStr != "" {
			lenParam, err := strconv.Atoi(lenParamStr)
			if err != nil {
				returnJson(w, http.StatusBadRequest, []byte(fmt.Sprintf(`{"error": "invalid int len %s"}`, lenParamStr)))
				return
			} else {
				if lenParam > short.Conf.Maxlen {
					returnJson(w, http.StatusBadRequest, []byte(fmt.Sprintf(`{"error": "len exceeds the limit %d > %d"}`, lenParam,  short.Conf.Maxlen)))
					return
				}
				ln = lenParam
			}
		}
		switch shortener {
		case "generator":
			err := short.Random(&url, ln)
			if err != nil {
				log.Error(err)
				// todo does not have to be internal
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		case "hash":
			err := short.Hash(&url, hashAlgo, ln)
			if err != nil {
				log.Error(err)
				// todo does not have to be internal
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	}
	su, err := model.UrlSerialize(url)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(su))
	if err != nil {
		log.Warnf("response write failed %s", err.Error())
	}
}

func UrlsShortUrlDelete(w http.ResponseWriter, r *http.Request) {
	p := mux.Vars(r)
	sUrl := p["short_url"]
	err := bck.Remove(sUrl)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func UrlsShortUrlGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	p := mux.Vars(r)
	sUrl := p["short_url"]
	v, err := bck.GetRaw(sUrl)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(v))
	if err != nil {
		log.Warnf("response write failed %s", err.Error())
	}
}

func UrlsShortUrlPut(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	p := mux.Vars(r)
	sUrl := p["short_url"]
	url := model.Url{}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		returnJson(w, http.StatusBadRequest, []byte(`{"error": "malformed request body"}`))
		return
	}
	err = json.Unmarshal(b, &url)
	if err != nil {
		returnJson(w, http.StatusBadRequest, []byte(`{"error": "malformed json body"}`))
		return
	}
	// todo validate struct
	if url.ShortUrl == "" {
		url.ShortUrl = sUrl // todo this is not needed and probably not a great idea
		if !short.IsValidShort(url.ShortUrl) {
			returnJson(w, http.StatusBadRequest, []byte(fmt.Sprintf(`{"error": "invalid short_url %s"}`, url.ShortUrl)))
			return
		}
	}
	err = storeUrl(bck, sUrl, url)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func storeUrl(bck *backend.ShushBackend, sUrl string, url model.Url) error {
	return bck.Set(sUrl, url, url.Expires())
}
