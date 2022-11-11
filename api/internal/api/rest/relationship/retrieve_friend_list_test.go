package relationship

import "testing"

func TestRetrieveFriendList(t *testing.T) {
	type retrieveFriendList struct {
	}
	type arg struct {
		retrieveFriendList           retrieveFriendList
		retrieveFriendListMockCalled bool
		givenBody                    string
		expRs                        string
		expHTTPCode                  int
	}
}
