SITE_BACKEND_NAME=backend
SITE_FRONTEND_NAME=http-server

pushd backend
go build
ps -ef | grep "${SITE_BACKEND}" | grep -v grep | awk '{print $2}' | xargs kill
nohup ./backend &
popd 

cd ..
pushd dist
ps -ef | grep "${SITE_FRONTEND}" | grep -v grep | awk '{print $2}' | xargs kill
nohup http-server -p 80 &
popd

# exit ssh
exit