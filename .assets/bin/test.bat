@ ECHO OFF

for /f "" %%G in ('go list ./... ^| find /i /v "/vendor/"') do (go test -race %%G & IF ERRORLEVEL == 1 EXIT 1)