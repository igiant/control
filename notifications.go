package control

import "encoding/json"

type NotificationType string

const (
	NotificationUpdate                NotificationType = "NotificationUpdate"
	NotificationDump                  NotificationType = "NotificationDump"
	NotificationLowMemory             NotificationType = "NotificationLowMemory"
	NotificationDomains               NotificationType = "NotificationDomains"
	NotificationSubWillExpire         NotificationType = "NotificationSubWillExpire"
	NotificationSubExpired            NotificationType = "NotificationSubExpired"
	NotificationLicWillExpire         NotificationType = "NotificationLicWillExpire"
	NotificationLicExpired            NotificationType = "NotificationLicExpired"
	NotificationBackupLine            NotificationType = "NotificationBackupLine"
	NotificationInterfaceSpeed        NotificationType = "NotificationInterfaceSpeed"
	NotificationSmtp                  NotificationType = "NotificationSmtp"
	NotificationLlbLine               NotificationType = "NotificationLlbLine"
	NotificationLlb                   NotificationType = "NotificationLlb"
	NotificationConnectionOnDemand    NotificationType = "NotificationConnectionOnDemand"
	NotificationConnectionFailover    NotificationType = "NotificationConnectionFailover"
	NotificationConnectionBalancing   NotificationType = "NotificationConnectionBalancing"
	NotificationConnectionPersistent  NotificationType = "NotificationConnectionPersistent"
	NotificationCertificateError      NotificationType = "NotificationCertificateError"
	NotificationCertificateWillExpire NotificationType = "NotificationCertificateWillExpire"
	NotificationCertificateExpired    NotificationType = "NotificationCertificateExpired"
	NotificationCaWillExpire          NotificationType = "NotificationCaWillExpire"
	NotificationCaExpired             NotificationType = "NotificationCaExpired"
	NotificationBackupFailed          NotificationType = "NotificationBackupFailed"
	NotificationPacketDump            NotificationType = "NotificationPacketDump"
	NotificationUnknown               NotificationType = "NotificationUnknown"
)

type NotificationTypeList []NotificationType

type NotificationSeverity string

const (
	NotificationWarning NotificationSeverity = "NotificationWarning"
	NotificationError   NotificationSeverity = "NotificationError"
)

type Notification struct {
	Type     NotificationType     `json:"type"`
	Severity NotificationSeverity `json:"severity"`
	Value    string               `json:"value"`
	Code     int                  `json:"code"`
}

type NotificationList []Notification

// NotificationsGet - Returns list of notifications without filtered (cleared) messages
// When lastNotifications are the same as current notifications, method waits until timeout occurs and than returns
// Parameters
//	lastNotifications - notifications returned by last call or empty list
//	timeout - how long should engine wait for notifications change (in seconds)
// Return
//	notifications - list of notifications
func (s *ServerConnection) NotificationsGet(lastNotifications NotificationList, timeout int) (NotificationList, error) {
	params := struct {
		LastNotifications NotificationList `json:"lastNotifications"`
		Timeout           int              `json:"timeout"`
	}{lastNotifications, timeout}
	data, err := s.CallRaw("Notifications.get", params)
	if err != nil {
		return nil, err
	}
	notifications := struct {
		Result struct {
			Notifications NotificationList `json:"notifications"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &notifications)
	return notifications.Result.Notifications, err
}

// NotificationsClear - Clears defined notification for current user
// Parameters
//	notification - one of the notifications returned by get
func (s *ServerConnection) NotificationsClear(notification Notification) error {
	params := struct {
		Notification Notification `json:"notification"`
	}{notification}
	_, err := s.CallRaw("Notifications.clear", params)
	return err
}
