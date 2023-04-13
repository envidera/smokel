#!/usr/bin/env bash
# github.com/anstk - MIT LICENSE

#----------------------------------------------------------------------
# Backup this project to the backup folder in tar.gz format
# If something really goes wrong, you can restore these backups
# 
# EXCLUDE FROM BACKUP
#   bin
#   build
#   tmp
#
# INCLUDE IN BACKUP
#   *    (everything)
#   .*   (everything hidden)
#----------------------------------------------------------------------


backupPath="tmp/backup"


#----------------------------------------------------------------------
source _.sh
#----------------------------------------------------------------------


# mkdir only if not exists
mkdir -p "$backupPath"

# Format: YEAR MONTH DAY - HOUR MINUTE SECOND
DATE=$(date +%Y-%m-%d-%H:%M:%S)

tar -czf "$backupPath/${DATE}.tar.gz" --exclude='bin' --exclude='build' --exclude='tmp' * .*  


#----------------------------------------------------------------------
echoFinish $backupPath
#----------------------------------------------------------------------