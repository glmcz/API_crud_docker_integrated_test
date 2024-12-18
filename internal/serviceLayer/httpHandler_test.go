package serviceLayer

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io"
	"net/http"
	"simpleCloudService/internal/model"
	"testing"

	testClient "simpleCloudService/pkg/http"
)

func testData() []struct {
	userReq model.User
} {
	testCases := []struct {
		userReq model.User
	}{ // create a new []
		{
			userReq: model.User{
				ID:          uuid.MustParse("224e9a8e-5571-48d3-9da4-c18a1974e561"),
				Name:        "Franta",
				Email:       "franta@skocDoPole.com",
				DateOfBirth: "1987-08-24", // ISO 8601
			},
		},
		{
			userReq: model.User{
				ID:          uuid.MustParse("774e9a8e-5571-48d3-9da4-c18a1974e565"),
				Name:        "Vitezoslav",
				Email:       "vitezoMir@gmail.com",
				DateOfBirth: "2024-01-25",
			},
		},
		{
			userReq: model.User{
				ID:          uuid.MustParse("784e9a8e-5571-48d3-9da4-c18a1974e564"),
				Name:        "Olmik",
				Email:       "olmik@yahoo.com",
				DateOfBirth: "2000-05-07",
			},
		},
	}
	return testCases
}

func PostEndpointWithHTTPClient(t *testing.T) {
	testData := testData()
	client := &testClient.Client{}
	for _, user := range testData {
		request, err := client.ServerPostRequest("3000", user.userReq, "save")
		if err != nil {
			t.Errorf("Error while posting user: %v", err)
		}
		if request.StatusCode == http.StatusOK {
			fmt.Printf("Successfully saved user [%s]\n", user.userReq.Name)
		} else {
			// need to add proper response that user already exist.
			fmt.Printf("Failed to save user [%s]\n", user.userReq.Name)
		}

		err = request.Body.Close()
		if err != nil {
			t.Errorf("failed to close response body %v", err)
		}
	}
}

func GetEndpointWithHTTPClient(t *testing.T) {
	testData := testData()
	client := &testClient.Client{}
	for _, user := range testData {
		response, err := client.ServerGetRequest("3000", user.userReq.ID.String())
		if err != nil {
			return
		}
		if response.StatusCode == http.StatusOK {
			var serverUser model.User
			if err := readResponse(response.Body, &serverUser); err != nil {
				t.Errorf("failed to read response body %v", err)
			}
			if user.userReq.Name == serverUser.Name {
				fmt.Printf("Successfully get user [%s]\n", serverUser.Name)
			}
		} else {
			// TODO handle non-existing users in DB in case of parallel testing
			// need to add proper response that user already exist.
			t.Errorf("Failed to get user [%s]\n", user.userReq.Name)
		}

		err = response.Body.Close()
		if err != nil {
			t.Errorf("failed to close response body %v", err)
		}
	}
}

func TestHTTPRequest(t *testing.T) {
	t.Run("Post", PostEndpointWithHTTPClient)
	t.Run("Get", GetEndpointWithHTTPClient)
}

func readResponse(body io.ReadCloser, user *model.User) error {
	byteBody, err := io.ReadAll(body)
	if err != nil {
		return fmt.Errorf("failed to read response body %v", err)
	}
	_ = json.Unmarshal(byteBody, &user)
	return nil
}
