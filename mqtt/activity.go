package mqtt

import (
	"time"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/eclipse/paho.mqtt.golang"
)

// MQTTActivity implements a simple MQTT send.
type MQTTActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new MQTTActivity.
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MQTTActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata.
func (a *MQTTActivity) Metadata() *activity.Metadata {
	return a.metadata
}


// Eval implements activity.Activity.Eval
func (a *MQTTActivity) Eval(context activity.Context) (done bool, err error)  {
	// TODO: The client ID should probably be defined as a flow input parameter?
	opts := mqtt.NewClientOptions().AddBroker("tcp://localhost:1883").SetClientID("flogo")
	opts.SetKeepAlive(2 * time.Second)
	opts.SetPingTimeout(1 * time.Second)

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		err := activity.NewError(token.Error().Error(), "", nil)
		return false, err
	}

	topic := context.GetInput("topic").(string)
	/* if token := c.Subscribe(topic, 0, nil); token.Wait() && token.Error() != nil {
		err := activity.NewError(token.Error().Error(), "", nil)
		return false, err
	} */
	msg := context.GetInput("message").(string)
	if token := c.Publish(topic, 0, false, msg); token.Wait() && token.Error() != nil {
		err := activity.NewError(token.Error().Error(), "", nil)
		return false, err
	}

	return true, nil
}
