package afcclient

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

//GetAllUsers returns list of Users
func (c Client) GetAllUsers() ([]User, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/Users/getAllUsers", c.HostURL), nil)
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

	return Users, nil
}

//CreateUser will create an User
func (c *Client) CreateUser(User User) (*User, error) {
	avg, err := json.Marshal(User)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/Users/createNewUser", c.HostURL), strings.NewReader(string(avg)))
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

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/Users/updateUserByName", c.HostURL), strings.NewReader(string(avg)))
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
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/Users/deleteUserByName", c.HostURL), http.NoBody)
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
