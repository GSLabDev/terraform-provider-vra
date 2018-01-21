Generate JSON data

  To generate JSON data we need to request template of required blueprint, after getting that template for first time we can use it for multiple times by changing required fields.

1) Get token:-

    method - Post

    url - https://+Host+/identity/api/tokens/

    headers – Accept : applcation/json, Content-Type : application/json

    data - {"username":"+ userName +","password":"+ password +","tenant":"+ tenant +"}

If the response status code is 200 then take the id field from response body as authorization token($Token) which is required for next operations.

2) Get blueprint ID :-

   method - Get

   url - https://+Host+/catalog-service/api/consumer/entitledCatalogItems/?$filter=name eq 'blueprint_name' 

   headers – Accept : applcation/json, Content-Type : application/json and Authorization: Bearer $Token

 Get the id field from response body (Blueprint\_ID) used to get template and to the send request
 
3) Get Template :-

   method - Get

   url – https://+Host+/catalogservice/api/consumer/entitledCatalogItems/Blueprint_ID/requests/template

   headers – Accept : applcation/json, Content-Type : application/json and Authorization: Bearer $Token

After getting template we can use it for multiple times by changing required field, depends on the blueprint
i.e If you have created blueprint to create instance on vCenter, then you can change name of that instance, machine resources, storage, network configuration etc.
You can check that template by sending request

4) Send Request :-

    method- Post

    url - https://+Host+/catalog-service/api/consumer/entitledCatalogItems/Blueprint_ID/requests

    headers – Accept : applcation/json, Content-Type : application/json and Authorization: Bearer $Token