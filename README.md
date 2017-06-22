vRealize Automation:-
      
        vRealize Automation provides UI where developers,users can request IT 
        services and manage IT resources. vRA provider is created newly. It 
        utilizes go library. It uses rest apis of vRA and send http request.   


Resources:-

Execute Blueprint:

     -It executes any blueprint present in VRA.
     -Blueprint name is given in configuaration file.
     -Timeout should be in seconds and its optional i.e user can gives in seconds 
     otherwise its default value is 50 sec.
     -data.json and main.tf should be placed in same folder
     -To execute blueprint json data should pass from user in .json file, to generate 
     that json data we have to set value in vra blueprint and send http request
     to get that template using postman or restclient ,so we get the json format and 
     we can use it anytime only by changing the values in that json data.

The main.tf file contains the microservices of how to call the providers and resources. We need to specify required details for resource creation in this file.