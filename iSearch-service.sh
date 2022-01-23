#!/bin/sh
# copy to /usr/local/etc/rc.d

case "$1" in
	start)
		/volume1/linux/bin/iSearch/iSearch-backend &
	;;
	stop)
		pkill iSearch-backend
	;;
	*)
		echo "Usage: $0 [start|stop]"
		exit 1
	;;
esac
exit 0
