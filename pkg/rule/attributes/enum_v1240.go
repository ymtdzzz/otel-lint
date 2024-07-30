// Code generated by otel-lint. DO NOT EDIT.
package attributes

import (
	"slices"

	"github.com/ymtdzzz/otel-lint/pkg/rule"
)

var RulesEnumv1240 = []*rule.AttributeRule{
	ruleEnumAndroidStatev1240,
	ruleEnumAwsEcsLaunchtypev1240,
	ruleEnumDbCosmosdbConnectionModev1240,
	ruleEnumDiskIoDirectionv1240,
	ruleEnumFaasTriggerv1240,
	ruleEnumGraphqlOperationTypev1240,
	ruleEnumIosStatev1240,
	ruleEnumJvmMemoryTypev1240,
	ruleEnumJvmThreadStatev1240,
	ruleEnumLogIostreamv1240,
	ruleEnumMessagingRocketmqConsumptionModelv1240,
	ruleEnumMessagingRocketmqMessageTypev1240,
	ruleEnumNetworkIoDirectionv1240,
	ruleEnumOpentracingRefTypev1240,
	ruleEnumOtelStatusCodev1240,
	ruleEnumStatev1240,
	ruleEnumSystemFilesystemStatev1240,
	ruleEnumSystemNetworkStatev1240,
	ruleEnumSystemPagingDirectionv1240,
	ruleEnumSystemPagingStatev1240,
	ruleEnumSystemPagingTypev1240,
}

var ruleEnumAndroidStatev1240 = &rule.AttributeRule{
	Name:         "enum.android.state",
	Title:        "This attribute represents the state the application has transitioned into at the occurrence of the event.",
	Check:        checkEnumAndroidStatev1240,
	Severity:     rule.SeverityError,
	Stability:    rule.StabilityExperimental,
	Source:       "",
}

func checkEnumAndroidStatev1240(a *rule.SignalAttributes) bool {
  val, ok := a.Get("android.state")
  if !ok {
    return true
  }
  return slices.Contains([]string{ "created", "background", "foreground" }, val.Str())
}

var ruleEnumAwsEcsLaunchtypev1240 = &rule.AttributeRule{
	Name:         "enum.aws.ecs.launchtype",
	Title:        "The [launch type](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/launch_types.html) for an ECS task.",
	Check:        checkEnumAwsEcsLaunchtypev1240,
	Severity:     rule.SeverityError,
	Stability:    rule.StabilityExperimental,
	Source:       "",
}

func checkEnumAwsEcsLaunchtypev1240(a *rule.SignalAttributes) bool {
  val, ok := a.Get("aws.ecs.launchtype")
  if !ok {
    return true
  }
  return slices.Contains([]string{ "ec2", "fargate" }, val.Str())
}

var ruleEnumDbCosmosdbConnectionModev1240 = &rule.AttributeRule{
	Name:         "enum.db.cosmosdb.connection_mode",
	Title:        "Cosmos client connection mode.",
	Check:        checkEnumDbCosmosdbConnectionModev1240,
	Severity:     rule.SeverityError,
	Stability:    rule.StabilityExperimental,
	Source:       "",
}

func checkEnumDbCosmosdbConnectionModev1240(a *rule.SignalAttributes) bool {
  val, ok := a.Get("db.cosmosdb.connection_mode")
  if !ok {
    return true
  }
  return slices.Contains([]string{ "gateway", "direct" }, val.Str())
}

var ruleEnumDiskIoDirectionv1240 = &rule.AttributeRule{
	Name:         "enum.disk.io.direction",
	Title:        "The disk IO operation direction.",
	Check:        checkEnumDiskIoDirectionv1240,
	Severity:     rule.SeverityError,
	Stability:    rule.StabilityExperimental,
	Source:       "",
}

func checkEnumDiskIoDirectionv1240(a *rule.SignalAttributes) bool {
  val, ok := a.Get("disk.io.direction")
  if !ok {
    return true
  }
  return slices.Contains([]string{ "read", "write" }, val.Str())
}

var ruleEnumFaasTriggerv1240 = &rule.AttributeRule{
	Name:         "enum.faas.trigger",
	Title:        "Type of the trigger which caused this function invocation.",
	Check:        checkEnumFaasTriggerv1240,
	Severity:     rule.SeverityError,
	Stability:    rule.StabilityExperimental,
	Source:       "",
}

func checkEnumFaasTriggerv1240(a *rule.SignalAttributes) bool {
  val, ok := a.Get("faas.trigger")
  if !ok {
    return true
  }
  return slices.Contains([]string{ "datasource", "http", "pubsub", "timer", "other" }, val.Str())
}

var ruleEnumGraphqlOperationTypev1240 = &rule.AttributeRule{
	Name:         "enum.graphql.operation.type",
	Title:        "The type of the operation being executed.",
	Check:        checkEnumGraphqlOperationTypev1240,
	Severity:     rule.SeverityError,
	Stability:    rule.StabilityExperimental,
	Source:       "",
}

func checkEnumGraphqlOperationTypev1240(a *rule.SignalAttributes) bool {
  val, ok := a.Get("graphql.operation.type")
  if !ok {
    return true
  }
  return slices.Contains([]string{ "query", "mutation", "subscription" }, val.Str())
}

var ruleEnumIosStatev1240 = &rule.AttributeRule{
	Name:         "enum.ios.state",
	Title:        "This attribute represents the state the application has transitioned into at the occurrence of the event.",
	Check:        checkEnumIosStatev1240,
	Severity:     rule.SeverityError,
	Stability:    rule.StabilityExperimental,
	Source:       "",
}

func checkEnumIosStatev1240(a *rule.SignalAttributes) bool {
  val, ok := a.Get("ios.state")
  if !ok {
    return true
  }
  return slices.Contains([]string{ "active", "inactive", "background", "foreground", "terminate" }, val.Str())
}

var ruleEnumJvmMemoryTypev1240 = &rule.AttributeRule{
	Name:         "enum.jvm.memory.type",
	Title:        "The type of memory.",
	Check:        checkEnumJvmMemoryTypev1240,
	Severity:     rule.SeverityError,
	Stability:    rule.StabilityStable,
	Source:       "",
}

func checkEnumJvmMemoryTypev1240(a *rule.SignalAttributes) bool {
  val, ok := a.Get("jvm.memory.type")
  if !ok {
    return true
  }
  return slices.Contains([]string{ "heap", "non_heap" }, val.Str())
}

var ruleEnumJvmThreadStatev1240 = &rule.AttributeRule{
	Name:         "enum.jvm.thread.state",
	Title:        "State of the thread.",
	Check:        checkEnumJvmThreadStatev1240,
	Severity:     rule.SeverityError,
	Stability:    rule.StabilityStable,
	Source:       "",
}

func checkEnumJvmThreadStatev1240(a *rule.SignalAttributes) bool {
  val, ok := a.Get("jvm.thread.state")
  if !ok {
    return true
  }
  return slices.Contains([]string{ "new", "runnable", "blocked", "waiting", "timed_waiting", "terminated" }, val.Str())
}

var ruleEnumLogIostreamv1240 = &rule.AttributeRule{
	Name:         "enum.log.iostream",
	Title:        "The stream associated with the log. See below for a list of well-known values.",
	Check:        checkEnumLogIostreamv1240,
	Severity:     rule.SeverityError,
	Stability:    rule.StabilityExperimental,
	Source:       "",
}

func checkEnumLogIostreamv1240(a *rule.SignalAttributes) bool {
  val, ok := a.Get("log.iostream")
  if !ok {
    return true
  }
  return slices.Contains([]string{ "stdout", "stderr" }, val.Str())
}

var ruleEnumMessagingRocketmqConsumptionModelv1240 = &rule.AttributeRule{
	Name:         "enum.messaging.rocketmq.consumption_model",
	Title:        "Model of message consumption. This only applies to consumer spans.",
	Check:        checkEnumMessagingRocketmqConsumptionModelv1240,
	Severity:     rule.SeverityError,
	Stability:    rule.StabilityExperimental,
	Source:       "",
}

func checkEnumMessagingRocketmqConsumptionModelv1240(a *rule.SignalAttributes) bool {
  val, ok := a.Get("messaging.rocketmq.consumption_model")
  if !ok {
    return true
  }
  return slices.Contains([]string{ "clustering", "broadcasting" }, val.Str())
}

var ruleEnumMessagingRocketmqMessageTypev1240 = &rule.AttributeRule{
	Name:         "enum.messaging.rocketmq.message.type",
	Title:        "Type of message.",
	Check:        checkEnumMessagingRocketmqMessageTypev1240,
	Severity:     rule.SeverityError,
	Stability:    rule.StabilityExperimental,
	Source:       "",
}

func checkEnumMessagingRocketmqMessageTypev1240(a *rule.SignalAttributes) bool {
  val, ok := a.Get("messaging.rocketmq.message.type")
  if !ok {
    return true
  }
  return slices.Contains([]string{ "normal", "fifo", "delay", "transaction" }, val.Str())
}

var ruleEnumNetworkIoDirectionv1240 = &rule.AttributeRule{
	Name:         "enum.network.io.direction",
	Title:        "The network IO operation direction.",
	Check:        checkEnumNetworkIoDirectionv1240,
	Severity:     rule.SeverityError,
	Stability:    rule.StabilityExperimental,
	Source:       "",
}

func checkEnumNetworkIoDirectionv1240(a *rule.SignalAttributes) bool {
  val, ok := a.Get("network.io.direction")
  if !ok {
    return true
  }
  return slices.Contains([]string{ "transmit", "receive" }, val.Str())
}

var ruleEnumOpentracingRefTypev1240 = &rule.AttributeRule{
	Name:         "enum.opentracing.ref_type",
	Title:        "Parent-child Reference type",
	Check:        checkEnumOpentracingRefTypev1240,
	Severity:     rule.SeverityError,
	Stability:    rule.StabilityExperimental,
	Source:       "",
}

func checkEnumOpentracingRefTypev1240(a *rule.SignalAttributes) bool {
  val, ok := a.Get("opentracing.ref_type")
  if !ok {
    return true
  }
  return slices.Contains([]string{ "child_of", "follows_from" }, val.Str())
}

var ruleEnumOtelStatusCodev1240 = &rule.AttributeRule{
	Name:         "enum.otel.status_code",
	Title:        "Name of the code, either OK or ERROR. MUST NOT be set if the status code is UNSET.",
	Check:        checkEnumOtelStatusCodev1240,
	Severity:     rule.SeverityError,
	Stability:    rule.StabilityExperimental,
	Source:       "",
}

func checkEnumOtelStatusCodev1240(a *rule.SignalAttributes) bool {
  val, ok := a.Get("otel.status_code")
  if !ok {
    return true
  }
  return slices.Contains([]string{ "OK", "ERROR" }, val.Str())
}

var ruleEnumStatev1240 = &rule.AttributeRule{
	Name:         "enum.state",
	Title:        "The state of a connection in the pool",
	Check:        checkEnumStatev1240,
	Severity:     rule.SeverityError,
	Stability:    rule.StabilityExperimental,
	Source:       "",
}

func checkEnumStatev1240(a *rule.SignalAttributes) bool {
  val, ok := a.Get("state")
  if !ok {
    return true
  }
  return slices.Contains([]string{ "idle", "used" }, val.Str())
}

var ruleEnumSystemFilesystemStatev1240 = &rule.AttributeRule{
	Name:         "enum.system.filesystem.state",
	Title:        "The filesystem state",
	Check:        checkEnumSystemFilesystemStatev1240,
	Severity:     rule.SeverityError,
	Stability:    rule.StabilityExperimental,
	Source:       "",
}

func checkEnumSystemFilesystemStatev1240(a *rule.SignalAttributes) bool {
  val, ok := a.Get("system.filesystem.state")
  if !ok {
    return true
  }
  return slices.Contains([]string{ "used", "free", "reserved" }, val.Str())
}

var ruleEnumSystemNetworkStatev1240 = &rule.AttributeRule{
	Name:         "enum.system.network.state",
	Title:        "A stateless protocol MUST NOT set this attribute",
	Check:        checkEnumSystemNetworkStatev1240,
	Severity:     rule.SeverityError,
	Stability:    rule.StabilityExperimental,
	Source:       "",
}

func checkEnumSystemNetworkStatev1240(a *rule.SignalAttributes) bool {
  val, ok := a.Get("system.network.state")
  if !ok {
    return true
  }
  return slices.Contains([]string{ "close", "close_wait", "closing", "delete", "established", "fin_wait_1", "fin_wait_2", "last_ack", "listen", "syn_recv", "syn_sent", "time_wait" }, val.Str())
}

var ruleEnumSystemPagingDirectionv1240 = &rule.AttributeRule{
	Name:         "enum.system.paging.direction",
	Title:        "The paging access direction",
	Check:        checkEnumSystemPagingDirectionv1240,
	Severity:     rule.SeverityError,
	Stability:    rule.StabilityExperimental,
	Source:       "",
}

func checkEnumSystemPagingDirectionv1240(a *rule.SignalAttributes) bool {
  val, ok := a.Get("system.paging.direction")
  if !ok {
    return true
  }
  return slices.Contains([]string{ "in", "out" }, val.Str())
}

var ruleEnumSystemPagingStatev1240 = &rule.AttributeRule{
	Name:         "enum.system.paging.state",
	Title:        "The memory paging state",
	Check:        checkEnumSystemPagingStatev1240,
	Severity:     rule.SeverityError,
	Stability:    rule.StabilityExperimental,
	Source:       "",
}

func checkEnumSystemPagingStatev1240(a *rule.SignalAttributes) bool {
  val, ok := a.Get("system.paging.state")
  if !ok {
    return true
  }
  return slices.Contains([]string{ "used", "free" }, val.Str())
}

var ruleEnumSystemPagingTypev1240 = &rule.AttributeRule{
	Name:         "enum.system.paging.type",
	Title:        "The memory paging type",
	Check:        checkEnumSystemPagingTypev1240,
	Severity:     rule.SeverityError,
	Stability:    rule.StabilityExperimental,
	Source:       "",
}

func checkEnumSystemPagingTypev1240(a *rule.SignalAttributes) bool {
  val, ok := a.Get("system.paging.type")
  if !ok {
    return true
  }
  return slices.Contains([]string{ "major", "minor" }, val.Str())
}
