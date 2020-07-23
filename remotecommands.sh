pushd backend
go build
ps -ef | grep "backend" | grep -v grep | awk '{print $2}' | xargs kill
nohup ./backend &
popd 

# exit ssh
exit
