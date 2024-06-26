// Code AutoGenerated; DO NOT EDIT.

package types

// WriteAccessAllowed Represents a service message about a user allowing a bot to write messages after adding it to the attachment menu, launching a Web App from a link, or accepting an explicit request from a Web App sent by the method requestWriteAccess.
type WriteAccessAllowed struct {
	FromAttachmentMenu bool   `json:"from_attachment_menu,omitempty"`
	FromRequest        bool   `json:"from_request,omitempty"`
	WebAppName         string `json:"web_app_name,omitempty"`
}
