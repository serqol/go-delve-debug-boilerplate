#! /bin/bash

go build -gcflags="all=-N -l" -o /src/app
dlv --listen=:40000 --log --headless=true --api-version=2 --accept-multiclient exec /src/app