#!/bin/bash

# Give the SQS queue name that want to create for SQS Clipboard as the only command-line argument

qName=$1
now=$(date +"%Y%m%d%H%M%S")
sName="sqsClipboard${now}"
templateName="cloud_formation_template.json"
if [[ $qName != *.fifo ]] ; then
	echo "queue name must end in: .fifo"
	exit 1
fi

echo ""
echo "[queue name] $qName"
echo "[stack name] $sName"
echo ""

aws cloudformation create-stack --stack-name ${sName} --parameters ParameterKey=qName,ParameterValue=${qName} --template-body "file://./${templateName}"

i=0
while [ $i -lt 15 ] ; do
    qURL=$(aws cloudformation describe-stacks --stack-name ${sName} --query "Stacks[0].Outputs[?OutputKey=='QueueURL'].OutputValue" --output text 2> /dev/null)
    if [ "${qURL}" != "None" ] ; then
        echo "[queue url] $qURL"
        echo ""
        echo "Please add this to your '.bash_profile' file:"
        echo
        echo "export SQS_CLIPBOARD_URL='${qURL}'"
        echo
        echo "See also:"
        echo "https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-envvars.html#envvars-set"
        break
    else
        sleep 2
        let i=${i}+1
        echo "waiting for stack creation to complete..."
    fi
done

