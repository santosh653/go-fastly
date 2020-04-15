package fastly

import "testing"

func TestClient_GetPlatformPrivateKeys(t *testing.T) {
	t.Parallel()

	var err error
	record(t, "platform_tls/platform_private_keys", func(c *Client) {
		_, err = c.GetPlatformPrivateKeys()
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_GetPlatformPrivateKey(t *testing.T) {
	t.Parallel()

	var err error
	record(t, "platform_tls/platform_private_key", func(c *Client) {
		_, err = c.GetPlatformPrivateKey(&GetPlatformPrivateKeyInput{
			ID: "PRIVATE_KEY_ID",
		})
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_CreatePlatformPrivateKey(t *testing.T) {
	t.Parallel()

	var err error
	record(t, "platform_tls/create_private_key", func(c *Client) {
		_, err = c.CreatePlatformPrivateKey(&CreatePlatformPrivateKeyInput{
			Data: CreatePlatformPrivateKeyData{
				Type: "tls_private_key",
				Attributes: CreatePlatformPrivateKeyAttributes{
					Key:  "-----BEGIN PRIVATE KEY-----\n...\n-----END PRIVATE KEY-----\n",
					Name: "My private key",
				},
			},
		})
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_CreateBulkCertificate(t *testing.T) {
	t.Parallel()

	var err error
	record(t, "platform_tls/create_bulk_certificate", func(c *Client) {
		_, err = c.CreateBulkCertificate(&CreateBulkCertificatesInput{
			Data: CreateBulkCertificatesData{
				Type: "tls_bulk_certificate",
				Attributes: CreateBulkCertificatesAttributes{
					CertBlob:          "-----BEGIN CERTIFICATE-----\n...\n-----END CERTIFICATE-----\n",
					IntermediatesBlob: "-----BEGIN CERTIFICATE-----\n...\n-----END CERTIFICATE-----\n",
				},
				Relationships: CreateBulkCertificatesRelationships{
					TLSConfigurations: CreateTLSConfigurations{
						Data: []CreateTLSConfiguration{
							{
								Type: "tls_configuration",
								ID:   "TLS_CONFIGURATION_ID",
							},
						},
					},
				},
			},
		})
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_DeletePrivateKey(t *testing.T) {
	t.Parallel()

	var err error
	record(t, "platform_tls/delete_private_key", func(c *Client) {
		err = c.DeletePrivateKey(&DeletePrivateKeyInput{
			ID: "PRIVATE_KEY_ID",
		})
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_DeleteBulkCertificate(t *testing.T) {
	t.Parallel()

	var err error
	record(t, "platform_tls/delete_bulk_certificate", func(c *Client) {
		err = c.DeleteBulkCertificate(&DeleteBulkCertificateInput{
			ID: "BULK_CERTIFICATE_ID",
		})
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_GetBulkCertificates(t *testing.T) {
	t.Parallel()

	var err error
	record(t, "platform_tls/platform_bulk_certificates", func(c *Client) {
		_, err = c.GetBulkCertificates(&GetBulkCertificateInput{})
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_GetBulkCertificate(t *testing.T) {
	t.Parallel()

	var err error
	record(t, "platform_tls/platform_bulk_certificate", func(c *Client) {
		_, err = c.GetBulkCertificate(&GetBulkCertificateInput{
			ID: "BULK_CERTIFICATE_ID",
		})
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_UpdateBulkCertificate(t *testing.T) {
	t.Parallel()

	var err error
	record(t, "platform_tls/update_bulk_certificate", func(c *Client) {
		_, err = c.UpdateBulkCertificate(&UpdateBulkCertificateInput{
			Data: UpdateBulkCertificateData{
				ID:                "CERTIFICATE_ID",
				Type:              "tls_bulk_certificate",
				CertBlob:          "-----BEGIN CERTIFICATE-----\n...\n-----END CERTIFICATE-----\n",
				IntermediatesBlob: "-----BEGIN CERTIFICATE-----\n...\n-----END CERTIFICATE-----\n",
			},
		})
	})
	if err != nil {
		t.Fatal(err)
	}
}
