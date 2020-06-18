package aliyuncvc

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

// CreateUserEvaluations invokes the aliyuncvc.CreateUserEvaluations API synchronously
// api document: https://help.aliyun.com/api/aliyuncvc/createuserevaluations.html
func (client *Client) CreateUserEvaluations(request *CreateUserEvaluationsRequest) (response *CreateUserEvaluationsResponse, err error) {
	response = CreateCreateUserEvaluationsResponse()
	err = client.DoAction(request, response)
	return
}

// CreateUserEvaluationsWithChan invokes the aliyuncvc.CreateUserEvaluations API asynchronously
// api document: https://help.aliyun.com/api/aliyuncvc/createuserevaluations.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateUserEvaluationsWithChan(request *CreateUserEvaluationsRequest) (<-chan *CreateUserEvaluationsResponse, <-chan error) {
	responseChan := make(chan *CreateUserEvaluationsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateUserEvaluations(request)
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

// CreateUserEvaluationsWithCallback invokes the aliyuncvc.CreateUserEvaluations API asynchronously
// api document: https://help.aliyun.com/api/aliyuncvc/createuserevaluations.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateUserEvaluationsWithCallback(request *CreateUserEvaluationsRequest, callback func(response *CreateUserEvaluationsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateUserEvaluationsResponse
		var err error
		defer close(result)
		response, err = client.CreateUserEvaluations(request)
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

// CreateUserEvaluationsRequest is the request struct for api CreateUserEvaluations
type CreateUserEvaluationsRequest struct {
	*requests.RpcRequest
	Memo        string           `position:"Query" name:"Memo"`
	Description string           `position:"Query" name:"Description"`
	CreateDate  requests.Integer `position:"Query" name:"CreateDate"`
	MemberUUID  string           `position:"Query" name:"MemberUUID"`
	UserId      string           `position:"Query" name:"UserId"`
	Evaluation  string           `position:"Query" name:"Evaluation"`
	Score       string           `position:"Query" name:"Score"`
	MeetingUUID string           `position:"Query" name:"MeetingUUID"`
	AppId       string           `position:"Query" name:"AppId"`
}

// CreateUserEvaluationsResponse is the response struct for api CreateUserEvaluations
type CreateUserEvaluationsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateCreateUserEvaluationsRequest creates a request to invoke CreateUserEvaluations API
func CreateCreateUserEvaluationsRequest() (request *CreateUserEvaluationsRequest) {
	request = &CreateUserEvaluationsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aliyuncvc", "2019-09-19", "CreateUserEvaluations", "aliyuncvc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateUserEvaluationsResponse creates a response to parse from CreateUserEvaluations response
func CreateCreateUserEvaluationsResponse() (response *CreateUserEvaluationsResponse) {
	response = &CreateUserEvaluationsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}