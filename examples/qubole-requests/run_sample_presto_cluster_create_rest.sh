X_AUTH_TOKEN=
curl -X POST -H "X-AUTH-TOKEN:$X_AUTH_TOKEN" -H "Content-Type:application/json" -H "Accept: application/json" -d @create-presto-cluster.json https://us.qubole.com/api/v1.3/clusters