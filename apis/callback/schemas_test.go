package callback

import (
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMipushCallbackData(t *testing.T) {

	// https://dev.mi.com/distribute/doc/details?pId=1558#_36
	jsonString := `{
        "msgId1":{"param":"param","type": 1, "targets":"alias1,alias2,alias3", "jobkey": "123" ,"barStatus":"Enable","timestamp":1324167800000},
        "msgId2":{"param":"param","type": 2, "targets":"alias1,alias2,alias3", "jobkey": "456", "barStatus": "Enable", "timestamp": 1524187800000},
        "msgId3":{"param":"param","type":16,"targets":"alias1,alias2,alias3","barStatus":"Unknown","timestamp":1572228055643},
        "msgId4":{"param":"param","type":16,"targets":"regId1,regId2,regId3","barStatus":"Unknown","errorCode":1,"timestamp":1572228055643,"replaceTarget":{"regId1":"otherRegId"}},
        "msgId5":{"param":"param","type":64,"targets":"regId1,regId2,regId3", "barStatus":"Unknown","timestamp":1572228055643},
        "msgId6": {"extra":{"ack":"当日已送达数","quota":"当日可以下发总数"},"type":128,"targets":"alias","timestamp":1585203103625},
        "msgId7": {"extra":{"device_acked":"当日单设备已接收数","device_quota":"当日单设备可以下发总数"},"type":128,"targets":"alias","timestamp":1585203104390},
        "msgId8": {"param":"param","type":1024,"targets":"regId1,regId2,regId3", "timestamp":1572228055643}
	}`

	var data map[string]MipushCallbackData

	err := json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		t.Error(err)
	}

	assert.EqualValues(t, "param", data["msgId1"].Param)
	assert.EqualValues(t, 1, data["msgId1"].Type)
	assert.EqualValues(t, "alias1,alias2,alias3", data["msgId1"].Targets)
}
