package trademark

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

// QueryMaterialList invokes the trademark.QueryMaterialList API synchronously
// api document: https://help.aliyun.com/api/trademark/querymateriallist.html
func (client *Client) QueryMaterialList(request *QueryMaterialListRequest) (response *QueryMaterialListResponse, err error) {
	response = CreateQueryMaterialListResponse()
	err = client.DoAction(request, response)
	return
}

// QueryMaterialListWithChan invokes the trademark.QueryMaterialList API asynchronously
// api document: https://help.aliyun.com/api/trademark/querymateriallist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryMaterialListWithChan(request *QueryMaterialListRequest) (<-chan *QueryMaterialListResponse, <-chan error) {
	responseChan := make(chan *QueryMaterialListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryMaterialList(request)
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

// QueryMaterialListWithCallback invokes the trademark.QueryMaterialList API asynchronously
// api document: https://help.aliyun.com/api/trademark/querymateriallist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryMaterialListWithCallback(request *QueryMaterialListRequest, callback func(response *QueryMaterialListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryMaterialListResponse
		var err error
		defer close(result)
		response, err = client.QueryMaterialList(request)
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

// QueryMaterialListRequest is the request struct for api QueryMaterialList
type QueryMaterialListRequest struct {
	*requests.RpcRequest
	Name       string           `position:"Query" name:"Name"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	CardNumber string           `position:"Query" name:"CardNumber"`
	Type       requests.Integer `position:"Query" name:"Type"`
	Region     requests.Integer `position:"Query" name:"Region"`
	PageNum    requests.Integer `position:"Query" name:"PageNum"`
	Status     requests.Integer `position:"Query" name:"Status"`
}

// QueryMaterialListResponse is the response struct for api QueryMaterialList
type QueryMaterialListResponse struct {
	*responses.BaseResponse
	RequestId      string                  `json:"RequestId" xml:"RequestId"`
	TotalItemNum   int                     `json:"TotalItemNum" xml:"TotalItemNum"`
	CurrentPageNum int                     `json:"CurrentPageNum" xml:"CurrentPageNum"`
	PageSize       int                     `json:"PageSize" xml:"PageSize"`
	TotalPageNum   int                     `json:"TotalPageNum" xml:"TotalPageNum"`
	Data           DataInQueryMaterialList `json:"Data" xml:"Data"`
}

// CreateQueryMaterialListRequest creates a request to invoke QueryMaterialList API
func CreateQueryMaterialListRequest() (request *QueryMaterialListRequest) {
	request = &QueryMaterialListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Trademark", "2018-07-24", "QueryMaterialList", "trademark", "openAPI")
	return
}

// CreateQueryMaterialListResponse creates a response to parse from QueryMaterialList response
func CreateQueryMaterialListResponse() (response *QueryMaterialListResponse) {
	response = &QueryMaterialListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}