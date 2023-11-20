package credentials

import "github.com/alibabacloud-go/tea/tea"

// StsTokenCredential is a kind of credentials
type StsTokenCredential struct {
	AccessKeyId     string
	AccessKeySecret string
	SecurityToken   string
}

func newStsTokenCredential(accessKeyId, accessKeySecret, securityToken string) *StsTokenCredential {
	return &StsTokenCredential{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
	}
}

func (s *StsTokenCredential) GetCredential() (*CredentialModel, error) {
	credential := &CredentialModel{
		AccessKeyId:     tea.String(s.AccessKeyId),
		AccessKeySecret: tea.String(s.AccessKeySecret),
		SecurityToken:   tea.String(s.SecurityToken),
		Type:            tea.String("sts"),
	}
	return credential, nil
}

// GetAccessKeyId reutrns  StsTokenCredential's AccessKeyId
func (s *StsTokenCredential) GetAccessKeyId() (*string, error) {
	return tea.String(s.AccessKeyId), nil
}

// GetAccessSecret reutrns  StsTokenCredential's AccessKeySecret
func (s *StsTokenCredential) GetAccessKeySecret() (*string, error) {
	return tea.String(s.AccessKeySecret), nil
}

// GetSecurityToken reutrns  StsTokenCredential's SecurityToken
func (s *StsTokenCredential) GetSecurityToken() (*string, error) {
	return tea.String(s.SecurityToken), nil
}

// GetBearerToken is useless StsTokenCredential
func (s *StsTokenCredential) GetBearerToken() *string {
	return tea.String("")
}

// GetType reutrns  StsTokenCredential's type
func (s *StsTokenCredential) GetType() *string {
	return tea.String("sts")
}
