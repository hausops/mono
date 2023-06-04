dapr.rm-logs:
	find . -type f -path '**/.dapr/logs/*.log' -exec rm {} +
