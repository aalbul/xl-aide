#!/bin/sh

echo "Copying files from [${deployed.file}] to [${deployed.targetPath}]"
cd ${deployed.targetPath}; tar xfz ${deployed.file}
res=$?
if [ $res != 0 ] ; then
        exit $res
fi


${deployed.symphoneHome}/bin/cache -invalidate

${deployed.tripwireHome}/bin/updateDatabase

echo "Setting permissions on [${deployed.targetPath}] to [${deployed.permissions}]"
chmod -R ${deployed.permissions} ${deployed.targetPath}
res=$?
if [ $res != 0 ] ; then
        exit $res
fi
