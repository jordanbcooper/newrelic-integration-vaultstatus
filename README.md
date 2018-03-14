# newrelic-integration-vaultstatus

Be sure to set environment variable for VAULT_URL:

export VAULT_URL=https://DNS_NAME/v1/sys/health

Sample output:

{  
   "name":"com.org.vaultstatus",
   "protocol_version":"1",
   "integration_version":"0.1.0",
   "metrics":[  
      {  
         "Sealed":"false",
         "event_type":"VaultStatus"
      }
   ],
   "inventory":{  
      "cluster":{  
         "clusterid":"abc12345-ab12-ab12-12ab-123456789012",
         "clustername":"vault-cluster-123456"
      }
   },
   "events":[  

   ]
}
