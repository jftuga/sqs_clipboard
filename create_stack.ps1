$qName=$args[0]
if($qName -notmatch "\.fifo$") {
    "queue name must end in .fifo"
    exit
}
$now = get-date -format "yyyyMMddHHmmss"
$sName="sqsClipboard$now"
$templateName="cloud_formation_template.json"

""
"[queue name] $qName"
"[stack name] $sName"

aws cloudformation create-stack --stack-name $sName --parameters ParameterKey=qName,ParameterValue=$qName --template-body "file://./$templateName"

#$sName="sqsClipboard20201129215855"

$i=0
while($i -lt 15) {
    $qURL=(aws cloudformation describe-stacks --stack-name $sName --query "Stacks[0].Outputs[?OutputKey=='QueueURL'].OutputValue" --output text) -join "`n"
    if($qURL.length -gt 30) {
        "[queue url] $qURL"
        ""
        "Please update your environment by running:"
        ""
        "setx SQS_CLIPBOARD_URL $qURL"
        ""
        "See also:"
        "https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-envvars.html#envvars-set"
        ""
        break
    }
    Start-Sleep 2
    $i += 1
    "waiting for stack creation to complete..."
}