#!/bin/bash

# clean up before the test
rm -rf victoria-metrics-data
go build ./app/victoria-metrics/

# start victoria
./victoria-metrics -search.disableCache=true &
VICTORIA_PID=$!

sleep 10
echo Importing first datapoint
curl -X POST http://localhost:8428/api/v1/import -d '{"metric":{"__name__":"sample_metric"},"values":[1],"timestamps":[1666915196457]}'

sleep 30
echo Importing second datapoint
curl -X POST http://localhost:8428/api/v1/import -d '{"metric":{"__name__":"sample_metric"},"values":[2],"timestamps":[1666915211457]}'

sleep 45
curl http://127.0.0.1:8428/prometheus/api/v1/query_range -d 'query=sample_metric' -d 'start=1666908000' -d 'end=1666994400' -d 'step=39' -d 'nocache=1' -d 'trace=1' | jq

read -p "Press enter to end the test"
kill $VICTORIA_PID
