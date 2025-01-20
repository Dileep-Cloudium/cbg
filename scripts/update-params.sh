#!/bin/bash

# Example usage:
# ./scripts/update-params.sh -p "arcadia-develop" -r "us-west-2" -s "develop"
# ./scripts/update-params.sh -p "arcadia-test" -r "us-west-2" -s "test"
# ./scripts/update-params.sh -p "arcadia-prod" -r "us-west-2" -s "prod"

while getopts p:r:a:s: flag
do
    case "${flag}" in
        p) profile=${OPTARG};;
        r) region=${OPTARG};;
        s) stage=${OPTARG};;
        *) echo "Not valid";;
    esac
done

if [[ $stage == 'develop' ]]; then
    securityGroupIds="sg-04f65b2ef3057d7e8"
    subnetIds="subnet-080aae4acb3a424da,subnet-032a0d5dc445733a9,subnet-0cf7436b715c9eeca"
    domainName="cbgapps-develop.io"
    certificateArn="arn:aws:acm:us-west-2:241533163333:certificate/f5224de7-81e3-4a88-af64-631a27d266ed"
    hostedZoneId="Z029729718BRLAXVLKV59"
    graphqlEndpoint="https://slipspace-coreapi-np.hasura.app/v1/graphql"
    graphqlApiKey="${GRAPHQL_API_KEY:-???}"

elif [[ $stage == 'test' ]]; then
    securityGroupIds="sg-02874700334706269"
    subnetIds="subnet-050546de0f19f234b,subnet-08bf1cb4f0f6472a8,subnet-0c30132370e610d49"
    domainName="cbgapps-test.io"
    certificateArn="arn:aws:acm:us-west-2:841162688595:certificate/9eb6971e-b29d-42ad-8d44-9262df512d98"
    hostedZoneId="Z01179455UQEKRJ5JSPR"
    graphqlEndpoint="https://slipspace-coreapi-np.hasura.app/v1/graphql"
    graphqlApiKey="${GRAPHQL_API_KEY:-???}"

elif [[ $stage == 'prod' ]]; then
    securityGroupIds="sg-01f8dafaa154738b1"
    subnetIds="subnet-04a324fc6cff9e3a0,subnet-0af30639904594ce3,subnet-037cc3626252eb77a"
    domainName="cbgapps.io"
    certificateArn="arn:aws:acm:us-west-2:650251721418:certificate/86b99465-a7c1-4104-a4e9-903f2e422b80"
    hostedZoneId="Z011569026N0YZR0F891F"
    graphqlEndpoint="https://slipspace-coreapi-p.hasura.app/v1/graphql"
    graphqlApiKey="${GRAPHQL_API_KEY:-???}"

else
    echo "Not valid account or stage"
fi

if  [[ $graphqlApiKey == '???' ]]; then
    echo "API KEYS AND SECRETS CANNOT BE ???."
    echo "SET THEM CORRECTLY IN THE SCRIPT, THEN TRY EXECUTING AGAIN."
    exit 1
fi

aws ssm put-parameter --profile "$profile" \
    --region "$region" \
    --name "/$stage/service/core/nexusgate-pbmapi/config/SECURITY_GROUP_IDS" \
    --type "String" \
    --value "$securityGroupIds" \
    --overwrite

aws ssm put-parameter --profile "$profile" \
    --region "$region" \
    --name "/$stage/service/core/nexusgate-pbmapi/config/SUBNET_IDS" \
    --type "String" \
    --value "$subnetIds" \
    --overwrite

aws ssm put-parameter --profile "$profile" \
    --region "$region" \
    --name "/$stage/service/core/nexusgate-pbmapi/config/DOMAIN_NAME" \
    --type "String" \
    --value "$domainName" \
    --overwrite

aws ssm put-parameter --profile "$profile" \
    --region "$region" \
    --name "/$stage/service/core/nexusgate-pbmapi/config/CERTIFICATE_ARN" \
    --type "String" \
    --value "$certificateArn" \
    --overwrite

aws ssm put-parameter --profile "$profile" \
    --region "$region" \
    --name "/$stage/service/core/nexusgate-pbmapi/config/HOSTED_ZONE_ID" \
    --type "String" \
    --value "$hostedZoneId" \
    --overwrite

aws ssm put-parameter --profile "$profile" \
    --region "$region" \
    --name "/$stage/service/core/nexusgate-pbmapi/config/GRAPHQL_ENDPOINT" \
    --type "String" \
    --value "$graphqlEndpoint" \
    --overwrite

aws ssm put-parameter --profile "$profile" \
    --region "$region" \
    --name "/$stage/service/core/nexusgate-pbmapi/config/GRAPHQL_API_KEY" \
    --type "SecureString" \
    --value "$graphqlApiKey" \
    --overwrite
