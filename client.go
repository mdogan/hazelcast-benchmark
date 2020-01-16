package main

import (
    "net"
    "strconv"

    "github.com/hazelcast/hazelcast-go-client"
    "github.com/hazelcast/hazelcast-go-client/config/property"
    "github.com/hazelcast/hazelcast-go-client/core/logger"
)

func newClient() (hazelcast.Client, error) {
    config := hazelcast.NewConfig()
    config.SetProperty(property.LoggingLevel.Name(), logger.ErrorLevel)
    config.GroupConfig().SetName(clusterName)
    networkConfig := config.NetworkConfig()
    networkConfig.AddAddress(net.JoinHostPort(host, strconv.Itoa(port)))
    return hazelcast.NewClientWithConfig(config)
}
