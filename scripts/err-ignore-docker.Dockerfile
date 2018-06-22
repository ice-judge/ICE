FROM plugins/docker:17

ENTRYPOINT [": ||", "/usr/local/bin/dockerd-entrypoint.sh", ": ||", "/bin/drone-docker" ]


