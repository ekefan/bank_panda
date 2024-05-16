package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	mockdb "github.com/ekefan/bank_panda/db/mock"
	db "github.com/ekefan/bank_panda/db/sqlc"
	util "github.com/ekefan/bank_panda/utils"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)


func randomAccount() db.Account {
	return db.Account{
		ID: util.RandomInt(1, 1000),
		Owner: util.RandomOwner(),
		Balance: util.RandomBalance(),
		Currency: util.RandomCurrency(),
	}
}
func TestGetAccountAPI(t *testing.T) {
	account := randomAccount()

	// creating a table of test cases to cover
	testCases := []struct{
		name string
		accountID int64
		buildStubs func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:"ok",
			accountID : account.ID,
			buildStubs: func(mockStore *mockdb.MockStore) {
				mockStore.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account.ID)).Times(1).Return(account, nil) //input and output params should match the ones described in the Querier interface
				},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			requireBodyMatchAccount(t, recorder.Body, account)		
				},
		},
	}

	for i := range testCases {
		tc := testCases[i]
  
		t.Run(tc.name, func(t *testing.T){
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockStore := mockdb.NewMockStore(ctrl)
			tc.buildStubs(mockStore)

		// //build stubs
		// mockStore.EXPECT().
		// GetAccount(gomock.Any(), gomock.Eq(account.ID)).
		// Times(1).
		// Return(account, nil) //input and output params should match the ones described in the Querier interface


		// start test server and send request
			server := NewServer(mockStore)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/accounts/%v", account.ID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
			// require.Equal(t, http.StatusOK, recorder.Code)
			// requireBodyMatchAccount(t, recorder.Body, account)
			})
	}

}


func requireBodyMatchAccount(t *testing.T, body *bytes.Buffer, account db.Account) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)
	var getAccount db.Account
	err = json.Unmarshal(data, &getAccount)
	require.NoError(t, err)
	require.Equal(t, account, getAccount)
}