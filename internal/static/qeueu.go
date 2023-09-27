package static

const (
	QUEUE_EXCHANGE_DEFAULT = ""
	QUEUE_EXCHANGE_TOPIC   = "miduoduo-dev-topic"
	QUEUE_EXCHANGE_FANOUT  = "miduoduo-dev-fanout"

	QUEUE_NAME_KASUMI_TOKEN_USAGE            = "kasumi_token_usage"
	QUEUE_NAME_KASUMI_REQUIRE_APP_BOARDCAST  = "kasumi_require_app_info"
	QUEUE_NAME_KASUMI_SELECT_PROXY_BOARDCAST = "kasumi_proxy_select_boardcast"
	QUEUE_NAME_KASUMI_INFO_APP_BOARDCAST     = "kasumi_info_app_boardcast"
	QUEUE_NAME_MISTRIPE_WEBHOOK              = "miduoduo_stripe_webhook"
	QUEUE_NAME_MIDUODUO_RPC                  = "miduoduo_rpc"
)
