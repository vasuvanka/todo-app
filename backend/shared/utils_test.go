package shared_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vasuvanka/todo-app/backend/shared"
)


func TestCreateJWT(t *testing.T) {
	jwt, err := shared.CreateJWT("id","test","user")
	assert.Nil(t,err,"error should be nil")
	assert.NotEqual(t,jwt,"","jwt should not be empty")
}

func TestGetTokenFromHeader(t *testing.T){
	var r = http.Request{
		Header: map[string][]string{
			"Authorization": {"bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTQyNDc3OTAsImlkIjoiNjAzMzZmMWE1Njk1Y2EyZTJiNzEzNTVjIiwibmFtZSI6IlRlc3QgVXNlciIsInJvbGUiOiJ1c2VyIn0.COKPpjb_UpLRZgPf7GSxmyEN8p7i_BrWO7rzukXcdvI"},
		},
	}
	
	token, err := shared.GetTokenFromHeader(&r)
	assert.Nil(t,err,"error should be nil")
	assert.Equal(t,token,"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTQyNDc3OTAsImlkIjoiNjAzMzZmMWE1Njk1Y2EyZTJiNzEzNTVjIiwibmFtZSI6IlRlc3QgVXNlciIsInJvbGUiOiJ1c2VyIn0.COKPpjb_UpLRZgPf7GSxmyEN8p7i_BrWO7rzukXcdvI","Token mismatch")
}

func TestGetTokenFromHeaderFail(t *testing.T){
	var r = http.Request{
		Header: map[string][]string{
			"Authorization": {"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTQyNDc3OTAsImlkIjoiNjAzMzZmMWE1Njk1Y2EyZTJiNzEzNTVjIiwibmFtZSI6IlRlc3QgVXNlciIsInJvbGUiOiJ1c2VyIn0.COKPpjb_UpLRZgPf7GSxmyEN8p7i_BrWO7rzukXcdvI"},
		},
	}
	
	_, err := shared.GetTokenFromHeader(&r)
	assert.NotNil(t,err,"error should not be nil")
	assert.Error(t,err,"Unauthorized access")
}

func TestGetTokenFromHeaderFailForNoAuthHeader(t *testing.T){
	var r = http.Request{
		Header: map[string][]string{},
	}
	
	_, err := shared.GetTokenFromHeader(&r)
	assert.Error(t,err,"Unauthorized access")
}