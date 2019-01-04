#!/bin/bash
set -e

if [ -z "$1" ]; then
   echo "invalid args. please input zimzua-api"
   exit 1
fi

if [[ $EUID -ne 0 ]]; then
   echo "setup.sh must be run as root"
   exit 1
fi

if [ "$(arch)" != "x86_64" ]; then
  echo "$(arch) architecture not support"
  exit 1
fi

if [ ! -d /etc/rsyslog.d ]; then
  echo "/etc/rsyslog.d directory not found"
  exit 1
fi

LOG_PATH=/var/log/$1.log
echo "========================================================"
echo "Please input $1 Configuration"
echo "========================================================"
if [ "$1" == "nsom-web" ]; then
    echo "example)"
    echo "BASE_DIR : /opt/zimzua-api"
    echo "DEVELOPMENT_MODE : false"
    echo "========================================================="
    echo ""
    echo -n "BASE_DIR : "
    read INSTALL_DIR
    echo -n "DEVELOPMENT_MODE : "
    read DEV_MODE

    # chmod
    chmod +x zimzua-api/*.sh
    chmod +x zimzua-api/zimzua-api

    # copy files
    mkdir -p $INSTALL_DIR
    echo "completed mkdir($INSTALL_DIR)"
    cp -R $1/* $INSTALL_DIR/
    echo "completed copy files"

    # write execute shell file
    EXEC_PATH=$INSTALL_DIR/$1.sh
    echo "#!/bin/bash" > $EXEC_PATH
    echo "export DEVELOPEMENT_MODE='$DEV_MODE'" >> $EXEC_PATH
    echo "cd /opt/$1" >> $EXEC_PATH
    echo "$INSTALL_DIR/$1 >> $LOG_PATH" >> $EXEC_PATH
else
    echo "invalid argument : $1"
    exit 1
fi

# chmod
chmod +x *.sh
echo "set configuration complete"

# service config file setting
mkdir -p $1-config
echo "completed mkdir $1-config"
cp config/* $1-config/
echo "completed copy config files"
sed -i -- "s~changeme~$EXEC_PATH~g" $1-config/*.conf
echo "completed replace $1 path"

cd $(dirname $0)
if [ -d /etc/systemd/system ] && [ "x$(pidof systemd)" != "x"  ]; then
  # Systemd (CentOS >= 7, Ubuntu >= 15.04)
  echo "========================================================"
  echo "Systemd Setting"
  echo "========================================================"
  cp $1-config/systemd.conf /etc/systemd/system/$1.service
  echo "copy complete service configuration file"
  systemctl enable $1
  echo "systemctl enable complete"
  systemctl start $1
  echo "started $1"
  exit 0
fi

if [ -d /etc/init ]; then
  # Upstart (CentOS 6, Ubuntu >= 6.10)
  echo "========================================================"
  echo "Upstart Setting"
  echo "========================================================"
  if [ -f /etc/init.d/sshd ]; then
    cp $1-config/upstart-centos.conf /etc/init/$1.conf
    echo "copy complete upstart configuration file"
  else
    cp $1-config/upstart-ubuntu.conf /etc/init/$1.conf
    echo "copy complete upstart configuration file"
  fi
  
  start $1
  echo "started $1"
  cp $1-config/rsyslog.conf /etc/rsyslog.d/90-$1.conf
  echo "copy complete rsyslog configuration file"
  /etc/init.d/rsyslog restart
  echo "restarted rsyslog"
  exit 0
fi

# CentOS < 6, old Ubuntu < 6.10
echo "systemd or upstart not installed"
exit 1

