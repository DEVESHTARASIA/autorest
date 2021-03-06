package

// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
    "github.com/Azure/go-autorest/autorest"
    "github.com/Azure/go-autorest/autorest/azure"
    "net/http"
    "io"
)

// TextGroupClient is the client for the TextGroup methods of the  service.
type TextGroupClient struct {
    ManagementClient
}
// NewTextGroupClient creates an instance of the TextGroupClient client.
func NewTextGroupClient() TextGroupClient {
        return NewTextGroupClientWithBaseURI(DefaultBaseURI, )
        }

// NewTextGroupClientWithBaseURI creates an instance of the TextGroupClient client.
    func NewTextGroupClientWithBaseURI(baseURI string, ) TextGroupClient {
        return TextGroupClient{ NewWithBaseURI(baseURI, )}
    }

// AMethod sends the a method request.
//
// textParameter is a text stream. textParameter will be closed upon successful return. Callers should ensure closure
// when receiving an error.contentType is the content type of the image.
func (client TextGroupClient) AMethod(textParameter io.ReadCloser, contentType string) (result autorest.Response, err error) {
    req, err := client.AMethodPreparer(textParameter, contentType)
    if err != nil {
        err = autorest.NewErrorWithError(err, ".TextGroupClient", "AMethod", nil , "Failure preparing request")
        return
    }

    resp, err := client.AMethodSender(req)
    if err != nil {
        result.Response = resp
        err = autorest.NewErrorWithError(err, ".TextGroupClient", "AMethod", resp, "Failure sending request")
        return
    }

    result, err = client.AMethodResponder(resp)
    if err != nil {
        err = autorest.NewErrorWithError(err, ".TextGroupClient", "AMethod", resp, "Failure responding to request")
    }

    return
}

// AMethodPreparer prepares the AMethod request.
func (client TextGroupClient) AMethodPreparer(textParameter io.ReadCloser, contentType string) (*http.Request, error) {
    preparer := autorest.CreatePreparer(
                        autorest.AsPost(),
                        autorest.WithBaseURL(client.BaseURI),
                        autorest.WithPath("/ProcessImage/FunctionTA"),
                        autorest.WithFile(textParameter),
                        autorest.WithHeader("Content-Type",autorest.String(contentType)))
    return preparer.Prepare(&http.Request{})
}

// AMethodSender sends the AMethod request. The method will close the
// http.Response Body if it receives an error.
func (client TextGroupClient) AMethodSender(req *http.Request) (*http.Response, error) {
    return autorest.SendWithSender(client, req)
}

// AMethodResponder handles the response to the AMethod request. The method always
// closes the http.Response Body.
func (client TextGroupClient) AMethodResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
            resp,
            client.ByInspecting(),
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByClosing())
    result.Response = resp
    return
}

// BMethod sends the b method request.
//
// textParameter is a text stream. textParameter will be closed upon successful return. Callers should ensure closure
// when receiving an error.contentType is the content type of the image.
func (client TextGroupClient) BMethod(textParameter io.ReadCloser, contentType string) (result autorest.Response, err error) {
    req, err := client.BMethodPreparer(textParameter, contentType)
    if err != nil {
        err = autorest.NewErrorWithError(err, ".TextGroupClient", "BMethod", nil , "Failure preparing request")
        return
    }

    resp, err := client.BMethodSender(req)
    if err != nil {
        result.Response = resp
        err = autorest.NewErrorWithError(err, ".TextGroupClient", "BMethod", resp, "Failure sending request")
        return
    }

    result, err = client.BMethodResponder(resp)
    if err != nil {
        err = autorest.NewErrorWithError(err, ".TextGroupClient", "BMethod", resp, "Failure responding to request")
    }

    return
}

// BMethodPreparer prepares the BMethod request.
func (client TextGroupClient) BMethodPreparer(textParameter io.ReadCloser, contentType string) (*http.Request, error) {
    preparer := autorest.CreatePreparer(
                        autorest.AsPost(),
                        autorest.WithBaseURL(client.BaseURI),
                        autorest.WithPath("/ProcessImage/FunctionTB"),
                        autorest.WithFile(textParameter),
                        autorest.WithHeader("Content-Type",autorest.String(contentType)))
    return preparer.Prepare(&http.Request{})
}

// BMethodSender sends the BMethod request. The method will close the
// http.Response Body if it receives an error.
func (client TextGroupClient) BMethodSender(req *http.Request) (*http.Response, error) {
    return autorest.SendWithSender(client, req)
}

// BMethodResponder handles the response to the BMethod request. The method always
// closes the http.Response Body.
func (client TextGroupClient) BMethodResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
            resp,
            client.ByInspecting(),
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByClosing())
    result.Response = resp
    return
}

