Generate JSON data

  To generate JSON data we need to request template of required blueprint, after getting that template for first time we can use it for multiple times by changing required fields.
1) Get token:-

    method - Post

    url - https://+Host+/identity/api/tokens/

    headers – Accept : applcation/json, Content-Type : application/json

    data - {&quot;username&quot;:&quot;+ userName +&quot;,&quot;password&quot;:&quot;+ password +&quot;,&quot;tenant&quot;:&quot;+ tenant +&quot;}
If the response status code is 200 then take the id field from response body as authorization token($Token) which is required for next operations.

2) Get blueprint ID

   method - Get

   url - https://+Host+/catalog-service/api/consumer/entitledCatalogItems/?$filter=name eq

          &#39;blueprint\_name&#39;

   headers – Accept : applcation/json, Content-Type : application/json,

                   Authorization: Bearer $Token

 get the id field from response body (Blueprint\_ID) used to get template and to the send request
3) Get Template :

   method - Get

   url – https://+Host+/catalogservice/api/consumer/entitledCatalogItems/Blueprint\_ID/

            requests/template

   headers – Accept : applcation/json, Content-Type : application/json,

                   Authorization: Bearer $Token

After getting template we can use it for multiple times by changing required field, depends on the blueprint
i.e If you have created blueprint to create instance on vCenter, then you can change name of that instance, machine resources, storage, network configuration etc.

You can check that template by sending request

4) Send Request :

    method- Post

    url - https://+Host+/catalog-service/api/consumer/entitledCatalogItems/Blueprint\_ID/requests

    headers – Accept : applcation/json, Content-Type : application/json

                     Authorization: Bearer $Token

    data -Template
Json Data for blueprint rhel7

{

"type": "com.vmware.vcac.catalog.domain.request.CatalogItemProvisioningRequest",

"catalogItemId": "f48dbb26-ce8e-4ca3-9e2c-edc4bc4cb4a8",

"requestedFor": "vraadmin@vsphere.local",

"businessGroupId": "3a3c8770-fec1-42e2-b41b-77c3d84681da",

"description": null,

"reasons": null,

"data": {

&quot;Payal\_R&quot;: {

  &quot;componentTypeId&quot;: &quot;com.vmware.csp.component.cafe.composition&quot;,

  &quot;componentId&quot;: null,

  &quot;classId&quot;: &quot;Blueprint.Component.Declaration&quot;,

  &quot;typeFilter&quot;: &quot;rhel7\*Payal\_R&quot;,

  &quot;data&quot;: {

    &quot;Extensibility.Lifecycle.Properties.VMPSMasterWorkflow32.BuildingMachine &quot;: &quot;\_\_\*,\*&quot;,

    &quot;Extensibility.Lifecycle.Properties.VMPSMasterWorkflow32.Requested&quot;: &quot;\_\_\*,\*&quot;,

    &quot;\_allocation&quot;: {

      &quot;componentTypeId&quot;: &quot;com.vmware.csp.iaas.blueprint.service&quot;,

      &quot;componentId&quot;: null,

      &quot;classId&quot;: &quot;Infrastructure.Compute.Machine.Allocation&quot;,

      &quot;typeFilter&quot;: null,

      &quot;data&quot;: {

        &quot;machines&quot;: [

          {

            &quot;componentTypeId&quot;: &quot;com.vmware.csp.iaas.blueprint.service&quot;,

            &quot;componentId&quot;: null,

            &quot;classId&quot;: &quot;Infrastructure.Compute.Machine.Allocation.Machine&quot;,

            &quot;typeFilter&quot;: null,

            &quot;data&quot;: {

              &quot;machine\_id&quot;: &quot;&quot;,

              &quot;nics&quot;: [

                {

                  &quot;componentTypeId&quot;: &quot;com.vmware.csp.iaas.blueprint.service&quot;,

                  &quot;componentId&quot;: null,

                  &quot;classId&quot;: &quot;Infrastructure.Compute.Machine.Nic&quot;,

                  &quot;typeFilter&quot;: null,

                  &quot;data&quot;: {

                    &quot;address&quot;: &quot;&quot;,

                    &quot;assignment\_type&quot;: &quot;Static&quot;,

                    &quot;external\_address&quot;: &quot;&quot;,

                    &quot;id&quot;: null,

                    &quot;load\_balancing&quot;: null,

                    &quot;network&quot;: null,

                    &quot;network\_profile&quot;: null

                  }

                }

              ]

            }

          }

        ]

      }

    },

    &quot;\_cluster&quot;: 1,

    &quot;\_hasChildren&quot;: false,

    &quot;cpu&quot;: 4,

    &quot;datacenter\_location&quot;: null,

    &quot;description&quot;: &quot;rhelOS-7 server&quot;,

    &quot;disks&quot;: [

      {

        &quot;componentTypeId&quot;: &quot;com.vmware.csp.iaas.blueprint.service&quot;,

        &quot;componentId&quot;: null,

        &quot;classId&quot;: &quot;Infrastructure.Compute.Machine.MachineDisk&quot;,

        &quot;typeFilter&quot;: null,

        &quot;data&quot;: {

          &quot;capacity&quot;: 100,

          &quot;custom\_properties&quot;: null,

          &quot;id&quot;: 1502433487366,

          &quot;initial\_location&quot;: &quot;&quot;,

          &quot;is\_clone&quot;: true,

          &quot;label&quot;: &quot;Hard disk 1&quot;,

          &quot;storage\_reservation\_policy&quot;: &quot;&quot;,

          &quot;userCreated&quot;: false,

          &quot;volumeId&quot;: 0

        }

      }

    ],

    &quot;guest\_customization\_specification&quot;: &quot;rhel7&quot;,

    &quot;max\_network\_adapters&quot;: -1,

    &quot;max\_per\_user&quot;: 0,

    &quot;max\_volumes&quot;: 60,

    &quot;memory&quot;: 4096,

    &quot;nics&quot;: [

      {

        &quot;componentTypeId&quot;: &quot;com.vmware.csp.iaas.blueprint.service&quot;,

        &quot;componentId&quot;: null,

        &quot;classId&quot;: &quot;Infrastructure.Compute.Machine.Nic&quot;,

        &quot;typeFilter&quot;: null,

        &quot;data&quot;: {

          &quot;address&quot;: &quot;&quot;,

          &quot;assignment\_type&quot;: &quot;Static&quot;,

          &quot;external\_address&quot;: &quot;&quot;,

          &quot;id&quot;: null,

          &quot;load\_balancing&quot;: null,

          &quot;network&quot;: null,

          &quot;network\_profile&quot;: null

        }

      }

    ],

    &quot;os\_arch&quot;: &quot;x86\_64&quot;,

    &quot;os\_distribution&quot;: null,

    &quot;os\_type&quot;: &quot;Linux&quot;,

    &quot;os\_version&quot;: null,

    &quot;property\_groups&quot;: null,

    &quot;reservation\_policy&quot;: null,

    &quot;security\_groups&quot;: [],

    &quot;security\_tags&quot;: [],

    &quot;source\_machine\_external\_snapshot&quot;: null,

    &quot;source\_machine\_vmsnapshot&quot;: null,

    &quot;storage&quot;: 100

  }

},

&quot;\_leaseDays&quot;: 1,

&quot;\_number\_of\_instances&quot;: 1
}

}