# newrelic-integration-vaultstatus

Be sure to set environment variable for VAULT_URL:

export VAULT_URL=https://DNS_NAME/v1/sys/health

Sample output:

{
        "name": "com.org.vaultstatus",
        "protocol_version": "1",
        "integration_version": "0.1.0",
        "metrics": [
                {
                        "Sealed": "false",
                        "event_type": "VaultStatus"
                }
        ],
        "inventory": {
                "cluster": {
                        "clusterid": "fad215c5-afc2-fab0-12f1-2af59c9448a8",
                        "clustername": "vault-cluster-3f1d247a"
                }
        },
        "events": []
}
