package api

// import (
// 	"net/http/httptest"
// 	"testing"

// 	mockdb "github.com/akmal4410/simple_bank/db/mock"
// 	db "github.com/akmal4410/simple_bank/db/sqlc"
// 	"github.com/golang/mock/gomock"
// )

// func TestCreateTransfer(t *testing.T) {
// 	amount := int64(10)

// 	user1, _ := randomUser(t)
// 	user2, _ := randomUser(t)
// 	user3, _ := randomUser(t)

// 	account1 := randomAccount(user1.UserName)
// 	account2 := randomAccount(user2.UserName)
// 	account3 := randomAccount(user3.UserName)

// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockStore := mockdb.NewMockStore(ctrl)
// 	// build stubs---
// 	mockStore.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account1.ID)).Times(1).Return(account1, nil)
// 	mockStore.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account2.ID)).Times(1).Return(account2, nil)

// 	arg := db.TransferTxParams{
// 		FromAccountID: account1.ID,
// 		ToAccountID:   account2.ID,
// 		Amount:        amount,
// 	}
// 	mockStore.EXPECT().TransferTx(gomock.Any(), gomock.Eq(arg)).Times(1)

// 	server := newTestServer(t, mockStore)
// 	recorder := httptest.NewRecorder()
// }
