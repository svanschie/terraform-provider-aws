---
subcategory: "API Gateway"
layout: "aws"
page_title: "AWS: aws_api_gateway_api_key"
description: |-
  Get information on an API Gateway REST API Key
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_api_gateway_api_key

Use this data source to get the name and value of a pre-existing API Key, for
example to supply credentials for a dependency microservice.

## Example Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.data_aws_api_gateway_api_key import DataAwsApiGatewayApiKey
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        DataAwsApiGatewayApiKey(self, "my_api_key",
            id="ru3mpjgse6"
        )
```

## Argument Reference

* `id` - (Required) ID of the API Key to look up.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `id` - Set to the ID of the API Key.
* `name` - Set to the name of the API Key.
* `value` - Set to the value of the API Key.
* `created_date` - Date and time when the API Key was created.
* `last_updated_date` - Date and time when the API Key was last updated.
* `description` - Description of the API Key.
* `enabled` - Whether the API Key is enabled.
* `tags` - Map of tags for the resource.

<!-- cache-key: cdktf-0.18.0 input-17297a64af50979746acd59feb4bc637e441edcf9aac1743e07b3d713dd934c5 -->