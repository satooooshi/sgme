# SGME API


<a name="overview"></a>
## Overview
this is SGME API


### Version information
*Version* : 1.0


### Contact information
*Contact* : Satoshi AIKAWA  
*Contact Email* : support@swagger.io


### License information
*License* : Satoshi AIKAWA  
*License URL* : http://www.apache.org/licenses/LICENSE-2.0.html  
*Terms of service* : null


### URI scheme
*Host* : 0.0.0.0:8080  
*BasePath* : /




<a name="paths"></a>
## Paths

<a name="addserviceapidocurl-ns-svcname-apidocurl-get"></a>
### addServiceApidocUrl
```
GET /addServiceApidocUrl/{ns}/{svcname}/{apidocurl}
```


#### Description
addServiceApidocUrl


#### Parameters

|Type|Name|Description|Schema|
|---|---|---|---|
|**Path**|**apidocurl**  <br>*required*|api document url|string|
|**Path**|**ns**  <br>*required*|namespace|string|
|**Path**|**svcname**  <br>*required*|service name|string|


#### Responses

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|ok|No Content|


#### Consumes

* `application/json`


#### Produces

* `application/json`


<a name="discovergateway-ns-get"></a>
### discoverGateway
```
GET /discoverGateway/{ns}
```


#### Description
discoverGateway


#### Parameters

|Type|Name|Description|Schema|
|---|---|---|---|
|**Path**|**ns**  <br>*required*|namespace|string|


#### Responses

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|ok|No Content|


#### Consumes

* `application/json`


#### Produces

* `application/json`


<a name="discovergatewayport-ns-get"></a>
### discoverGatewayPort
```
GET /discoverGatewayPort/{ns}
```


#### Description
discoverGatewayPort


#### Parameters

|Type|Name|Description|Schema|
|---|---|---|---|
|**Path**|**ns**  <br>*required*|namespace|string|


#### Responses

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|ok|No Content|


#### Consumes

* `application/json`


#### Produces

* `application/json`


<a name="discoverservice-ns-svcname-get"></a>
### discoverService
```
GET /discoverService/{ns}/{svcname}
```


#### Description
discoverService


#### Parameters

|Type|Name|Description|Schema|
|---|---|---|---|
|**Path**|**ns**  <br>*required*|namespace|string|
|**Path**|**svcname**  <br>*required*|service name|string|


#### Responses

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|ok|No Content|


#### Consumes

* `application/json`


#### Produces

* `application/json`


<a name="discoverserviceapidocurl-ns-svcname-get"></a>
### discoverServiceApidocUrl
```
GET /discoverServiceApidocUrl/{ns}/{svcname}
```


#### Description
discoverServiceApidocUrl


#### Parameters

|Type|Name|Description|Schema|
|---|---|---|---|
|**Path**|**ns**  <br>*required*|namespace|string|
|**Path**|**svcname**  <br>*required*|service name|string|


#### Responses

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|ok|No Content|


#### Consumes

* `application/json`


#### Produces

* `application/json`


<a name="discoverservices-ns-get"></a>
### discoverServices
```
GET /discoverServices{ns}
```


#### Description
discoverServices


#### Parameters

|Type|Name|Description|Schema|
|---|---|---|---|
|**Path**|**ns**  <br>*required*|namespace|string|


#### Responses

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|ok|No Content|


#### Consumes

* `application/json`


#### Produces

* `application/json`


<a name="registergateway-ns-ip-get"></a>
### registerGateway
```
GET /registerGateway/{ns}/{ip}
```


#### Description
registerGateway


#### Parameters

|Type|Name|Description|Schema|
|---|---|---|---|
|**Path**|**ip**  <br>*required*|gateway external ip address|string|
|**Path**|**ns**  <br>*required*|namespace|string|


#### Responses

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|ok|No Content|


#### Consumes

* `application/json`


#### Produces

* `application/json`


<a name="registerservice-ns-svcname-get"></a>
### registerService
```
GET /registerService/{ns}/{svcname}
```


#### Description
registerService


#### Parameters

|Type|Name|Description|Schema|
|---|---|---|---|
|**Path**|**ns**  <br>*required*|namespace|string|
|**Path**|**svcname**  <br>*required*|service name|string|


#### Responses

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|ok|No Content|


#### Consumes

* `application/json`


#### Produces

* `application/json`


<a name="test-get"></a>
### GET /test/

#### Description
テスト用APIの詳細


#### Parameters

|Type|Name|Description|Schema|
|---|---|---|---|
|**Query**|**none**  <br>*optional*|必須ではありません。|string|


#### Responses

|HTTP Code|Schema|
|---|---|
|**200**|No Content|


#### Consumes

* `application/x-json-stream`







