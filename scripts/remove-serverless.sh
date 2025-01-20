#!/bin/bash

# Example usage:
# ./scripts/remove-serverless.sh -p "arcadia" -e "develop" -c "CBG" -r "us-west-2"

while getopts p:e:c:r:a: flag
do
    case "${flag}" in
        p) CBG_ACCOUNT_SET=${OPTARG};;
        e) CBG_ENVIRONMENT=${OPTARG};;
        c) CBG_CUSTOMER=${OPTARG};;
        r) AWS_REGION=${OPTARG};;
        *) echo "Not valid";;
    esac
done

# Pass in temporary environment variables to CLI.
CBG_ACCOUNT_SET="$CBG_ACCOUNT_SET" \
CBG_ENVIRONMENT="$CBG_ENVIRONMENT" \
CBG_CUSTOMER="$CBG_CUSTOMER" \
AWS_REGION="$AWS_REGION" \
AWS_PROFILE="$CBG_ACCOUNT_SET-$CBG_ENVIRONMENT" sls remove \
    --region "$AWS_REGION" \
    --stage "$CBG_ENVIRONMENT"
