package afcclient

//external binary to talk to AFC
import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

//GetAllUsers returns list of Users
func (c Client) GetAllUsers() ([]User, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/users", c.HostURL), nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprint("Bearer", c.Token))
	//req.Header.Set("Authorization", "Bearer b7dcf700f34a22fc9c20b0ebb40f57d44ce941c103c4496a2c5dd0e57e1cb37c5eb850b3ebc7427ac47e6895364274bb")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprint("Bearer", c.Token))

	if err != nil {
		return nil, err
	}

	body, err := c.DoRequest(req)
	if err != nil {
		return nil, err
	}

	var Users []User
	err = json.Unmarshal(body, &Users)
	if err != nil {
		return nil, err
	}
	fmt.Println("response Headers:", body)
	return Users, nil
}

//CreateUser will create a User
func (c *Client) CreateUser(User User) (*User, error) {
	avg, err := json.Marshal(User)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/users", c.HostURL), strings.NewReader(string(avg)))
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}

	body, err := c.DoRequest(req)
	if err != nil {
		return nil, err
	}
	var insertedID InsertedResult
	err = json.Unmarshal(body, &insertedID)
	if err != nil {
		return nil, err
	}
	User.ID = insertedID.InsertedID
	return &User, nil
}

//UpdateUserByName will update an User
func (c *Client) UpdateUserByName(User User) (*UpdateResult, error) {
	avg, err := json.Marshal(User)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/api/v1/users/updateUserByName", c.HostURL), strings.NewReader(string(avg)))
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}

	body, err := c.DoRequest(req)
	if err != nil {
		return nil, err
	}

	var updateResult UpdateResult
	err = json.Unmarshal(body, &updateResult)
	if err != nil {
		return nil, err
	}

	return &updateResult, nil
}

//DeleteUserByName will delete an User
func (c *Client) DeleteUserByName(UserName string) (*DeleteResult, error) {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/api/v1/users/deleteUserByName", c.HostURL), http.NoBody)
	req.URL.Query().Add("name", UserName)
	if err != nil {
		return nil, err
	}

	body, err := c.DoRequest(req)
	if err != nil {
		return nil, err
	}

	var deleteResult DeleteResult
	err = json.Unmarshal(body, &deleteResult)
	if err != nil {
		return nil, err
	}

	return &deleteResult, nil
}
