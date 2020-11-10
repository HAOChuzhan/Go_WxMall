package v1

//用户登录
/*
func UserLogin(c *gin.Context) {
	state1 := c.DefaultPostForm("state", "")
	code := c.DefaultPostForm("code", "")

	state2 := make(map[string]interface{})
	if err := json.Unmarshal([]byte(state1), &state2); err != nil {
		return
	}

	//err := json.NewDecoder(state1).Decode(&state2)
	branch_id := state2["branch_id"]

	admin := admin.Admin{
		BranchID: branch_id.(string),
	}
	login_result := admin.AdminLogin(code, branch_id)

}
*/
