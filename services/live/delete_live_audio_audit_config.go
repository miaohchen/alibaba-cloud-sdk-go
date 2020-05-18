package live

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

// DeleteLiveAudioAuditConfig invokes the live.DeleteLiveAudioAuditConfig API synchronously
// api document: https://help.aliyun.com/api/live/deleteliveaudioauditconfig.html
func (client *Client) DeleteLiveAudioAuditConfig(request *DeleteLiveAudioAuditConfigRequest) (response *DeleteLiveAudioAuditConfigResponse, err error) {
	response = CreateDeleteLiveAudioAuditConfigResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteLiveAudioAuditConfigWithChan invokes the live.DeleteLiveAudioAuditConfig API asynchronously
// api document: https://help.aliyun.com/api/live/deleteliveaudioauditconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteLiveAudioAuditConfigWithChan(request *DeleteLiveAudioAuditConfigRequest) (<-chan *DeleteLiveAudioAuditConfigResponse, <-chan error) {
	responseChan := make(chan *DeleteLiveAudioAuditConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteLiveAudioAuditConfig(request)
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

// DeleteLiveAudioAuditConfigWithCallback invokes the live.DeleteLiveAudioAuditConfig API asynchronously
// api document: https://help.aliyun.com/api/live/deleteliveaudioauditconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteLiveAudioAuditConfigWithCallback(request *DeleteLiveAudioAuditConfigRequest, callback func(response *DeleteLiveAudioAuditConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteLiveAudioAuditConfigResponse
		var err error
		defer close(result)
		response, err = client.DeleteLiveAudioAuditConfig(request)
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

// DeleteLiveAudioAuditConfigRequest is the request struct for api DeleteLiveAudioAuditConfig
type DeleteLiveAudioAuditConfigRequest struct {
	*requests.RpcRequest
	AppName    string           `position:"Query" name:"AppName"`
	StreamName string           `position:"Query" name:"StreamName"`
	DomainName string           `position:"Query" name:"DomainName"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
}

// DeleteLiveAudioAuditConfigResponse is the response struct for api DeleteLiveAudioAuditConfig
type DeleteLiveAudioAuditConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteLiveAudioAuditConfigRequest creates a request to invoke DeleteLiveAudioAuditConfig API
func CreateDeleteLiveAudioAuditConfigRequest() (request *DeleteLiveAudioAuditConfigRequest) {
	request = &DeleteLiveAudioAuditConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DeleteLiveAudioAuditConfig", "live", "openAPI")
	return
}

// CreateDeleteLiveAudioAuditConfigResponse creates a response to parse from DeleteLiveAudioAuditConfig response
func CreateDeleteLiveAudioAuditConfigResponse() (response *DeleteLiveAudioAuditConfigResponse) {
	response = &DeleteLiveAudioAuditConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
