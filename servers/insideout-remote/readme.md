# InsideOut (Riley)

AI infrastructure design agent by Luther Systems. Designs, prices, and deploys production-ready cloud infrastructure on AWS or GCP from plain English.

- **Documentation**: https://insideout.luthersystems.com
- **Discord**: https://insideout.luthersystems.com/discord
- **Pricing**: https://insideout.luthersystems.com/pricing
- **Provider**: [Luther Systems](https://luthersystems.com)

## What it does

Describe your app in natural language and Riley:

1. Designs the cloud architecture (components, networking, security)
2. Generates production-ready, modular Terraform configurations
3. Estimates monthly costs broken down by component
4. Validates your cloud credentials
5. Deploys to AWS or GCP via Terraform with live log streaming
6. Inspects and manages live infrastructure after deployment

## Supported clouds and components

**AWS**: VPC, ALB, ECS, EKS, RDS, ElastiCache, CloudFront, S3, Route53, ACM, WAF, SES, Lambda, API Gateway, DynamoDB, Cognito, Bedrock, OpenSearch, SQS

**GCP**: VPC, Cloud Run, GKE, Cloud SQL, Memorystore, Cloud CDN, Cloud Storage, Cloud DNS

## Auth

No auth required for design and pricing. Cloud credentials (AWS IAM role or GCP service account) are only required at deploy time and stay encrypted on the backend.
