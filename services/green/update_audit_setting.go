package green

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// UpdateAuditSetting invokes the green.UpdateAuditSetting API synchronously
// api document: https://help.aliyun.com/api/green/updateauditsetting.html
func (client *Client) UpdateAuditSetting(request *UpdateAuditSettingRequest) (response *UpdateAuditSettingResponse, err error) {
	response = CreateUpdateAuditSettingResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateAuditSettingWithChan invokes the green.UpdateAuditSetting API asynchronously
// api document: https://help.aliyun.com/api/green/updateauditsetting.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateAuditSettingWithChan(request *UpdateAuditSettingRequest) (<-chan *UpdateAuditSettingResponse, <-chan error) {
	responseChan := make(chan *UpdateAuditSettingResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateAuditSetting(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// UpdateAuditSettingWithCallback invokes the green.UpdateAuditSetting API asynchronously
// api document: https://help.aliyun.com/api/green/updateauditsetting.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateAuditSettingWithCallback(request *UpdateAuditSettingRequest, callback func(response *UpdateAuditSettingResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateAuditSettingResponse
		var err error
		defer close(result)
		response, err = client.UpdateAuditSetting(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// UpdateAuditSettingRequest is the request struct for api UpdateAuditSetting
type UpdateAuditSettingRequest struct {
	*requests.RpcRequest
	Seed       string `position:"Query" name:"Seed"`
	AuditRange string `position:"Query" name:"AuditRange"`
	SourceIp   string `position:"Query" name:"SourceIp"`
	Callback   string `position:"Query" name:"Callback"`
}

// UpdateAuditSettingResponse is the response struct for api UpdateAuditSetting
type UpdateAuditSettingResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateAuditSettingRequest creates a request to invoke UpdateAuditSetting API
func CreateUpdateAuditSettingRequest() (request *UpdateAuditSettingRequest) {
	request = &UpdateAuditSettingRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Green", "2017-08-23", "UpdateAuditSetting", "green", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateAuditSettingResponse creates a response to parse from UpdateAuditSetting response
func CreateUpdateAuditSettingResponse() (response *UpdateAuditSettingResponse) {
	response = &UpdateAuditSettingResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
