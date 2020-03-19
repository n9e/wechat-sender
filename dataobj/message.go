package dataobj

import "time"

type Message struct {
	Tos              []string `json:"tos"`
	Event            *Event   `json:"event"`
	ClaimLink        string   `json:"claim_link"`
	StraLink         string   `json:"stra_link"`
	EventLink        string   `json:"event_link"`
	Bindings         []string `json:"bindings"`
	NotifyType       string   `json:"notify_type"`
	Metrics          []string `json:"metrics"`
	ReadableEndpoint string   `json:"readable_endpoint"`
	ReadableTags     string   `json:"readable_tags"`
	IsUpgrade        bool     `json:"is_upgrade"`
}

type Event struct {
	Id            int64     `json:"id"`
	Sid           int64     `json:"sid"`
	Sname         string    `json:"sname"`
	NodePath      string    `json:"node_path"`
	Endpoint      string    `json:"endpoint"`
	EndpointAlias string    `json:"endpoint_alias"`
	Priority      int       `json:"priority"`
	EventType     string    `json:"event_type"` // alert|recovery
	Category      int       `json:"category"`
	Status        uint16    `json:"status"`
	HashId        uint64    `json:"hashid"  xorm:"hashid"`
	Etime         int64     `json:"etime"`
	Value         string    `json:"value"`
	Info          string    `json:"info"`
	Created       time.Time `json:"created" xorm:"created"`
	Detail        string    `json:"detail"`
	Users         string    `json:"users"`
	Groups        string    `json:"groups"`
	Nid           int64     `json:"nid"`
	NeedUpgrade   int       `json:"need_upgrade"`
	AlertUpgrade  string    `json:"alert_upgrade"`
}

type EventDetail struct {
	Metric     string              `json:"metric"`
	Tags       map[string]string   `json:"tags"`
	Points     []*EventDetailPoint `json:"points"`
	PredPoints []*EventDetailPoint `json:"pred_points,omitempty"` // 预测值, 预测值不为空时, 现场值对应的是实际值
}

type EventDetailPoint struct {
	Timestamp int64   `json:"timestamp"`
	Value     float64 `json:"value"`
}
