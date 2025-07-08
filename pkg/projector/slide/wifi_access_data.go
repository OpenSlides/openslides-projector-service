package slide

import (
	"context"
	"fmt"
	"strings"
)

func WifiAccessDataSlideHandler(ctx context.Context, req *projectionRequest) (map[string]any, error) {
	if req.ContentObjectID == nil {
		return nil, fmt.Errorf("no meeting id provided for slide")
	}

	var wlanData struct {
		SSID       string
		Password   string
		Encryption string
		QrString   string
	}
	req.Fetch.Meeting_UsersPdfWlanSsid(*req.ContentObjectID).Lazy(&wlanData.SSID)
	req.Fetch.Meeting_UsersPdfWlanPassword(*req.ContentObjectID).Lazy(&wlanData.Password)
	req.Fetch.Meeting_UsersPdfWlanEncryption(*req.ContentObjectID).Lazy(&wlanData.Encryption)
	if err := req.Fetch.Execute(ctx); err != nil {
		return nil, fmt.Errorf("could not fetch wlan data")
	}

	if wlanData.SSID != "" && (wlanData.Encryption == "" || wlanData.Encryption == "nopass" || wlanData.Password != "") {
		wlanData.QrString = `WIFI:S:` + escapeSpecialCharactersForWiFiConfig(wlanData.SSID) + `;`
		wlanData.QrString += `T:` + wlanData.Encryption + `;`
		if wlanData.Password != "" {
			wlanData.QrString += `P:` + escapeSpecialCharactersForWiFiConfig(wlanData.Password) + `;`
		}
		wlanData.QrString += `;`
	}

	return map[string]any{
		"WlanData": wlanData,
	}, nil
}

func escapeSpecialCharactersForWiFiConfig(str string) string {
	symbols := []string{"\\", ";", ",", "\"", ":"}
	for _, symbol := range symbols {
		str = strings.ReplaceAll(str, symbol, "\\"+symbol)
	}

	return str
}
