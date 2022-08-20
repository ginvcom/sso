#!/bin/bash
arg = $1
if arg -z
cd api
goctl api ts -api ssoms.api -dir ../vue/src/api -caller ssoms -webapi @/config -unwrap
goctl api go -api ssoms.api -dir .

cd model
goctl model mysql ddl -src ./sql/*.sql -dir .

cd auth/api
goctl api ts -api auth.api -dir ../../ssoms/vue/src/api -caller auth -webapi @/config -unwrap 