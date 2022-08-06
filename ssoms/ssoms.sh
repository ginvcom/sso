#!/bin/bash
arg = $1
if arg -z
cd api
goctl api ts -api ssoms.api -dir ./ssoms/vue/src/api -caller ssoms -webapi ../utils/ajax -unwrap
goctl api go -api ssoms.api -dir .

cd model
goctl model mysql ddl -src ./sql/*.sql -dir .