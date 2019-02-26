package settings

import (
    "github.com/scryInfo/iscap/demo/src/sdk/core/ethereum/events"
)

type ScryInfo struct {
    Chain Chain  `yaml:"Chain"`
    Log Log      `yaml:"Log"`
}

type Chain struct {
    LastEvent events.Event  `yaml:"LastEvent"`

}

type Log struct {
    Dir string
    File string
}

