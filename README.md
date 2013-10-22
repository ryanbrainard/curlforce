Salesforce API Quickstart
=========================

A quick way of getting started with the Salesforce API without the OAuth headaches.

This app takes care of OAuth and provides the environment and commands to jump into using [`curl`](http://curl.haxx.se/docs/manpage.html) or [`restforce`](https://github.com/ejholmes/restforce) with the Salesforce API.

Usage
-----

1. Go to https://salesforce-api-quickstart.herokuapp.com.

2. Log into Salesforce and grant access.

3. Get the envionment and commands to get started:

  ```
  export SFDC_INSTANCE_URL='https://na4.salesforce.com'
  export SFDC_ACCESS_TOKEN='00D700000001234!A5667kzzqh0bxkUtGMJkAMmHPvCMlvMLMErojvh3zxSSG0PoLm.u6Vbt8HP2LdKFp0JuPmCGIwroFlCNhuzFJk_MmRBJY'
  
  curl -H 'X-PrettyPrint: 1' -H "Authorization: Bearer $SFDC_ACCESS_TOKEN" $SFDC_INSTANCE_URL/services/data
  
  sudo gem install restforce
  irb
  require 'restforce'
  client = Restforce.new(:oauth_token => ENV['SFDC_ACCESS_TOKEN'], :instance_url  => ENV['SFDC_INSTANCE_URL'])
  ```

4. Paste what you need into your Terminal.
