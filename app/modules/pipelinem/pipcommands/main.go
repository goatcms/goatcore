package pipcommands

const (
	// PipClear is a pip:clear command help
	PipClear = "clear current pipeline context"
	// PipRun is a pip:run command help
	PipRun = "[name, --sandbox=terminal/docker:image, --body=required, [--wait=task1,task2], [--lock=resource1,resource2]] Run script "
	// PipSummary is a pip:summary command help
	PipSummary = "Show execution summary"
	// PipLogs is a pip:logs command help
	PipLogs = "Show execution logs"
	// PipWait is a pip:wait command help
	PipWait = "Wait for all tasks in context"
)
