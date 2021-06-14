package iota

/*
#cgo CFLAGS: -I./include -DIOTA_STREAMS_CHANNELS_CLIENT
#cgo LDFLAGS: -L./include -liota_streams_c
#include <channels.h>
*/
import "C"
import (
	"bytes"
	"fmt"
	"github.com/project-alvarium/Alvarium-Server/configuration"
	"net/http"
	"time"
)

const PAYLOAD_LENGTH=1024

type Subscriber struct {
	subscriber *C.subscriber_t
	keyload *C.message_links_t
}

func NewSubscriber(nodeConfig configuration.NodeConfig, subConfig configuration.SubConfig) Subscriber {
	// Generate Transport client
	transport := C.tsp_client_new_from_url(C.CString(nodeConfig.Url))
	C.tsp_client_set_mwm(transport, C.uchar(nodeConfig.Mwm))

	// Generate Subscriber instance
	sub := Subscriber {
		C.sub_new(C.CString(subConfig.Seed), C.CString(subConfig.Encoding), PAYLOAD_LENGTH, transport),
		nil,
	}

	// Process announcement message
	address := C.address_from_string(C.CString(subConfig.AnnAddress))
	C.sub_receive_announce(sub.subscriber, address)

	// Fetch sub link and pk for subscription
	subLink := C.sub_send_subscribe(sub.subscriber, address)
	subPk := C.sub_get_public_key(sub.subscriber)

	subIdStr := C.get_address_id_str(subLink)
	subPkStr := C.public_key_to_string(subPk)

	fmt.Println("Sending subscription request... ", C.GoString(subIdStr))
	sendSubscriptionIdToAuthor(
		configuration.AuthConsoleUrl,
		subscriptionRequestBody(C.GoString(subIdStr), C.GoString(subPkStr)))

	// Free generated c strings from mem
	C.drop_str(subIdStr)
	C.drop_str(subPkStr)

	return sub
}

func (sub *Subscriber) InsertKeyload(keyload *C.message_links_t) {
	s := sub
	s.keyload = keyload
	*sub = *s
}

func (sub *Subscriber) SendMessage(message TangleMessage) {
	messageBytes := C.CBytes([]byte(message.message))
	messageLen := len(message.message)

	C.sub_send_signed_packet(
		sub.subscriber,
		*sub.keyload,
		nil, 0,
		(*C.uchar) (messageBytes), C.size_t(messageLen))
}

func (sub *Subscriber) Drop() {
	C.sub_drop(sub.subscriber)
	C.drop_links(*sub.keyload)
}

func (sub *Subscriber) AwaitKeyload() {
	exists := false
	for exists == false {
		// Gen next message ids to look for existing messages
		msgIds := C.sub_gen_next_msg_ids(sub.subscriber)
		// Search for keyload message from these ids and try to process it
		processed := C.sub_receive_keyload_from_ids(sub.subscriber, msgIds)
		// Free memory for c msgids object
		C.drop_next_msg_ids(msgIds)

		if processed != nil {
			// Store keyload links for attaching messages to
			sub.InsertKeyload(processed)
			exists = true
		} else {
			// Loop until keyload is found and processed
			time.Sleep(time.Second)
		}
	}
}

func subscriptionRequestBody(msgid string, pk string) []byte {
	body := "{ \"msgid\": \"" + msgid + "\", \"pk\": \"" + pk + "\" }"
	return []byte(body)
}

func sendSubscriptionIdToAuthor(url string, body []byte) {
	data := bytes.NewReader(body)
	req, err := http.NewRequest("POST", url + "/subscribe", data)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
}
