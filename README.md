Curlforce
=========

A quick way to get started using [`curl`](http://curl.haxx.se/docs/manpage.html) with the [Salesforce REST API](http://www.salesforce.com/us/developer/docs/api_rest/) minus the OAuth headaches.

Usage
-----

1. Go to https://curlforce.herokuapp.com

2. Log into Salesforce and grant access

3. Get your environment and `curl` command all ready to go:

  ```
  export SFDC_INSTANCE_URL='https://na4.salesforce.com'
  export SFDC_ACCESS_TOKEN='00D700000001234!A5667kzzqh0bxkUtGMJkAMmHPvCMlvMLMErojvh3zxSSG0PoLm.u6Vbt8HP2LdKFp0JuPmCGIwroFlCNhuzFJk_MmRBJY'
  curl -H 'X-PrettyPrint: 1' -H "Authorization: Bearer $SFDC_ACCESS_TOKEN" $SFDC_INSTANCE_URL/services/data
  ```

Just paste that into your terminal to start using the API.
